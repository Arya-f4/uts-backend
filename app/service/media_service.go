package service

import (
	"encoding/base64"
	"errors"
	"golang-train/app/model"
	"golang-train/app/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mediaService struct {
	fotoRepo       repository.MediaRepository
	sertifikatRepo repository.MediaRepository
}

// NewMediaService menginisialisasi servis dengan kedua repositori
func NewMediaService(fotoRepo, sertifikatRepo repository.MediaRepository) MediaService {
	return &mediaService{
		fotoRepo:       fotoRepo,
		sertifikatRepo: sertifikatRepo,
	}
}

// getRepoForMediaType memilih repositori yang tepat berdasarkan parameter rute
func (s *mediaService) getRepoForMediaType(mediaType string) (repository.MediaRepository, error) {
	if mediaType == "foto" {
		return s.fotoRepo, nil
	}
	if mediaType == "sertifikat" {
		return s.sertifikatRepo, nil
	}
	return nil, errors.New("tipe media tidak valid")
}

// UploadMedia menangani logika upload file
func (s *mediaService) UploadMedia(c *fiber.Ctx) error {
	targetUserIDStr := c.Params("id")
	mediaType := c.Params("mediaType")

	// Keamanan: Pastikan pengguna hanya mengunggah ke profil mereka sendiri
	// (Admin dapat diperbolehkan di masa depan, tetapi ini adalah default yang aman)
	requestingUserIDStr := c.Locals("user_id").(string)
	if requestingUserIDStr != targetUserIDStr {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Anda hanya dapat mengunggah file ke profil Anda sendiri"})
	}

	// Pilih repositori yang benar
	repo, err := s.getRepoForMediaType(mediaType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var req model.UploadMediaRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body JSON tidak valid"})
	}

	if req.Data == "" || req.ContentType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Field 'data' (base64) dan 'content_type' diperlukan"})
	}

	// Konversi UserID string ke ObjectID
	targetUserID, err := primitive.ObjectIDFromHex(targetUserIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID tidak valid"})
	}

	media := &model.Media{
		UserID:      targetUserID,
		Data:        req.Data, // Simpan string Base64 murni
		ContentType: req.ContentType,
		Size:        int64(len(req.Data)), // Ukuran string Base64
	}

	// Buat atau perbarui file
	savedMedia, err := repo.UpsertByUserID(c.Context(), media)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Kembalikan respons yang bersih tanpa data Base64
	return c.Status(fiber.StatusCreated).JSON(model.UploadMediaResponse{
		UserID:      savedMedia.UserID.Hex(),
		ContentType: savedMedia.ContentType,
		Size:        savedMedia.Size,
		CreatedAt:   savedMedia.CreatedAt,
	})
}

// GetMedia menangani logika download file
func (s *mediaService) GetMedia(c *fiber.Ctx) error {
	targetUserIDStr := c.Params("id")
	mediaType := c.Params("mediaType")

	// Pilih repositori
	repo, err := s.getRepoForMediaType(mediaType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	targetUserID, err := primitive.ObjectIDFromHex(targetUserIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID tidak valid"})
	}

	// Ambil data media dari DB
	media, err := repo.GetByUserID(c.Context(), targetUserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Decode string Base64 kembali menjadi data biner
	decodedData, err := base64.StdEncoding.DecodeString(media.Data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal men-decode file"})
	}

	// Set header Content-Type yang benar dan kirim data binernya
	c.Set("Content-Type", media.ContentType)
	return c.Send(decodedData)
}
