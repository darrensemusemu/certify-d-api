// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Application liveness probe
	// (GET /health/alive)
	GetServerAlive(w http.ResponseWriter, r *http.Request)
	// Application accepting requests
	// (GET /health/ready)
	GetServerReady(w http.ResponseWriter, r *http.Request)
	// Get a list stores
	// (GET /store)
	GetStore(w http.ResponseWriter, r *http.Request, params GetStoreParams)
	// Create a store
	// (POST /store)
	AddStore(w http.ResponseWriter, r *http.Request)
	// Delete store
	// (DELETE /store/{storeId})
	RemoveStoreById(w http.ResponseWriter, r *http.Request, storeId openapi_types.UUID)
	// Upload file
	// (POST /store/{storeId}/file)
	RemoveStoreFiles(w http.ResponseWriter, r *http.Request, storeId openapi_types.UUID)
	// Delete file
	// (DELETE /store/{storeId}/file/{fileId})
	RemoveStoreFileById(w http.ResponseWriter, r *http.Request, storeId openapi_types.UUID, fileId openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetServerAlive operation middleware
func (siw *ServerInterfaceWrapper) GetServerAlive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetServerAlive(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetServerReady operation middleware
func (siw *ServerInterfaceWrapper) GetServerReady(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetServerReady(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetStore operation middleware
func (siw *ServerInterfaceWrapper) GetStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetStoreParams

	// ------------- Required query parameter "storeRef" -------------
	if paramValue := r.URL.Query().Get("storeRef"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "storeRef"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "storeRef", r.URL.Query(), &params.StoreRef)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "storeRef", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetStore(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// AddStore operation middleware
func (siw *ServerInterfaceWrapper) AddStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddStore(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RemoveStoreById operation middleware
func (siw *ServerInterfaceWrapper) RemoveStoreById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "storeId" -------------
	var storeId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "storeId", chi.URLParam(r, "storeId"), &storeId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "storeId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RemoveStoreById(w, r, storeId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RemoveStoreFiles operation middleware
func (siw *ServerInterfaceWrapper) RemoveStoreFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "storeId" -------------
	var storeId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "storeId", chi.URLParam(r, "storeId"), &storeId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "storeId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RemoveStoreFiles(w, r, storeId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// RemoveStoreFileById operation middleware
func (siw *ServerInterfaceWrapper) RemoveStoreFileById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "storeId" -------------
	var storeId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "storeId", chi.URLParam(r, "storeId"), &storeId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "storeId", Err: err})
		return
	}

	// ------------- Path parameter "fileId" -------------
	var fileId openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "fileId", chi.URLParam(r, "fileId"), &fileId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "fileId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RemoveStoreFileById(w, r, storeId, fileId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/health/alive", wrapper.GetServerAlive)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/health/ready", wrapper.GetServerReady)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/store", wrapper.GetStore)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/store", wrapper.AddStore)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/store/{storeId}", wrapper.RemoveStoreById)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/store/{storeId}/file", wrapper.RemoveStoreFiles)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/store/{storeId}/file/{fileId}", wrapper.RemoveStoreFileById)
	})

	return r
}