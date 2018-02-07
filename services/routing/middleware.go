package routing

import (
	"log"
	"net/http"
	"time"

	"github.com/timkellogg/five_three_one/config"
)

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
