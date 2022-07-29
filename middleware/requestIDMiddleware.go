package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// ContextKey is used for context.Context value. The value requires a key that is not primitive type.
type ContextKey string

// ContextKeyRequestID is the ContextKey for RequestID
const ContextKeyRequestID ContextKey = "requestID"

func RequestIDMiddleware(next http.Handler) http.Handler {
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
