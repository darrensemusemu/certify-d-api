package http

import (
	"net/http"

	"github.com/darrensemusemu/certify-d-api/service.user/internal/user"
	"github.com/darrensemusemu/certify-d-api/service.user/pkg/api"
)

// Handler implements api.ServerInterface
var _ api.ServerInterface = (*Handler)(nil)

//
type Handler struct {
	userSvc user.Service
}

// Creates a new Handler
func NewHandler(userSvc user.Service) (*Handler, error) {
	h := &Handler{
		userSvc: userSvc,
	}
	return h, nil
}

func (h *Handler) HandleGetServerAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleGetServerReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// ctx := r.Context()

	// 	var userExists bool
	// 	if userExists, err = userSvc.Exists(ctx, req.ID); err != nil {
	// 		httpresult.ServeJSONProblem(http.StatusBadRequest, err)(w, r)
	// 		l.Errorw("SignUpUser err: could validate if user user exists")
	// 		return
	// 	}

	// 	usr := user.User{ID: req.ID, Role: req.Role}
	// 	if !userExists {
	// 		usr, err = userSvc.Add(ctx, usr)
	// 	} else {
	// 		usr, err = userSvc.Update(ctx, user.User{ID: req.ID, Role: req.Role})
	// 	}

	// 	if err != nil {
	// 		httpresult.ServeJSONProblem(http.StatusBadRequest, err)(w, r)
	// 		l.Errorw("SignUpUser err: could add/update user")
	// 		return
	// 	}

	// 	w.WriteHeader(http.StatusCreated)
	// 	httpresult.ServeJSON(usr)(w, r)
}
