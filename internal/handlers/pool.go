package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"net/http"
)

// PoolHandler - Main Node API route handler
func PoolHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node API\n"))
}

func PoolRetrievePublicKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	poolAddress := vars["poolAddress"]

	publicKey, err := blockchain.PoolRetrievePublicKey(poolAddress)
	if err != nil {
		ErrorHandler(w, r, "Could not retrieve Pool's Public Key", err, http.StatusUnprocessableEntity)
		return
	}

	response := fmt.Sprintf("{\"publicKey\": \"%s\"}", publicKey)
	ResponseHandler(w, r, "null", response)
}

//func NodeRetrieveDataHandler(w http.ResponseWriter, r *http.Request) {
//nodeData, _ := blockchain.NodeRetrieveData()
//jsonResponse := nodeData.String()

//ResponseHandler(w, r, "null", jsonResponse)
//}

//func NodeSetDataHandler(w http.ResponseWriter, r *http.Request) {
//auth := r.Header.Get("X-Authorization")
//decoder := json.NewDecoder(r.Body)
//var data blockchain.NodeData
//err := decoder.Decode(&data)

//if err != nil {
//ErrorHandler(w, r, "Passphrase `passphrase` not included or invalid in request", err, http.StatusBadRequest)
//}

//transaction, _ := blockchain.NodeSetData(auth, &data)
//TransactionHandler(w, r, "null", transaction)
//}
