package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/darrensemusemu/certify-d-api/common/pkg/middleware/jwt"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	os.Setenv(jwt.EnvVarJWKSUrl, "http://oathkeeper-api:4456/.well-known/jwks.json")

	r.Get("/health/alive", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		autHeader := r.Header.Get("Authorization")
		jwtB64 := strings.Split(autHeader, " ")[1]
		claims := jwt.NewClaims()
		err := jwt.Validate(jwtB64, claims)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		if err != nil {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println(err)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r.Header.Get("Authorization")))
	})

	http.ListenAndServe(":8080", r)
}
