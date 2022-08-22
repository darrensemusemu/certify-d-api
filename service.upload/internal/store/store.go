package store

import "github.com/darrensemusemu/certify-d-api/service.upload/internal/file"

// Properties of a Store (Collection of related blobs)
type Store struct {
	// Assigned unique identifier
	ID string
	// Reference unique to a store
	StoreRef string
	// Files available in the store
	Files []file.File
}
