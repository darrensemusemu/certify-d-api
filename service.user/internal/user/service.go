package user

import (
	"context"
	"fmt"
)

// Provide user related operations
type Service interface {
	// Add a new user
	Add(ctx context.Context, user User) (User, error)
	// Check user account exists
	Exists(ctx context.Context, id string) (bool, error)
	// Gets a user by id
	GetById(ctx context.Context, id string) (User, error)
	// Update an existing user
	Update(ctx context.Context, user User) (User, error)
}

type service struct {
	r Repository
}

// Create a new User Service
func NewService(r Repository) (Service, error) {
	return service{r}, nil
}

// Handle adding a user to repository
func (s service) Add(ctx context.Context, user User) (User, error) {
	if user.ID == "" || user.Role == "" {
		return user, fmt.Errorf("add user err: id/role not provided")
	}
	user, err := s.r.AddUser(ctx, user)
	if err != nil {
		return user, fmt.Errorf("add user err: %v", err)
	}
	return user, err
}

// Check user exists
func (s service) Exists(ctx context.Context, id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("user exists err: id not provided")
	}
	exists, err := s.r.UserExists(ctx, id)
	if err != nil {
		return false, fmt.Errorf("user exists err: %v", err)
	}
	return exists, err
}

// Handle Get User by id
func (s service) GetById(ctx context.Context, id string) (User, error) {
	if id == "" {
		return User{}, fmt.Errorf("get userby id err: id not provided")
	}
	user, err := s.r.GetUserById(ctx, id)
	if err != nil {
		return user, fmt.Errorf("get userby id err: %v", err)

	}
	return user, err
}

// Handle updating user in repository
func (s service) Update(ctx context.Context, user User) (User, error) {
	if user.ID == "" || user.Role == "" {
		return user, fmt.Errorf("update user err: id/role not provided")
	}
	user, err := s.r.UpdateUser(ctx, user)
	if err != nil {
		return user, fmt.Errorf("update user err: %v", err)
	}
	return user, err
}
