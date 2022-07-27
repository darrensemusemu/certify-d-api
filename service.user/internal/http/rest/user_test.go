package rest_test

import (
	"testing"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/http/rest"
	"github.com/go-playground/validator"
	"github.com/gofrs/uuid"
	"github.com/matryer/is"
)

func TestSignUpRequestValidation(t *testing.T) {
	is := is.New(t)

	uuidv4, err := uuid.NewV4()
	is.NoErr(err)

	id := uuidv4.String()
	req := &rest.SignUpRequest{
		ID:   id,
		Role: "asdsad",
	}

	validate := validator.New()
	err = validate.Struct(req)
	is.NoErr(err)
}
