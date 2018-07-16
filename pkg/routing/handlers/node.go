package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gorilla/mux"
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

	nodeAddress := vars["nodeAddress"]
	poolAddress := vars["poolAddress"]

	auth := r.Header.Get("X-Authorization")
	transaction, err := blockchain.NodeApplyToPool(auth, nodeAddress, poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Could not apply to pool", err, http.StatusBadRequest)
		return
	}

	ResponseHandler(w, r, "null", true, nil, nil, transaction)
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
