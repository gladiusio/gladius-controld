package handlers

import (
	"net/http"
)

// StatusHandler Main Status API route handler
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Status Page\n"))
}

// StatusTxHandler - Checks Status of txHash
func StatusTxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Status txHash Page\n"))
}
