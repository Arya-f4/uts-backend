package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	NIM       string             `json:"nim" bson:"nim"`
	Nama      string             `json:"nama" bson:"nama"`
	Jurusan   string             `json:"jurusan" bson:"jurusan"`
	Angkatan  int                `json:"angkatan" bson:"angkatan"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateMahasiswaRequest struct {
	NIM      string `json:"nim"`
	Nama     string `json:"nama"`
	Jurusan  string `json:"jurusan"`
	Angkatan int    `json:"angkatan"`
	Email    string `json:"email"`
}

type UpdateMahasiswaRequest struct {
	Nama     string `json:"nama"`
	Jurusan  string `json:"jurusan"`
	Angkatan int    `json:"angkatan"`
	Email    string `json:"email"`
}
