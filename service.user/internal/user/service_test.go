package user_test

import (
	"context"
	"testing"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/matryer/is"
)

type R struct{}

func (r R) AddUser(ctx context.Context, u user.User) (user.User, error) {
	return u, nil
}

func (r R) UserExists(ctx context.Context, id string) (bool, error) {
	return true, nil
}

func (r R) GetUserById(ctx context.Context, id string) (user.User, error) {
	return user.User{}, nil
}

func (r R) UpdateUser(ctx context.Context, u user.User) (user.User, error) {
	return u, nil
}

func TestNewService(t *testing.T) {
	is := is.New(t)

	testRepo := R{}
	svc, err := user.NewService(testRepo)
	is.NoErr(err)
	is.True(svc != nil)
}

func TestServiceAdd(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	testRepo := R{}
	svc, err := user.NewService(testRepo)
	is.NoErr(err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Add(ctx, tt.User)
			is.True(tt.expectErr == (err != nil))
		})
	}
}

func TestServiceUpdate(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	testRepo := R{}
	svc, err := user.NewService(testRepo)
	is.NoErr(err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Update(ctx, tt.User)
			is.True(tt.expectErr == (err != nil))
		})
	}
}

func TestServiceUserExists(t *testing.T) {
	// TODO: implement test"
}

func TestServiceGetUserByIdt(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	testRepo := R{}
	svc, err := user.NewService(testRepo)
	is.NoErr(err)

	type args struct{ id string }

	var tests = []struct {
		name      string
		expectErr bool
		args      args
	}{
		{
			name:      "user id empty",
			expectErr: true,
			args:      args{id: ""},
		},
		{
			name:      "user id provided",
			expectErr: false,
			args:      args{id: "1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.GetById(ctx, tt.args.id)
			is.True(tt.expectErr == (err != nil))
		})
	}
}

var tests = []struct {
	name      string
	expectErr bool
	user.User
}{
	{
		name:      "user struct empty",
		expectErr: true,
		User:      user.User{},
	},
	{
		name:      "user id zero value",
		expectErr: true,
		User: user.User{
			ID:          "",
			Role:        "customer",
			Permissions: []string{},
		},
	},
	{
		name:      "user customer zero value",
		expectErr: true,
		User: user.User{
			ID:          "1",
			Role:        "",
			Permissions: []string{},
		},
	},
	{
		name:      "user customer zero value",
		expectErr: true,
		User: user.User{
			ID:          "1",
			Role:        "",
			Permissions: []string{},
		},
	},
	{
		name:      "user fields valid",
		expectErr: false,
		User: user.User{
			ID:          "1",
			Role:        "customer",
			Permissions: []string{},
		},
	},
}
