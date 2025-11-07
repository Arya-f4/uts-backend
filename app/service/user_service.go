package service

import (
	"context"
	"errors"
	"golang-train/app/repository"
	"slices"

	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(db *mongo.Database) UserService {
	return &userService{
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *userService) DeleteUser(ctx context.Context, requestingUserID string, requestingUserRoles []string, targetUserID string) error {
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
func (s *userService) RestoreUser(ctx context.Context, requestingUserID string, requestingUserRoles []string, targetUserID string) error {
	isAdmin := slices.Contains(requestingUserRoles, "admin")

	// Admin-only check is simpler for restore
	if !isAdmin {
		return errors.New("forbidden: Anda tidak memiliki izin untuk memulihkan pengguna ini")
	}

	err := s.userRepo.Restore(ctx, targetUserID)
	if err != nil {
		return err
	}

	return nil
}
