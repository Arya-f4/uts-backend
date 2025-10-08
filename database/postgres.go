
package database

import (
	"context"
	"golang-train/config"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresConnection(cfg *config.Config) *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v\n", err)
	}

	log.Println("Berhasil terhubung ke database.")
	return dbPool
}

