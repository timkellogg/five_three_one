package main

import (
	"log"
	"net/http"
	"os"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/services/authentication"

	"github.com/timkellogg/five_three_one/services/session"

	"github.com/timkellogg/five_three_one/services/database"

	_ "github.com/lib/pq"
	"github.com/timkellogg/five_three_one/handlers"
	"github.com/timkellogg/five_three_one/services/router"
)

var context = config.ApplicationContext{
	Database: database.NewDatabase().Store,
	Session:  session.NewSession().Memcache,
	Auth:     authentication.AuthService{},
}

var routes = router.Routes{
	router.Route{"Info", "GET", "/api/info", handlers.InfoShow},
	router.Route{"Users Create", "POST", "/api/users/create", handlers.UsersCreate},
}

func main() {
	config.LoadEnvironment()
	config.PerformEnvChecks(context)

	notFoundHandler := handlers.Errors404
	router := router.NewRouter(&context, routes, notFoundHandler)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
