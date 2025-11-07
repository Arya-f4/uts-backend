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

type pekerjaanRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewPekerjaanRepository(db *mongo.Database) PekerjaanRepository {
	return &pekerjaanRepository{
		db:         db,
		collection: db.Collection("pekerjaan"),
	}
}

func (r *pekerjaanRepository) Create(ctx context.Context, p *model.Pekerjaan) (*model.Pekerjaan, error) {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.IsDeleted = nil // Ensure is_deleted is nil on create

	_, err := r.collection.InsertOne(ctx, p)
	return p, err
}

// Helper for FindAll and FindAllDeleted
func (r *pekerjaanRepository) findAllWithFilter(ctx context.Context, params model.PaginationParams, filter bson.M) (*model.PaginationResult[model.Pekerjaan], error) {
	if params.Search != "" {
		// MongoDB doesn't support $regex in $lookup easily.
		// We'll search on pekerjaan fields first.
		// A more complex search involving alumni name would require an aggregation pipeline.
		regex := bson.M{"$regex": params.Search, "$options": "i"}
		searchFilter := []bson.M{
			{"nama_perusahaan": regex},
			{"posisi_jabatan": regex},
			{"bidang_industri": regex},
		}
		
		// If filter already has $or or other keys, merge carefully
		if existingOr, ok := filter["$or"].([]bson.M); ok {
			filter["$and"] = []bson.M{
				{"$or": existingOr},
				{"$or": searchFilter},
			}
			delete(filter, "$or")
		} else if len(filter) > 0 {
			filter["$and"] = []bson.M{
				filter,
				{"$or": searchFilter},
			}
		} else {
			filter["$or"] = searchFilter
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
			validCols := map[string]string{"nama_perusahaan": "nama_perusahaan", "posisi_jabatan": "posisi_jabatan", "tanggal_mulai_kerja": "tanggal_mulai_kerja", "created_at": "created_at"}
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

	var pekerjaanList []model.Pekerjaan
	if err = cursor.All(ctx, &pekerjaanList); err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(params.Limit)))
	if lastPage < 1 && total > 0 {
		lastPage = 1
	}

	return &model.PaginationResult[model.Pekerjaan]{
		Data:     pekerjaanList,
		Total:    total,
		Page:     params.Page,
		Limit:    params.Limit,
		LastPage: lastPage,
	}, nil
}


func (r *pekerjaanRepository) FindAll(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error) {
	// Filter for non-deleted records
	filter := bson.M{"is_deleted": bson.M{"$exists": false}}
	return r.findAllWithFilter(ctx, params, filter)
}

func (r *pekerjaanRepository) FindAllDeleted(ctx context.Context, params model.PaginationParams) (*model.PaginationResult[model.Pekerjaan], error) {
	// Filter for soft-deleted records
	filter := bson.M{"is_deleted": bson.M{"$exists": true}}
	return r.findAllWithFilter(ctx, params, filter)
}

func (r *pekerjaanRepository) FindByID(ctx context.Context, id string) (*model.Pekerjaan, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID tidak valid")
	}

	var p model.Pekerjaan
	filter := bson.M{"_id": objID, "is_deleted": bson.M{"$exists": false}}
	err = r.collection.FindOne(ctx, filter).Decode(&p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("pekerjaan tidak ditemukan")
		}
		return nil, err
	}
	return &p, nil
}

func (r *pekerjaanRepository) Update(ctx context.Context, id string, p *model.Pekerjaan) (*model.Pekerjaan, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID tidak valid")
	}

	update := bson.M{
		"$set": bson.M{
			"nama_perusahaan":      p.NamaPerusahaan,
			"posisi_jabatan":       p.PosisiJabatan,
			"bidang_industri":      p.BidangIndustri,
			"lokasi_kerja":         p.LokasiKerja,
			"gaji_range":           p.GajiRange,
			"tanggal_mulai_kerja":  p.TanggalMulaiKerja,
			"tanggal_selesai_kerja": p.TanggalSelesaiKerja,
			"status_pekerjaan":     p.StatusPekerjaan,
			"deskripsi_pekerjaan":  p.DeskripsiPekerjaan,
			"updated_at":           time.Now(),
		},
	}

	filter := bson.M{"_id": objID}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("pekerjaan tidak ditemukan")
	}

	p.ID = objID
	return p, nil
}

func (r *pekerjaanRepository) Delete(ctx context.Context, id string) error {
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

func (r *pekerjaanRepository) SoftDelete(ctx context.Context, id string) error {
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
		return errors.New("pekerjaan tidak ditemukan atau sudah dihapus")
	}
	return nil
}

func (r *pekerjaanRepository) Restore(ctx context.Context, id string) error {
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
		return errors.New("pekerjaan tidak ditemukan di data yang dihapus")
	}
	return nil
}
