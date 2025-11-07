package helper

import (
	"golang-train/app/model"
	"golang-train/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlumniHelper struct {
	alumniService service.AlumniService
}

func NewAlumniHelper(s service.AlumniService) *AlumniHelper {
	return &AlumniHelper{alumniService: s}
}

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

func (h *AlumniHelper) GetAlumniByID(c *fiber.Ctx) error {
	id := c.Params("id") // ID is now a string
	alumni, err := h.alumniService.GetAlumniByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(alumni)
}

func (h *AlumniHelper) UpdateAlumni(c *fiber.Ctx) error {
	id := c.Params("id") // ID is now a string

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

func (h *AlumniHelper) DeleteAlumni(c *fiber.Ctx) error {
	id := c.Params("id") // ID is now a string
	if err := h.alumniService.DeleteAlumni(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
