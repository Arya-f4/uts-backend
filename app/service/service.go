package service

import (
	"context"
	"golang-train/app/model"

	"github.com/gofiber/fiber/v2" // Tambahkan import fiber
)

// ... (Interface AuthService, UserService, AlumniService, MahasiswaService tetap sama) ...
type AuthService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.User, error)
	Login(ctx context.Context, req *model.LoginRequest) (string, error)
}

type UserService interface {
	DeleteUser(ctx context.Context, requesterID int, requesterRoles []string, targetUserID int) error
	RestoreUser(ctx context.Context, requesterID int, requesterRoles []string, targetUserID int) error
}

type AlumniService interface {
	CreateAlumni(ctx context.Context, req *model.CreateAlumniRequest) (*model.Alumni, error)
	GetAllAlumni(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error)
	GetAlumniByID(ctx context.Context, id int) (*model.Alumni, error)
	UpdateAlumni(ctx context.Context, id int, req *model.UpdateAlumniRequest) (*model.Alumni, error)
	DeleteAlumni(ctx context.Context, id int) error
}

type MahasiswaService interface {
	CreateMahasiswa(ctx context.Context, req *model.CreateMahasiswaRequest) (*model.Mahasiswa, error)
	GetAllMahasiswa(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Mahasiswa], error)
	GetMahasiswaByID(ctx context.Context, id int) (*model.Mahasiswa, error)
	UpdateMahasiswa(ctx context.Context, id int, req *model.UpdateMahasiswaRequest) (*model.Mahasiswa, error)
	DeleteMahasiswa(ctx context.Context, id int) error
}

// Interface PekerjaanService diperbarui untuk menangani fiber.Ctx secara langsung
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
