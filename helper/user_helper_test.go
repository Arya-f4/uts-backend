package helper

import (
	"golang-train/app/service"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewUserHelper(t *testing.T) {
	type args struct {
		s service.UserService
	}
	tests := []struct {
		name string
		args args
		want *UserHelper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserHelper(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserHelper_DeleteUser(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *UserHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.DeleteUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserHelper.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserHelper_RestoreUser(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *UserHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.RestoreUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserHelper.RestoreUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
