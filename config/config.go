package config

import "github.com/prstephens/go-webapi/helpers"

// AppConfig is the Application configuration struct
type AppConfig struct {
	Name        string
	Version     string
	Environment string
	LogFile     string
}

// HTTPConfig is the Application HTTP configuration
type HTTPConfig struct {
	Content string
	Problem string
}

// Config is the Configuration struct
type Config struct {
	App  AppConfig
	HTTP HTTPConfig
}

// New returns a new Config Struct
func NewApplicationConfig() *Config {
	return &Config{
		App: AppConfig{
			Name:        helpers.Env("APP_NAME", "Go Web API"),
			Version:     helpers.Env("APP_VERSION", "v1.0"),
			Environment: helpers.Env("ENV", "dev"),
			LogFile:     helpers.Env("LOG_FILE", "log.txt"),
		},
		HTTP: HTTPConfig{
			Content: helpers.Env("HTTP_CONTENT_TYPE", "application/json"),
			Problem: helpers.Env("HTTP_PROBLEM", "application/problem+json"),
		},
	}
}
