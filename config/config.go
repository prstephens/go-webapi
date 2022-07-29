package config

import "os"

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
			Name:        env("APP_NAME", "Go Web API"),
			Version:     env("APP_VERSION", "v1.0"),
			Environment: env("ENV", "dev"),
			LogFile:     env("LOG_FILE", "log.txt"),
		},
		HTTP: HTTPConfig{
			Content: env("HTTP_CONTENT_TYPE", "application/json"),
			Problem: env("HTTP_PROBLEM", "application/problem+json"),
		},
	}
}

// env is a simple helper function to read an environment variable or return a default value
func env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
