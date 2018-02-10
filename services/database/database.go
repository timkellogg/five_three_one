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

	name := os.Getenv("DB_NAME")
	databaseURL := os.Getenv("DATABASE_URL")

	d.Store, err = sql.Open("postgres", databaseURL)
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
