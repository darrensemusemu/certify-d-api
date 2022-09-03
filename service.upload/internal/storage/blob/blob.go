//
package blob

import (
	"context"
	"io"
	"time"
)

// Provides blob related operations
type Service interface {
	// Creates a new object
	CreateObject(ctx context.Context, objectName string, reader io.Reader) (err error)
	// Deletes an object
	DeleteObject(ctx context.Context, objectName string) (err error)
	// Return a slice of urls to access objects for limited time
	GetObjects(ctx context.Context, prefix string, expires time.Time) ([]string, error)
}
