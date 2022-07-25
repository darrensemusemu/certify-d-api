package rest

import (
	"net/http"
	"time"

	"github.com/darrensemusemu/certify-d-api/common/pkg/middleware/jwt"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

// Configures http server rooutes and middleware
func Handler(s *server.Server) http.Handler {
	r := chi.NewRouter()
	validate = validator.New()

	// r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 30))

	// Public routes
	r.Group(func(r chi.Router) {

		// Health checks
		r.Route("/health", func(r chi.Router) {
			r.Get("/alive", checkServerAlive())
			r.Get("/ready", checkServerReady())
		})
	})

	// Private routes
	r.Group(func(r chi.Router) {
		// User routes
		r.Use(jwt.KratosAuthClaims(s.Logger))
		r.Post("/signup", SignUpUser(s.UserService, s.Logger))
	})

	return r
}
