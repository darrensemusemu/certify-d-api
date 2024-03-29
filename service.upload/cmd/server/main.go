package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/darrensemusemu/certify-d-api/common/pkg/env"
	"github.com/darrensemusemu/certify-d-api/common/pkg/server"
	httpUpload "github.com/darrensemusemu/certify-d-api/service.upload/internal/http"
	"github.com/darrensemusemu/certify-d-api/service.upload/internal/storage/blob"
	"github.com/darrensemusemu/certify-d-api/service.upload/internal/storage/db"
	"github.com/darrensemusemu/certify-d-api/service.upload/internal/store"
	"github.com/darrensemusemu/certify-d-api/service.upload/pkg/api"
)

type config struct {
	Port int
}

func main() {
	cfg := config{}
	pflag.IntVar(&cfg.Port, "port", 8080, "service port number")
	pflag.Parse()

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func initViper() error {
	viper.BindPFlags(pflag.CommandLine)
	viper.BindEnv("env")
	viper.BindEnv("GOOGLE_APPLICATION_CREDENTIALS")
	viper.BindEnv("gs_bucket")
	viper.BindEnv("svc_name")

	viper.SetDefault("env", env.Development)
	viper.SetDefault("svc_name", "service.upload")

	if viper.GetString("env") != env.Development {
		return nil
	}

	viper.SetConfigName("server-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("service.upload/config")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("config file: %w", err)
	}

	return nil
}

func run(cfg config) error {
	ctx := context.Background()

	if err := initViper(); err != nil {
		return err
	}

	repo, err := db.NewPostgresDB("")
	defer func() {
		if tempErr := repo.Close(); tempErr != nil {
			err = tempErr
		}
	}()
	if err != nil {
		return err
	}

	storeSvc, err := store.NewService(repo)
	if err != nil {
		return err
	}

	gsSvc, err := blob.NewGoogleStorage(ctx, viper.GetString("gs_bucket"))
	if err != nil {
		return err
	}

	restHandler, err := httpUpload.NewHandler(storeSvc, gsSvc)
	if err != nil {
		return err
	}

	s, err := server.NewHttpServer()
	if err != nil {
		return err
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		return err
	}

	s.RegisterRoutes(func(r chi.Router) http.Handler {
		r.Mount("/v1", func() http.Handler {
			r.Use(middleware.OapiRequestValidator(swagger))
			s := api.HandlerFromMux(restHandler, r)
			return s
		}())
		return r
	})

	if err = s.RunHttpServer(); err != nil {
		return err
	}
	return nil
}
