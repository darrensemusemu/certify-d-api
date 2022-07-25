package server_test

import (
	"context"
	"testing"

	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/server"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/matryer/is"
)

type userRepo struct{}

func (r *userRepo) AddUser(ctx context.Context, u user.User) (user.User, error) {
	return u, nil
}
func (r *userRepo) UserExists(ctx context.Context, id string) (bool, error) {
	return true, nil
}
func (r *userRepo) GetUserById(ctx context.Context, id string) (user.User, error) {
	return user.User{}, nil
}
func (r *userRepo) UpdateUser(ctx context.Context, u user.User) (user.User, error) {
	return u, nil
}

func TestServerNew(t *testing.T) {
	is := is.New(t)
	l, err := logger.New("testservice")
	is.NoErr(err)

	type args struct {
		tUserRepo user.Repository
		tLogger   *logger.Logger
	}
	tests := []struct {
		name         string
		expectErr    bool
		expectServer bool
		args         args
	}{
		{
			name:         "check no user repo params",
			expectErr:    true,
			expectServer: false,
			args:         args{tUserRepo: nil, tLogger: nil},
		},
		{
			name:         "check nil logger",
			expectErr:    false,
			expectServer: true,
			args:         args{tUserRepo: &userRepo{}, tLogger: nil},
		},
		{
			name:         "pass valid params",
			expectErr:    false,
			expectServer: true,
			args:         args{tUserRepo: &userRepo{}, tLogger: l},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := server.New(tt.args.tUserRepo, tt.args.tLogger)
			is.True(tt.expectErr == (err != nil))
			is.True(tt.expectServer == (s != nil))
		})
	}
}
