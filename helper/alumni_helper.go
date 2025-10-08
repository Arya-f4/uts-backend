
// File Path: helper/alumni_helper.go
package helper

import (
	"golang-train/app/model"
	"golang-train/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AlumniHelper menangani logika request-response untuk alumni.
type AlumniHelper struct {
	alumniService service.AlumniService
}

// NewAlumniHelper membuat instance baru dari AlumniHelper.
func NewAlumniHelper(s service.AlumniService) *AlumniHelper {
	return &AlumniHelper{alumniService: s}
}

// CreateAlumni menangani permintaan pembuatan alumni baru.
func (h *AlumniHelper) CreateAlumni(c *fiber.Ctx) error {
	var req model.CreateAlumniRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	alumni, err := h.alumniService.CreateAlumni(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(alumni)
}

// GetAllAlumni menangani permintaan untuk mendapatkan semua alumni.
func (h *AlumniHelper) GetAllAlumni(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sort := c.Query("sort", "created_at:desc")
	search := c.Query("search", "")

	params := model.PaginationParams{
		Page:   page,
		Limit:  limit,
		Sort:   sort,
		Search: search,
	}

	result, err := h.alumniService.GetAllAlumni(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

// GetAlumniByID menangani permintaan untuk mendapatkan alumni berdasarkan ID.
func (h *AlumniHelper) GetAlumniByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	alumni, err := h.alumniService.GetAlumniByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumni)
}

// UpdateAlumni menangani permintaan pembaruan data alumni.
func (h *AlumniHelper) UpdateAlumni(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var req model.UpdateAlumniRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	alumni, err := h.alumniService.UpdateAlumni(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumni)
}

// DeleteAlumni menangani permintaan penghapusan data alumni.
func (h *AlumniHelper) DeleteAlumni(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	if err := h.alumniService.DeleteAlumni(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}


