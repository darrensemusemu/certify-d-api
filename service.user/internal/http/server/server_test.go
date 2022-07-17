package server_test

import (
	"testing"

	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/server"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/matryer/is"
)

type userRepo struct{}

func (r *userRepo) AddUser(user.User) error    { return nil }
func (r *userRepo) UpdateUser(user.User) error { return nil }

func TestServerNew(t *testing.T) {
	is := is.New(t)
	l, err := logger.New("testservice")
	is.NoErr(err)

	tests := []struct {
		name         string
		expectErr    bool
		expectServer bool
		server       *server.Server
	}{
		{
			name:         "check zero value params",
			expectErr:    true,
			expectServer: false,
			server:       &server.Server{},
		},
		{
			name:         "check nil logger",
			expectErr:    false,
			expectServer: true,
			server: &server.Server{
				UserR:  &userRepo{},
				Logger: nil,
			},
		},
		{
			name:         "pass valid params",
			expectErr:    false,
			expectServer: true,
			server: &server.Server{
				UserR:  &userRepo{},
				Logger: l,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.New(tt.server.UserR, tt.server.Logger)
			is.True(tt.expectErr == (err != nil))
			is.True(tt.expectServer == (s != nil))
		})
	}
}
