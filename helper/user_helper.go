package helper

import (
	"golang-train/app/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserHelper struct {
	userService service.UserService
}

func NewUserHelper(s service.UserService) *UserHelper {
	return &UserHelper{userService: s}
}

func (h *UserHelper) DeleteUser(c *fiber.Ctx) error {
	targetUserID := c.Params("id")
	if targetUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID pengguna tidak valid"})
	}

	requestingUserID := c.Locals("user_id").(string) // Now a string
	requestingUserRoles := c.Locals("roles").([]string)

	err := h.userService.DeleteUser(c.Context(), requestingUserID, requestingUserRoles, targetUserID)
	if err != nil {
		if strings.Contains(err.Error(), "forbidden") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
func (h *UserHelper) RestoreUser(c *fiber.Ctx) error {
	targetUserID := c.Params("id")
	if targetUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID pengguna tidak valid"})
	}

	requestingUserID := c.Locals("user_id").(string) // Now a string
	requestingUserRoles := c.Locals("roles").([]string)

	err := h.userService.RestoreUser(c.Context(), requestingUserID, requestingUserRoles, targetUserID)
	if err != nil {
		if strings.Contains(err.Error(), "forbidden") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
