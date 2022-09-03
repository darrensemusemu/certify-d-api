package blob_test

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/darrensemusemu/certify-d-api/service.upload/internal/storage/blob"
	"github.com/gofrs/uuid"
	"github.com/matryer/is"
)

func TestNewGoogleStorage(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	tests := []struct {
		name           string
		giveBucketName string
		wantErr        bool
	}{
		{
			name:           "check no bucket name",
			giveBucketName: "",
			wantErr:        true,
		},
		{
			name:           "check invalid bucket",
			giveBucketName: "bkt_does_not_exist",
			wantErr:        true,
		},
		{
			name:           "check valid bucket",
			giveBucketName: "certify-d_uploads",
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs, err := blob.NewGoogleStorage(ctx, tt.giveBucketName)
			is.True(tt.wantErr == (err != nil))
			if tt.wantErr {
				return
			}
			is.True(gs != nil)
		})
	}

}

func TestGSCreateDeleteObject(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./gs-service-account.json")
	f, err := os.Open("./data/test-file.txt")
	is.NoErr(err)
	// set object file path
	uuidV4, err := uuid.NewV4()
	is.NoErr(err)
	filenamePrefix := fmt.Sprintf("test/create-obj/%s/", uuidV4.String())

	tests := []struct {
		name        string
		giveObjName string
		giveReader  io.Reader
		wantErr     bool
	}{
		{
			name:        "check no obj name",
			giveObjName: "",
			giveReader:  nil,
			wantErr:     true,
		},
		{
			name:        "check no io.Reader interface",
			giveObjName: filenamePrefix + "test-file.txt",
			giveReader:  nil,
			wantErr:     true,
		},
		{
			name:        "check valid params",
			giveObjName: filenamePrefix + "test-file.txt",
			giveReader:  f,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs, err := blob.NewGoogleStorage(ctx, "certify-d_uploads")
			is.NoErr(err)
			err = gs.CreateObject(ctx, tt.giveObjName, tt.giveReader)
			is.True(tt.wantErr == (err != nil))
			if tt.wantErr {
				return
			}
			err = gs.DeleteObject(ctx, tt.giveObjName)
			is.NoErr(err)
		})
	}
}

func TestGSGetObjects(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./gs-service-account.json")
	// create gs client
	gs, err := blob.NewGoogleStorage(ctx, "certify-d_uploads")
	is.NoErr(err)
	// create test objects
	uuidV4, err := uuid.NewV4()
	is.NoErr(err)
	filenamePrefix := "test/get-objects/" + uuidV4.String()
	f, err := os.Open("./data/test-file.txt")
	is.NoErr(err)
	var savedFiles []string
	for i := 1; i <= 3; i++ {
		filename := fmt.Sprintf("%s/%d.txt", filenamePrefix, i)
		savedFiles = append(savedFiles, filename)
		err = gs.CreateObject(ctx, filename, f)
		is.NoErr(err)
		defer func() {
			err = gs.DeleteObject(ctx, filename)
			is.NoErr(err)
		}()
	}

	tests := []struct {
		name           string
		givePrefix     string
		giveExpires    time.Time
		wantErr        bool
		wantNumObjects int
	}{
		{
			name:           "check invalid object path",
			givePrefix:     "test/non_existent_obj.txt",
			giveExpires:    time.Now().Add(3 * time.Minute),
			wantErr:        false,
			wantNumObjects: 0,
		},
		{
			name:           "check valid object path",
			givePrefix:     filenamePrefix,
			giveExpires:    time.Now().Add(3 * time.Minute),
			wantErr:        false,
			wantNumObjects: len(savedFiles),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			names, err := gs.GetObjects(ctx, tt.givePrefix, tt.giveExpires)
			is.True(tt.wantErr == (err != nil))
			if tt.wantErr {
				return
			}
			is.Equal(len(names), tt.wantNumObjects)
		})
	}
}
