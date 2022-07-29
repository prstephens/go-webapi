package hello

import (
	"net/http"

	"github.com/gorilla/mux"
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

// Response is the Hello Response
type Response struct {
	Message string `json:"message"`
}

// Handle will handle the incoming request
func (handler *Handler) Handle(response http.ResponseWriter, request *http.Request) {

	handler.app.Logger.Info("Hello Handler Dispatched.")

	vars := mux.Vars(request)

	if vars["name"] == "" {
		helpers.SendResponse(
			response,
			http.StatusBadRequest,
			Response{
				Message: "name not given",
			},
			handler.app.Config.HTTP.Problem,
		)

		return
	}

	helpers.SendResponse(
		response,
		http.StatusOK,
		Response{
			Message: "Hello " + vars["name"],
		},
		handler.app.Config.HTTP.Content,
	)

	return
}
