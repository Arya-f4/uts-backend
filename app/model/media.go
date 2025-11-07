package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Media adalah struktur generik untuk menyimpan file Base64 di koleksi 'foto' atau 'sertifikat'
type Media struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"` // Referensi ke collection 'users'
	Data        string             `json:"-" bson:"data"`          // Data file dalam Base64 (dihilangkan dari JSON response)
	ContentType string             `json:"content_type" bson:"content_type"`
	Size        int64              `json:"size" bson:"size"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

// UploadMediaRequest adalah DTO untuk request body saat mengunggah
type UploadMediaRequest struct {
	Data        string `json:"data"`         // Diharapkan string Base64 murni
	ContentType string `json:"content_type"` // misal: "image/png" atau "application/pdf"
}

// UploadMediaResponse adalah DTO untuk respons setelah berhasil mengunggah
type UploadMediaResponse struct {
	UserID      string    `json:"user_id"`
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
}
