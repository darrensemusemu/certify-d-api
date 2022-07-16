package server

import (
	"fmt"

	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
)

// Properties of an http server
type Server struct {
	UserR  user.Repository
	Logger *logger.Logger
}

// Creates a new server
func New(uR user.Repository, l *logger.Logger) (*Server, error) {
	if uR == nil {
		return nil, fmt.Errorf("new server err: user service nil")
	}

	var err error
	if l == nil {
		l, err = logger.New("service.user")
	}
	if err != nil {
		return nil, err
	}

	s := &Server{
		UserR:  uR,
		Logger: l,
	}
	return s, nil
}
