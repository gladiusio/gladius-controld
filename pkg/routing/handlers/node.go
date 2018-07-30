package handlers

import (
	"encoding/json"
	"net/http"

	"bytes"
	"io/ioutil"
	"time"

	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gorilla/mux"
)

func PoolResponseForAddress(poolAddress string, ga *blockchain.GladiusAccountManager) (blockchain.PoolResponse, error) {
	poolData, err := blockchain.PoolRetrievePublicData(poolAddress, ga)
	poolResponse := blockchain.PoolResponse{Address: poolAddress, Data: poolData}
	if err != nil {
		return blockchain.PoolResponse{}, err
	}

	return poolResponse, nil
}

// New Routes
func NodeNewApplicationHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		poolAddress := vars["poolAddress"]

		poolResponse, err := PoolResponseForAddress(poolAddress, ga)
		if err != nil {
			ErrorHandler(w, r, "Pool data could not be found for Pool: "+poolAddress, err, http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var requestPayload models.NodeRequestPayload
		err = decoder.Decode(&requestPayload)

		// IP Address is detected from the server
		requestPayload.IPAddress = ""

		address, err := ga.GetAccountAddress()
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve account wallet address", err, http.StatusBadRequest)
			return
		}

		requestPayload.Wallet = address.String()

		application, err := sendRequest(http.MethodPost, poolResponse.Data.URL+"applications/new", requestPayload)

		var defaultResponse response.DefaultResponse
		json.Unmarshal([]byte(application), &defaultResponse)
		ResponseHandler(w, r, "null", true, nil, defaultResponse.Response, nil)
	}
}

func NodeViewApplicationHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		poolAddress := vars["poolAddress"]

		poolResponse, err := PoolResponseForAddress(poolAddress, ga)
		if err != nil {
			ErrorHandler(w, r, "Pool data could not be found for Pool: "+poolAddress, err, http.StatusBadRequest)
			return
		}

		address, err := ga.GetAccountAddress()
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve account wallet address", err, http.StatusBadRequest)
			return
		}

		applicationResponse, err := sendRequest(http.MethodGet, poolResponse.Data.URL+"applications/view/"+address.String(), nil)
		var defaultResponse response.DefaultResponse
		json.Unmarshal([]byte(applicationResponse), &defaultResponse)
		ResponseHandler(w, r, "null", true, nil, defaultResponse.Response, nil)
	}
}

func NodeViewAllApplicationsHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		poolArrayResponse, err := blockchain.MarketPools(true, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusBadRequest)
			return
		}

		address, err := blockchain.NewGladiusAccountManager().GetAccountAddress()
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve account wallet address", err, http.StatusBadRequest)
			return
		}

		var responses []interface{}

		for _, poolResponse := range poolArrayResponse.Pools {
			//poolResponse.Data.URL
			if poolResponse.Data.URL != "" {
				applicationResponse, err := sendRequest(http.MethodGet, poolResponse.Data.URL+"applications/view/"+address.String(), nil)

				if err == nil {
					var responseStruct response.DefaultResponse
					json.Unmarshal([]byte(applicationResponse), &responseStruct)
					responses = append(responses, responseStruct.Response)
				}
			}
		}

		ResponseHandler(w, r, "null", true, nil, responses, nil)
	}
}

// For control over HTTP client headers,
// redirect policy, and other settings,
// create an HTTP client
var client = &http.Client{
	Timeout: time.Second * 10, //10 second timeout
}

// SendRequest - custom function to make sending api requests less of a pain
// in the arse.
func sendRequest(requestType, url string, data interface{}) (string, error) {

	b := bytes.Buffer{}

	// if data present, turn it into a bytesBuffer(jsonPayload)
	if data != nil {
		jsonPayload, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		b = *bytes.NewBuffer(jsonPayload)
	}

	// Build the request
	req, err := http.NewRequest(requestType, url, &b)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "gladius-controld")
	req.Header.Set("Content-Type", "application/json")

	// Send the request via a client
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	// read the body of the response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", err
	}

	// Defer the closing of the body
	defer res.Body.Close()

	return string(body), nil //tx
}
