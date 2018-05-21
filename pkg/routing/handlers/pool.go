package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gorilla/mux"
)

// PoolHandler - Main Node API route handler
func PoolHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node API\n"))
}

func PoolRetrievePublicDataHandler(w http.ResponseWriter, r *http.Request) {

}

func PoolSetPublicDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]
	auth := r.Header.Get("X-Authorization")
	decoder := json.NewDecoder(r.Body)
	var data PoolPublicData
	err := decoder.Decode(&data)

	transaction, err := blockchain.PoolSetPublicData(auth, poolAddress, data.String())
	if err != nil {
		ErrorHandler(w, r, "Could not set Pool's public data", err, http.StatusUnprocessableEntity)
		return
	}

	TransactionHandler(w, r, "\"Public data set, pending transaction\"", transaction)
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

func PoolRetrieveNodes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	poolAddress := common.HexToAddress(vars["poolAddress"])
	nodeAddresses, _ := blockchain.Nodes(poolAddress.String())

	response := "["

	for _, nodeAddress := range *nodeAddresses {
		nodeApplication, err := blockchain.NodeRetrieveApplication(&nodeAddress, &poolAddress)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve application", err, http.StatusUnprocessableEntity)
			return
		}
		response += nodeApplication.String() + ","
	}
	strings.TrimRight(response, ",")
	response += "]"
	ResponseHandler(w, r, "null", response)
}

func PoolRetrieveNodeApplication(w http.ResponseWriter, r *http.Request) {
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
