package main

import (
	"fmt"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/darrensemusemu/certify-d-api/common/pkg/env"
	"github.com/darrensemusemu/certify-d-api/common/pkg/middleware/jwt"
	"github.com/darrensemusemu/certify-d-api/common/pkg/server"
	httpuser "github.com/darrensemusemu/certify-d-api/service.user/internal/http"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/storage/db"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/darrensemusemu/certify-d-api/service.user/pkg/api"
)

type config struct {
	Port int
}

func main() {
	cfg := config{}
	pflag.IntVar(&cfg.Port, "port", 8080, "port number for service")
	pflag.Parse()

	err := run(cfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func initViper() error {
	viper.BindPFlags(pflag.CommandLine)
	viper.BindEnv("db_conn")
	viper.BindEnv("env")
	viper.BindEnv(jwt.EnvVarJWKSUrl)
	viper.BindEnv("svc_name")

	viper.SetDefault("env", env.Development)
	viper.SetDefault("svc_name", "user")

	if viper.GetString("env") != env.Development {
		return nil
	}

	viper.SetConfigName("server-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("service.user/config")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fatal error config file: %w", err)
	}

	return nil
}

func run(cfg config) (err error) {
	if err := initViper(); err != nil {
		return err
	}

	repo, err := db.NewPostgresDB(viper.GetString("db_conn"))
	defer func() {
		if tempErr := repo.Close(); tempErr != nil {
			err = tempErr
		}
	}()
	if err != nil {
		return err
	}

	userSvc, err := user.NewService(repo)
	if err != nil {
		return err
	}

	restHandler, err := httpuser.NewHandler(userSvc)
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
			r.Use(jwt.KratosAuthMiddleware)
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
