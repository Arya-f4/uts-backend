package service

import (
	"fmt"
	"golang-train/app/model"
	"golang-train/app/repository"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pekerjaanService struct {
	repo repository.PekerjaanRepository
}

func NewPekerjaanService(repo repository.PekerjaanRepository) PekerjaanService {
	return &pekerjaanService{repo: repo}
}

func (s *pekerjaanService) CreatePekerjaan(c *fiber.Ctx) error {
	var req model.CreatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	alumniObjID, err := primitive.ObjectIDFromHex(req.AlumniID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Alumni ID tidak valid"})
	}

	tanggalMulai, err := time.Parse("2006-01-02", req.TanggalMulaiKerja)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("format tanggal mulai kerja tidak valid: %v", err)})
	}
	var tanggalSelesai *time.Time
	if req.TanggalSelesaiKerja != nil && *req.TanggalSelesaiKerja != "" {
		parsed, err := time.Parse("2006-01-02", *req.TanggalSelesaiKerja)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("format tanggal selesai kerja tidak valid: %v", err)})
		}
		tanggalSelesai = &parsed
	}

	pekerjaan := &model.Pekerjaan{
		AlumniID:            alumniObjID,
		NamaPerusahaan:      req.NamaPerusahaan,
		PosisiJabatan:       req.PosisiJabatan,
		BidangIndustri:      req.BidangIndustri,
		LokasiKerja:         req.LokasiKerja,
		GajiRange:           req.GajiRange,
		TanggalMulaiKerja:   tanggalMulai,
		TanggalSelesaiKerja: tanggalSelesai,
		StatusPekerjaan:     req.StatusPekerjaan,
		DeskripsiPekerjaan:  req.DeskripsiPekerjaan,
	}

	createdPekerjaan, err := s.repo.Create(c.Context(), pekerjaan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createdPekerjaan)
}

func (s *pekerjaanService) GetAllPekerjaan(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	params := model.PaginationParams{
		Page:   page,
		Limit:  limit,
		Sort:   c.Query("sort", "created_at:desc"),
		Search: c.Query("search", ""),
	}

	result, err := s.repo.FindAll(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (s *pekerjaanService) GetAllPekerjaanDeleted(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	params := model.PaginationParams{
		Page:   page,
		Limit:  limit,
		Sort:   c.Query("sort", "created_at:desc"),
		Search: c.Query("search", ""),
	}

	result, err := s.repo.FindAllDeleted(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (s *pekerjaanService) GetPekerjaanByID(c *fiber.Ctx) error {
	id := c.Params("id") // ID is now a string
	pekerjaan, err := s.repo.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pekerjaan)
}

func (s *pekerjaanService) UpdatePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id") // ID is now a string

	var req model.UpdatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	// We fetch first to get the existing AlumniID, or let the repo handle a partial update
	// For this migration, we'll pass a model with only the updated fields.
	// The repository implementation will use $set, so AlumniID won't be touched.

	tanggalMulai, err := time.Parse("2006-01-02", req.TanggalMulaiKerja)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("format tanggal mulai kerja tidak valid: %v", err)})
	}
	var tanggalSelesai *time.Time
	if req.TanggalSelesaiKerja != nil && *req.TanggalSelesaiKerja != "" {
		parsed, err := time.Parse("2006-01-02", *req.TanggalSelesaiKerja)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("format tanggal selesai kerja tidak valid: %v", err)})
		}
		tanggalSelesai = &parsed
	}

	// Create a Pekerjaan model with updated fields
	pekerjaan := &model.Pekerjaan{
		NamaPerusahaan:      req.NamaPerusahaan,
		PosisiJabatan:       req.PosisiJabatan,
		BidangIndustri:      req.BidangIndustri,
		LokasiKerja:         req.LokasiKerja,
		GajiRange:           req.GajiRange,
		TanggalMulaiKerja:   tanggalMulai,
		TanggalSelesaiKerja: tanggalSelesai,
		StatusPekerjaan:     req.StatusPekerjaan,
		DeskripsiPekerjaan:  req.DeskripsiPekerjaan,
	}

	updatedPekerjaan, err := s.repo.Update(c.Context(), id, pekerjaan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updatedPekerjaan)
}

func (s *pekerjaanService) DeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := s.repo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pekerjaan dihapus secara permanen"})
}

func (s *pekerjaanService) SoftDeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := s.repo.SoftDelete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pekerjaan berhasil dihapus"})
}

func (s *pekerjaanService) RestorePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := s.repo.Restore(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pekerjaan berhasil dipulihkan"})
}
