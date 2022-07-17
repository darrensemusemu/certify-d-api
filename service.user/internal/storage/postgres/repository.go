package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/darrensemusemu/certify-d-api/service.user/pkg/models"
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

// Adds a new user to repository
func (s *storage) AddUser(ctx context.Context, user user.User) (user.User, error) {
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

// Gets a user given an id
func (s *storage) GetUserByID(ctx context.Context, id string) (user.User, error) {
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
func (s *storage) UpdateUser(ctx context.Context, user user.User) (user.User, error) {
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
