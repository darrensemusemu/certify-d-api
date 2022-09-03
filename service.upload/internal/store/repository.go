package store

import "context"

// Repo filter options type
type RepoStoreFilterType int

const (
	// No filter options
	RepoStoreFilterNone RepoStoreFilterType = iota
	// Filters by given ID
	RepoStoreFilterByID
	// Filters by given Ref
	RepoStoreFilterByRef
)

// Handle data access
type Repository interface {
	// Add new store
	AddStore(context.Context, Store) error
	// Get a single store
	GetAll(context.Context, RepoStoreFilterType, string) (Store, error)
	// Get a single store
	GetSingleStore(context.Context, RepoStoreFilterType, string) (Store, error)
	// Update a given store
	UpdateStore(context.Context, Store) error
	// Remove a store given an id
	RemoveStoreByID(context.Context, string) error
}

//
type BlobRepository interface {
	//
	GetFile(context.Context)
	//
	DeleteFile(context.Context)
}
