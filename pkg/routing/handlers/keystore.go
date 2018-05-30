package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/crypto"
)

type accountBody struct {
	Passphrase string `json:"passphrase"`
}

type pgp_struct struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Email   string `json:"email"`
}

func passphraseDecoder(w http.ResponseWriter, r *http.Request) (*accountBody, error) {
	decoder := json.NewDecoder(r.Body)
	var accountBody accountBody
	err := decoder.Decode(&accountBody)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &accountBody, nil
}

func KeystoreAccountCreationHandler(w http.ResponseWriter, r *http.Request) {
	wallet, err := passphraseDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusInternalServerError)
		return
	}
	ga := blockchain.NewGladiusAccountManager()

	_, err2 := ga.CreateAccount(wallet.Passphrase)
	if err2 != nil {
		ErrorHandler(w, r, "Account could not be created", err2, http.StatusInternalServerError)
		return
	}

	response := ga.AccountResponseFormatter()

	ResponseHandler(w, r, "null", response)
}

func KeystoreAccountRetrievalHandler(w http.ResponseWriter, r *http.Request) {
	ga := blockchain.NewGladiusAccountManager()

	response := ga.AccountResponseFormatter()
	ResponseHandler(w, r, "null", response)
}

func KeystoreAccountUnlockHandler(w http.ResponseWriter, r *http.Request) {
	accountBody, err := passphraseDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusInternalServerError)
		return
	}

	ga := blockchain.NewGladiusAccountManager()

	ga.UnlockAccount(accountBody.Passphrase)
	if err != nil {
		ErrorHandler(w, r, "Wallet could not be opened, passphrase could be incorrect", err, http.StatusBadRequest)
		return
	}

	response := ga.AccountResponseFormatter()
	ResponseHandler(w, r, "null", response)
}

func KeystorePGPPublicKeyRetrievalHandler(w http.ResponseWriter, r *http.Request) {
	publicKey, err := blockchain.GetPGPPublicKey()
	if err != nil {
		ErrorHandler(w, r, "Public key not found or cannot be read", err, http.StatusNotFound)
		return
	}
	ResponseHandler(w, r, "null", "{\"publicKey\": \""+publicKey+"\"}")
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
	println(path)

	if err != nil {
		ErrorHandler(w, r, "PGP key pair could not be created", err, http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("{\"created\": true}")
	ResponseHandler(w, r, "null", response)
}
