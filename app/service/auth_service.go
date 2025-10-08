

package service

import (
	"context"
	"errors"
	"golang-train/app/model"
	"golang-train/app/repository"
	"golang-train/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type authService struct {
	userRepo           repository.UserRepository
	jwtSecret          string
	jwtExpirationHours time.Duration
}

func NewAuthService(ur repository.UserRepository, secret string, exp time.Duration) AuthService {
	return &authService{
		userRepo:           ur,
		jwtSecret:          secret,
		jwtExpirationHours: exp,
	}
}

func (u *authService) Register(ctx context.Context, req *model.RegisterRequest) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	return u.userRepo.CreateUser(ctx, user, "user")
}

func (u *authService) Login(ctx context.Context, req *model.LoginRequest) (string, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("kredensial tidak valid")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return "", errors.New("kredensial tidak valid")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"roles":   user.Roles,
		"exp":     time.Now().Add(u.jwtExpirationHours).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


