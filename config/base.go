package config

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
