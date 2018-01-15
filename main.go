package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/rs/cors"

	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/timkellogg/five_three_one/config"
)

// Application - main container for application database, etc.
type Application struct {
	DB      *sql.DB
	Session *gsm.MemcacheStore
	Router  http.Handler
}

var app Application

// Initialize - start and boostrap application pieces like database, cache, etc.
func (a *Application) Initialize(c config.ApplicationConfiguration) {
	var err error
	connection := fmt.Sprintf("user=%s password=%s dbname=%s", c.DBUser, c.DBPass, c.DBName)

	a.DB, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	loggingLevel, err := strconv.Atoi(c.SessionLoggingLevel)
	if err != nil {
		log.Fatal(err)
	}

	router := config.NewRouter()

	memcacheClient := memcache.New(c.MemecachePort)
	a.Session = gsm.NewMemcacheStore(memcacheClient, c.MemecacheName, []byte(c.SessionSecret))
	a.Session.Logging = loggingLevel

	a.Router = cors.Default().Handler(router)
}

// Run - launch http server to listen on address
func (a *Application) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.Router))
}

func main() {
	a := Application{}
	a.Initialize(config.DevConfig)
	a.Run(":" + config.DevConfig.Port)
}
