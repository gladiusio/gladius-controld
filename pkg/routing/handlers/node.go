package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gorilla/mux"
)

// NodeHandler - Main Node API route handler
func NodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node API\n"))
}

func NodeRetrieveDataHandler(w http.ResponseWriter, r *http.Request) {
	nodeData, err := blockchain.NodeRetrieveData()
	if err != nil {
		ErrorHandler(w, r, "Node data could not be retrieved or data is not set", err, http.StatusNotFound)
	}

	jsonResponse := nodeData.String()

	ResponseHandler(w, r, "null", jsonResponse)
}

func NodeSetDataHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	decoder := json.NewDecoder(r.Body)
	var data blockchain.NodeData
	err := decoder.Decode(&data)

	if err != nil {
		ErrorHandler(w, r, "Passphrase `passphrase` not included or invalid in request", err, http.StatusBadRequest)
	}

	transaction, _ := blockchain.NodeSetData(auth, &data)
	TransactionHandler(w, r, "null", transaction)
}

func NodeApplyToPoolHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	nodeAddress := vars["nodeAddress"]
	poolAddress := vars["poolAddress"]

	auth := r.Header.Get("X-Authorization")
	transaction, err := blockchain.NodeApplyToPool(auth, nodeAddress, poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Could not apply to pool", err, http.StatusBadRequest)
		return
	}

	println(transaction)

	TransactionHandler(w, r, "null", transaction)
}

func NodeApplicationStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	nodeAddress := vars["nodeAddress"]
	poolAddress := vars["poolAddress"]

	status, err := blockchain.NodeApplicationStatus(nodeAddress, poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Could not find status for pool application", err, http.StatusBadRequest)
		return
	}

	var response string = "{ \"code\": " + status.String() + ", \"status\": "

	switch status.String() {
	// Unavailable
	case "0":
		response += "\"Unavailable\""
	// Approved
	case "1":
		response += "\"Approved\""
	// Rejected
	case "2":
		response += "\"Rejected\""
	// Pending
	case "3":
		response += "\"Pending\""
	}
	response += ",\"availableStatuses\": [{\"status\": \"Not Available\",\"code\": 0},{\"status\": \"Approved\",\"code\": 1},{\"status\": \"Rejected\",\"code\": 2},{\"status\": \"Pending\",\"code\": 3}]"

	response += "}"

	ResponseHandler(w, r, "null", response)
}
