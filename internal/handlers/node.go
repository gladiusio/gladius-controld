package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"net/http"
)

// NodeHandler - Main Node API route handler
func NodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Main Node API\n"))
}

func NodeRetrieveDataHandler(w http.ResponseWriter, r *http.Request) {
	nodeData, _ := blockchain.NodeRetrieveData()
	jsonResponse := nodeData.String()

	ResponseHandler(w, r, "null", jsonResponse)
}

func NodeSetDataHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Authorization")
	decoder := json.NewDecoder(r.Body)
	var data blockchain.NodeData
	err := decoder.Decode(&data)

	if err != nil {
		ErrorHandler(w, r, "Passphrase `passphrase` not included or invalid in request", err, http.StatusBadRequest)
	}

	transaction, _ := blockchain.NodeSetData(auth, &data)
	jsonResponse := fmt.Sprintf("{\"txHash\": \"0x%x\"}", transaction)
	ResponseHandler(w, r, "null", string(jsonResponse))
}
