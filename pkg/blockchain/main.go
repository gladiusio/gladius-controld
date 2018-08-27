package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

// ConnectClient - Main Connection function
func ConnectClient() *ethclient.Client {
	providerURL := viper.GetString("blockchain.provider")
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(providerURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	return conn
}
