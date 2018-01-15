package config

import (
	"os"
)

// DevConfig - app config for development env
var DevConfig = ApplicationConfiguration{
	devPort,
	devDBName,
	devDBUser,
	devDBPass,
	devMemecachePort,
	devMemecacheName,
	devSessionSecret,
	devSessionLoggingLevel,
}

var devPort = os.Getenv("DEV_PORT")

var devDBPass = os.Getenv("DEV_DB_PASS")

var devDBName = os.Getenv("DEV_DB_NAME")

var devDBUser = os.Getenv("DEV_DB_USER")

var devMemecachePort = os.Getenv("DEV_MEMECACHE_PORT")

var devMemecacheName = os.Getenv("DEV_MEMECACHE_NAME")

var devSessionSecret = os.Getenv("DEV_SESSION_SECRET")

var devSessionLoggingLevel = os.Getenv("DEV_SESSION_LOGGING_LEVEL")
