package response

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

type AddressHashes []common.Address

func (addressHashes AddressHashes) StringArray() []string {
	response := make([]string, len(addressHashes))

	for index, address := range addressHashes {
		response[index] = address.Str()
	}

	return response[:]
}

type AddressArray struct {
	Addresses []string `json:"addresses"`
}

type AddressResponse struct {
	Address common.Address `json:"address"`
}

type PublicKeyResponse struct {
	PublicKey string `json:"publicKey"`
}

type CreationResponse struct {
	Created bool `json:"created"`
}

type DefaultResponse struct {
	Message     string       `json:"message"`
	Success     bool         `json:"success"`
	Error       string       `json:"error"`
	Response    *interface{} `json:"response"`
	Transaction *TxHash      `json:"txHash"`
	Endpoint    string       `json:"endpoint"`
}

type TxHash struct {
	Value       string          `json:"value,omitempty"`
	Status      bool            `json:"status"`
	EndPoint    string          `json:"statusEndpoint,omitempty"`
	Complete    bool            `json:"complete"`
	Etherscan   etherscan       `json:"etherscan,omitempty"`
	Transaction json.RawMessage `json:"transaction,omitempty"`
	Receipt     json.RawMessage `json:"receipt,omitempty"`
}

type etherscan struct {
	Main    string `json:"main,omitempty"`
	Ropsten string `json:"ropsten,omitempty"`
}

func (response *DefaultResponse) FormatTransactionResponse(transaction string) {
	response.Transaction = &TxHash{
		EndPoint: "http://localhost:3000/api/status/tx/" + transaction,
		Value:    transaction,
		Etherscan: etherscan{
			//Main: "https://etherscan.io/tx/" + transaction,
			Ropsten: "https://ropsten.etherscan.io/tx/" + transaction,
		},
	}
}

type NodeApplication struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	PoolAddress string `json:"poolAddress"`
}

type NodePoolApplications struct {
	NodeApplications []NodeApplication `json:"applications"`
	Address          string            `json:"address"`
}
