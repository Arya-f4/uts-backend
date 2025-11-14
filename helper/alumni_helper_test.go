package helper

import (
	"golang-train/app/service"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewAlumniHelper(t *testing.T) {
	type args struct {
		s service.AlumniService
	}
	tests := []struct {
		name string
		args args
		want *AlumniHelper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlumniHelper(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlumniHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlumniHelper_CreateAlumni(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AlumniHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.CreateAlumni(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AlumniHelper.CreateAlumni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlumniHelper_GetAllAlumni(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AlumniHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAllAlumni(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AlumniHelper.GetAllAlumni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlumniHelper_GetAlumniByID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AlumniHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAlumniByID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AlumniHelper.GetAlumniByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlumniHelper_UpdateAlumni(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AlumniHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.UpdateAlumni(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AlumniHelper.UpdateAlumni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlumniHelper_DeleteAlumni(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *AlumniHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.DeleteAlumni(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AlumniHelper.DeleteAlumni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
