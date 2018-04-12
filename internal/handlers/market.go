package handlers

import (
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"net/http"
)

// MarketHandler - Main Market API route handler
func MarketHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Market API\n"))
}

// MarketPoolsHandler - Returns all Pools
func MarketPoolsHandler(w http.ResponseWriter, r *http.Request) {
	blockchain.MarketPools()
}

// MarketPoolsCreateHandler - Create a new Pool
func MarketPoolsCreateHandler(w http.ResponseWriter, r *http.Request) {

}
