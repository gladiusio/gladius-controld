package handlers

import (
		"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gorilla/mux"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
)

// StatusHandler Main Status API route handler
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Status Page\n"))
}

// StatusTxHandler - Checks Status of txHash
func StatusTxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	txHash := vars["tx"]
	transaction, isPending, err := blockchain.Tx(common.HexToHash(txHash))
	complete := !isPending

	if transaction == nil || err != nil {
		ErrorHandler(w, r, "Transaction not found", err, http.StatusNotFound)
		return
	}

	transactionJSON, err := transaction.MarshalJSON()
	if err != nil {
		ErrorHandler(w, r, "Transaction could not be formatted", err, http.StatusUnprocessableEntity)
		return
	}

	var receiptResponseJSON []byte
	var status bool

	if complete {
		receipt, err := blockchain.TxReceipt(common.HexToHash(txHash))
		if receipt == nil || err != nil {
			ErrorHandler(w, r, "Receipt not found", err, http.StatusNotFound)
			return
		}

		statusResponse := !(receipt.Status == 0)
		if statusResponse {
			status = true
		} else {
			status = false
		}

		receiptJSON, err := receipt.MarshalJSON()
		if err != nil {
			ErrorHandler(w, r, "Receipt could not be formatted", err, http.StatusUnprocessableEntity)
			return
		}

		receiptResponseJSON = receiptJSON
	} else {
		receiptResponseJSON = []byte("null")
	}

	w.WriteHeader(http.StatusOK)

	txHashResponse := response.TxHash{Value:txHash, Status:status, Complete:complete, Transaction: transactionJSON, Receipt:receiptResponseJSON}

	ResponseHandler(w, r, "null", true, nil, txHashResponse, transaction)
}
