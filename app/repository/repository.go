package repository

import (
	"context"
	"golang-train/app/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User, roleName string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	Delete(ctx context.Context, id int) error
	Restore(ctx context.Context, id int) error
}

type AlumniRepository interface {
	Create(ctx context.Context, alumni *model.Alumni) (*model.Alumni, error)
	FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error)
	FindAllDeleted(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error)
	FindByID(ctx context.Context, id int) (*model.Alumni, error)
	Update(ctx context.Context, alumni *model.Alumni) (*model.Alumni, error)
	Delete(ctx context.Context, id int) error
}

type MahasiswaRepository interface {
	Create(ctx context.Context, mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error)
	FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Mahasiswa], error)
	FindByID(ctx context.Context, id int) (*model.Mahasiswa, error)
	Update(ctx context.Context, mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error)
	Delete(ctx context.Context, id int) error
}

type PekerjaanRepository interface {
	Create(ctx context.Context, pekerjaan *model.Pekerjaan) (*model.Pekerjaan, error)
	FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error)
	FindAllDeleted(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error)
	FindByID(ctx context.Context, id int) (*model.Pekerjaan, error)
	Update(ctx context.Context, pekerjaan *model.Pekerjaan) (*model.Pekerjaan, error)
	Delete(ctx context.Context, id int) error     // Hard delete
	SoftDelete(ctx context.Context, id int) error // Soft delete
	Restore(ctx context.Context, id int) error    // Restore soft-deleted
}
