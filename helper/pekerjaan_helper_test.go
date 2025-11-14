package helper

import (
	"golang-train/app/service"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewPekerjaanHelper(t *testing.T) {
	type args struct {
		s service.PekerjaanService
	}
	tests := []struct {
		name string
		args args
		want *PekerjaanHelper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPekerjaanHelper(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPekerjaanHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPekerjaanHelper_CreatePekerjaan(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.CreatePekerjaan(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.CreatePekerjaan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPekerjaanHelper_GetAllPekerjaan(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAllPekerjaan(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.GetAllPekerjaan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPekerjaanHelper_GetAllPekerjaanDeleted(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetAllPekerjaanDeleted(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.GetAllPekerjaanDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPekerjaanHelper_GetPekerjaanByID(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetPekerjaanByID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.GetPekerjaanByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPekerjaanHelper_UpdatePekerjaan(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.UpdatePekerjaan(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.UpdatePekerjaan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPekerjaanHelper_DeletePekerjaan(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.DeletePekerjaan(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.DeletePekerjaan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPekerjaanHelper_SoftDeletePekerjaan(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.SoftDeletePekerjaan(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.SoftDeletePekerjaan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPekerjaanHelper_RestorePekerjaan(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *PekerjaanHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.RestorePekerjaan(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PekerjaanHelper.RestorePekerjaan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
