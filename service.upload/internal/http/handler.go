package http

import (
	"net/http"

	"github.com/darrensemusemu/certify-d-api/service.upload/internal/store"
	"github.com/darrensemusemu/certify-d-api/service.upload/pkg/api"
	"github.com/deepmap/oapi-codegen/pkg/types"
)

// Handler implements api.ServerInterface
var _ api.ServerInterface = (*Handler)(nil)

//
// var _ http.Handler = (*Handler)(nil)

// Properties of http upload Handler
type Handler struct {
	storeSvc store.Service
}

// Creates a new Handler
func NewHandler(StoreSvc store.Service) (*Handler, error) {
	return &Handler{
		storeSvc: StoreSvc,
	}, nil
}

func (h *Handler) HandleGetServerAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleGetServerReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleGetStore(w http.ResponseWriter, r *http.Request, params api.HandleGetStoreParams) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleAddStore(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleDeleteStoreById(w http.ResponseWriter, r *http.Request, storeId types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleDeleteStoreFiles(w http.ResponseWriter, r *http.Request, storeId types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleDeleteStoreFileById(w http.ResponseWriter, r *http.Request, storeId types.UUID, fileId types.UUID) {
	w.WriteHeader(http.StatusOK)
}
