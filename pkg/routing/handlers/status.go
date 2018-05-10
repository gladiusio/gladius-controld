package handlers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"net/http"
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
	status := []byte("null")

	if complete {
		receipt, err := blockchain.TxReceipt(common.HexToHash(txHash))
		if receipt == nil || err != nil {
			ErrorHandler(w, r, "Receipt not found", err, http.StatusNotFound)
			return
		}

		statusResponse := !(receipt.Status == 0)
		if statusResponse {
			status = []byte("true")
		} else {
			status = []byte("false")
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

	response := fmt.Sprintf("{ \"txHash\": \"%s\", \"complete\": %t, \"status\": %s, \"transaction\": %s, \"receipt\": %s }", txHash, complete, status, transactionJSON, receiptResponseJSON)
	ResponseHandler(w, r, "null", response)
}
