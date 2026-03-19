// Package middleware provides HTTP middleware.
package middleware

import (
	"log"
	"net/http"
	"time"
)

// Log is a middleware that logs each incoming HTTP request.
// It prints the HTTP method, request URI, and the current timestamp before calling the next handler.
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Now().Format("2006-01-02 15:04:05"))
		next.ServeHTTP(w, r)
	})
}
