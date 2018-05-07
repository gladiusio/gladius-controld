package handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

// ResponseHandler - Default Response Handler
func ResponseHandler(w http.ResponseWriter, r *http.Request, m string, res string) {
	response := fmt.Sprintf("{ \"message\": %s, \"success\": true, \"error\": null, \"response\": %s, \"endpoint\": \"%s\" }", m, res, r.URL)

	txFormattedResponse := replaceTx(response)
	w.Write([]byte(txFormattedResponse))

	return
}

// Replaces any response
func replaceTx(payload string) string {
	re := regexp.MustCompile(`"txHash":\s"(0[xX][a-fA-F0-9]{64})"`)
	s := re.ReplaceAllString(payload, "\"txHash\": { \"value\": \"$1\", \"status\": \"http://localhost:3000/api/status/tx/$1\", \"etherscan\": { \"main\": \"https://etherscan.io/tx/$1\", \"ropsten\":\"https://ropsten.etherscan.io/tx/$1\"} }")

	return s
}
