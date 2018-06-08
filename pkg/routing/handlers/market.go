package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
)

// MarketHandler - Main Market API route handler
func MarketHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Market API\n"))
}

// MarketPoolsHandler - Returns all Pools
func MarketPoolsHandler(w http.ResponseWriter, r *http.Request) {
	poolsWithData, err := blockchain.MarketPoolsWithData()
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
		return
	}

	ResponseHandler(w, r, "null", poolsWithData)
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
