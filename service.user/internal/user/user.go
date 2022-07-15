package user

// Properties of User
type User struct {
	// User assigned unique identifier
	ID string `json:"id,omitempty"`
	// User assigned role
	Role string `json:"role,omitempty"`
	// user assigned permissions
	Permissions []string `json:"permissions,omitempty"`
}
