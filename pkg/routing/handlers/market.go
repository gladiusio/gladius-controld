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
	pools, err := blockchain.MarketPools()
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
		return
	}

	length := int(len(pools))
	response := make([]string, length)

	for i, pool := range pools {
		response[i] = pool.String()
	}

	jsonResponse, _ := json.Marshal(response)
	ResponseHandler(w, r, "null", string(jsonResponse))
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
