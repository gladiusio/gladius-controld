package handlers

import (
	"fmt"
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gorilla/mux"
)

// PoolHandler - Main Node API route handler
func PoolHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node API\n"))
}

func PoolRetrievePublicKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]

	publicKey, err := blockchain.PoolRetrievePublicKey(poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve Pool's Public Key", err, http.StatusUnprocessableEntity)
		return
	}

	response := fmt.Sprintf("{\"publicKey\": \"%s\"}", publicKey)
	ResponseHandler(w, r, "null", response)
}
