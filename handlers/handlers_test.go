package handlers

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/timkellogg/five_three_one/services/routing"

	"github.com/gorilla/mux"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/models"
	"github.com/timkellogg/five_three_one/services/authentication"
	"github.com/timkellogg/five_three_one/services/database"
	"github.com/timkellogg/five_three_one/services/session"
)

var context config.ApplicationContext
var router *mux.Router
var server *httptest.Server

var testUser = models.User{Email: "test@test.com", Password: "password"}
var routes = routing.Routes{
	routing.Route{Name: "Info", Method: "GET", Pattern: "/info", HandlerFunc: InfoShow},
	routing.Route{Name: "Users Create", Method: "POST", Pattern: "/users/create", HandlerFunc: UsersCreate},
	routing.Route{Name: "Users Show", Method: "GET", Pattern: "/users/me", HandlerFunc: UsersShow},
	routing.Route{Name: "Authorize", Method: "POST", Pattern: "/oauth/authorize", HandlerFunc: Authorize},
}

// need server to be able to keep track of cookies

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", "five_three_one_test")
	os.Setenv("DATABASE_URL", "postgres://tkellogg:password@localhost:5432/five_three_one_test?sslmode=disable")

	context = config.ApplicationContext{
		Database: database.NewDatabase().Store,
		Session:  session.NewSession().Memcache,
		Auth:     authentication.AuthService{},
	}
	defer context.Database.Close()

	router = routing.NewRouter(&context, routes, Errors404)
	server = httptest.NewServer(router)

	runTests := m.Run()
	os.Exit(runTests)
}

func SetAuthorized() {
	//

	//
}
