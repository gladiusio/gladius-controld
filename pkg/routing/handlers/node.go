package handlers

import (
	"encoding/json"
	"github.com/gladiusio/gladius-controld/pkg/p2p/message"
	"github.com/gladiusio/gladius-controld/pkg/p2p/signature"
	"github.com/gladiusio/gladius-controld/pkg/utils"
	"net/http"

	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gorilla/mux"
)

func PoolResponseForAddress(poolAddress string, ga *blockchain.GladiusAccountManager) (blockchain.PoolResponse, error) {
	poolUrl, err := blockchain.PoolRetrieveApplicationServerUrl(poolAddress, ga)
	poolResponse := blockchain.PoolResponse{Address: poolAddress, Url: poolUrl}
	if err != nil {
		return blockchain.PoolResponse{}, err
	}

	return poolResponse, nil
}

// New Routes
func NodeNewApplicationHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		vars := mux.Vars(r)
		poolAddress := vars["poolAddress"]

		poolResponse, err := PoolResponseForAddress(poolAddress, ga)
		if err != nil {
			ErrorHandler(w, r, "Pool data could not be found for Pool: "+poolAddress, err, http.StatusNotFound)
			return
		}

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var requestPayload models.NodeRequestPayload
		err = decoder.Decode(&requestPayload)

		// IP Address is detected from the server
		requestPayload.IPAddress = ""

		address, err := ga.GetAccountAddress()
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve account wallet address", err, http.StatusForbidden)
			return
		}

		requestPayload.Wallet = address.String()
		payload, err := json.Marshal(requestPayload)
		if err != nil {
			ErrorHandler(w, r, "Could not create payload string", err, http.StatusInternalServerError)
			return
		}

		unsignedMessage := message.New(payload)
		signedMessage, err := signature.CreateSignedMessage(unsignedMessage, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not create signed message, account could be locked", err, http.StatusForbidden)
			return
		}

		application, err := utils.SendRequest(http.MethodPost, poolResponse.Url+"applications/new", signedMessage)
		//application, err := sendRequest(http.MethodPost, "http://localhost:3333/api/applications/new", signedMessage)
		if err != nil {
			ErrorHandler(w, r, "Could not submit application to " + poolResponse.Address, err, http.StatusBadGateway)
			return
		}

		var defaultResponse response.DefaultResponse
		json.Unmarshal([]byte(application), &defaultResponse)
		ResponseHandler(w, r, defaultResponse.Message, defaultResponse.Success, &defaultResponse.Error, defaultResponse.Response, nil)
	}
}

func NodeViewApplicationHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		vars := mux.Vars(r)
		poolAddress := vars["poolAddress"]

		poolResponse, err := PoolResponseForAddress(poolAddress, ga)
		//_, err := PoolResponseForAddress(poolAddress, ga)
		if err != nil {
			ErrorHandler(w, r, "Pool data could not be found for Pool: "+poolAddress, err, http.StatusBadRequest)
			return
		}

		unsignedMessage := message.NewBlankMessage()
		signedMessage, err := signature.CreateSignedMessage(unsignedMessage, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not create signed message, account could be locked", err, http.StatusForbidden)
			return
		}

		applicationResponse, err := utils.SendRequest(http.MethodPost, poolResponse.Url+"applications/view", signedMessage)
		//applicationResponse, err := sendRequest(http.MethodPost, "http://localhost:3333/api/applications/view", signedMessage)
		if err != nil {
			ErrorHandler(w, r, "Could not view application", err, http.StatusForbidden)
			return
		}

		var defaultResponse response.DefaultResponse
		json.Unmarshal([]byte(applicationResponse), &defaultResponse)
		ResponseHandler(w, r, defaultResponse.Message, defaultResponse.Success, &defaultResponse.Error, defaultResponse.Response, nil)
	}
}

func NodeViewAllApplicationsHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		poolArrayResponse, err := blockchain.MarketPools(true, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusServiceUnavailable)
			return
		}

		unsignedMessage := message.NewBlankMessage()
		signedMessage, err := signature.CreateSignedMessage(unsignedMessage, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not create signed message, account could be locked", err, http.StatusForbidden)
			return
		}

		var responses = make([]interface{}, 0)

		for _, poolResponse := range poolArrayResponse.Pools {
			//poolResponse.Data.URL
			if poolResponse.Url != "" {
				applicationResponse, err := utils.SendRequest(http.MethodPost, poolResponse.Url+"applications/view", signedMessage)

				if err == nil {
					var responseStruct response.DefaultResponse
					json.Unmarshal([]byte(applicationResponse), &responseStruct)
					if responseStruct.Success {
						responses = append(responses, responseStruct.Response)
					}
				}
			}
		}

		ResponseHandler(w, r, "null", true, nil, responses, nil)
	}
}