package routes

import (
	"net/http"

	"github.com/alistairpialek/api-go/v1/log"
	"github.com/alistairpialek/api-go/v1/util"
)

const (
	// HealthEndpoint route.
	HealthEndpoint string = "/health"
)

// GetHealth returns a 200 http status code.
func GetHealth(w http.ResponseWriter, r *http.Request) {
	log.ResponseMessage(w, log.RouteResponse{
		Endpoint:      HealthEndpoint,
		StatusCode:    http.StatusOK,
		Duration:      util.TimeSinceRequestStart(),
		ClientMessage: "",
	})
}
