package jwt

import (
	"fmt"
	"os"
	"sync"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	kratosClient "github.com/ory/kratos-client-go"
)

// JWKS url envirmoment variable
const EnvVarJWKSUrl = "JWKS_URL"

// JWT session claims
type claims struct {
	// Kratos user session
	Session kratosClient.Session `json:"session,omitempty"`

	jwt.RegisteredClaims
	m sync.Mutex
}

// Check that jwt are valid
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

// Creates a new JWT Claim
func NewClaims() *claims {
	return &claims{}
}

// Retrieve a JSON Web Key (JWKS)
func (s *claims) GetJWKS() (jwt.Keyfunc, error) {
	jwksUrl := os.Getenv(EnvVarJWKSUrl)
	if jwksUrl == "" {
		return nil, ErrJwksUrlNotFound
	}
	jwks, err := keyfunc.Get(jwksUrl, keyfunc.Options{})
	if err != nil {
		return nil, fmt.Errorf("fetch JWKS errror: %w", err)
	}
	return jwks.Keyfunc, nil
}

// Parse and validate base 64 encoded jwt
func (s *claims) Parse(jwtB64 string, keyFunc jwt.Keyfunc) error {
	s.m.Lock()
	defer s.m.Unlock()

	token, err := jwt.ParseWithClaims(jwtB64, s, keyFunc)
	if err != nil {
		return fmt.Errorf("pass jwt: %w", err)
	}
	claims, ok := token.Claims.(*claims)
	if !ok && !token.Valid {
		_ = claims
		return err
	}
	s = claims
	_ = s
	return err
}
