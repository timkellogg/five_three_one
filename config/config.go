package config

import (
	"database/sql"
	"flag"
	"fmt"
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

// PerformEnvChecks - make sure the application deps are running
func (c *ApplicationContext) PerformEnvChecks() {
	// checks if db is the correct one
	// var name string
	// c.Database.QueryRow("SELECT current_database()").Scan(&name)
	// fmt.Println("db name:" + name)

	err := c.Database.Ping()
	if err != nil {
		log.Fatalf("Database environment check failed: %s", err)
	}

	r, err := c.Database.Exec("SELECT * FROM users;")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tables: %v", r)
}

// TruncateDBTables - removes all data from tables
func (c *ApplicationContext) TruncateDBTables(tables []string) {
	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s;", table)
		c.Database.Exec(query)
	}
}
