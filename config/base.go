package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/rs/cors"
)

// Application - main container for application database, etc.
type Application struct {
	DB      *sql.DB
	Session *gsm.MemcacheStore
	Router  http.Handler
}

// ApplicationConfiguration - server environment details
type ApplicationConfiguration struct {
	Port                string
	DBName              string
	DBUser              string
	DBPass              string
	MemecachePort       string
	MemecacheName       string
	SessionSecret       string
	SessionLoggingLevel string
}

// Initialize - start and boostrap application pieces like database, cache, etc.
func (a *Application) Initialize(c ApplicationConfiguration) {
	var err error
	var connection string

	if c.DBPass == "" {
		connection = fmt.Sprintf("dbname=%s", c.DBName)
	} else {
		connection = fmt.Sprintf("user=%s password=%s dbname=%s", c.DBUser, c.DBPass, c.DBName)
	}

	a.DB, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	defer a.DB.Close()

	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	loggingLevel, err := strconv.Atoi("1")
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter()

	memcacheClient := memcache.New(c.MemecachePort)
	a.Session = gsm.NewMemcacheStore(memcacheClient, c.MemecacheName, []byte(c.SessionSecret))
	a.Session.Logging = loggingLevel

	a.Router = cors.Default().Handler(router)
}

// Run - launch http server to listen on address
func (a *Application) Run(address string) {
	log.Fatal(http.ListenAndServe(":"+address, a.Router))
}
