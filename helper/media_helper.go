package helper

import (
	"golang-train/app/service"

	"github.com/gofiber/fiber/v2"
)

type MediaHelper struct {
	mediaService service.MediaService
}

func NewMediaHelper(s service.MediaService) *MediaHelper {
	return &MediaHelper{mediaService: s}
}

// UploadMedia meneruskan request upload ke servis
func (h *MediaHelper) UploadMedia(c *fiber.Ctx) error {
	return h.mediaService.UploadMedia(c)
}

// GetMedia meneruskan request get ke servis
func (h *MediaHelper) GetMedia(c *fiber.Ctx) error {
	return h.mediaService.GetMedia(c)
}
