package handlers

import (
	"encoding/json"
	"net/http"

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

func PoolSetBlockchainDataHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		poolAddress := vars["poolAddress"]

		auth := r.Header.Get("X-Authorization")
		decoder := json.NewDecoder(r.Body)
		var data blockchain.PoolPublicData
		err := decoder.Decode(&data)

		jsonPayload, err := json.Marshal(data)
		if err != nil {
			ErrorHandler(w, r, "Could not decode request into JSON", err, http.StatusNotFound)
			return
		}

		transaction, err := blockchain.PoolSetPublicData(auth, poolAddress, string(jsonPayload))
		if err != nil {
			ErrorHandler(w, r, "Could not set Pool's public data", err, http.StatusUnprocessableEntity)
			return
		}

		ResponseHandler(w, r, "Public data set, pending transaction", true, nil, nil, transaction)
	}
}