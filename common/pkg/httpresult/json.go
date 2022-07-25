package httpresult

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/moogar0880/problems"
)

//
func ServeJSON(v interface{}) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(v); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}

//
func ServeJSONProblem(statusCode int, err error) http.HandlerFunc {
	svc := os.Getenv("svc")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := problems.NewStatusProblem(statusCode)
		p.Detail = err.Error()
		p.Instance = fmt.Sprintf("/api/%s%v", svc, r.RequestURI)
		p.Type = fmt.Sprintf("/probs/%s%v", svc, r.RequestURI)

		w.WriteHeader(p.Status)
		w.Header().Set("Content-Type", "application/problem+json")
		if err := json.NewEncoder(w).Encode(p); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
