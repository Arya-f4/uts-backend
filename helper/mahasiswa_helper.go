
package helper

import (
	"golang-train/app/model"
	"golang-train/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MahasiswaHelper struct {
	mahasiswaService service.MahasiswaService
}

func NewMahasiswaHelper(ms service.MahasiswaService) *MahasiswaHelper {
	return &MahasiswaHelper{mahasiswaService: ms}
}

func (h *MahasiswaHelper) CreateMahasiswa(c *fiber.Ctx) error {
	var req model.CreateMahasiswaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	mahasiswa, err := h.mahasiswaService.CreateMahasiswa(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(mahasiswa)
}

func (h *MahasiswaHelper) GetAllMahasiswa(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sort := c.Query("sort", "created_at:desc")
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 { // Batasi limit untuk mencegah query yang berlebihan
		limit = 100
	}

	params := model.PaginationParams{
		Page:   page,
		Limit:  limit,
		Sort:   sort,
		Search: search,
	}

	result, err := h.mahasiswaService.GetAllMahasiswa(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *MahasiswaHelper) GetMahasiswaByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	mahasiswa, err := h.mahasiswaService.GetMahasiswaByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(mahasiswa)
}

func (h *MahasiswaHelper) UpdateMahasiswa(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var req model.UpdateMahasiswaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	mahasiswa, err := h.mahasiswaService.UpdateMahasiswa(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(mahasiswa)
}

func (h *MahasiswaHelper) DeleteMahasiswa(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	if err := h.mahasiswaService.DeleteMahasiswa(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}


