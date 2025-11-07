package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RunMigrations ensures all necessary collections have the correct indexes.
func RunMigrations(db *mongo.Database) {
	log.Println("Menjalankan migrasi (pembuatan-indeks) MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// === Indeks untuk Collection 'users' ===
	usersCollection := db.Collection("users")
	userIndexes := []mongo.IndexModel{
		{
			// Unique index for email
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_user_email_unique"),
		},
		{
			// Index for soft delete queries
			Keys:    bson.D{{Key: "is_deleted", Value: 1}},
			Options: options.Index().SetName("idx_user_is_deleted"),
		},
	}
	createIndexes(ctx, usersCollection, "users", userIndexes)

	// === Indeks untuk Collection 'alumni' ===
	alumniCollection := db.Collection("alumni")
	alumniIndexes := []mongo.IndexModel{
		{
			// Unique index for NIM
			Keys:    bson.D{{Key: "nim", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_alumni_nim_unique"),
		},
		{
			// Unique index for email
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_alumni_email_unique"),
		},
		{
			// Index for searching by name
			Keys:    bson.D{{Key: "nama", Value: 1}},
			Options: options.Index().SetName("idx_alumni_nama"),
		},
	}
	createIndexes(ctx, alumniCollection, "alumni", alumniIndexes)

	// === Indeks untuk Collection 'mahasiswa' ===
	mahasiswaCollection := db.Collection("mahasiswa")
	mahasiswaIndexes := []mongo.IndexModel{
		{
			// Unique index for NIM
			Keys:    bson.D{{Key: "nim", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_mahasiswa_nim_unique"),
		},
		{
			// Unique index for email
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_mahasiswa_email_unique"),
		},
		{
			// Index for searching by name
			Keys:    bson.D{{Key: "nama", Value: 1}},
			Options: options.Index().SetName("idx_mahasiswa_nama"),
		},
	}
	createIndexes(ctx, mahasiswaCollection, "mahasiswa", mahasiswaIndexes)

	// === Indeks untuk Collection 'pekerjaan' ===
	pekerjaanCollection := db.Collection("pekerjaan")
	pekerjaanIndexes := []mongo.IndexModel{
		{
			// Index for finding jobs by alumni
			Keys:    bson.D{{Key: "alumni_id", Value: 1}},
			Options: options.Index().SetName("idx_pekerjaan_alumni_id"),
		},
		{
			// Index for searching by company name
			Keys:    bson.D{{Key: "nama_perusahaan", Value: 1}},
			Options: options.Index().SetName("idx_pekerjaan_nama_perusahaan"),
		},
		{
			// Index for searching by position
			Keys:    bson.D{{Key: "posisi_jabatan", Value: 1}},
			Options: options.Index().SetName("idx_pekerjaan_posisi_jabatan"),
		},
		{
			// Index for soft delete queries
			Keys:    bson.D{{Key: "is_deleted", Value: 1}},
			Options: options.Index().SetName("idx_pekerjaan_is_deleted"),
		},
	}
	createIndexes(ctx, pekerjaanCollection, "pekerjaan", pekerjaanIndexes)

	// === Indeks untuk Collection 'foto' (BARU) ===
	fotoCollection := db.Collection("foto")
	fotoIndexes := []mongo.IndexModel{
		{
			// Indeks unik untuk user_id, memastikan 1 foto per user
			Keys:    bson.D{{Key: "user_id", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_foto_user_id_unique"),
		},
	}
	createIndexes(ctx, fotoCollection, "foto", fotoIndexes)

	// === Indeks untuk Collection 'sertifikat' (BARU) ===
	sertifikatCollection := db.Collection("sertifikat")
	sertifikatIndexes := []mongo.IndexModel{
		{
			// Indeks unik untuk user_id, memastikan 1 sertifikat per user
			Keys:    bson.D{{Key: "user_id", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("idx_sertifikat_user_id_unique"),
		},
	}
	createIndexes(ctx, sertifikatCollection, "sertifikat", sertifikatIndexes)

	log.Println("Migrasi MongoDB selesai.")
}

// createIndexes is a helper function to create indexes for a collection
func createIndexes(ctx context.Context, collection *mongo.Collection, collectionName string, indexes []mongo.IndexModel) {
	if len(indexes) == 0 {
		log.Printf("Tidak ada indeks yang didefinisikan untuk collection: %s", collectionName)
		return
	}

	indexNames, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("Gagal membuat indeks untuk collection %s: %v", collectionName, err)
	} else {
		log.Printf("Indeks berhasil dibuat untuk collection %s: %v", collectionName, indexNames)
	}
}
