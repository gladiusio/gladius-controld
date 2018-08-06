package handlers

import (
	"encoding/json"
	"github.com/gladiusio/gladius-application-server/pkg/controller"
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
		defer r.Body.Close()

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

func PoolRetrievePendingPoolConfirmationApplicationsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := controller.Initialize(nil)
		if err != nil {
			ErrorHandler(w, r, "Could not establish database connection", err, http.StatusInternalServerError)
			return
		}

		defer db.Close()

		profiles, err := controller.NodesPendingPoolConfirmation(db)

		if err != nil {
			ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, profiles, nil)
	}
}

func PoolRetrievePendingNodeConfirmationApplicationsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := controller.Initialize(nil)
		if err != nil {
			ErrorHandler(w, r, "Could not establish database connection", err, http.StatusInternalServerError)
			return
		}

		defer db.Close()

		profiles, err := controller.NodesPendingNodeConfirmation(db)

		if err != nil {
			ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, profiles, nil)
	}
}

func PoolRetrieveApprovedApplicationsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := controller.Initialize(nil)
		if err != nil {
			ErrorHandler(w, r, "Could not establish database connection", err, http.StatusInternalServerError)
			return
		}

		defer db.Close()

		profiles, err := controller.NodesAccepted(db)

		if err != nil {
			ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, profiles, nil)
	}
}

func PoolRetrieveRejectedApplicationsHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := controller.Initialize(nil)
		if err != nil {
			ErrorHandler(w, r, "Could not establish database connection", err, http.StatusInternalServerError)
			return
		}
		
		defer db.Close()

		profiles, err := controller.NodesRejected(db)

		if err != nil {
			ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, profiles, nil)
	}
}
