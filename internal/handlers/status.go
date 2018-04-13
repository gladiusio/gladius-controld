package handlers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/nfeld9807/rest-api/internal/blockchain"
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

	if !isPending {
		receipt, err := blockchain.TxReceipt(common.HexToHash(txHash))
		if receipt == nil || err != nil {
			ErrorHandler(w, r, "Receipt not found", err, http.StatusNotFound)
			return
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

	response := fmt.Sprintf("{ \"isPending\": %t, \"transaction\": %s, \"receipt\": %s }", isPending, transactionJSON, receiptResponseJSON)
	ResponseHandler(w, r, "null", response)
}
