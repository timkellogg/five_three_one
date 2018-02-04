package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/timkellogg/five_three_one/api/controllers"
	"github.com/timkellogg/five_three_one/api/middlewares"
	"github.com/timkellogg/five_three_one/services/authentication"
	"github.com/timkellogg/five_three_one/services/database"
	"github.com/timkellogg/five_three_one/services/session"
)

var routes = Routes{
	Route{"Info", "GET", "/api/info", middlewares.SetHeaders(controllers.InfoShow)},
	Route{"Users Create", "POST", "/api/users/create", middlewares.SetHeaders(controllers.UsersCreate)},
}

// Store - storage objects
type Store struct {
	Database database.Database
	Session  session.Session
	Auth     authentication.AuthService
}

func main() {
	loadEnvironment()

	router := NewRouter(routes)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func loadEnvironment() {
	var environmentFile string
	environment := flag.String("environment", "development", "Indicates the application environment")

	environmentFile = ".env." + *environment
	err := godotenv.Load(environmentFile)
	if err != nil {
		log.Fatal(err)
	}
}
