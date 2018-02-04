package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Route - application endpoint accessible through public http methods
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - collection of endpoints
type Routes []Route

// Cors - Cors with logging
type Cors struct {
	Log *log.Logger
}

// NewRouter establishes the root application router
func NewRouter(routes Routes, notFoundHandler http.HandlerFunc) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = notFoundHandler

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logRoute(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func logRoute(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
