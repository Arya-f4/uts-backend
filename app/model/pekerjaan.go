package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pekerjaan struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	AlumniID            primitive.ObjectID `json:"alumni_id" bson:"alumni_id"`
	NamaPerusahaan      string             `json:"nama_perusahaan" bson:"nama_perusahaan"`
	PosisiJabatan       string             `json:"posisi_jabatan" bson:"posisi_jabatan"`
	BidangIndustri      string             `json:"bidang_industri" bson:"bidang_industri"`
	LokasiKerja         string             `json:"lokasi_kerja" bson:"lokasi_kerja"`
	GajiRange           *string            `json:"gaji_range,omitempty" bson:"gaji_range,omitempty"`
	TanggalMulaiKerja   time.Time          `json:"tanggal_mulai_kerja" bson:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time         `json:"tanggal_selesai_kerja,omitempty" bson:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string             `json:"status_pekerjaan" bson:"status_pekerjaan"`
	DeskripsiPekerjaan  *string            `json:"deskripsi_pekerjaan,omitempty" bson:"deskripsi_pekerjaan,omitempty"`
	CreatedAt           time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at" bson:"updated_at"`
	IsDeleted           *time.Time         `json:"is_deleted,omitempty" bson:"is_deleted,omitempty"`
}

type CreatePekerjaanRequest struct {
	AlumniID            string  `json:"alumni_id"` // Now a string to accept hex ID
	NamaPerusahaan      string  `json:"nama_perusahaan"`
	PosisiJabatan       string  `json:"posisi_jabatan"`
	BidangIndustri      string  `json:"bidang_industri"`
	LokasiKerja         string  `json:"lokasi_kerja"`
	GajiRange           *string `json:"gaji_range"`
	TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja"`
	StatusPekerjaan     string  `json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string `json:"deskripsi_pekerjaan"`
}

type UpdatePekerjaanRequest struct {
	NamaPerusahaan      string  `json:"nama_perusahaan"`
	PosisiJabatan       string  `json:"posisi_jabatan"`
	BidangIndustri      string  `json:"bidang_industri"`
	LokasiKerja         string  `json:"lokasi_kerja"`
	GajiRange           *string `json:"gaji_range"`
	TanggalMulaiKerja   string  `json:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *string `json:"tanggal_selesai_kerja"`
	StatusPekerjaan     string  `json:"status_pekerjaan"`
	DeskripsiPekerjaan  *string `json:"deskripsi_pekerjaan"`
}
