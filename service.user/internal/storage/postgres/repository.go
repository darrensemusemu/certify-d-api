package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
)

type storage struct {
	DB *sql.DB
}

// Create a postgres strorage from a connection string
func NewWithConnString(connString string) (*storage, error) {
	if connString == "" {
		return nil, fmt.Errorf("new storage err: conn string cannot be emplty")
	}

	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}

	s := &storage{
		DB: db,
	}
	return s, nil
}

// Add a new user to repository
func (s *storage) AddUser(ctx context.Context, user user.User) error {
	return nil
}

// Updates an existing user to repository
func (s *storage) UpdateUser(ctx context.Context, user user.User) error {
	return nil
}
