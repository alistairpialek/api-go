package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Setup(t *testing.T) (statusCode int, respMetrics calculateResponse) {
	fileContents, err := os.Open("../tests/metrics-input.json")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", CalculateEndpoint, fileContents)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCalculate)

	handler.ServeHTTP(rr, req)

	respBody, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(respBody, &respMetrics)
	if err != nil {
		t.Fatal(err)
	}

	return rr.Code, respMetrics
}

func TestPostCalculationMeta(t *testing.T) {
	respStatusCode, respMetrics := Setup(t)

	// Check the status code is what we expect.
	if status := respStatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that the number of suggestions is what we expect.
	respMetricsLength := len(respMetrics.Services)
	if respMetricsLength != 3 {
		t.Errorf("handler returned unexpected number of suggestions: got %v want %v", respMetricsLength, 3)
	}
}

func TestPostCalculationRequests(t *testing.T) {
	_, respMetrics := Setup(t)

	// Check that requests are being calculated correctly.
	respRequestValue := respMetrics.Services[0].CPU.Request
	if respRequestValue != 1.50 {
		t.Errorf("handler returned unexpected cpu suggestion: got %v want %v", respRequestValue, 1.50)
	}

	respRequestValue = respMetrics.Services[0].Memory.Request
	if respRequestValue != 1.50 {
		t.Errorf("handler returned unexpected cpu suggestion: got %v want %v", respRequestValue, 1.50)
	}
}

func TestPostCalculationLimits(t *testing.T) {
	_, respMetrics := Setup(t)

	// Check that limits are being calculated correctly.
	respLimitValue := respMetrics.Services[0].CPU.Limit
	if respLimitValue != 3 {
		t.Errorf("handler returned unexpected cpu suggestion: got %v want %v", respLimitValue, 3)
	}

	respLimitValue = respMetrics.Services[0].Memory.Limit
	if respLimitValue != 3 {
		t.Errorf("handler returned unexpected cpu suggestion: got %v want %v", respLimitValue, 3)
	}
}
