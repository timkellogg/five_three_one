package config

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timkellogg/five_three_one/app/controllers"
	"github.com/timkellogg/five_three_one/app/middlewares"
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

	// Error handlers
	router.NotFoundHandler = middlewares.SetHeaders(controllers.Errors404)

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
