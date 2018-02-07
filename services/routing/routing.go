package routing

import (
	"net/http"

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
		route := route // make a copy of the route for use in the lambda

		router.Methods(route.Method).
			PathPrefix("/api").
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				logRoute(setJSONHeader(route.HandlerFunc), route.Name)(context, w, r)
			})
	}

	return router
}
