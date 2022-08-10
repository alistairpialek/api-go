package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alistairpialek/api-go/v1/utils"
)

const (
	// MetadataEndpoint route.
	MetadataEndpoint string = "/metadata"
)

type metadataResponse struct {
	GitCommit   string `json:"git_commit"`
	Invocations int    `json:"invocations"`
}

// GetMetadata returns information about the API.
func GetMetadata(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	response, err := json.Marshal(metadataResponse{
		GitCommit:   utils.GitCommit,
		Invocations: utils.Invocations,
	})

	timeElapsed := time.Since(timeStart)

	if err != nil {
		responseError(w, routeResponse{
			endpoint:      MetadataEndpoint,
			statusCode:    http.StatusInternalServerError,
			duration:      timeElapsed,
			clientMessage: jsonError("metadata currently unavailable"),
		}, err)
		return
	}

	responseMessage(w, routeResponse{
		endpoint:      MetadataEndpoint,
		statusCode:    http.StatusOK,
		duration:      timeElapsed,
		clientMessage: string(response),
	})
}
