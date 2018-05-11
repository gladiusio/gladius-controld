package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/crypto"
	"github.com/gorilla/mux"
)

type wallet_struct struct {
	Passphrase string `json:"passphrase"`
}

type pgp_struct struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Email   string `json:"email"`
}

func passphraseDecoder(w http.ResponseWriter, r *http.Request) (wallet_struct, error) {
	decoder := json.NewDecoder(r.Body)
	var wallet wallet_struct
	err := decoder.Decode(&wallet)

	if err != nil {
		return wallet_struct{}, err
	}

	defer r.Body.Close()

	return wallet, nil
}

func KeystoreCreationHandler(w http.ResponseWriter, r *http.Request) {
	wallet, err := passphraseDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusInternalServerError)
		return
	}

	account, err := blockchain.CreateAccount(wallet.Passphrase)
	if err != nil {
		ErrorHandler(w, r, "Wallet could not be created", err, http.StatusInternalServerError)
		return
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
	walletStruct, err := passphraseDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusInternalServerError)
		return
	}
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
		if index < len(wallets)-1 {
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
		return
	}

	path, err := crypto.CreateKeyPair(pgpStruct.Name, pgpStruct.Comment, pgpStruct.Email)
	if err != nil {
		ErrorHandler(w, r, "PGP key pair could not be created", err, http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("{\"path\": \"%s\"}", path)
	ResponseHandler(w, r, "null", response)
}
