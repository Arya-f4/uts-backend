
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

func (s *alumniService) GetAlumniByID(ctx context.Context, id int) (*model.Alumni, error) {
	return s.alumniRepo.FindByID(ctx, id)
}

func (s *alumniService) UpdateAlumni(ctx context.Context, id int, req *model.UpdateAlumniRequest) (*model.Alumni, error) {
	alumni, err := s.alumniRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	alumni.Nama = req.Nama
	alumni.Jurusan = req.Jurusan
	alumni.Angkatan = req.Angkatan
	alumni.TahunLulus = req.TahunLulus
	alumni.Email = req.Email
	alumni.NoTelepon = req.NoTelepon
	alumni.Alamat = req.Alamat

	return s.alumniRepo.Update(ctx, alumni)
}

func (s *alumniService) DeleteAlumni(ctx context.Context, id int) error {
	return s.alumniRepo.Delete(ctx, id)
}


