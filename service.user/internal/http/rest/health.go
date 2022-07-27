package rest

import "net/http"

// TODO: health checks

// Checks if server is running
func checkServerAlive() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

// Checks if server if ready to handle requests
func checkServerReady() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
