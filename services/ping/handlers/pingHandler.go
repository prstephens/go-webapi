package ping

import (
	"net/http"

	application "github.com/prstephens/go-webapi/app"
	"github.com/prstephens/go-webapi/helpers"
)

// Handler is the http.Handler for this request
type Handler struct {
	app *application.Application
}

// NewHandler will create a new Handler to handle this request
func NewHandler(app *application.Application) *Handler {
	return &Handler{app}
}

// Response is the Ping Response
type Response struct {
	Message string `json:"message"`
}

// Handle will handle the incoming request
func (handler *Handler) Handle(response http.ResponseWriter, request *http.Request) {

	handler.app.Logger.Info("Ping Handler Dispatched.")

	helpers.SendResponse(
		response,
		http.StatusOK,
		Response{
			Message: "Service Online",
		},
		handler.app.Config.HTTP.Content,
	)

	return
}
