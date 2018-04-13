package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nfeld9807/rest-api/internal/blockchain/generated"
	"log"
)

// ConnectMarket - Connect and return configured market
func ConnectMarket() *generated.Market {

	conn := ConnectClient()

	market, err := generated.NewMarket(common.HexToAddress("0xc4dfb5c9e861eeae844795cfb8d30b77b78bbc38"), conn)
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
	market := ConnectMarket()

	pools, err := market.GetAllPools(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve pools: %v", err)
	}

	for _, pool := range pools {
		fmt.Println("Pool: ", pool.String())
	}
}
