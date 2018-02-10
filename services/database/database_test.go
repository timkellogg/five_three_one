package database

import (
	"os"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	os.Setenv("DB_NAME", "five_three_one_test")
	os.Setenv("DATABASE_URL", "postgres://tkellogg:password@localhost:5432/five_three_one_development?sslmode=disable")

	database := NewDatabase()

	name := database.Name
	if name != "five_three_one_test" {
		t.Errorf("Database failed to set name to five_three_one. It was %s instead", name)
	}

	err := database.Store.Ping()
	if err != nil {
		t.Errorf("Database failed to connect: %s", err)
	}
}
