package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/darrensemusemu/certify-d-api/common/pkg/env"
	"github.com/darrensemusemu/certify-d-api/common/pkg/middleware/jwt"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/rest"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/server"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/storage/postgres"
)

type config struct {
	Port int
}

func main() {
	cfg := config{}
	pflag.IntVar(&cfg.Port, "port", 8080, "port number for service")
	pflag.Parse()

	if err := initViper(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

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

	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %w", err)
	}

	return nil
}

func run(cfg config) error {
	repo, err := postgres.NewWithConnString(viper.GetString("db_conn"))
	defer func() {
		err = repo.Close()
		if err != nil {
			log.Fatalf("close userrepo err: %v", err)
		}
	}()
	if err != nil {
		return err
	}

	server, err := server.New(repo, nil)
	if err != nil {
		return err
	}

	restHandler := rest.Handler(server)
	log.Printf("Running server on port: %v", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), restHandler)
	if err != nil {
		return err
	}

	return nil
}
