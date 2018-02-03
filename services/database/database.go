package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Database - persistant RDBMS
type Database struct {
	Name  string
	Store *sql.DB
}

// NewDatabase - creates new persistant storage
func NewDatabase() *Database {
	d := &Database{}

	var err error
	var connection string

	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	if pass == "" {
		connection = fmt.Sprintf("dbname=%s", name)
	} else {
		connection = fmt.Sprintf("user=%s password=%s dbname=%s", user, pass, name)
	}

	d.Store, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println(err)
	}

	err = d.Store.Ping()
	if err != nil {
		fmt.Println(err)
	}

	d.Name = name

	return d
}
