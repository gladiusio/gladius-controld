package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
)

// ResponseHandler - Default Response Handler
func ResponseHandler(w http.ResponseWriter, r *http.Request, m string, success bool, err *string, res interface{}, transaction *types.Transaction) {
	errorString := ""

	if err != nil {
		errorString = *err
	}

	responseStruct := response.DefaultResponse{
		Message:     m,
		Success:     success,
		Error:       errorString,
		Response:    &res,
		Transaction: nil,
		Endpoint:    r.URL.String(),
	}

	if transaction != nil {
		responseStruct.FormatTransactionResponse(transaction.Hash().String())
	}

	responseJSON, parseErr := json.Marshal(responseStruct)

	if parseErr != nil {
		ErrorHandler(w, r, "Could not parse response JSON", parseErr, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(responseJSON))
	return
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New(r.URL.String() + " not found in available routes")
	ErrorHandler(w, r, "Invalid request, check parameters and try again", err, http.StatusNotFound)
	return
}

// ErrorHandler - Default Error Handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, m string, e error, statusCode int) {
	w.WriteHeader(statusCode)

	var err string
	if e != nil {
		err = e.Error()
	} else {
		err = "Error message could not be parsed"
	}

	ResponseHandler(w, r, m, false, &err, nil, nil)

	return
}

func AccountNotFoundErrorHandler(w http.ResponseWriter, r *http.Request, ga *blockchain.GladiusAccountManager) error {
	if !ga.HasAccount() {
		err := errors.New("account not found")
		ErrorHandler(w, r, "Account not found, please create an account", err, http.StatusBadRequest)
		return err
	}

	return nil
}

func AccountUnlockedErrorHandler(w http.ResponseWriter, r *http.Request, ga *blockchain.GladiusAccountManager) error {
	if !ga.Unlocked() {
		err := errors.New("wallet locked")
		ErrorHandler(w, r, "Wallet could not be opened, passphrase is incorrect", err, http.StatusMethodNotAllowed)
		return err
	}
	return nil
}

// Account Manager Error Handler, checks required account permissions prior to accessing API endpoints
func AccountErrorHandler(w http.ResponseWriter, r *http.Request, ga *blockchain.GladiusAccountManager) error {
	err := AccountNotFoundErrorHandler(w, r, ga)
	if err != nil {
		return err
	}

	err = AccountUnlockedErrorHandler(w, r, ga)
	if err != nil {
		return err
	}

	return nil
}
