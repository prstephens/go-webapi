package main

import (
	"github.com/joho/godotenv"
	application "github.com/prstephens/go-webapi/app"
	"github.com/prstephens/go-webapi/services"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	app := application.Boot()

	services.RegisterServices(app)

	go func() {
		app.Run()
	}()

	app.Logger.Info("Web server started on port: " + app.Server.Addr)

	// Wait for termination signals and shut down gracefully
	application.WaitForShutdown(app)
}
