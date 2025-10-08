package helper

import (
	"golang-train/app/service"

	"github.com/gofiber/fiber/v2"
)

type PekerjaanHelper struct {
	pekerjaanService service.PekerjaanService
}

func NewPekerjaanHelper(s service.PekerjaanService) *PekerjaanHelper {
	return &PekerjaanHelper{pekerjaanService: s}
}

// Semua fungsi di bawah ini sekarang hanya meneruskan panggilan ke service layer,
// karena service layer kini bertanggung jawab menangani fiber.Ctx secara langsung.

func (h *PekerjaanHelper) CreatePekerjaan(c *fiber.Ctx) error {
	return h.pekerjaanService.CreatePekerjaan(c)
}

func (h *PekerjaanHelper) GetAllPekerjaan(c *fiber.Ctx) error {
	return h.pekerjaanService.GetAllPekerjaan(c)
}

func (h *PekerjaanHelper) GetAllPekerjaanDeleted(c *fiber.Ctx) error {
	return h.pekerjaanService.GetAllPekerjaanDeleted(c)
}

func (h *PekerjaanHelper) GetPekerjaanByID(c *fiber.Ctx) error {
	return h.pekerjaanService.GetPekerjaanByID(c)
}

func (h *PekerjaanHelper) UpdatePekerjaan(c *fiber.Ctx) error {
	return h.pekerjaanService.UpdatePekerjaan(c)
}

func (h *PekerjaanHelper) DeletePekerjaan(c *fiber.Ctx) error {
	return h.pekerjaanService.DeletePekerjaan(c)
}

func (h *PekerjaanHelper) SoftDeletePekerjaan(c *fiber.Ctx) error {
	return h.pekerjaanService.SoftDeletePekerjaan(c)
}

func (h *PekerjaanHelper) RestorePekerjaan(c *fiber.Ctx) error {
	return h.pekerjaanService.RestorePekerjaan(c)
}
