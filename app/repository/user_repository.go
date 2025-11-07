package repository

import (
	"context"
	"errors"
	"golang-train/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		db:         db,
		collection: db.Collection("users"),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User, roleName string) (*model.User, error) {
	// In MongoDB, we don't have separate role tables by default
	// We just add the role to the user document
	user.ID = primitive.NewObjectID()
	user.Roles = []string{roleName}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsDeleted = nil

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		// Handle duplicate key error
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New("email sudah terdaftar")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	// Filter for non-deleted user
	filter := bson.M{"email": email, "is_deleted": bson.M{"$exists": false}}
	err := r.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("pengguna tidak ditemukan")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID tidak valid")
	}

	user := &model.User{}
	// Filter for non-deleted user
	filter := bson.M{"_id": objID, "is_deleted": bson.M{"$exists": false}}
	err = r.collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("pengguna tidak ditemukan")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID tidak valid")
	}

	filter := bson.M{"_id": objID, "is_deleted": bson.M{"$exists": false}}
	update := bson.M{
		"$set": bson.M{
			"is_deleted": time.Now(),
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("pengguna tidak ditemukan atau sudah dihapus")
	}
	return nil
}

func (r *userRepository) Restore(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID tidak valid")
	}

	filter := bson.M{"_id": objID, "is_deleted": bson.M{"$exists": true}}
	update := bson.M{
		"$unset": bson.M{"is_deleted": ""},
		"$set":   bson.M{"updated_at": time.Now()},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("pengguna tidak ditemukan atau sudah direstore")
	}
	return nil
}
