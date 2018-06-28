package handlers

import (
	"encoding/json"
		"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gorilla/mux"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
)

func PoolPublicDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]

	if r.Method == http.MethodGet {
		poolData, err := blockchain.PoolRetrievePublicData(poolAddress)

		if err != nil {
			ErrorHandler(w, r, "Could not retrieve Pool's public data", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, poolData, nil)
	}

	if r.Method == http.MethodPost {
		auth := r.Header.Get("X-Authorization")
		decoder := json.NewDecoder(r.Body)
		var data blockchain.PoolPublicData
		err := decoder.Decode(&data)

		jsonPayload, err := json.Marshal(data)
		if err != nil {
			ErrorHandler(w, r, "Could not decode request into JSON", err, http.StatusNotFound)
		}

		transaction, err := blockchain.PoolSetPublicData(auth, poolAddress, string(jsonPayload))
		if err != nil {
			ErrorHandler(w, r, "Could not set Pool's public data", err, http.StatusUnprocessableEntity)
			return
		}

		ResponseHandler(w, r, "Public data set, pending transaction", true, nil, nil, transaction)
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

	publicKeyResponse := response.PublicKeyResponse{PublicKey:publicKey}

	ResponseHandler(w, r, "null", true, nil, publicKeyResponse, nil)
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

	applications, err := blockchain.PoolNodesWithData(poolAddress, nodeAddresses, statusInt)

	ResponseHandler(w, r, "null", true, nil, applications, nil)
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

	ResponseHandler(w, r, "null", true, nil, nodeApplication, nil)
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

	ResponseHandler(w, r, "Public data set, pending transaction", true, nil, nil, transaction)
}
