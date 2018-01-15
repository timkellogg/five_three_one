package middlewares

import (
	"net/http"
)

// RequireAuth - makes sure user has correct token
func RequireAuth() {
	// check token
}

// SetHeaders - sets application headers on route
func SetHeaders(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
