package repository

import (
	"context"
	"errors"
	"golang-train/app/model"
	"math"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type alumniRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewAlumniRepository(db *mongo.Database) AlumniRepository {
	return &alumniRepository{
		db:         db,
		collection: db.Collection("alumni"),
	}
}

func (r *alumniRepository) Create(ctx context.Context, alumni *model.Alumni) (*model.Alumni, error) {
	alumni.ID = primitive.NewObjectID()
	alumni.CreatedAt = time.Now()
	alumni.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, alumni)
	if err != nil {
		return nil, err
	}
	return alumni, nil
}

func (r *alumniRepository) FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error) {
	filter := bson.M{}
	if params.Search != "" {
		regex := bson.M{"$regex": params.Search, "$options": "i"}
		filter["$or"] = []bson.M{
			{"nama": regex},
			{"nim": regex},
			{"jurusan": regex},
			{"email": regex},
		}
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	opts := options.Find()
	opts.SetLimit(int64(params.Limit))
	opts.SetSkip(int64((params.Page - 1) * params.Limit))

	sortColumn := "created_at"
	sortOrder := -1 // desc
	if params.Sort != "" {
		parts := strings.Split(params.Sort, ":")
		if len(parts) == 2 {
			validCols := map[string]string{"nama": "nama", "nim": "nim", "angkatan": "angkatan", "tahun_lulus": "tahun_lulus", "created_at": "created_at"}
			if col, ok := validCols[parts[0]]; ok {
				sortColumn = col
			}
			if strings.ToUpper(parts[1]) == "ASC" {
				sortOrder = 1
			}
		}
	}
	opts.SetSort(bson.D{{Key: sortColumn, Value: sortOrder}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var alumniList []model.Alumni
	if err = cursor.All(ctx, &alumniList); err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	return &model.PaginationResult[model.Alumni]{
		Data:     alumniList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}, nil
}

func (r *alumniRepository) FindByID(ctx context.Context, id string) (*model.Alumni, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID tidak valid")
	}

	var a model.Alumni
	filter := bson.M{"_id": objID}
	err = r.collection.FindOne(ctx, filter).Decode(&a)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("alumni tidak ditemukan")
		}
		return nil, err
	}
	return &a, nil
}

func (r *alumniRepository) Update(ctx context.Context, id string, alumni *model.Alumni) (*model.Alumni, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID tidak valid")
	}

	update := bson.M{
		"$set": bson.M{
			"nama":        alumni.Nama,
			"jurusan":     alumni.Jurusan,
			"angkatan":    alumni.Angkatan,
			"tahun_lulus": alumni.TahunLulus,
			"email":       alumni.Email,
			"no_telepon":  alumni.NoTelepon,
			"alamat":      alumni.Alamat,
			"updated_at":  time.Now(),
		},
	}

	filter := bson.M{"_id": objID}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("alumni tidak ditemukan")
	}

	alumni.ID = objID // Ensure the ID is set on the returned object
	return alumni, nil
}

func (r *alumniRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID tidak valid")
	}

	filter := bson.M{"_id": objID}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("tidak ada baris yang ditemukan untuk dihapus")
	}
	return nil
}

// FindAllDeleted is not implemented for Alumni as it's a hard delete.
func (r *alumniRepository) FindAllDeleted(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Alumni], error) {
	return &model.PaginationResult[model.Alumni]{
		Data: []model.Alumni{},
	}, nil
}
