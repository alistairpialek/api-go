package utils

import (
	"net/http"
)

// GitCommit is set at build via -ldflags.
var GitCommit string

// Invocations is the number of times the API has been called.
var Invocations int = 0

// MetricsMiddleware runs before HTTP handlers.
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Invocations++
		next.ServeHTTP(w, r)
	})
}
