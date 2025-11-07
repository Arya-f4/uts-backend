package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"-" bson:"password_hash"`
	Roles        []string           `json:"roles" bson:"roles"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	IsDeleted    *time.Time         `json:"is_deleted,omitempty" bson:"is_deleted,omitempty"`
}

type Role struct {
	ID   int    `json:"id"` // This model seems unused by MongoDB, but kept for reference
	Name string `json:"name"`
}
