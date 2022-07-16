package user

import "fmt"

// Provide user related operations
type Service interface {
	// Add a new user
	Add(User) error
	// Update an existing user
	Update(User) error
}

type service struct {
	r Repository
}

// Create a new User Service
func NewService(r Repository) (Service, error) {
	return service{r}, nil
}

// Handle adding a user to repository
func (s service) Add(user User) error {
	if user.ID == "" || user.Role == "" {
		return fmt.Errorf("add user err: id/role not provided")
	}
	err := s.r.AddUser(user)
	return err
}

// Handle updating user in repository
func (s service) Update(user User) error {
	if user.ID == "" || user.Role == "" {
		return fmt.Errorf("update user err: id/role not provided")
	}
	err := s.r.UpdateUser(user)
	return err
}
