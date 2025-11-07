package database

import (
	"context"
	"golang-train/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDBConnection(cfg *config.Config) (*mongo.Database, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDBURI))
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke MongoDB: %v\n", err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Tidak dapat ping MongoDB: %v\n", err)
	}

	log.Println("Berhasil terhubung ke MongoDB.")
	db := client.Database(cfg.MongoDBName)

	// Return DB and a disconnect function
	disconnect := func() {
		log.Println("Memutus koneksi MongoDB...")
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error saat memutus koneksi MongoDB: %v", err)
		}
	}

	return db, disconnect
}
