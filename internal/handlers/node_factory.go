package handlers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"net/http"
)

// NodeFactoryHandler - Main Node API route handler
func NodeFactoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node Factory API\n"))
}

// NodeFactoryHandler - Main Node API route handler
func NodeFactoryNodeAddressHandler(w http.ResponseWriter, r *http.Request) {
	accountAddress := r.URL.Query().Get("account")
	if accountAddress == "" {
		accountAddress = blockchain.GetDefaultAccountAddress().String()
	}
	nodeAddress, err := blockchain.NodeForAccount(common.HexToAddress(accountAddress))

	if err != nil {
		ErrorHandler(w, r, "Could not retrieve Node for account", err, http.StatusNotFound)
	}

	jsonResponse := fmt.Sprintf("0x%x", nodeAddress)
	ResponseHandler(w, r, "null", string(jsonResponse))
}

func NodeFactoryCreateNodeHandler(w http.ResponseWriter, r *http.Request) {
	txHash, err := blockchain.CreateNode()

	if err != nil {
		ErrorHandler(w, r, "Could not create Node for account", err, http.StatusNotFound)
	}

	ResponseHandler(w, r, "null", "{ \"txHash\": \""+txHash+"\"}")
}
