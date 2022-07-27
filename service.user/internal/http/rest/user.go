package rest

import (
	"encoding/json"
	"net/http"

	"github.com/darrensemusemu/certify-d-api/common/pkg/httpresult"
	"github.com/darrensemusemu/certify-d-api/common/pkg/logger"
	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
)

type SignUpRequest struct {
	ID   string `json:"id,omitempty" validate:"required,uuid4"`
	Role string `json:"role,omitempty" validate:"required"`
}

// Add user http handler
func SignUpUser(userSvc user.Service, l *logger.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req *SignUpRequest

		var err error
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpresult.ServeJSONProblem(http.StatusInternalServerError, err)(w, r)
			l.Errorw("SignUpUser err: could not decode request body")
			return
		}

		if err = validate.Struct(req); err != nil {
			httpresult.ServeJSONProblem(http.StatusBadRequest, err)(w, r)
			l.Errorw("SignUpUser err: could validate request body values")
			return
		}

		var userExists bool
		if userExists, err = userSvc.Exists(ctx, req.ID); err != nil {
			httpresult.ServeJSONProblem(http.StatusBadRequest, err)(w, r)
			l.Errorw("SignUpUser err: could validate if user user exists")
			return
		}

		usr := user.User{ID: req.ID, Role: req.Role}
		if !userExists {
			usr, err = userSvc.Add(ctx, usr)
		} else {
			usr, err = userSvc.Update(ctx, user.User{ID: req.ID, Role: req.Role})
		}

		if err != nil {
			httpresult.ServeJSONProblem(http.StatusBadRequest, err)(w, r)
			l.Errorw("SignUpUser err: could add/update user")
			return
		}

		w.WriteHeader(http.StatusCreated)
		httpresult.ServeJSON(usr)(w, r)
	})
}
