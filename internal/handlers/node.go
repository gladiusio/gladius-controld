package handlers

import (
	//"encoding/json"
	//"fmt"
	//"github.com/nfeld9807/rest-api/internal/blockchain"
	"net/http"
)

// NodeHandler - Main Node API route handler
func NodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node API\n"))
}

//// MarketPoolsHandler - Returns all Pools
//func MarketPoolsHandler(w http.ResponseWriter, r *http.Request) {
//pools, err := blockchain.MarketPools()
//if err != nil {
//ErrorHandler(w, r, "Could not retrieve pools", err, http.StatusNotFound)
//}

//length := int(len(pools))
//response := make([]string, length)

//for i, pool := range pools {
//response[i] = pool.String()
//}

//jsonResponse, _ := json.Marshal(response)
//ResponseHandler(w, r, "null", string(jsonResponse))
//}

//// MarketPoolsCreateHandler - Create a new Pool
//func MarketPoolsCreateHandler(w http.ResponseWriter, r *http.Request) {
//transaction, err := blockchain.MarketCreatePool("test")
//if err != nil {
//ErrorHandler(w, r, "Could not build pool creation transaction", err, http.StatusNotFound)
//}

//jsonResponse := fmt.Sprintf("0x%x", transaction)
//ResponseHandler(w, r, "null", string(jsonResponse))
//}
