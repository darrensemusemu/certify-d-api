package blob

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

var (
	// google storage client application creds
	gsJsonKey []byte
	// get storage client application creds once
	gsJsonKeyOnce sync.Once
	// error from get storage client application creds once
	gsJsonKeyErr error
)

// Checks gcloudStorage implements Service
var _ Service = (*gcloudStorage)(nil)

type gcloudStorage struct {
	bktName string
	client  *storage.Client
}

// Creates a new google storage handler
func NewGoogleStorage(ctx context.Context, bktName string) (*gcloudStorage, error) {
	if bktName == "" {
		return nil, fmt.Errorf("new gcloud storage: bucket name not provided")
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("new gcloud storage: %w", err)
	}
	_, err = client.Bucket(bktName).Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("new gcloud storage: %w", err)
	}
	gs := &gcloudStorage{
		bktName: bktName,
		client:  client,
	}
	return gs, nil
}

// Close storage client
func (gs *gcloudStorage) Close(ctx context.Context) error {
	return gs.client.Close()
}

// Create an object
func (gs *gcloudStorage) CreateObject(ctx context.Context, objectName string, reader io.Reader) (err error) {
	if objectName == "" {
		return fmt.Errorf("create gs object: obj name empty")
	}
	if reader == nil {
		return fmt.Errorf("create gs object: readers param nil")
	}
	bucket := gs.client.Bucket(gs.bktName)
	wr := bucket.Object(objectName).NewWriter(ctx)
	defer func() {
		if tempErr := wr.Close(); tempErr != nil {
			err = tempErr
		}
	}()
	io.Copy(wr, reader)
	return err
}

// Delete an object
func (gs *gcloudStorage) DeleteObject(ctx context.Context, objectName string) (err error) {
	if objectName == "" {
		return fmt.Errorf("delete gs object: obj name empty")
	}
	obj := gs.client.Bucket(gs.bktName).Object(objectName)
	if err := obj.Delete(ctx); err != nil {
		return fmt.Errorf("delete gs object: %w", err)
	}

	if _, err = obj.Attrs(ctx); !errors.Is(err, storage.ErrObjectNotExist) {
		return fmt.Errorf("delete gs object: %w", err)
	}
	return nil
}

// Gets objects in bucket given prefix
func (gs *gcloudStorage) GetObjects(ctx context.Context, prefix string, expires time.Time) ([]string, error) {
	// get objects in bucket iterator
	var objNames []string
	sQuery := &storage.Query{Prefix: prefix}
	it := gs.client.Bucket(gs.bktName).Objects(ctx, sQuery)
	for {
		objAttrs, err := it.Next()
		_ = objAttrs
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objNames = append(objNames, objAttrs.Name)
	}
	// get signed urls
	eg, _ := errgroup.WithContext(ctx)
	signedUrls := make([]string, len(objNames))
	for i, objN := range objNames {
		workerCount := i
		workerObjName := objN
		eg.Go(func() error {
			url, err := GSSignedURL(gs.bktName, workerObjName, expires)
			if err != nil {
				return err
			}
			signedUrls[workerCount] = url
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return signedUrls, nil
}

func gsJsonKeyDo() ([]byte, error) {
	gsJsonKeyOnce.Do(func() {
		gsJsonKey, gsJsonKeyErr = ioutil.ReadFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		if gsJsonKeyErr != nil {
			return
		}
	})
	return gsJsonKey, gsJsonKeyErr
}

// Get a signed url valid accessible for limited time
func GSSignedURL(bktName, objName string, expires time.Time) (string, error) {
	key, err := gsJsonKeyDo()
	if err != nil {
		return "", err
	}
	conf, err := google.JWTConfigFromJSON(key)
	if err != nil {
		return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}
	opts := &storage.SignedURLOptions{
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        expires,
	}
	url, err := storage.SignedURL(bktName, objName, opts)
	if err != nil {
		return "", err
	}
	return url, nil
}
