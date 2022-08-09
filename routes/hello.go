package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// HelloEndpoint route.
	HelloEndpoint string = "/hello"
)

type helloResponse struct {
	Message string `json:"message"`
}

// GetHello returns a greeting.
func GetHello(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	response, err := json.Marshal(helloResponse{
		Message: fmt.Sprintf("Hello world, the time is currently %s", time.Now().UTC()),
	})

	timeElapsed := time.Since(timeStart)

	if err != nil {
		responseError(w, routeResponse{
			endpoint:      HelloEndpoint,
			statusCode:    http.StatusInternalServerError,
			duration:      timeElapsed,
			clientMessage: jsonError("time currently unavailable"),
		}, err)
		return
	}

	responseMessage(w, routeResponse{
		endpoint:      HelloEndpoint,
		statusCode:    http.StatusOK,
		duration:      timeElapsed,
		clientMessage: string(response),
	})
}
