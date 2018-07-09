package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/crypto"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
			)

type accountBody struct {
	Passphrase string `json:"passphrase"`
}

func passphraseDecoder(w http.ResponseWriter, r *http.Request) (*accountBody, error) {
	decoder := json.NewDecoder(r.Body)
	var ab accountBody
	err := decoder.Decode(&ab)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &ab, nil
}

func KeystoreAccountCreationHandler(w http.ResponseWriter, r *http.Request) {
	wallet, err := passphraseDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusInternalServerError)
		return
	}
	ga := blockchain.NewGladiusAccountManager()

	_, err = ga.CreateAccount(wallet.Passphrase)
	if err != nil {
		ErrorHandler(w, r, "Account could not be created", err, http.StatusInternalServerError)
		return
	}

	address, err := ga.GetAccountAddress()
	if err != nil {
		ErrorHandler(w, r, "Account address could not be retrieved", err, http.StatusInternalServerError)
		return
	}

	addressResponse := response.AddressResponse{Address: *address}

	ResponseHandler(w, r, "null", true, nil, addressResponse, nil)
}

func KeystoreAccountRetrievalHandler(w http.ResponseWriter, r *http.Request) {
	ga := blockchain.NewGladiusAccountManager()
	address, err := ga.GetAccountAddress()
	if err != nil {
		ErrorHandler(w, r, "Account address could not be retrieved", err, http.StatusInternalServerError)
		return
	}

	addressResponse := response.AddressResponse{Address: *address}

	ResponseHandler(w, r, "null", true, nil, addressResponse, nil)
}

func KeystoreAccountUnlockHandler(w http.ResponseWriter, r *http.Request) {
	accountBody, err := passphraseDecoder(w, r)
	if err != nil {
		ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusInternalServerError)
		return
	}

	ga := blockchain.NewGladiusAccountManager()

	success, err := ga.UnlockAccount(accountBody.Passphrase)
	if success == false || err != nil {
		ErrorHandler(w, r, "Wallet could not be opened, passphrase could be incorrect", err, http.StatusBadRequest)
		return
	}

	address, err := ga.GetAccountAddress()
	if err != nil {
		ErrorHandler(w, r, "Account address could not be retrieved", err, http.StatusInternalServerError)
		return
	}

	addressResponse := response.AddressResponse{Address: *address}

	ResponseHandler(w, r, "null", true, nil, addressResponse, nil)
}

func KeystorePGPPublicKeyRetrievalHandler(w http.ResponseWriter, r *http.Request) {
	publicKey, err := blockchain.GetPGPPublicKey()
	if err != nil {
		ErrorHandler(w, r, "Public key not found or cannot be read", err, http.StatusNotFound)
		return
	}

	publicKeyResponse := response.PublicKeyResponse{PublicKey: publicKey}

	ResponseHandler(w, r, "null", true, nil, publicKeyResponse, nil)
}

type pgpStruct struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Email   string `json:"email"`
}

func KeystorePGPCreationHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ps pgpStruct
	err := decoder.Decode(&ps)

	if err != nil {
		ErrorHandler(w, r, "Request invalid, body is missing either `name`, `comment`, and/or `email`", err, http.StatusBadRequest)
		return
	}

	path, err := crypto.CreateKeyPair(ps.Name, ps.Comment, ps.Email)
	println(path)

	if err != nil {
		ErrorHandler(w, r, "PGP key pair could not be created", err, http.StatusInternalServerError)
		return
	}

	creationResponse := response.CreationResponse{Created: true}

	ResponseHandler(w, r, "null", true, nil, creationResponse, nil)
}
