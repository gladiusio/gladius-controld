package handlers

import (
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"net/http"
	"github.com/gorilla/mux"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"encoding/json"
)

func AccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]
	address := vars["address"]
	var symbolEnum blockchain.BalanceType

	if symbol == "gla" {
		symbolEnum = blockchain.GLA
	} else if symbol == "eth" {
		symbolEnum = blockchain.ETH
	} else {
		symbolNotFoundErr := errors.New("symbol not found for " + symbol)
		ErrorHandler(w, r, "Symbol not supported at this time, try `eth` or `gla`", symbolNotFoundErr, http.StatusNotFound)
		return
	}

	balance, err := blockchain.GetAccountBalance(common.HexToAddress(address), blockchain.BalanceType(symbolEnum))

	if err != nil {
		ErrorHandler(w, r, "Could not retrieve balance for " + address, err, http.StatusInternalServerError)
		return
	}

	ResponseHandler(w, r, "null", true, nil, balance, nil)
}

func AccountTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	decoder := json.NewDecoder(r.Body)
	var options blockchain.TransactionOptions
	err := decoder.Decode(&options)

	transactions, err := blockchain.GetAccountTransactions(common.HexToAddress(address), options)

	if err != nil {
		ErrorHandler(w, r, "Could not retrieve transactions for " + address, err, http.StatusInternalServerError)
		return
	}

	ResponseHandler(w, r, "null", true, nil, transactions.Transactions, nil)
}