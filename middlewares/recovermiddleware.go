package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// recoverMiddleware is a middleware that recovers from panics, logs the error, and writes an internal server error response.
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recovered from panic:", err)
				debug.PrintStack()
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Internal Server Error: %v", err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
