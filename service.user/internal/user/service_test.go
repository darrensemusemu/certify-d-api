package user_test

import (
	"testing"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/matryer/is"
)

type R struct{}

func (r R) AddUser(user.User) error {
	return nil
}

func (r R) UpdateUser(user.User) error {
	return nil
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

	testRepo := R{}
	svc, err := user.NewService(testRepo)
	is.NoErr(err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Add(tt.User)
			t.Log(err)
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
