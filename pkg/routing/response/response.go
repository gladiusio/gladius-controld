package response

import (
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
	Addresses   []string `json:"addresses"`
}