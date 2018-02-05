package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/timkellogg/five_three_one/config"
)

// Route - application endpoint accessible through public http methods
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc ContextHandlerFunc
}

// Routes - collection of endpoints
type Routes []Route

// ContextHandlerFunc - placeholder
type ContextHandlerFunc func(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request)

func (context *ContextHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context.ServeHTTP(w, r)
}

// NewRouter establishes the root application router
func NewRouter(context *config.ApplicationContext, routes Routes, notFoundHandler http.HandlerFunc) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = notFoundHandler

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			// TODO: fix HandlerFunc. Right now, it is overriding previous routes and setting a single handler for all
			// this means that the last route is the only router with a handler
			HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				logRoute(setJSONHeader(route.HandlerFunc), route.Name)(context, w, r)
			})
	}

	return router
}

func logRoute(inner ContextHandlerFunc, name string) ContextHandlerFunc {
	return func(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner(c, w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	}
}

func setJSONHeader(inner ContextHandlerFunc) ContextHandlerFunc {
	return func(c *config.ApplicationContext, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		inner(c, w, r)
	}
}
