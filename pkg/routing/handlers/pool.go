package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gorilla/mux"
)

// PoolHandler - Main Node API route handler
func PoolHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node API\n"))
}

func PoolPublicDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]

	if r.Method == http.MethodGet {
		poolData, err := blockchain.PoolRetrievePublicData(poolAddress)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve Pool's public data", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", poolData.String())
	}

	if r.Method == http.MethodPost {
		auth := r.Header.Get("X-Authorization")
		decoder := json.NewDecoder(r.Body)
		var data blockchain.PoolPublicData
		err := decoder.Decode(&data)

		transaction, err := blockchain.PoolSetPublicData(auth, poolAddress, data.String())
		if err != nil {
			ErrorHandler(w, r, "Could not set Pool's public data", err, http.StatusUnprocessableEntity)
			return
		}

		TransactionHandler(w, r, "\"Public data set, pending transaction\"", transaction)
	}
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

func PoolRetrieveNodesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	poolAddress := common.HexToAddress(vars["poolAddress"])
	nodeAddresses, _ := blockchain.PoolNodes(poolAddress.String())
	status := vars["status"]
	statusInt, err := blockchain.ApplicationStatusFromString(status)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusUnprocessableEntity)
		return
	}

	response, err := blockchain.PoolNodesWithData(poolAddress, nodeAddresses, statusInt)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusUnprocessableEntity)
		return
	}

	ResponseHandler(w, r, "null", *response)
}

func PoolRetrieveNodeApplicationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	poolAddress := common.HexToAddress(vars["poolAddress"])
	nodeAddress := common.HexToAddress(vars["nodeAddress"])

	nodeApplication, err := blockchain.NodeRetrieveApplication(&nodeAddress, &poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve application", err, http.StatusUnprocessableEntity)
		return
	}
	ResponseHandler(w, r, "null", nodeApplication.String())
}

func PoolUpdateNodeStatusHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]
	nodeAddress := vars["nodeAddress"]
	status := vars["status"]
	var statusInt int
	if status == "approve" {
		statusInt = 1
	} else {
		statusInt = 2
	}

	transaction, err := blockchain.PoolUpdateNodeStatus(auth, poolAddress, nodeAddress, statusInt)
	if err != nil {
		ErrorHandler(w, r, "Could not set Pool's public data", err, http.StatusUnprocessableEntity)
		return
	}

	TransactionHandler(w, r, "\"Public data set, pending transaction\"", transaction)
}
