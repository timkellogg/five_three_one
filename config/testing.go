package config

import "os"

// TestConfig - app config for testing env
var TestConfig = ApplicationConfiguration{
	testPort,
	testDBName,
	testDBUser,
	testDBPass,
	testMemecachePort,
	testMemecacheName,
	testSessionSecret,
	testSessionLoggingLevel,
}

var testPort = os.Getenv("TEST_PORT")

var testDBPass = os.Getenv("TEST_DB_PASS")

var testDBName = os.Getenv("TEST_DB_NAME")

var testDBUser = os.Getenv("TEST_DB_USER")

var testMemecachePort = os.Getenv("TEST_MEMECACHE_PORT")

var testMemecacheName = os.Getenv("TEST_MEMECACHE_NAME")

var testSessionSecret = os.Getenv("TEST_SESSION_SECRET")

var testSessionLoggingLevel = os.Getenv("TEST_SESSION_LOGGING_LEVEL")
