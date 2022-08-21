package store

import (
	"context"
)

// Provide store related operations
type Service interface {
	// Add a new store
	Add(context.Context, Store) Store
	// Get store given an ID
	GetByID(context.Context, string) Store
	// Get store given an ID
	GetByRef(context.Context, string) Store
	// Removes a store given ID
	RemoveByID(context.Context, string)
}
