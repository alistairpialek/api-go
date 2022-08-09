package routes

import (
	"net/http"
	"time"
)

const (
	// HealthEndpoint route.
	HealthEndpoint string = "/health"
)

// GetHealth returns a 200 http status code.
func GetHealth(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	timeElapsed := time.Since(timeStart)

	responseMessage(w, routeResponse{
		endpoint:      HealthEndpoint,
		statusCode:    http.StatusOK,
		duration:      timeElapsed,
		clientMessage: "",
	})
}
