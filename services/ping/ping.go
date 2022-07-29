package ping

import (
	"net/http"

	application "github.com/prstephens/go-webapi/app"
	ping "github.com/prstephens/go-webapi/services/ping/handlers"
)

type Mddleware func(http.Handler) http.Handler

// BuildPingService is responsible for setting up the Ping Context for our Domain
func BuildPingService(app *application.Application) {

	app.Logger.Info("Building ping service...")

	// Create our Handler
	handler := ping.NewHandler(app)

	// Create a sub router for this service
	router := app.Router.Methods(http.MethodGet).Subrouter()

	// Register our service routes
	router.HandleFunc("/ping", handler.Handle).Name("ping:handle")
}
