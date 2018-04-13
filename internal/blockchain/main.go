package blockchain

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// ConnectClient - Main Connection function
func ConnectClient() *ethclient.Client {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("https://ropsten.infura.io/tjqLYxxGIUp0NylVCiWw")
	//conn, err := ethclient.Dial("/home/nate/.ethereum/testnet/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	return conn
}
