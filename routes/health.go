package routes

import (
	"net/http"
	"time"
)

const (
	HealthEndpoint string = "/health"
)

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
