package config

import (
	"log"
	"net/http"

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
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
