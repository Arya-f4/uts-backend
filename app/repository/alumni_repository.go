// File Path: app/repository/alumni_repository.go
package repository

import (
	"context"
	"errors"
	"fmt"
	"golang-train/app/model"
	"math"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// alumniRepository mengimplementasikan interface AlumniRepository.
type alumniRepository struct {
	db *pgxpool.Pool
}

// NewAlumniRepository adalah constructor untuk alumniRepository.

// Create menyimpan data alumni baru ke database.
func (r *alumniRepository) Create(ctx context.Context, alumni *model.Alumni) (*model.Alumni, error) {
	query := `INSERT INTO alumni (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
              RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, alumni.NIM, alumni.Nama, alumni.Jurusan, alumni.Angkatan, alumni.TahunLulus, alumni.Email, alumni.NoTelepon, alumni.Alamat).Scan(&alumni.ID, &alumni.CreatedAt, &alumni.UpdatedAt)
	return alumni, err
}

// FindAll mengambil daftar alumni dengan paginasi, sorting, dan searching.
func (r *alumniRepository) FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error) {
	var args []interface{}
	var whereClauses []string
	argID := 1

	baseQuery := `SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at FROM alumni`
	countQuery := `SELECT COUNT(id) FROM alumni`

	// Menambahkan klausa WHERE untuk fungsionalitas pencarian.
	if params.Search != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("(nama ILIKE $%d OR nim ILIKE $%d OR jurusan ILIKE $%d OR email ILIKE $%d)", argID, argID, argID, argID))
		args = append(args, "%"+params.Search+"%")
		argID++
	}

	if len(whereClauses) > 0 {
		whereSQL := " WHERE " + strings.Join(whereClauses, " AND ")
		baseQuery += whereSQL
		countQuery += whereSQL
	}

	// Mendapatkan total jumlah data untuk paginasi.
	var total int64
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, err
	}

	// Menambahkan klausa ORDER BY untuk sorting.
	validSortColumns := map[string]string{"nama": "nama", "nim": "nim", "angkatan": "angkatan", "tahun_lulus": "tahun_lulus", "created_at": "created_at"}
	sortColumn := "created_at"
	sortOrder := "DESC"
	if params.Sort != "" {
		parts := strings.Split(params.Sort, ":")
		if mappedCol, ok := validSortColumns[strings.ToLower(parts[0])]; ok {
			sortColumn = mappedCol
		}
		if len(parts) > 1 && strings.ToUpper(parts[1]) == "ASC" {
			sortOrder = "ASC"
		}
	}
	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortColumn, sortOrder)

	// Menambahkan klausa LIMIT dan OFFSET untuk paginasi.
	offset := (params.Page - 1) * params.Limit
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, params.Limit, offset)

	// Menjalankan query utama untuk mendapatkan data.
	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	alumniList := []model.Alumni{}
	for rows.Next() {
		var a model.Alumni
		if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		alumniList = append(alumniList, a)
	}

	// Menghitung halaman terakhir.
	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	return &model.PaginationResult[model.Alumni]{
		Data:     alumniList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}, nil
}

// FindByID mencari satu alumni berdasarkan ID.
func (r *alumniRepository) FindByID(ctx context.Context, id int) (*model.Alumni, error) {
	var a model.Alumni
	query := `SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at FROM alumni WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan, &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("alumni tidak ditemukan")
		}
		return nil, err
	}
	return &a, nil
}

// Update memperbarui data alumni di database.
func (r *alumniRepository) Update(ctx context.Context, alumni *model.Alumni) (*model.Alumni, error) {
	query := `UPDATE alumni SET nama=$1, jurusan=$2, angkatan=$3, tahun_lulus=$4, email=$5, no_telepon=$6, alamat=$7, updated_at=NOW()
              WHERE id=$8 RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, alumni.Nama, alumni.Jurusan, alumni.Angkatan, alumni.TahunLulus, alumni.Email, alumni.NoTelepon, alumni.Alamat, alumni.ID).Scan(&alumni.UpdatedAt)
	return alumni, err
}

// Delete menghapus data alumni dari database.
func (r *alumniRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM alumni WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("tidak ada baris yang ditemukan untuk dihapus")
	}
	return nil
}
