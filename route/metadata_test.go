package routes

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestGetMetadata(t *testing.T) {
	req, err := http.NewRequest("GET", MetadataEndpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMetadata)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `"git_commit":.+`
	if matched, _ := regexp.MatchString(expected, rr.Body.String()); !matched {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Check the response body is what we expect.
	expected = `"invocations":.+`
	if matched, _ := regexp.MatchString(expected, rr.Body.String()); !matched {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
