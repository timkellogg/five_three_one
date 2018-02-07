package main

import (
	"log"
	"net/http"
	"os"

	"github.com/timkellogg/five_three_one/services/routing"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/services/authentication"

	"github.com/timkellogg/five_three_one/services/session"

	"github.com/timkellogg/five_three_one/services/database"

	_ "github.com/lib/pq"
	"github.com/timkellogg/five_three_one/handlers"
)

var context = config.ApplicationContext{
	Database: database.NewDatabase().Store,
	Session:  session.NewSession().Memcache,
	Auth:     authentication.AuthService{},
}

var routes = routing.Routes{
	routing.Route{"Info", "GET", "/info", handlers.InfoShow},
	routing.Route{"Users Create", "POST", "/users/create", handlers.UsersCreate},
}

func main() {
	config.LoadEnvironment()
	context.PerformEnvChecks()

	router := routing.NewRouter(&context, routes, handlers.Errors404)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
