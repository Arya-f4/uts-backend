package service

import (
	"context"
	"golang-train/app/model"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.User, error)
	Login(ctx context.Context, req *model.LoginRequest) (string, error)
}

type UserService interface {
	DeleteUser(ctx context.Context, requesterID string, requesterRoles []string, targetUserID string) error
	RestoreUser(ctx context.Context, requesterID string, requesterRoles []string, targetUserID string) error
}

type AlumniService interface {
	CreateAlumni(ctx context.Context, req *model.CreateAlumniRequest) (*model.Alumni, error)
	GetAllAlumni(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error)
	GetAlumniByID(ctx context.Context, id string) (*model.Alumni, error)
	UpdateAlumni(ctx context.Context, id string, req *model.UpdateAlumniRequest) (*model.Alumni, error)
	DeleteAlumni(ctx context.Context, id string) error
}

type MahasiswaService interface {
	CreateMahasiswa(ctx context.Context, req *model.CreateMahasiswaRequest) (*model.Mahasiswa, error)
	GetAllMahasiswa(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Mahasiswa], error)
	GetMahasiswaByID(ctx context.Context, id string) (*model.Mahasiswa, error)
	UpdateMahasiswa(ctx context.Context, id string, req *model.UpdateMahasiswaRequest) (*model.Mahasiswa, error)
	DeleteMahasiswa(ctx context.Context, id string) error
}

type PekerjaanService interface {
	CreatePekerjaan(c *fiber.Ctx) error
	GetAllPekerjaan(c *fiber.Ctx) error
	GetAllPekerjaanDeleted(c *fiber.Ctx) error
	GetPekerjaanByID(c *fiber.Ctx) error
	UpdatePekerjaan(c *fiber.Ctx) error
	DeletePekerjaan(c *fiber.Ctx) error
	SoftDeletePekerjaan(c *fiber.Ctx) error
	RestorePekerjaan(c *fiber.Ctx) error
}

// Interface baru untuk servis media
type MediaService interface {
	UploadMedia(c *fiber.Ctx) error
	GetMedia(c *fiber.Ctx) error
}
