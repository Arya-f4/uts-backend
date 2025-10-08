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

type pekerjaanRepository struct {
	db *pgxpool.Pool
}

// NewPekerjaanRepository creates a new instance of PekerjaanRepository.
func NewPekerjaanRepository(db *pgxpool.Pool) PekerjaanRepository {
	return &pekerjaanRepository{db: db}
}

// Create inserts a new pekerjaan record into the database.
func (r *pekerjaanRepository) Create(ctx context.Context, p *model.Pekerjaan) (*model.Pekerjaan, error) {
	query := `INSERT INTO pekerjaan (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
              RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(ctx, query, p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan, p.DeskripsiPekerjaan).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}

// FindAll retrieves a paginated list of non-deleted pekerjaan records.
func (r *pekerjaanRepository) FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error) {
	var args []interface{}
	var whereClauses []string
	argID := 1

	baseQuery := `SELECT p.id, p.alumni_id, p.nama_perusahaan, p.posisi_jabatan, p.bidang_industri, p.lokasi_kerja, p.gaji_range, p.tanggal_mulai_kerja, p.tanggal_selesai_kerja, p.status_pekerjaan, p.deskripsi_pekerjaan, p.created_at, p.updated_at FROM pekerjaan p`
	countQuery := `SELECT COUNT(p.id) FROM pekerjaan p`

	// Always filter for non-deleted records
	whereClauses = append(whereClauses, "p.is_deleted IS NULL")

	if params.Search != "" {
		baseQuery += " JOIN alumni a ON p.alumni_id = a.id"
		countQuery += " JOIN alumni a ON p.alumni_id = a.id"
		// Add search conditions for multiple fields
		searchClause := fmt.Sprintf("(p.nama_perusahaan ILIKE $%d OR p.posisi_jabatan ILIKE $%d OR p.bidang_industri ILIKE $%d OR a.nama ILIKE $%d)", argID, argID, argID, argID)
		whereClauses = append(whereClauses, searchClause)
		args = append(args, "%"+params.Search+"%")
		argID++
	}

	if len(whereClauses) > 0 {
		whereSQL := " WHERE " + strings.Join(whereClauses, " AND ")
		baseQuery += whereSQL
		countQuery += whereSQL
	}

	var total int64
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, err
	}

	validSortColumns := map[string]string{"nama_perusahaan": "p.nama_perusahaan", "posisi_jabatan": "p.posisi_jabatan", "tanggal_mulai_kerja": "p.tanggal_mulai_kerja", "created_at": "p.created_at"}
	sortColumn := "p.created_at"
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

	var pekerjaanList []model.Pekerjaan
	for rows.Next() {
		var p model.Pekerjaan
		if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pekerjaanList = append(pekerjaanList, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	return &model.PaginationResult[model.Pekerjaan]{
		Data:     pekerjaanList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}, nil
}

// FindByID finds a single pekerjaan by its ID.
func (r *pekerjaanRepository) FindByID(ctx context.Context, id int) (*model.Pekerjaan, error) {
	var p model.Pekerjaan
	query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at 
			  FROM pekerjaan WHERE id = $1 AND is_deleted IS NULL`
	err := r.db.QueryRow(ctx, query, id).Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("pekerjaan tidak ditemukan")
		}
		return nil, err
	}
	return &p, nil
}

// Update modifies an existing pekerjaan record.
func (r *pekerjaanRepository) Update(ctx context.Context, p *model.Pekerjaan) (*model.Pekerjaan, error) {
	query := `UPDATE pekerjaan SET nama_perusahaan=$1, posisi_jabatan=$2, bidang_industri=$3, lokasi_kerja=$4, gaji_range=$5, tanggal_mulai_kerja=$6, tanggal_selesai_kerja=$7, status_pekerjaan=$8, deskripsi_pekerjaan=$9, updated_at=NOW()
              WHERE id=$10 RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan, p.DeskripsiPekerjaan, p.ID).Scan(&p.UpdatedAt)
	return p, err
}

// Delete permanently removes a record from the database.
func (r *pekerjaanRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM pekerjaan WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("tidak ada baris yang ditemukan untuk dihapus")
	}
	return nil
}

// SoftDelete marks a record as deleted by setting the 'is_deleted' timestamp.
func (r *pekerjaanRepository) SoftDelete(ctx context.Context, id int) error {
	query := `UPDATE pekerjaan SET is_deleted = NOW() WHERE id = $1 AND is_deleted IS NULL`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("pekerjaan tidak ditemukan atau sudah dihapus")
	}
	return nil
}

// Restore reverts a soft-deleted record by setting 'is_deleted' back to NULL.
func (r *pekerjaanRepository) Restore(ctx context.Context, id int) error {
	query := `UPDATE pekerjaan SET is_deleted = NULL, updated_at = NOW() WHERE id = $1 AND is_deleted IS NOT NULL`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("pekerjaan tidak ditemukan di data yang dihapus")
	}
	return nil
}

// FindAllDeleted retrieves a paginated list of soft-deleted pekerjaan records.
func (r *pekerjaanRepository) FindAllDeleted(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error) {
	var args []interface{}
	var whereClauses []string
	argID := 1

	baseQuery := `SELECT p.id, p.alumni_id, p.nama_perusahaan, p.posisi_jabatan, p.bidang_industri, p.lokasi_kerja, p.gaji_range, p.tanggal_mulai_kerja, p.tanggal_selesai_kerja, p.status_pekerjaan, p.deskripsi_pekerjaan, p.created_at, p.updated_at FROM pekerjaan p`
	countQuery := `SELECT COUNT(p.id) FROM pekerjaan p`

	// Always filter for soft-deleted records
	whereClauses = append(whereClauses, "p.is_deleted IS NOT NULL")

	if params.Search != "" {
		baseQuery += " JOIN alumni a ON p.alumni_id = a.id"
		countQuery += " JOIN alumni a ON p.alumni_id = a.id"
		searchClause := fmt.Sprintf("(p.nama_perusahaan ILIKE $%d OR p.posisi_jabatan ILIKE $%d OR p.bidang_industri ILIKE $%d OR a.nama ILIKE $%d)", argID, argID, argID, argID)
		whereClauses = append(whereClauses, searchClause)
		args = append(args, "%"+params.Search+"%")
		argID++
	}

	if len(whereClauses) > 0 {
		whereSQL := " WHERE " + strings.Join(whereClauses, " AND ")
		baseQuery += whereSQL
		countQuery += whereSQL
	}

	var total int64
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, err
	}

	validSortColumns := map[string]string{"nama_perusahaan": "p.nama_perusahaan", "posisi_jabatan": "p.posisi_jabatan", "tanggal_mulai_kerja": "p.tanggal_mulai_kerja", "created_at": "p.created_at"}
	sortColumn := "p.created_at"
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

	var pekerjaanList []model.Pekerjaan
	for rows.Next() {
		var p model.Pekerjaan
		if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		pekerjaanList = append(pekerjaanList, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	return &model.PaginationResult[model.Pekerjaan]{
		Data:     pekerjaanList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}, nil
}
