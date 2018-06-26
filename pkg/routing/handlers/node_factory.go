package handlers

import (
	"errors"
		"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
		)

// NodeFactoryHandler - Main Node API route handler
func NodeFactoryNodeAddressHandler(w http.ResponseWriter, r *http.Request) {
	accountAddress := r.URL.Query().Get("account")
	ga := blockchain.NewGladiusAccountManager()
	if accountAddress == "" {
		accountAddress = ga.GetAccountAddress().String()
	}
	nodeAddress, err := blockchain.NodeForAccount(common.HexToAddress(accountAddress))

	if err != nil {
		ErrorHandler(w, r, "Could not retrieve Node for account", err, http.StatusNotFound)
		return
	}

	if nodeAddress.String() == "0x0000000000000000000000000000000000000000" {
		ErrorHandler(w, r, "Account does not have an associated node", errors.New("account has not created a node"), http.StatusNotFound)
		return
	}

	nodeData, err := blockchain.NodeRetrieveDataForAddress(*nodeAddress)
	var nodeResponse blockchain.NodeResponse
	if err == nil {
		nodeResponse = blockchain.NodeResponse{Address: nodeAddress.String(), Data: nodeData}
	} else {
		nodeResponse = blockchain.NodeResponse{Address: nodeAddress.String(), Data: nil}
	}

	ResponseHandler(w, r, "null", true, nil, nodeResponse, nil)
}

func NodeFactoryCreateNodeHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	transaction, err := blockchain.CreateNode(auth)

	if transaction == nil || err != nil {
		ErrorHandler(w, r, "Could not create Node for account", err, http.StatusNotFound)
		return
	}

	ResponseHandler(w, r, "null", true, nil, nil, transaction)
}
