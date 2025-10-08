
package helper

import (
	"golang-train/app/model"
	"golang-train/app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHelper struct {
	authService service.AuthService
}

func NewAuthHelper(as service.AuthService) *AuthHelper {
	return &AuthHelper{authService: as}
}

func (h *AuthHelper) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	user, err := h.authService.Register(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHelper) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tidak dapat mem-parsing JSON"})
	}

	token, err := h.authService.Login(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(model.LoginResponse{Token: token})
}


