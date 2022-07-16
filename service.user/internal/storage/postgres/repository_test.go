package postgres_test

import (
	"testing"

	"github.com/matryer/is"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/storage/postgres"
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
