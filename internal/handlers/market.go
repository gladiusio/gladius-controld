package handlers

import (
	"encoding/json"
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"net/http"
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
	}

	length := int(len(pools))
	response := make([]string, length)

	for i, pool := range pools {
		response[i] = pool.String()
	}

	jsonResponse, _ := json.Marshal(response)
	ResponseHandler(w, r, "null", string(jsonResponse))
}

// MarketPoolsCreateHandler - Create a new Pool
func MarketPoolsCreateHandler(w http.ResponseWriter, r *http.Request) {

}
