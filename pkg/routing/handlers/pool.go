package handlers

import (
		"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
		"github.com/gorilla/mux"
)

func PoolPublicDataHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//poolAddress := vars["poolAddress"]
	//
	//if r.Method == http.MethodGet {
	//	poolData, err := blockchain.PoolRetrievePublicData(poolAddress)
	//
	//	if err != nil {
	//		ErrorHandler(w, r, "Could not retrieve Pool's public data", err, http.StatusNotFound)
	//		return
	//	}
	//
	//	ResponseHandler(w, r, "null", true, nil, poolData, nil)
	//}
	//
	//if r.Method == http.MethodPost {
	//	auth := r.Header.Get("X-Authorization")
	//	decoder := json.NewDecoder(r.Body)
	//	var data blockchain.PoolPublicData
	//	err := decoder.Decode(&data)
	//
	//	jsonPayload, err := json.Marshal(data)
	//	if err != nil {
	//		ErrorHandler(w, r, "Could not decode request into JSON", err, http.StatusNotFound)
	//		return
	//	}
	//
	//	transaction, err := blockchain.PoolSetPublicData(auth, poolAddress, string(jsonPayload))
	//	if err != nil {
	//		ErrorHandler(w, r, "Could not set Pool's public data", err, http.StatusUnprocessableEntity)
	//		return
	//	}
	//
	//	ResponseHandler(w, r, "Public data set, pending transaction", true, nil, nil, transaction)
	//}
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
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusUnprocessableEntity)
		return
	}

	ResponseHandler(w, r, "null", true, nil, applications, nil)
}