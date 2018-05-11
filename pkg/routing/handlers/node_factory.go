package handlers

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
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
		return
	}

	jsonResponse := fmt.Sprintf("0x%x", nodeAddress)
	ResponseHandler(w, r, "null", "\""+string(jsonResponse)+"\"")
}

func NodeFactoryCreateNodeHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	transaction, err := blockchain.CreateNode(auth)

	if transaction == nil || err != nil {
		ErrorHandler(w, r, "Could not create Node for account", err, http.StatusNotFound)
		return
	}

	TransactionHandler(w, r, "null", transaction)
}
