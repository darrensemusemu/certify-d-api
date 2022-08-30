package blob_test

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/darrensemusemu/certify-d-api/service.upload/internal/storage/blob"
	"github.com/matryer/is"
)

func TestGoogleStorage(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/darren/Documents/Github/certify-d-api/service.upload/config/gs-service-account.json")

	gs, err := blob.NewGoogleStorage(ctx, "certify-d_uploads")
	is.NoErr(err)
	defer func() {
		err = gs.Close(ctx)
		is.NoErr(err)
	}()

	// gs.
	// _ = gs

	// tests := []struct {
	// 	name        string
	// 	giveObjName string
	// 	giveReader  []io.Reader
	// 	wantErr     bool
	// }{
	// 	{
	// 		name:        "",
	// 		giveObjName: "",
	// 		giveReader:  nil,
	// 		wantErr:     false,
	// 	},
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {

	// 	})
	// }

}

type testFileReader struct{}

var _ io.Reader = (*testFileReader)(nil)

func (t *testFileReader) Read(p []byte) (n int, err error) {
	b := []byte("file written")
	return len(b), nil
}

func TestGSCreateObject(t *testing.T) {

	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/darren/Documents/Github/certify-d-api/service.upload/config/gs-service-account.json")

	tests := []struct {
		name        string
		giveObjName string
		giveReader  []io.Reader
		wantErr     bool
	}{
		{
			name:        "",
			giveObjName: "",
			giveReader:  nil,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}

func TestXxx(t *testing.T) {
	at := make([]byte, 6, 6)
	aa(at)
	fmt.Printf("%v", string(at))
}

func aa(ar []byte) {
	// _ = ar
	copy(ar, []byte("s"))
	// z = []byte("s")
	// z := []byte("s")[0]
	// ar = append(ar, z[0])
	// _ = z

}
