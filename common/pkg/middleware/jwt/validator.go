package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// JWKS could not retrieve from URL
var ErrJwksUrlNotFound = errors.New("jwks url found err: failed to get JWKS_URL env variable")

// Validates a JWT session
type JWTValidator interface {
	// Retrieves JSON Web Keys JWKS
	GetJWKS() (jwt.Keyfunc, error)
	// Parses and validates a jwt
	Parse(jwtB64 string, keyFunc jwt.Keyfunc) error
}

// Check that jwt is valid
func Validate(jwtB64 string, v JWTValidator) error {
	keyFunc, err := v.GetJWKS()
	if err != nil {
		return fmt.Errorf("jwt claims err: %w", err)
	}
	err = v.Parse(jwtB64, keyFunc)
	if err != nil {
		return fmt.Errorf("verify jwt token err: %w", err)
	}
	return nil
}
