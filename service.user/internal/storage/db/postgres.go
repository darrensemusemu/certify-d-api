package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/darrensemusemu/certify-d-api/service.user/pkg/models"
)

// Connection string not provided
var ErrDBConnEmpty = errors.New("new db: conn string empty")

// postgresDB implements store.Repository
var _ user.Repository = (*postgresDB)(nil)

type postgresDB struct {
	DB *sql.DB
}

// Create a postgres storage from a connection string
func NewPostgresDB(connString string) (*postgresDB, error) {
	if connString == "" {
		return nil, ErrDBConnEmpty
	}
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}
	s := &postgresDB{
		DB: db,
	}
	return s, nil
}

// Storage close connection
func (s *postgresDB) Close() error {
	err := s.DB.Close()
	return err
}

// Adds a new user to repository
func (s *postgresDB) AddUser(ctx context.Context, user user.User) (user.User, error) {
	role, err := models.Roles(models.RoleWhere.Slug.EQ(user.Role)).One(ctx, s.DB)
	if err != nil {
		return user, fmt.Errorf("add user err: %v", err)
	}

	if user.ID == "" {
		idGen, err := uuid.NewV4()
		if err != nil {
			return user, fmt.Errorf("add user err: %v", err)
		}
		user.ID = idGen.String()
	}

	newUser := models.User{
		ID:     user.ID,
		RoleID: role.ID,
	}

	err = newUser.Insert(ctx, s.DB, boil.Infer())
	if err != nil {
		return user, fmt.Errorf("add user err: %v", err)
	}

	err = newUser.Reload(ctx, s.DB)
	if err != nil {
		return user, fmt.Errorf("add user err: reload %v", err)
	}

	user.ID = newUser.ID
	return user, nil
}

// Check if user exists in repo
func (s *postgresDB) UserExists(ctx context.Context, id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("user exists err: no id was provided")
	}

	exists, err := models.Users(models.UserWhere.ID.EQ(id)).Exists(ctx, s.DB)
	if err != nil {
		return false, fmt.Errorf("user exist err: %v", err)
	}

	return exists, nil
}

// Gets a user given an id
func (s *postgresDB) GetUserById(ctx context.Context, id string) (user.User, error) {
	user := user.User{}
	if id == "" {
		return user, fmt.Errorf("get user by id err: id provided not valid")
	}

	dbUser, err := models.Users(
		models.UserWhere.ID.EQ(id),
		qm.Load(models.UserRels.Role),
	).One(ctx, s.DB)

	if err != nil {
		return user, fmt.Errorf("get user by id err: %v", err)
	}

	user.ID = dbUser.ID
	user.Role = dbUser.R.Role.Slug
	user.Permissions = make([]string, 0)
	return user, nil
}

// Updates an existing user to repository
func (s *postgresDB) UpdateUser(ctx context.Context, user user.User) (user.User, error) {
	if user.ID == "" {
		return user, fmt.Errorf("update user err: user id not provided")
	}

	if user.Role == "" {
		return user, fmt.Errorf("update user err: user role not provided")
	}

	role, err := models.Roles(models.RoleWhere.Slug.EQ(user.Role)).One(ctx, s.DB)
	if err != nil {
		return user, fmt.Errorf("update user err: %v", err)
	}

	updatedUser := models.User{
		ID:     user.ID,
		RoleID: role.ID,
	}

	_, err = updatedUser.Update(ctx, s.DB, boil.Infer())
	if err != nil {
		return user, fmt.Errorf("update user err: %v", err)
	}

	return user, nil
}
