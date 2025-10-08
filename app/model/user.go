package model

import "time"

type User struct {
	ID           int        `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	Roles        []string   `json:"roles"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	IsDeleted    *time.Time `json:"is_deleted,omitempty"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
