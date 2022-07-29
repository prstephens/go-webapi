package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prstephens/go-webapi/config"
	"github.com/prstephens/go-webapi/logger"
	"go.uber.org/zap"
)

// Application is our general purpose Application struct
type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *config.Config
}

func Boot() *Application {

	appConfig := config.NewApplicationConfig()
	logger := logger.Initialize(appConfig)
	router := mux.NewRouter()

	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	return &Application{
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      corsHandler(requestIDMiddleware(router)), // use more handlers here to chain common things
			IdleTimeout:  120 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
		},
		Router: router,
		Logger: logger,
		Config: appConfig,
	}
}

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID ContextKey = "requestID"

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		id := uuid.New()

		ctx = context.WithValue(ctx, ContextKeyRequestID, id.String())

		r = r.WithContext(ctx)

		fmt.Printf("Incomming request %s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, id.String())

		next.ServeHTTP(w, r)

		fmt.Printf("Finished handling http req. %s", id.String())
	})
}

// Run will run the Application server
func (app *Application) Run() {
	err := app.Server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

// WaitForShutdown is a graceful way to handle server shutdown events
func WaitForShutdown(application *Application) {
	// Create a channel to listen for OS signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal to our channel
	<-interruptChan

	application.Logger.Info("Received shutdown signal, gracefully terminating")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	application.Server.Shutdown(ctx)
	os.Exit(0)
}
