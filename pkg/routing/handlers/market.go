package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	)

// MarketPoolsHandler - Returns all Pools
func MarketPoolsHandler(w http.ResponseWriter, r *http.Request) {
	poolsWithData, err := blockchain.MarketPools(true)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
		return
	}
	ResponseHandler(w, r, "null", true, nil, poolsWithData, nil)
}

func MarketPoolsOwnedHandler(w http.ResponseWriter, r *http.Request) {
	pools, err := blockchain.MarketPoolsOwnedByUser(true)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
		return
	}

	ResponseHandler(w, r, "null", true, nil, pools, nil)
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

	ResponseHandler(w, r, "null", true, nil, nil, transaction)
}
