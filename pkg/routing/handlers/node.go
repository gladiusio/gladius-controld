package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gorilla/mux"
)

func NodeRetrieveDataHandler(w http.ResponseWriter, r *http.Request) {
	nodeAddress, err := blockchain.NodeOwnedByUser()
	if err != nil {
		ErrorHandler(w, r, "Node not found for user", err, http.StatusNotFound)
	}
	nodeData, err := blockchain.NodeRetrieveDataForAddress(*nodeAddress)
	if err != nil {
		ErrorHandler(w, r, "Node data could not be retrieved or data is not set", err, http.StatusNotFound)
	}

	nodeResponse := blockchain.NodeResponse{Address:nodeAddress.String(), Data:nodeData}

	ResponseHandler(w, r, "null", true, nil, nodeResponse, nil)
}

func NodeSetDataHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	decoder := json.NewDecoder(r.Body)
	var data blockchain.NodeData
	err := decoder.Decode(&data)

	if err != nil {
		ErrorHandler(w, r, "Passphrase `passphrase` not included or invalid in request", err, http.StatusBadRequest)
	}

	transaction, err := blockchain.NodeSetData(auth, &data)
	if err != nil {
		ErrorHandler(w, r, "Node data could not be set", err, http.StatusBadRequest)
		return
	}

	ResponseHandler(w, r, "null", true, nil, nil, transaction)
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

	ResponseHandler(w, r, "null", true, nil, nil, transaction)
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

	var response = "{ \"code\": " + status.String() + ", \"status\": "

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
	//TODO
	response += ",\"availableStatuses\": [{\"status\": \"Not Available\",\"code\": 0},{\"status\": \"Approved\",\"code\": 1},{\"status\": \"Rejected\",\"code\": 2},{\"status\": \"Pending\",\"code\": 3}]"

	response += "}"

	ResponseHandler(w, r, "null", true, nil, nil, nil)
}
