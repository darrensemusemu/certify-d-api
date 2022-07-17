package postgres_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/matryer/is"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/role"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/storage/postgres"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/darrensemusemu/certify-d-api/service.user/pkg/models"
)

var connString string = "postgres://user_service:user_service@localhost:5432/certify_d"

func TestNewWithConnString(t *testing.T) {
	is := is.New(t)

	type args = struct {
		connString string
	}

	tests := []struct {
		name          string
		expectErr     bool
		expectStorage bool
		args          args
	}{
		{
			name:          "check empty conn string",
			expectErr:     true,
			expectStorage: false,
			args:          args{connString: ""},
		},
		{
			name:          "check connection string",
			expectErr:     false,
			expectStorage: true,
			args:          args{connString: connString},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := postgres.NewWithConnString(tt.args.connString)
			is.True(tt.expectErr == (err != nil))
			is.True(tt.expectStorage == (s != nil))

			if !tt.expectStorage {
				return
			}

			err = s.DB.Ping()
			is.NoErr(err)
		})
	}
}

func TestAddUser(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	s, err := postgres.NewWithConnString(connString)
	is.NoErr(err)
	idGen, err := uuid.NewV4()
	is.NoErr(err)
	tempUserId := idGen.String()

	tests := []struct {
		name      string
		expectErr bool
		args      user.User
	}{
		{
			name:      "empty role field",
			expectErr: true,
			args: user.User{
				ID:          "",
				Role:        "",
				Permissions: []string{},
			},
		},
		{
			name:      "add user with no id",
			expectErr: false,
			args: user.User{
				ID:          "",
				Role:        role.TesterSlug,
				Permissions: []string{},
			},
		},
		{
			name:      "add user with id",
			expectErr: false,
			args: user.User{
				ID:          tempUserId,
				Role:        role.TesterSlug,
				Permissions: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := s.AddUser(ctx, tt.args)
			is.True(tt.expectErr == (err != nil))

			if tt.expectErr {
				return
			}

			is.True(u.ID != "")
			is.Equal(u.Role, tt.args.Role)

			delCount, err := models.Users(models.UserWhere.ID.EQ(u.ID)).DeleteAll(ctx, s.DB)
			is.NoErr(err)
			is.Equal(delCount, int64(1))
		})
	}
}

func TestGetUserByID(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	s, err := postgres.NewWithConnString(connString)
	is.NoErr(err)

	idGen, err := uuid.NewV4()
	is.NoErr(err)
	tempUserId := idGen.String()

	tempUser := user.User{
		ID:          tempUserId,
		Role:        role.TesterSlug,
		Permissions: []string{},
	}

	tempUser, err = s.AddUser(ctx, tempUser)
	is.NoErr(err)

	tests := []struct {
		name      string
		expectErr bool
		args      string
	}{
		{
			name:      "id field not provided",
			expectErr: true,
			args:      "",
		},
		{
			name:      "invalid user id",
			expectErr: true,
			args:      "23",
		},
		{
			name:      "valid user id",
			expectErr: false,
			args:      tempUser.ID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := s.GetUserByID(ctx, tt.args)
			is.True(tt.expectErr == (err != nil))

			if tt.expectErr {
				return
			}

			is.True(u.ID == tempUser.ID)
			is.True(u.Role == tempUser.Role)
			is.True(reflect.DeepEqual(u, tempUser))
		})
	}
	_, err = models.Users(models.UserWhere.ID.EQ(tempUser.ID)).DeleteAll(ctx, s.DB)
	is.NoErr(err)
}

func TestUpdateUser(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	s, err := postgres.NewWithConnString(connString)
	is.NoErr(err)

	idGen, err := uuid.NewV4()
	is.NoErr(err)
	tempUserId := idGen.String()

	tempUser := user.User{
		ID:          tempUserId,
		Role:        role.CustomerSlug,
		Permissions: []string{},
	}

	tempUser, err = s.AddUser(ctx, tempUser)
	is.NoErr(err)

	tests := []struct {
		name      string
		expectErr bool
		args      user.User
	}{
		{
			name:      "id/role field not provided",
			expectErr: true,
			args: user.User{
				ID:          tempUser.ID,
				Role:        "",
				Permissions: []string{},
			},
		},
		{
			name:      "role field not provided",
			expectErr: true,
			args: user.User{
				ID:          "",
				Role:        role.TesterSlug,
				Permissions: []string{},
			},
		},
		{
			name:      "update user fields",
			expectErr: false,
			args: user.User{
				ID:          tempUser.ID,
				Role:        role.TesterSlug,
				Permissions: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := s.UpdateUser(ctx, tt.args)
			is.True(tt.expectErr == (err != nil))

			if tt.expectErr {
				return
			}

			is.True(u.ID == tt.args.ID)
			is.True(u.ID == tempUser.ID)
			is.True(u.Role == tt.args.Role)
			is.True(u.Role != tempUser.Role)
		})
	}
	_, err = models.Users(models.UserWhere.ID.EQ(tempUser.ID)).DeleteAll(ctx, s.DB)
	is.NoErr(err)
}
