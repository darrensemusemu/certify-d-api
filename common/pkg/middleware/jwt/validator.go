package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// JWKS could not reticved from URL
var ErrJwksUrlNotFound = errors.New("jwks url found err: failed to get JWKS_URL env variable")

// Validates a JWT session
type JWTValidator interface {
	// Retrieves JSON Web Keys JWKS
	GetJWKS() (jwt.Keyfunc, error)
	// Parses and validates a jwt
	Parse(jwtB64 string, keyFunc jwt.Keyfunc) error
}
