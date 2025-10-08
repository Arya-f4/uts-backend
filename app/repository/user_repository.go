package repository

import (
	"context"
	"errors"
	"golang-train/app/model"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User, roleName string) (*model.User, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	userSQL := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err = tx.QueryRow(ctx, userSQL, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var roleID int
	roleSQL := `SELECT id FROM roles WHERE name = $1`
	err = tx.QueryRow(ctx, roleSQL, roleName).Scan(&roleID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("peran tidak ditemukan")
		}
		return nil, err
	}

	userRoleSQL := `INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)`
	_, err = tx.Exec(ctx, userRoleSQL, user.ID, roleID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	user.Roles = []string{roleName}
	return user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	query := `
		SELECT u.id, u.email, u.password_hash, u.created_at, u.updated_at, array_agg(r.name) as roles
		FROM users u
		LEFT JOIN user_roles ur ON u.id = ur.user_id
		LEFT JOIN roles r ON ur.role_id = r.id
		WHERE u.email = $1 AND u.is_deleted IS NULL
		GROUP BY u.id`
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.Roles)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("pengguna tidak ditemukan")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user := &model.User{}
	query := `
		SELECT u.id, u.email, u.password_hash, u.created_at, u.updated_at, array_agg(r.name) as roles
		FROM users u
		LEFT JOIN user_roles ur ON u.id = ur.user_id
		LEFT JOIN roles r ON ur.role_id = r.id
		WHERE u.id = $1 AND u.is_deleted IS NULL
		GROUP BY u.id`
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.Roles)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("pengguna tidak ditemukan")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	query := `UPDATE users SET is_deleted = NOW() WHERE id = $1 AND is_deleted IS NULL`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("pengguna tidak ditemukan atau sudah dihapus")
	}
	return nil
}
func (r *userRepository) Restore(ctx context.Context, id int) error {
	query := `UPDATE users SET is_deleted = NULL WHERE id = $1 `
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("pengguna tidak ditemukan atau sudah direstore")
	}
	return nil
}
