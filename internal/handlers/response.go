package handlers

import (
	"fmt"
	"net/http"
)

// ResponseHandler - Default Response Handler
func ResponseHandler(w http.ResponseWriter, r *http.Request, m string, res string) {
	response := fmt.Sprintf("{ \"message\": %s, \"success\": true, \"error\": null, \"response\": %s, \"endpoint\": \"%s\" }", m, res, r.URL)
	w.Write([]byte(response))

	return
}
