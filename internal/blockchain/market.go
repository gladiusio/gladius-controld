package blockchain

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// main function
func temp() *Market {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("https://ropsten.infura.io/tjqLYxxGIUp0NylVCiWw")
	//conn, err := ethclient.Dial("/home/nate/.ethereum/testnet/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Instantiate the contract and display its name
	market, err := NewMarket(common.HexToAddress("0xc4dfb5c9e861eeae844795cfb8d30b77b78bbc38"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Market contract: %v", err)
	}

	owner, err := market.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve owner: %v", err)
	}

	fmt.Println("Owner: ", owner.String())

	return market
}

// MarketPools - List all available market pools
func MarketPools() {
	market := temp()

	pools, err := market.GetAllPools(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve pools: %v", err)
	}

	for _, pool := range pools {
		fmt.Println("Pool: ", pool.String())
	}
}
