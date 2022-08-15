package log

import (
	"log"
	"time"
)

// RouteResponse is a logging object that ensures all fields we want to log are required
// and can be presented in a particular way.
type RouteResponse struct {
	Endpoint      string
	StatusCode    int
	Duration      time.Duration
	ClientMessage string
}

// accessLog prints a standard access log format.
// TODO: send this to a particular access log file or stdout.
func accessLog(r RouteResponse) {
	log.Printf("access: %s, %d, %s\n", r.Endpoint, r.StatusCode, r.Duration)
}

// errorLog prints a standard error log format.
// TODO: send this to particular error log file or stderr.
func errorLog(r RouteResponse, err error) {
	log.Printf("error: %s, %d, %s, %s\n", r.Endpoint, r.StatusCode, r.Duration, err)
}
