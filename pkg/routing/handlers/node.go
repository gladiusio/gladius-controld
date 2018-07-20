package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gorilla/mux"
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"time"
	"bytes"
		"io/ioutil"
)

func NodeRetrieveDataHandler(w http.ResponseWriter, r *http.Request) {
	nodeAddress, err := blockchain.NodeOwnedByUser()
	if err != nil {
		ErrorHandler(w, r, "Node not found for user", err, http.StatusNotFound)
	}
	nodeData, err := blockchain.NodeRetrieveDataForAddress(*nodeAddress)
	if err != nil {
		ErrorHandler(w, r, "Node data could not be retrieved or data is not set", err, http.StatusNotFound)
	}

	nodeResponse := blockchain.NodeResponse{Address: nodeAddress.String(), Data: nodeData}

	ResponseHandler(w, r, "null", true, nil, nodeResponse, nil)
}

func NodeSetDataHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	decoder := json.NewDecoder(r.Body)
	var data blockchain.NodeData
	err := decoder.Decode(&data)

	if err != nil {
		ErrorHandler(w, r, "Passphrase `passphrase` not included or invalid in request", err, http.StatusBadRequest)
	}

	transaction, err := blockchain.NodeSetData(auth, &data)
	if err != nil {
		ErrorHandler(w, r, "Node data could not be set", err, http.StatusBadRequest)
		return
	}

	ResponseHandler(w, r, "null", true, nil, nil, transaction)
}

func NodeApplyToPoolHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	poolAddress := vars["poolAddress"]

	poolData, err := blockchain.PoolRetrievePublicData(poolAddress)
	poolResponse := blockchain.PoolResponse{poolAddress, poolData}
	if err != nil {
		ErrorHandler(w, r, "Pool data could not be found for Pool: " + poolAddress, err, http.StatusBadRequest)
		return
	}

	poolURL := poolResponse.Data.URL
	println(poolURL)

	decoder := json.NewDecoder(r.Body)
	var requestPayload models.NodeRequestPayload
	err = decoder.Decode(&requestPayload)

	requestPayload.IPAddress = r.RemoteAddr

	address, _ := blockchain.NewGladiusAccountManager().GetAccountAddress()

	requestPayload.Wallet = address.String()

	if err != nil {
		ErrorHandler(w, r, "Could not decode request payload", err, http.StatusBadRequest)
	}

	sendRequest(http.MethodPost, poolResponse.Data.URL + "application/new", requestPayload)
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

func NodeApplicationStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	nodeAddress := vars["nodeAddress"]
	poolAddress := vars["poolAddress"]

	status, err := blockchain.NodeApplicationStatus(nodeAddress, poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Could not find status for pool application", err, http.StatusBadRequest)
		return
	}

	statusString, err := blockchain.ApplicationStatusFromInt(int(status.Uint64()))
	if err != nil {
		ErrorHandler(w, r, "Could not find status for pool application", err, http.StatusBadRequest)
		return
	}

	statusResponse := response.NodeApplication{Status: statusString, Code: int(status.Uint64())}

	ResponseHandler(w, r, "null", true, nil, statusResponse, nil)
}

func NodePoolApplications(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nodeAddress := vars["nodeAddress"]

	address := common.HexToAddress(nodeAddress)

	pools, err := blockchain.NodePools(&address)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve applications", err, http.StatusBadRequest)
		return
	}

	var applications []response.NodeApplication

	for _, pool := range pools {
		status, err := blockchain.NodeApplicationStatus(nodeAddress, pool.String())
		if err != nil {
			ErrorHandler(w, r, "Could not find status for pool application", err, http.StatusBadRequest)
			return
		}

		statusString, err := blockchain.ApplicationStatusFromInt(int(status.Uint64()))
		if err != nil {
			ErrorHandler(w, r, "Could not find status for pool application", err, http.StatusBadRequest)
			return
		}

		app := response.NodeApplication{Status: statusString, Code: int(status.Uint64()), PoolAddress: pool.String()}

		applications = append(applications, app)
	}

	applicationsResponse := response.NodePoolApplications{NodeApplications: applications, Address: nodeAddress}

	ResponseHandler(w, r, "null", true, nil, applicationsResponse, nil)
}
