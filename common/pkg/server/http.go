package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

// Creates a new http server
func NewHttpServer() (*Server, error) {
	return &Server{Router: nil}, nil
}

// Registers API routes and middleware
func (s *Server) RegisterRoutes(createAPIHandler func(chi.Router) http.Handler) {
	rootRouter := chi.NewRouter()

	apiRouter := chi.NewRouter()
	setMiddleware(apiRouter)

	rootRouter.Mount("/api", createAPIHandler(apiRouter))
	s.Router = rootRouter
}

// Runs http server
func (s *Server) RunHttpServer() error {
	port := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	fmt.Fprintf(os.Stdout, "Starting HTTP server on PORT %v\n", port)
	if err := http.ListenAndServe(port, s.Router); err != nil {
		return err
	}
	return nil
}

func setMiddleware(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	// router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	// addCorsMiddleware(router)
	// addAuthMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}
