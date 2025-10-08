package service

import (
	"context"
	"errors"
	"golang-train/app/repository"
	"slices"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userService struct {
	db       *pgxpool.Pool // Diperlukan untuk repository
	userRepo repository.UserRepository
}

func NewUserService(db *pgxpool.Pool) UserService {
	return &userService{
		db:       db,
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *userService) DeleteUser(ctx context.Context, requestingUserID int, requestingUserRoles []string, targetUserID int) error {
	isAdmin := slices.Contains(requestingUserRoles, "admin")

	if !isAdmin && requestingUserID != targetUserID {
		return errors.New("forbidden: Anda tidak memiliki izin untuk menghapus pengguna ini")
	}

	err := s.userRepo.Delete(ctx, targetUserID)
	if err != nil {
		return err
	}

	return nil
}
func (s *userService) RestoreUser(ctx context.Context, requestingUserID int, requestingUserRoles []string, targetUserID int) error {
	isAdmin := slices.Contains(requestingUserRoles, "admin")

	if !isAdmin && requestingUserID != targetUserID {
		return errors.New("forbidden: Anda tidak memiliki izin untuk menghapus pengguna ini")
	}

	err := s.userRepo.Restore(ctx, targetUserID)
	if err != nil {
		return err
	}

	return nil
}
