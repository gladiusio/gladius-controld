package handlers

import (
	"fmt"
	//"github.com/ethereum/go-ethereum/common"
	"net/http"
	"regexp"

	"github.com/ethereum/go-ethereum/core/types"
)

// ResponseHandler - Default Response Handler
func ResponseHandler(w http.ResponseWriter, r *http.Request, m string, res string) {
	response := fmt.Sprintf("{ \"message\": %s, \"success\": true, \"error\": null, \"response\": %s, \"endpoint\": \"%s\" }", m, res, r.URL)

	txFormattedResponse := replaceTx(response)
	w.Write([]byte(txFormattedResponse))

	return
}

func TransactionHandler(w http.ResponseWriter, r *http.Request, m string, transaction *types.Transaction) {
	transactionJSON, err := transaction.MarshalJSON()
	if err != nil {
		ErrorHandler(w, r, "Could not parse transaction JSON", err, http.StatusInternalServerError)
		return
	}

	txHash := transaction.Hash().String()

	response := fmt.Sprintf("{ \"txPayload\": %s, \"txHash\": \"%s\" }", string(transactionJSON), txHash)

	ResponseHandler(w, r, m, response)
}

// ErrorHandler - Default Error Handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, m string, e error, statusCode int) {
	w.WriteHeader(statusCode)

	response := fmt.Sprintf("{ \"message\": \"%s\", \"success\": false, \"error\": \"%s\", \"response\": null, \"endpoint\": \"%s\" }", m, e, r.URL)
	w.Write([]byte(response))

	return
}

// Replaces any response
func replaceTx(payload string) string {
	re := regexp.MustCompile(`"txHash":\s"(0[xX][a-fA-F0-9]{64})"`)
	s := re.ReplaceAllString(payload, "\"txHash\": { \"value\": \"$1\", \"status\": \"http://localhost:3000/api/status/tx/$1\", \"etherscan\": { \"main\": \"https://etherscan.io/tx/$1\", \"ropsten\":\"https://ropsten.etherscan.io/tx/$1\"} }")

	return s
}
