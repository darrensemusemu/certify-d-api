package server

import (
	"net/http"
)

// Properties of server
type Server struct {
	// Handles http requests
	Router http.Handler
}
