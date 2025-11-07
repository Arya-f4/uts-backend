package repository

import (
	"context"
	"errors"
	"golang-train/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mediaRepository mengimplementasikan MediaRepository
type mediaRepository struct {
	collection *mongo.Collection
}

// NewFotoRepository membuat repositori baru untuk koleksi 'foto'
func NewFotoRepository(db *mongo.Database) MediaRepository {
	return &mediaRepository{
		collection: db.Collection("foto"),
	}
}

// NewSertifikatRepository membuat repositori baru untuk koleksi 'sertifikat'
func NewSertifikatRepository(db *mongo.Database) MediaRepository {
	return &mediaRepository{
		collection: db.Collection("sertifikat"),
	}
}

// UpsertByUserID membuat entri media baru atau memperbarui yang sudah ada untuk UserID
func (r *mediaRepository) UpsertByUserID(ctx context.Context, media *model.Media) (*model.Media, error) {
	media.CreatedAt = time.Now() // Selalu perbarui timestamp

	filter := bson.M{"user_id": media.UserID}
	update := bson.M{
		"$set": bson.M{
			"data":         media.Data,
			"content_type": media.ContentType,
			"size":         media.Size,
			"created_at":   media.CreatedAt,
		},
	}
	opts := options.Update().SetUpsert(true)

	result, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	if result.UpsertedID != nil {
		media.ID = result.UpsertedID.(primitive.ObjectID)
	}

	return media, nil
}

// GetByUserID mengambil media berdasarkan UserID
func (r *mediaRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) (*model.Media, error) {
	var media model.Media
	filter := bson.M{"user_id": userID}

	err := r.collection.FindOne(ctx, filter).Decode(&media)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("file tidak ditemukan")
		}
		return nil, err
	}
	return &media, nil
}
