package handlers

import (
	"net/http"
)

// APIHandler - Main API route handler
func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main API\n"))
}
