package handlers

import (
	"encoding/json"
	"net/http"

		"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gorilla/mux"
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"time"
	"bytes"
	"io/ioutil"
	)

func poolResponseForAddress(poolAddress string) (blockchain.PoolResponse, error) {
	poolData, err := blockchain.PoolRetrievePublicData(poolAddress)
	poolResponse := blockchain.PoolResponse{poolAddress, poolData}
	if err != nil {
		return blockchain.PoolResponse{}, err
	}

	return poolResponse, nil
}

// New Routes
func NodeNewApplicationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]

	poolResponse, err := poolResponseForAddress(poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Pool data could not be found for Pool: " + poolAddress, err, http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestPayload models.NodeRequestPayload
	err = decoder.Decode(&requestPayload)

	requestPayload.IPAddress = r.RemoteAddr

	address, err := blockchain.NewGladiusAccountManager().GetAccountAddress()
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve account wallet address", err, http.StatusBadRequest)
		return
	}

	requestPayload.Wallet = address.String()

	sendRequest(http.MethodPost, poolResponse.Data.URL + "application/new", requestPayload)
}

func NodeViewApplicationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]

	poolResponse, err := poolResponseForAddress(poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Pool data could not be found for Pool: " + poolAddress, err, http.StatusBadRequest)
		return
	}

	address, err := blockchain.NewGladiusAccountManager().GetAccountAddress()
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve account wallet address", err, http.StatusBadRequest)
		return
	}

	applicationResponse, err := sendRequest(http.MethodGet, poolResponse.Data.URL + "application/view/" + address.String(), nil)
	w.Write([]byte(applicationResponse))
}

func NodeViewAllApplicationsHandler(w http.ResponseWriter, r *http.Request) {
	poolArrayResponse, err := blockchain.MarketPools(true)
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
			applicationResponse, err := sendRequest(http.MethodGet, poolResponse.Data.URL + "application/view/" + address.String(), nil)

			if err == nil {
				var responseStruct response.DefaultResponse
				json.Unmarshal([]byte(applicationResponse), &responseStruct)
				responses  = append(responses, responseStruct.Response)
			}
		}
	}

	ResponseHandler(w, r, "null", true, nil, responses, nil)
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

	req.Header.Set("User-Agent", "gladius-cli")
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

