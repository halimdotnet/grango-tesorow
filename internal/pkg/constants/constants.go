package constants

import "time"

// App environment config
const (
	AppEnv         = "APP_ENV"
	ConfigPath     = "./configs"
	EnvProduction  = "production"
	EnvStaging     = "staging"
	EnvDevelopment = "development"
)

// HTTP server
const (
	ReadTimeout     = 15 * time.Second
	WriteTimeout    = 15 * time.Second
	IdleTimeout     = 60 * time.Second
	ShutdownTimeout = 30 * time.Second
	MaxHeaderBytes  = 1 << 20
)
