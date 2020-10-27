package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

// addHTTPDeprecationHeaders is a mux middleware function for adding HTTP
// Deprecation headers to a http handler
func addHTTPDeprecationHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Deprecation", "true")
		w.Header().Set("Link", "<https://docs.cosmos.network/v0.40/interfaces/rest.html>; rel=\"deprecation\"")
		h.ServeHTTP(w, r)
	})
}

// WithHTTPDeprecationHeaders returns a new *mux.Router, identical to its input
// but with the addition of HTTP Deprecation headers. This is used to mark legacy
// amino REST endpoints as deprecated in the REST API.
func WithHTTPDeprecationHeaders(r *mux.Router) *mux.Router {
	subRouter := r.NewRoute().Subrouter()
	subRouter.Use(addHTTPDeprecationHeaders)
	return subRouter
}
