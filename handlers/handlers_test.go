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

var testUser = models.User{Email: "test@test.com"}
var routes = routing.Routes{
	routing.Route{"Info", "GET", "/info", InfoShow},
	routing.Route{"Users Create", "POST", "/users/create", UsersCreate},
}

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", "five_three_one_test")
	os.Setenv("DB_USER", "")
	os.Setenv("DB_PASS", "")

	context = config.ApplicationContext{
		Database: database.NewDatabase().Store,
		Session:  session.NewSession().Memcache,
		Auth:     authentication.AuthService{},
	}
	defer context.Database.Close()

	context.TruncateDBTables(tables())

	router = routing.NewRouter(&context, routes, Errors404)
	server = httptest.NewServer(router)

	runTests := m.Run()
	os.Exit(runTests)
}

func tables() []string {
	return []string{"users"}
}
