package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"net/http"
	"strconv"
)

type wallet_struct struct {
	Passphrase string `json:"passphrase"`
}

func passphraseDecoder(w http.ResponseWriter, r *http.Request) wallet_struct {
	decoder := json.NewDecoder(r.Body)
	var wallet wallet_struct
	err := decoder.Decode(&wallet)

	if err != nil {
		ErrorHandler(w, r, "Passphrase `passphrase` not included or invalid in request", err, http.StatusBadRequest)
	}

	defer r.Body.Close()

	return wallet
}

func KeystoreCreationHandler(w http.ResponseWriter, r *http.Request) {
	wallet := passphraseDecoder(w, r)

	account, err := blockchain.CreateAccount(wallet.Passphrase)
	if err != nil {
		ErrorHandler(w, r, "Wallet could not be created", err, http.StatusInternalServerError)
	}

	response := blockchain.AccountResponseFormatter(&account)

	ResponseHandler(w, r, "null", response)
}

func KeystoreWalletRetrievalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	indexString := vars["index"]
	index, _ := strconv.Atoi(indexString)

	wallet := blockchain.Wallets()[index]

	response := blockchain.WalletResponseFormatter(wallet)
	ResponseHandler(w, r, "null", response)
}

func KeystoreWalletOpenHandler(w http.ResponseWriter, r *http.Request) {
	walletStruct := passphraseDecoder(w, r)
	vars := mux.Vars(r)
	indexString := vars["index"]
	index, _ := strconv.Atoi(indexString)

	wallet := blockchain.OpenWallet(index, walletStruct.Passphrase)

	response := blockchain.WalletResponseFormatter(wallet)
	ResponseHandler(w, r, "null", response)
}

func KeystoreWalletsRetrievalHandler(w http.ResponseWriter, r *http.Request) {
	wallets := blockchain.Wallets()

	response := "["

	for index, wallet := range wallets {
		response += blockchain.WalletResponseFormatter(wallet)
		if index < len(wallets) {
			response += ","
		}
	}

	response += "]"

	ResponseHandler(w, r, "null", response)
}
