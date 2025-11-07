package repository

import (
	"context"
	"golang-train/app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User, roleName string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	Delete(ctx context.Context, id string) error
	Restore(ctx context.Context, id string) error
}

type AlumniRepository interface {
	Create(ctx context.Context, alumni *model.Alumni) (*model.Alumni, error)
	FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error)
	FindAllDeleted(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error) // Note: Soft delete not implemented for Alumni
	FindByID(ctx context.Context, id string) (*model.Alumni, error)
	Update(ctx context.Context, id string, alumni *model.Alumni) (*model.Alumni, error)
	Delete(ctx context.Context, id string) error
}

type MahasiswaRepository interface {
	Create(ctx context.Context, mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error)
	FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Mahasiswa], error)
	FindByID(ctx context.Context, id string) (*model.Mahasiswa, error)
	Update(ctx context.Context, id string, mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error)
	Delete(ctx context.Context, id string) error
}

type PekerjaanRepository interface {
	Create(ctx context.Context, pekerjaan *model.Pekerjaan) (*model.Pekerjaan, error)
	FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error)
	FindAllDeleted(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error)
	FindByID(ctx context.Context, id string) (*model.Pekerjaan, error)
	Update(ctx context.Context, id string, pekerjaan *model.Pekerjaan) (*model.Pekerjaan, error)
	Delete(ctx context.Context, id string) error     // Hard delete
	SoftDelete(ctx context.Context, id string) error // Soft delete
	Restore(ctx context.Context, id string) error    // Restore soft-deleted
}

// Interface baru untuk repositori media
// Satu interface generik dapat digunakan oleh kedua repositori (foto dan sertifikat)
type MediaRepository interface {
	// UpsertByUserID akan membuat atau memperbarui media untuk user_id tertentu
	UpsertByUserID(ctx context.Context, media *model.Media) (*model.Media, error)
	// GetByUserID mengambil media berdasarkan user_id
	GetByUserID(ctx context.Context, userID primitive.ObjectID) (*model.Media, error)
}
