package routes

import (
	"encoding/json"
	"net/http"

	"github.com/alistairpialek/api-go/v1/log"
	"github.com/alistairpialek/api-go/v1/util"
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
	response, err := json.Marshal(metadataResponse{
		GitCommit:   util.GitCommit,
		Invocations: util.Invocations,
	})

	if err != nil {
		log.ResponseError(w, log.RouteResponse{
			Endpoint:      MetadataEndpoint,
			StatusCode:    http.StatusInternalServerError,
			Duration:      util.TimeSinceRequestStart(),
			ClientMessage: log.JSONError("metadata currently unavailable"),
		}, err)
		return
	}

	log.ResponseMessage(w, log.RouteResponse{
		Endpoint:      MetadataEndpoint,
		StatusCode:    http.StatusOK,
		Duration:      util.TimeSinceRequestStart(),
		ClientMessage: string(response),
	})
}
