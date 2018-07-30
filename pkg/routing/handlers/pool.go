package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gorilla/mux"
)

func PoolPublicDataHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		poolAddress := vars["poolAddress"]

		poolResponse, err := PoolResponseForAddress(poolAddress, ga)

		if err != nil {
			ErrorHandler(w, r, "Pool data could not be found for Pool: "+poolAddress, err, http.StatusBadRequest)
			return
		}

		poolInformationResponse, err := sendRequest(http.MethodGet, poolResponse.Data.URL+"server/info", nil)
		var defaultResponse response.DefaultResponse
		json.Unmarshal([]byte(poolInformationResponse), &defaultResponse)

		ResponseHandler(w, r, "null", true, nil, defaultResponse.Response, nil)
	}
}

func PoolRetrieveNodesHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		poolAddress := common.HexToAddress(vars["poolAddress"])
		nodeAddresses, _ := blockchain.PoolNodes(poolAddress.String(), ga)
		status := vars["status"]
		statusInt, err := blockchain.ApplicationStatusFromString(status)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusUnprocessableEntity)
			return
		}

		applications, err := blockchain.PoolNodesWithData(poolAddress, nodeAddresses, statusInt, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusUnprocessableEntity)
			return
		}

		ResponseHandler(w, r, "null", true, nil, applications, nil)
	}
}
