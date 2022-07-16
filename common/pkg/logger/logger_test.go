package logger_test

import (
	"testing"

	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/matryer/is"
)

func TestNew(t *testing.T) {
	is := is.New(t)
	l, err := logger.New("service.user")
	is.NoErr(err)
	is.True(l != nil)
}
