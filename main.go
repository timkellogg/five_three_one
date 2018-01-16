package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/timkellogg/five_three_one/config"
)

func main() {
	loadEnvironment()

	applicationConfig := config.ApplicationConfiguration{
		Port:                os.Getenv("PORT"),
		DBName:              os.Getenv("DB_NAME"),
		DBUser:              os.Getenv("DB_USER"),
		DBPass:              os.Getenv("DB_PASS"),
		MemecachePort:       os.Getenv("MEMECACHE_PORT"),
		MemecacheName:       os.Getenv("MEMECACHE_NAME"),
		SessionSecret:       os.Getenv("SESSION_SECRET"),
		SessionLoggingLevel: os.Getenv("SESSION_LOGGING_LEVEL"),
	}

	application := config.Application{}
	application.Initialize(applicationConfig)
	application.Run(applicationConfig.Port)
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
