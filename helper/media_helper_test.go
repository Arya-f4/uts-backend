package helper

import (
	"golang-train/app/service"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewMediaHelper(t *testing.T) {
	type args struct {
		s service.MediaService
	}
	tests := []struct {
		name string
		args args
		want *MediaHelper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMediaHelper(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMediaHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMediaHelper_UploadMedia(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *MediaHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.UploadMedia(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MediaHelper.UploadMedia() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMediaHelper_GetMedia(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		h       *MediaHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.GetMedia(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MediaHelper.GetMedia() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
