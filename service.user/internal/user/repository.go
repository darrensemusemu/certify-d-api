package user

import "context"

// Provides access to storgae
type Repository interface {
	// Add a user to the provided repository
	AddUser(context.Context, User) (User, error)
	// Checks if a user exists given a user id
	UserExists(ctx context.Context, id string) (bool, error)
	// Get a single user given the user id
	GetUserById(ctx context.Context, id string) (User, error)
	// Update a user in provided repository
	UpdateUser(context.Context, User) (User, error)
}
