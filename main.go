package main

import (
	"log"
	"net/http"
	"os"

	"github.com/timkellogg/five_three_one/services/authentication"
	"github.com/timkellogg/five_three_one/services/database"
	"github.com/timkellogg/five_three_one/services/routing"
	"github.com/timkellogg/five_three_one/services/session"

	"github.com/timkellogg/five_three_one/config"

	"github.com/timkellogg/five_three_one/handlers"
)

var context = config.ApplicationContext{}

var routes = routing.Routes{
	routing.Route{"Info", "GET", "/info", handlers.InfoShow},
	routing.Route{"Users Create", "POST", "/users/create", handlers.UsersCreate},
	routing.Route{"Users Show", "GET", "/users/me", handlers.UsersShow},
	routing.Route{"Login", "POST", "/auth/login", handlers.Login},
}

func main() {
	config.LoadEnvironment()

	context.Database = database.NewDatabase().Store
	context.Session = session.NewSession().Memcache
	context.Auth = authentication.AuthService{}

	context.PerformEnvChecks()

	router := routing.NewRouter(&context, routes, handlers.Errors404)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
