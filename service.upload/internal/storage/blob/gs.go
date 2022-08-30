package blob

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

var (
	//
	ErrGSBucketNameEmpty = errors.New("google storage: bucket name not provided")
)

type gcloudStorage struct {
	bucketName string
	client     *storage.Client
}

//
func NewGoogleStorage(ctx context.Context, bucketName string) (*gcloudStorage, error) {
	if bucketName == "" {
		return nil, ErrGSBucketNameEmpty
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("new gcloud storage: %w", err)
	}

	_, err = client.Bucket(bucketName).Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("new gcloud storage: %w", err)
	}

	gs := &gcloudStorage{
		bucketName: bucketName,
		client:     client,
	}

	return gs, nil
}

//
func (gs *gcloudStorage) Close(ctx context.Context) error {
	return gs.client.Close()
}

//
func (gs *gcloudStorage) CreateObject(ctx context.Context, objName string, readers ...io.Reader) (url string, err error) {
	if objName == "" {
		return url, fmt.Errorf("create gs object: obj name empty")
	}

	if readers == nil {
		return url, fmt.Errorf("create gs object: readers param nil")
	}

	bucket := gs.client.Bucket(gs.bucketName)
	wr := bucket.Object(objName).NewWriter(ctx)
	defer func() {
		if tempErr := wr.Close(); tempErr != nil {
			err = tempErr
		}
	}()

	// io.Copy(wr, reader)

	attr, err := bucket.Attrs(ctx)
	if err != nil {
		return url, err
	}

	fmt.Printf("%+v", attr)
	return "", err
}
