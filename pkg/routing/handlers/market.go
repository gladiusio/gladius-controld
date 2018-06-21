package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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

type AddressArray []common.Address

func (addressArray AddressArray) String() string {
	response := "["

	for _, address := range addressArray {
		response += "\"" + address.String() + "\"" + ","
	}

	response = strings.TrimRight(response, ",")
	response += "]"

	return response
}

func MarketPoolsOwnedHandler(w http.ResponseWriter, r *http.Request) {
	pools, err := blockchain.MarketPoolsOwnedByUser()
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
		return
	}

	var poolsArray AddressArray = pools

	ResponseHandler(w, r, "null", poolsArray.String())
}

type poolData struct {
	publicKey string `json:"publicKey"`
}

// MarketPoolsCreateHandler - Create a new Pool
func MarketPoolsCreateHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	decoder := json.NewDecoder(r.Body)
	var data poolData
	err := decoder.Decode(&data)

	transaction, err := blockchain.MarketCreatePool(auth, data.publicKey)
	if err != nil {
		ErrorHandler(w, r, "Could not build pool creation transaction", err, http.StatusNotFound)
		return
	}

	TransactionHandler(w, r, "null", transaction)
}
