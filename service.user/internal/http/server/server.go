package server

import (
	"fmt"

	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
)

// Properties of an http server
type Server struct {
	UserService user.Service
	Logger      *logger.Logger
}

type repository interface {
	user.Repository
}

// Creates a new server
func New(r repository, l *logger.Logger) (*Server, error) {
	if r == nil {
		return nil, fmt.Errorf("new server err: user service nil")
	}

	var err error
	if l == nil {
		l, err = logger.New("service.user")
	}
	if err != nil {
		return nil, err
	}

	userService, err := user.NewService(r)
	if err != nil {
		return nil, err
	}

	s := &Server{
		UserService: userService,
		Logger:      l,
	}
	return s, nil
}
