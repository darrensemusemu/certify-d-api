package jwt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

func testHelperClaimsValidator(t *testing.T, ts *httptest.Server) *Claims {
	t.Setenv(EnvVarJWKSUrl, fmt.Sprintf("%v/.well-known/jwks.json", ts.URL))
	return NewClaims()
}

func testHelperAuthServer() *httptest.Server {
	r := chi.NewRouter()
	l, _ := logger.New("KratosAuthClaims")

	// public routes
	r.Group(func(r chi.Router) {
		r.Get("/.well-known/jwks.json", func(w http.ResponseWriter, r *http.Request) {
			jwksString, err := os.ReadFile("./data/id_token.jwks.json")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(200)
			w.Write(jwksString)
		})

	})

	// test protected routes
	r.Group(func(r chi.Router) {
		r.Use(KratosAuthClaims(l))

		r.Get("/auth/kratos", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	return httptest.NewServer(r)
}

func testHelperOverrideJWTTime() {
	tt := time.Unix(1657496900, 0)
	jwt.TimeFunc = func() time.Time {
		return tt
	}

}

func testHelperResetJWTTime() {
	jwt.TimeFunc = time.Now
}
