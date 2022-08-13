package api

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3filter"
)

func init() {
	openapi3filter.RegisterBodyDecoder("image/png", openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder("image/jpeg", openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder("image/jpg", openapi3filter.FileBodyDecoder)
	openapi3filter.RegisterBodyDecoder("application/pdf", openapi3filter.FileBodyDecoder)
}

type HttpServer struct{}

func NewHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) GetServerAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) GetServerReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) GetStore(w http.ResponseWriter, r *http.Request, params GetStoreParams) {
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) AddStore(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) RemoveStoreById(w http.ResponseWriter, r *http.Request, storeId openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) RemoveStoreFiles(w http.ResponseWriter, r *http.Request, storeId openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (h HttpServer) RemoveStoreFileById(w http.ResponseWriter, r *http.Request, storeId openapi_types.UUID, fileId openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}