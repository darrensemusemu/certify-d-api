package jwt

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/MicahParks/keyfunc"
	"github.com/darrensemusemu/certify-d-api/common/pkg/env"
	"github.com/darrensemusemu/certify-d-api/common/pkg/httpresult"
	"github.com/golang-jwt/jwt/v4"
	kratosClient "github.com/ory/kratos-client-go"
	"github.com/spf13/viper"
)

func init() {
	viper.BindEnv(EnvVarJWKSUrl)
	viper.BindEnv(env.EnvVarKey)
}

// JWKS url envirmoment variable
const EnvVarJWKSUrl = "JWKS_URL"

type key int

const (
	// Claims context value key
	KratosClaimsCxtKey key = iota + 1
)

// Kratos claims implements interface
var _ JWTValidator = (*KratosClaims)(nil)

// JWT session claims
type KratosClaims struct {
	// Kratos user session
	Session kratosClient.Session `json:"session,omitempty"`
	jwt.RegisteredClaims

	mu sync.Mutex
}

// Creates a new JWT Claim
func NewKratosClaims() *KratosClaims {
	return &KratosClaims{
		Session:          kratosClient.Session{},
		RegisteredClaims: jwt.RegisteredClaims{},
		mu:               sync.Mutex{},
	}
}

// Retrieve a JSON Web Key (JWKS)
func (s *KratosClaims) GetJWKS() (jwt.Keyfunc, error) {
	jwksUrl := viper.GetString(EnvVarJWKSUrl)
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
func (s *KratosClaims) Parse(jwtB64 string, keyFunc jwt.Keyfunc) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	token, err := jwt.ParseWithClaims(jwtB64, s, keyFunc)
	if err != nil {
		return fmt.Errorf("pass jwt: %w", err)
	}
	claims, ok := token.Claims.(*KratosClaims)
	if !ok && !token.Valid {
		_ = claims
		return err
	}
	_ = claims
	s = claims
	return err
}

// HTTP middleware setting a value on the request context.
func KratosAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}
		jwtB64 := strings.Split(authHeader, " ")[1]
		claims := NewKratosClaims()
		err := Validate(jwtB64, claims)
		if err != nil {
			httpresult.ServeJSONProblem(http.StatusForbidden, err)(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), KratosClaimsCxtKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Retrieves kratos claims from ctx
func KratosClaimsFromCtx(ctx context.Context) (*KratosClaims, error) {
	value := ctx.Value(KratosClaimsCxtKey)
	claims, ok := value.(*KratosClaims)
	if !ok {
		return nil, fmt.Errorf("kratos claims ctx: err getting claims")
	}
	return claims, nil
}
