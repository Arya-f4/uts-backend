package service

import (
	"context"
	"golang-train/app/model"
	"golang-train/app/repository"
)

type alumniService struct {
	alumniRepo repository.AlumniRepository
}

func NewAlumniService(ar repository.AlumniRepository) AlumniService {
	return &alumniService{alumniRepo: ar}
}

func (s *alumniService) CreateAlumni(ctx context.Context, req *model.CreateAlumniRequest) (*model.Alumni, error) {
	alumni := &model.Alumni{
		NIM:        req.NIM,
		Nama:       req.Nama,
		Jurusan:    req.Jurusan,
		Angkatan:   req.Angkatan,
		TahunLulus: req.TahunLulus,
		Email:      req.Email,
		NoTelepon:  req.NoTelepon,
		Alamat:     req.Alamat,
	}
	return s.alumniRepo.Create(ctx, alumni)
}

func (s *alumniService) GetAllAlumni(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error) {
	return s.alumniRepo.FindAll(ctx, params)
}

func (s *alumniService) GetAlumniByID(ctx context.Context, id string) (*model.Alumni, error) {
	return s.alumniRepo.FindByID(ctx, id)
}

func (s *alumniService) UpdateAlumni(ctx context.Context, id string, req *model.UpdateAlumniRequest) (*model.Alumni, error) {
	// Find is not necessary here, repository update can handle it
	// Or, if we need to merge, we find first. Let's let repo handle it.
	alumni := &model.Alumni{
		Nama:       req.Nama,
		Jurusan:    req.Jurusan,
		Angkatan:   req.Angkatan,
		TahunLulus: req.TahunLulus,
		Email:      req.Email,
		NoTelepon:  req.NoTelepon,
		Alamat:     req.Alamat,
	}

	return s.alumniRepo.Update(ctx, id, alumni)
}

func (s *alumniService) DeleteAlumni(ctx context.Context, id string) error {
	return s.alumniRepo.Delete(ctx, id)
}
