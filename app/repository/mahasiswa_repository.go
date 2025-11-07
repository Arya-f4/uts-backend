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

type mahasiswaRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMahasiswaRepository(db *mongo.Database) MahasiswaRepository {
	return &mahasiswaRepository{
		db:         db,
		collection: db.Collection("mahasiswa"),
	}
}

func (r *mahasiswaRepository) Create(ctx context.Context, m *model.Mahasiswa) (*model.Mahasiswa, error) {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *mahasiswaRepository) FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Mahasiswa], error) {
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
			validCols := map[string]string{"nama": "nama", "nim": "nim", "angkatan": "angkatan", "created_at": "created_at"}
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

	var mahasiswaList []model.Mahasiswa
	if err = cursor.All(ctx, &mahasiswaList); err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	return &model.PaginationResult[model.Mahasiswa]{
		Data:     mahasiswaList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}, nil
}

func (r *mahasiswaRepository) FindByID(ctx context.Context, id string) (*model.Mahasiswa, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID tidak valid")
	}

	var m model.Mahasiswa
	filter := bson.M{"_id": objID}
	err = r.collection.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("mahasiswa tidak ditemukan")
		}
		return nil, err
	}
	return &m, nil
}

func (r *mahasiswaRepository) Update(ctx context.Context, id string, m *model.Mahasiswa) (*model.Mahasiswa, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID tidak valid")
	}

	update := bson.M{
		"$set": bson.M{
			"nama":       m.Nama,
			"jurusan":    m.Jurusan,
			"angkatan":   m.Angkatan,
			"email":      m.Email,
			"updated_at": time.Now(),
		},
	}

	filter := bson.M{"_id": objID}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("mahasiswa tidak ditemukan")
	}

	m.ID = objID
	return m, nil
}

func (r *mahasiswaRepository) Delete(ctx context.Context, id string) error {
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
