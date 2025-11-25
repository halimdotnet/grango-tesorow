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

// HTTP server and router
const (
	ReadTimeout         = 15 * time.Second
	WriteTimeout        = 15 * time.Second
	IdleTimeout         = 60 * time.Second
	ShutdownTimeout     = 30 * time.Second
	MaxHeaderBytes      = 1 << 20
	MaxRequestBodyBytes = 1 << 20
	MaxUploadSize       = 3 << 20

	RateLimitAttempt = 5
	RatelimitPeriod  = 60 * time.Minute
)
