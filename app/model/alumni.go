package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alumni struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	NIM        string             `json:"nim" bson:"nim"`
	Nama       string             `json:"nama" bson:"nama"`
	Jurusan    string             `json:"jurusan" bson:"jurusan"`
	Angkatan   int                `json:"angkatan" bson:"angkatan"`
	TahunLulus int                `json:"tahun_lulus" bson:"tahun_lulus"`
	Email      string             `json:"email" bson:"email"`
	NoTelepon  *string            `json:"no_telepon,omitempty" bson:"no_telepon,omitempty"`
	Alamat     *string            `json:"alamat,omitempty" bson:"alamat,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateAlumniRequest struct {
	NIM        string  `json:"nim"`
	Nama       string  `json:"nama"`
	Jurusan    string  `json:"jurusan"`
	Angkatan   int     `json:"angkatan"`
	TahunLulus int     `json:"tahun_lulus"`
	Email      string  `json:"email"`
	NoTelepon  *string `json:"no_telepon"`
	Alamat     *string `json:"alamat"`
}

type UpdateAlumniRequest struct {
	Nama       string  `json:"nama"`
	Jurusan    string  `json:"jurusan"`
	Angkatan   int     `json:"angkatan"`
	TahunLulus int     `json:"tahun_lulus"`
	Email      string  `json:"email"`
	NoTelepon  *string `json:"no_telepon"`
	Alamat     *string `json:"alamat"`
}
