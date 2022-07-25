package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/darrensemusemu/certify-d-api/common/pkg/middleware/jwt"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/rest"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/server"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/storage/postgres"
)

type config struct {
	dbConn  string
	jwksUrl string
	port    int
}

func main() {
	cfg := config{}
	flag.StringVar(&cfg.dbConn, "dbConn", "postgres://user_service:user_service@localhost:5432/certify_d", "datababse connection string")
	flag.StringVar(&cfg.jwksUrl, jwt.EnvVarJWKSUrl, "http://localhost:4456/.well-known/jwks.json", "json web keys (JWKS) url")
	flag.IntVar(&cfg.port, "port", 8080, "port number for service")
	flag.Parse()

	err := run(cfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

}

func run(cfg config) error {
	// TODO: use viper???
	err := os.Setenv(jwt.EnvVarJWKSUrl, cfg.jwksUrl)
	if err != nil {
		return err
	}

	err = os.Setenv("svc", "user")
	if err != nil {
		return err
	}

	userRepo, err := postgres.NewWithConnString(cfg.dbConn)
	if err != nil {
		return err
	}

	server, err := server.New(userRepo, nil)
	if err != nil {
		return err
	}

	restHandler := rest.Handler(server)
	err = http.ListenAndServe(fmt.Sprintf(":%v", cfg.port), restHandler)
	if err != nil {
		return err
	}

	return nil
}
