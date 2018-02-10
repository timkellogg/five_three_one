package models

import (
	"os"
	"testing"

	"github.com/timkellogg/five_three_one/config"
	"github.com/timkellogg/five_three_one/services/authentication"
	"github.com/timkellogg/five_three_one/services/database"
	"github.com/timkellogg/five_three_one/services/session"
)

var context config.ApplicationContext
var testUser = User{Email: "test@test.com"}

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", "five_three_one_test")
	os.Setenv("DATABASE_URL", "postgres://tkellogg:password@localhost:5432/five_three_one_test?sslmode=disable")

	context = config.ApplicationContext{
		Database: database.NewDatabase().Store,
		Session:  session.NewSession().Memcache,
		Auth:     authentication.AuthService{},
	}
	defer context.Database.Close()

	context.TruncateDBTables(tables())

	runTests := m.Run()
	os.Exit(runTests)
}

func tables() []string {
	return []string{"users"}
}
