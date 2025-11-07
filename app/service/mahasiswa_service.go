package service

import (
	"context"
	"golang-train/app/model"
	"golang-train/app/repository"
)

type mahasiswaService struct {
	mahasiswaRepo repository.MahasiswaRepository
}

func NewMahasiswaService(mr repository.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{mahasiswaRepo: mr}
}

func (s *mahasiswaService) CreateMahasiswa(ctx context.Context, req *model.CreateMahasiswaRequest) (*model.Mahasiswa, error) {
	mahasiswa := &model.Mahasiswa{
		NIM:      req.NIM,
		Nama:     req.Nama,
		Jurusan:  req.Jurusan,
		Angkatan: req.Angkatan,
		Email:    req.Email,
	}
	return s.mahasiswaRepo.Create(ctx, mahasiswa)
}

func (s *mahasiswaService) GetAllMahasiswa(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Mahasiswa], error) {
	return s.mahasiswaRepo.FindAll(ctx, params)
}

func (s *mahasiswaService) GetMahasiswaByID(ctx context.Context, id string) (*model.Mahasiswa, error) {
	return s.mahasiswaRepo.FindByID(ctx, id)
}

func (s *mahasiswaService) UpdateMahasiswa(ctx context.Context, id string, req *model.UpdateMahasiswaRequest) (*model.Mahasiswa, error) {
	mahasiswa := &model.Mahasiswa{
		Nama:     req.Nama,
		Jurusan:  req.Jurusan,
		Angkatan: req.Angkatan,
		Email:    req.Email,
	}

	return s.mahasiswaRepo.Update(ctx, id, mahasiswa)
}

func (s *mahasiswaService) DeleteMahasiswa(ctx context.Context, id string) error {
	return s.mahasiswaRepo.Delete(ctx, id)
}
