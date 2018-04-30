package handlers

import (
	"encoding/json"
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"log"
	"net/http"
)

type wallet_struct struct {
	Passphrase string `json:"passphrase"`
}

func KeystoreCreationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var wallet wallet_struct
	err := decoder.Decode(&wallet)

	if err != nil {
		ErrorHandler(w, r, "Passphrase `passphrase` not included or invalid in request", err, http.StatusBadRequest)
	}

	defer r.Body.Close()
	log.Println(wallet.Passphrase)

	response, err := blockchain.CreateKeyStore(wallet.Passphrase)
	if err != nil {
		ErrorHandler(w, r, "Wallet could not be created", err, http.StatusInternalServerError)
	}

	ResponseHandler(w, r, "null", response)
}
