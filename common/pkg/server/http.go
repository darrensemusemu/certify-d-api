package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

func RunHttpServer(createHandler func(chi.Router) http.Handler) error {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", createHandler(apiRouter))

	port := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	fmt.Fprintf(os.Stdout, "Starting HTTP server on PORT %v\n", port)

	err := http.ListenAndServe(port, rootRouter)
	if err != nil {
		logger.Log.Errorw(err.Error())
		return err
	}
	return nil
}

func setMiddlewares(router *chi.Mux) {
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
