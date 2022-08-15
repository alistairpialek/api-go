package log

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type routeError struct {
	Message string `json:"error"`
}

// ResponseError ensures that we log errors per a standard format and give users a helpful
// error message and status code so that they figure out what they did wrong.
func ResponseError(w http.ResponseWriter, r RouteResponse, err error) {
	errorLog(r, err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(r.StatusCode)
	fmt.Fprintln(w, r.ClientMessage)
}

// ResponseMessage ensures that we log access requests per a standard format.
func ResponseMessage(w http.ResponseWriter, r RouteResponse) {
	accessLog(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(r.StatusCode)
	fmt.Fprint(w, r.ClientMessage)
}

// JSONError formats an json error message.
func JSONError(format string, v ...any) string {
	body, _ := json.Marshal(routeError{
		Message: fmt.Sprintf(format, v...),
	})

	return string(body)
}
