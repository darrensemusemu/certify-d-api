package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/darrensemusemu/certify-d-api/service.upload/internal/store"
)

// Connection string not provided
var ErrDBConnEmpty = errors.New("new db: conn string empty")

// postgresDB implements store.Repository
var _ store.Repository = (*postgresDB)(nil)

type postgresDB struct {
	DB *sql.DB
}

// Create a postgres storage from a connection string
func NewPostgresDB(connString string) (*postgresDB, error) {
	return &postgresDB{}, nil
}

func (s *postgresDB) Close() error {
	return nil
}
func (s *postgresDB) AddStore(context.Context, store.Store) error {
	return nil
}
func (s *postgresDB) GetAll(context.Context, store.RepoStoreFilterType, string) (store.Store, error) {
	return store.Store{}, nil
}
func (s *postgresDB) GetSingleStore(context.Context, store.RepoStoreFilterType, string) (store.Store, error) {
	return store.Store{}, nil
}
func (s *postgresDB) UpdateStore(context.Context, store.Store) error {
	return nil
}
func (s *postgresDB) RemoveStoreByID(context.Context, string) error {
	return nil
}
