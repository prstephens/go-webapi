package services

import (
	application "github.com/prstephens/go-webapi/app"
	"github.com/prstephens/go-webapi/services/hello"
	"github.com/prstephens/go-webapi/services/ping"
)

func RegisterServices(app *application.Application) {

	ping.BuildPingService(app)
	hello.BuildHelloService(app)

}
