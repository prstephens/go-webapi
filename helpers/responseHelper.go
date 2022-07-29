package helpers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(responseWriter http.ResponseWriter, code int, payload interface{}, contentType string) {
	response, _ := json.Marshal(payload)

	responseWriter.Header().Set("Content-Type", contentType)
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}
