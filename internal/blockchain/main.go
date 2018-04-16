package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

var providerURL = "https://ropsten.infura.io/tjqLYxxGIUp0NylVCiWw"

// ConnectClient - Main Connection function
func ConnectClient() *ethclient.Client {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(providerURL)
	//conn, err := ethclient.Dial("/home/nate/.ethereum/testnet/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	return conn
}

// Nonce - get the current Nonce for an address
func Nonce(hash common.Hash) (count uint, _ error) {
	conn := ConnectClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return conn.TransactionCount(ctx, hash)
}
