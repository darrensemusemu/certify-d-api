package user

// Provides access to storgae
type Repository interface {
	// Add a user to the provided repository
	AddUser(user User) error
	// Update a user in provided repository
	UpdateUser(user User) error
}
