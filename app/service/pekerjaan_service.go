package service

import (
	"fmt"
	"golang-train/app/model"
	"golang-train/app/repository"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2" // Tambahkan import fiber
)

type pekerjaanService struct {
	repo repository.PekerjaanRepository
}

func NewPekerjaanService(repo repository.PekerjaanRepository) PekerjaanService {
	return &pekerjaanService{repo: repo}
}

// CreatePekerjaan sekarang menangani HTTP request dan response
func (s *pekerjaanService) CreatePekerjaan(c *fiber.Ctx) error {
	var req model.CreatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	// Logika bisnis tetap di sini
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
		AlumniID:            req.AlumniID,
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
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	pekerjaan, err := s.repo.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(pekerjaan)
}

func (s *pekerjaanService) UpdatePekerjaan(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var req model.UpdatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	pekerjaan, err := s.repo.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Logika bisnis pembaruan data
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

	pekerjaan.NamaPerusahaan = req.NamaPerusahaan
	pekerjaan.PosisiJabatan = req.PosisiJabatan
	pekerjaan.BidangIndustri = req.BidangIndustri
	pekerjaan.LokasiKerja = req.LokasiKerja
	pekerjaan.GajiRange = req.GajiRange
	pekerjaan.TanggalMulaiKerja = tanggalMulai
	pekerjaan.TanggalSelesaiKerja = tanggalSelesai
	pekerjaan.StatusPekerjaan = req.StatusPekerjaan
	pekerjaan.DeskripsiPekerjaan = req.DeskripsiPekerjaan

	updatedPekerjaan, err := s.repo.Update(c.Context(), pekerjaan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updatedPekerjaan)
}

func (s *pekerjaanService) DeletePekerjaan(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	if err := s.repo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pekerjaan dihapus secara permanen"})
}

func (s *pekerjaanService) SoftDeletePekerjaan(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	if err := s.repo.SoftDelete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pekerjaan berhasil dihapus"})
}

func (s *pekerjaanService) RestorePekerjaan(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	if err := s.repo.Restore(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pekerjaan berhasil dipulihkan"})
}
