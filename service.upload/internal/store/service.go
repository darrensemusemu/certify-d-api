package store

import (
	"context"
)

// Provide store related operations
type Service interface {
	// Add a new store
	Add(context.Context, Store) (Store, error)
	// Get store given an ID
	GetByID(context.Context, string) (Store, error)
	// Get store given an ID
	GetByRef(context.Context, string) (Store, error)
	// Removes a store given ID
	RemoveByID(context.Context, string) error
}

// service implements Service interface
var _ Service = (*service)(nil)

type service struct {
	r Repository
}

// Creates a new store service
func NewService(r Repository) (Service, error) {
	return service{r: r}, nil
}

// TODO: comments
func (s service) Add(context.Context, Store) (Store, error) {
	return Store{}, nil
}
func (s service) GetByID(context.Context, string) (Store, error) {
	return Store{}, nil
}
func (s service) GetByRef(context.Context, string) (Store, error) {
	return Store{}, nil
}
func (s service) RemoveByID(context.Context, string) error {
	return nil
}
