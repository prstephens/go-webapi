package hello

import (
	"net/http"

	application "github.com/prstephens/go-webapi/app"
	hello "github.com/prstephens/go-webapi/services/hello/handlers"
)

// BuildPingService is responsible for setting up the Ping Context for our Domain
func BuildHelloService(app *application.Application) {

	app.Logger.Info("Building hello service...")

	// Create our Handler
	handler := hello.NewHandler(app)

	// Create a sub router for this service
	router := app.Router.Methods(http.MethodGet).Subrouter()

	// Register our service routes
	router.HandleFunc("/hello", handler.Handle).Name("hello:handle")
	router.HandleFunc("/hello/{name}", handler.Handle).Name("hello:handle")

}
