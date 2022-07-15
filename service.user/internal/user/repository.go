package user

// Provides access to storgae
type Repository interface {
	// Add a user to the provided repository
	Add(user User) error
	// Update a user in provided repository
	Update(user User) error
}
