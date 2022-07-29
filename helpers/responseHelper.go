package helpers

import (
	"encoding/json"
	"net/http"
	"os"
)

func SendResponse(responseWriter http.ResponseWriter, code int, payload interface{}, contentType string) {
	response, _ := json.Marshal(payload)

	responseWriter.Header().Set("Content-Type", contentType)
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}

// env is a simple helper function to read an environment variable or return a default value
func Env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
