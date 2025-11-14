package helper

import (
	"golang-train/app/service"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewAuthHelper(t *testing.T) {
	type args struct {
		as service.AuthService
	}
	tests := []struct {
		name string
		args args
		want *AuthHelper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthHelper(tt.args.as); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthHelper_Register(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AuthHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Register(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHelper.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthHelper_Login(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AuthHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AuthHelper.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
