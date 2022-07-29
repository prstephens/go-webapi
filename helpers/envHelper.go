package helpers

import (
	"os"
)

// env is a simple helper function to read an environment variable or return a default value
func Env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
