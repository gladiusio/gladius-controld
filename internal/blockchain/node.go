package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nfeld9807/rest-api/internal/blockchain/generated"
	"log"
)

// ConnectNode - Connect and grab node
func ConnectNode(nodeAddress string) *generated.Node {

	conn := ConnectClient()

	node, err := generated.NewNode(common.HexToAddress(nodeAddress), conn)

	if err != nil {
		log.Fatalf("Failed to instantiate a Market contract: %v", err)
	}

	return node
}

// Encryption

// Encrypt against key

// Decryption

// GET / SET Data

// Data for pool

// Apply to pool

// Status for pool
