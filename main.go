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
	routing.Route{Name: "Info", Method: "GET", Pattern: "/info", HandlerFunc: handlers.InfoShow},
	routing.Route{Name: "Users Create", Method: "POST", Pattern: "/users/create", HandlerFunc: handlers.UsersCreate},
	routing.Route{Name: "Users Show", Method: "GET", Pattern: "/users/me", HandlerFunc: handlers.UsersShow},
	routing.Route{Name: "Authorize", Method: "POST", Pattern: "/oauth/login", HandlerFunc: handlers.Authorize},
	routing.Route{Name: "Token", Method: "POST", Pattern: "/oauth/token", HandlerFunc: handlers.Token},
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
