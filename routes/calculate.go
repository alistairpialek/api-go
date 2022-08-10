package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	// CalculateEndpoint route.
	CalculateEndpoint string = "/calculate"
)

type calculateResponse struct {
	Services []service `json:"services"`
}

type service struct {
	Name   string           `json:"name"`
	CPU    suggestedMetrics `json:"cpu"`
	Memory suggestedMetrics `json:"memory"`
}

type suggestedMetrics struct {
	Request float32 `json:"request"`
	Limit   float32 `json:"limit"`
}

type inputMetrics []struct {
	App    string  `json:"app"`
	CPU    float32 `json:"cpu usage (mcores)"`
	Memory float32 `json:"memory usage (MiB)"`
}

// PostCalculate takes service CPU and memory readings and makes suggestions about its request
// and limit values.
func PostCalculate(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()

	// Read the request body.
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responseError(w, routeResponse{endpoint: CalculateEndpoint,
			statusCode:    http.StatusInternalServerError,
			duration:      time.Since(timeStart),
			clientMessage: jsonError("invalid payload"),
		}, err)
		return
	}

	// Load the request into our data structure.
	var reqMetrics inputMetrics
	err = json.Unmarshal(reqBody, &reqMetrics)
	if err != nil {
		responseError(w, routeResponse{endpoint: CalculateEndpoint,
			statusCode:    http.StatusInternalServerError,
			duration:      time.Since(timeStart),
			clientMessage: jsonError("invalid payload"),
		}, err)
		return
	}

	// To make performing statistics on app names simpler, create a collection that is keyed
	// by app names so that each collection (app) can be worked with in isolation.
	metricsKeyed := make(map[string]inputMetrics)
	for _, metric := range reqMetrics {
		metricsKeyed[metric.App] = append(metricsKeyed[metric.App], metric)
	}

	// Metrics in, suggestions out.
	appSuggestions := calculateResourceSuggestions(metricsKeyed)
	response, err := json.Marshal(calculateResponse{
		Services: appSuggestions,
	})
	if err != nil {
		responseError(w, routeResponse{
			endpoint:      CalculateEndpoint,
			statusCode:    http.StatusInternalServerError,
			duration:      time.Since(timeStart),
			clientMessage: jsonError("suggestions currently unavailable"),
		}, err)
		return
	}

	responseMessage(w, routeResponse{
		endpoint:      CalculateEndpoint,
		statusCode:    http.StatusOK,
		duration:      time.Since(timeStart),
		clientMessage: string(response),
	})
}

// calculateResourceSuggestions takes point in time metrics from an app and returns suggestions
// on what the resource request and limits should be set to.
func calculateResourceSuggestions(metricsKeyed map[string]inputMetrics) (services []service) {
	for appName, app := range metricsKeyed {
		// Get the CPU and memory request averages. An average provides a good idea of what resources the app is actually using.
		var cpuTotal, memoryTotal float32

		for _, v := range app {
			cpuTotal += v.CPU
			memoryTotal += v.Memory
		}

		// Calculate the requests.
		appCount := float32(len(app))
		cpuAverage := cpuTotal / appCount
		memoryAverage := memoryTotal / appCount

		// Calculate the limits.
		additionalToleranceMult := float32(2.00)
		cpuMax := cpuAverage * additionalToleranceMult
		memoryMax := memoryAverage * additionalToleranceMult

		tempService := service{
			Name:   appName,
			CPU:    suggestedMetrics{Request: cpuAverage, Limit: cpuMax},
			Memory: suggestedMetrics{Request: memoryAverage, Limit: memoryMax},
		}

		services = append(services, tempService)
	}

	return services
}
