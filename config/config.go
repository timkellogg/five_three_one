package config

import (
	"database/sql"
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/timkellogg/five_three_one/services/authentication"

	"github.com/bradleypeabody/gorilla-sessions-memcache"
)

// ApplicationContext - resources needed by application
type ApplicationContext struct {
	Database *sql.DB
	Session  *gsm.MemcacheStore
	Auth     authentication.AuthService
}

// LoadEnvironment - loads appropriate env whether config is develoment, production, etc.
func LoadEnvironment() {
	var environmentFile string
	environment := flag.String("environment", "development", "Indicates the application environment")

	environmentFile = ".env." + *environment
	err := godotenv.Load(environmentFile)
	if err != nil {
		log.Fatal(err)
	}
}

// PerformEnvChecks - makes sure the application is up
func PerformEnvChecks(c ApplicationContext) {
	err := c.Database.Ping()
	if err != nil {
		log.Fatalf("Database environment check failed: %s", err)
	}
}
