
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

type mahasiswaRepository struct {
	db *pgxpool.Pool
}

func NewMahasiswaRepository(db *pgxpool.Pool) MahasiswaRepository {
	return &mahasiswaRepository{db: db}
}

func (r *mahasiswaRepository) Create(ctx context.Context, m *model.Mahasiswa) (*model.Mahasiswa, error) {
	query := `INSERT INTO mahasiswa (nim, nama, jurusan, angkatan, email)
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, m.NIM, m.Nama, m.Jurusan, m.Angkatan, m.Email).Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *mahasiswaRepository) FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Mahasiswa], error) {
	var args []interface{}
	var whereClauses []string
	argID := 1

	baseQuery := `SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at FROM mahasiswa`
	countQuery := `SELECT COUNT(id) FROM mahasiswa`

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

	var total int64
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	validSortColumns := map[string]string{
		"nama":       "nama",
		"nim":        "nim",
		"angkatan":   "angkatan",
		"created_at": "created_at",
	}
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

	offset := (params.Page - 1) * params.Limit
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, params.Limit, offset)

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mahasiswaList := []model.Mahasiswa{}
	for rows.Next() {
		var m model.Mahasiswa
		if err := rows.Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan, &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		mahasiswaList = append(mahasiswaList, m)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	return &model.PaginationResult[model.Mahasiswa]{
		Data:     mahasiswaList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}, nil
}

func (r *mahasiswaRepository) FindByID(ctx context.Context, id int) (*model.Mahasiswa, error) {
	var m model.Mahasiswa
	query := `SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at FROM mahasiswa WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan, &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("mahasiswa tidak ditemukan")
		}
		return nil, err
	}
	return &m, nil
}

func (r *mahasiswaRepository) Update(ctx context.Context, m *model.Mahasiswa) (*model.Mahasiswa, error) {
	query := `UPDATE mahasiswa SET nama=$1, jurusan=$2, angkatan=$3, email=$4, updated_at=NOW()
              WHERE id=$5 RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, m.Nama, m.Jurusan, m.Angkatan, m.Email, m.ID).Scan(&m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *mahasiswaRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM mahasiswa WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("tidak ada baris yang ditemukan untuk dihapus")
	}
	return nil
}


