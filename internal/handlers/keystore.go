package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nfeld9807/rest-api/internal/blockchain"
	"github.com/nfeld9807/rest-api/internal/crypto"
	"net/http"
	"strconv"
)

type wallet_struct struct {
	Passphrase string `json:"passphrase"`
}

type pgp_struct struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Email   string `json:"email"`
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

func KeystorePGPCreationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var pgpStruct pgp_struct
	err := decoder.Decode(&pgpStruct)

	if err != nil {
		ErrorHandler(w, r, "Request invalid, body is missing either `name`, `comment`, and/or `email`", err, http.StatusBadRequest)
	}

	path, err := crypto.CreateKeyPair(pgpStruct.Name, pgpStruct.Comment, pgpStruct.Email)
	if err != nil {
		ErrorHandler(w, r, "PGP key pair could not be created", err, http.StatusInternalServerError)
	}

	//encMessage, err := crypto.EncryptMessage("Test Message")

	//if err != nil {
	//log.Fatal(err)
	//}

	//println(encMessage)

	//decMessage, err := crypto.DecryptMessage(encMessage)

	//if err != nil {
	//log.Fatal(err)
	//}

	//println(decMessage)

	//wallet := passphraseDecoder(w, r)

	//account, err := blockchain.CreateAccount(wallet.Passphrase)
	//response := blockchain.AccountResponseFormatter(&account)
	response := fmt.Sprintf("{\"path\": \"%s\"}", path)
	ResponseHandler(w, r, "null", response)
}
