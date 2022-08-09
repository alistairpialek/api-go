package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type routeResponse struct {
	endpoint      string
	statusCode    int
	duration      time.Duration
	clientMessage string
}

type routeError struct {
	Message string `json:"error"`
}

// accessLog prints a standard access log format.
// TODO: send this to a particular access log file or stdout.
func accessLog(r routeResponse) {
	fmt.Printf("access: %s, %d, %s\n", r.endpoint, r.statusCode, r.duration)
}

// errorLog prints a standard error log format.
// TODO: send this to particular error log file or stderr.
func errorLog(r routeResponse, err error) {
	fmt.Printf("error: %s, %d, %s, %s\n", r.endpoint, r.statusCode, r.duration, err)
}

// responseError ensures that we log errors per a standard format and give users a helpful
// error message and status code so that they figure out what they did wrong.
func responseError(w http.ResponseWriter, r routeResponse, err error) {
	errorLog(r, err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(r.statusCode)
	fmt.Fprintln(w, r.clientMessage)
}

// responseMessage ensures that we log access requests per a standard format.
func responseMessage(w http.ResponseWriter, r routeResponse) {
	accessLog(r)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(r.statusCode)
	fmt.Fprint(w, r.clientMessage)
}

// jsonError formats an json error message.
func jsonError(format string, v ...any) string {
	body, _ := json.Marshal(routeError{
		Message: fmt.Sprintf(format, v...),
	})

	return string(body)
}

// // jsonMessage formats an json message.
// func jsonMessage(format string, v ...any) (string, error) {
// 	body, err := json.Marshal(fmt.Sprintf(format, v...))

// 	if err != nil {
// 		return "error formatting message", err
// 	}

// 	return string(body), nil
// }
