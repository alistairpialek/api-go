package util

import (
	"net/http"
	"time"
)

// RequestStartTime is started at the beginning of each request.
var RequestStartTime time.Time

// Invocations is the number of times the API has been called.
var Invocations int = 0

// MetricsMiddleware runs before HTTP handlers to .
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start the request start timer.
		RequestStartTime = time.Now()

		// Increment the number of invocations for each API call.
		Invocations++

		next.ServeHTTP(w, r)
	})
}

// TimeSinceRequestStart is a simple wrapper around time.Since() the request start.
func TimeSinceRequestStart() time.Duration {
	return time.Since(RequestStartTime)
}
