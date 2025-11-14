package helper

import (
	"golang-train/app/service"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewMahasiswaHelper(t *testing.T) {
	type args struct {
		ms service.MahasiswaService
	}
	tests := []struct {
		name string
		args args
		want *MahasiswaHelper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMahasiswaHelper(tt.args.ms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMahasiswaHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMahasiswaHelper_CreateMahasiswa(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *MahasiswaHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.CreateMahasiswa(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MahasiswaHelper.CreateMahasiswa() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMahasiswaHelper_GetAllMahasiswa(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *MahasiswaHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAllMahasiswa(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MahasiswaHelper.GetAllMahasiswa() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMahasiswaHelper_GetMahasiswaByID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *MahasiswaHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetMahasiswaByID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MahasiswaHelper.GetMahasiswaByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMahasiswaHelper_UpdateMahasiswa(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *MahasiswaHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.UpdateMahasiswa(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MahasiswaHelper.UpdateMahasiswa() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMahasiswaHelper_DeleteMahasiswa(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *MahasiswaHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.DeleteMahasiswa(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MahasiswaHelper.DeleteMahasiswa() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
