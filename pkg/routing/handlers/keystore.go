package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gladiusio/gladius-controld/pkg/blockchain"
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

func KeystoreAccountCreationHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		wallet, err := passphraseDecoder(w, r)
		if err != nil {
			ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusBadRequest)
			return
		}

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
}
func KeystoreAccountRetrievalHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountErrorHandler(w, r, ga)
		if err != nil {
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
}
func KeystoreAccountUnlockHandler(ga *blockchain.GladiusAccountManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AccountNotFoundErrorHandler(w, r, ga)
		if err != nil {
			return
		}

		accountBody, err := passphraseDecoder(w, r)
		if err != nil {
			ErrorHandler(w, r, "Could not find `passphrase` in request", err, http.StatusBadRequest)
			return
		}

		success, err := ga.UnlockAccount(accountBody.Passphrase)
		if success == false || err != nil {
			ErrorHandler(w, r, "Wallet could not be opened, passphrase is incorrect", err, http.StatusMethodNotAllowed)
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
}
