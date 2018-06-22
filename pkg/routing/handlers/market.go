package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
)

// MarketPoolsHandler - Returns all Pools
func MarketPoolsHandler(w http.ResponseWriter, r *http.Request) {
	poolsWithData, err := blockchain.MarketPoolsWithData()
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
		return
	}

	ResponseHandler(w, r, "null", poolsWithData)
}

func MarketPoolsOwnedHandler(w http.ResponseWriter, r *http.Request) {
	pools, err := blockchain.MarketPoolsOwnedByUser()
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
		return
	}

	var poolsArray response.AddressHashes = pools
	jsonPayload, err := json.Marshal(poolsArray)
	if err != nil {
		ErrorHandler(w, r, "Could not parse pools json", err, http.StatusNotFound)
		return
	}

	ResponseHandler(w, r, "null", string(jsonPayload))
}

type poolData struct {
	PublicKey string `json:"publicKey"`
}

// MarketPoolsCreateHandler - Create a new Pool
func MarketPoolsCreateHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	decoder := json.NewDecoder(r.Body)
	var data poolData
	err := decoder.Decode(&data)

	transaction, err := blockchain.MarketCreatePool(auth, data.PublicKey)
	if err != nil {
		ErrorHandler(w, r, "Could not build pool creation transaction", err, http.StatusNotFound)
		return
	}

	TransactionHandler(w, r, "null", transaction)
}
