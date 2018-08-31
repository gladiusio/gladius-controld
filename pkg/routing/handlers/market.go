package handlers

import (
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
)

// MarketPoolsHandler - Returns all Pools
func MarketPoolsHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		poolsWithData, err := blockchain.MarketPools(true, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, poolsWithData, nil)
	}
}

func MarketPoolsOwnedHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		pools, err := blockchain.MarketPoolsOwnedByUser(true, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, pools, nil)
	}
}

type poolData struct {
	PublicKey string `json:"publicKey"`
}

// MarketPoolsCreateHandler - Create a new Pool
func MarketPoolsCreateHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		auth := r.Header.Get("X-Authorization")

		transaction, err := blockchain.MarketCreatePool(auth, ga)
		if err != nil {
			ErrorHandler(w, r, "Could not build pool creation transaction", err, http.StatusNotFound)
			return
		}

		ResponseHandler(w, r, "null", true, nil, nil, transaction)
	}
}
