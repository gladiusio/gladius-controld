package handlers

import (
	"fmt"
	"net/http"
)

// ErrorHandler - Default Error Handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, m string, e error, statusCode int) {
	w.WriteHeader(statusCode)

	response := fmt.Sprintf("{ \"message\": \"%s\", \"success\": false, \"error\": \"%s\", \"response\": null, \"endpoint\": \"%s\" }", m, e, r.URL)
	w.Write([]byte(response))

	return
}
