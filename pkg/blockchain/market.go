package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
	"github.com/spf13/viper"
)

// ConnectMarket - Connect and return configured market
func ConnectMarket() *generated.Market {

	conn := ConnectClient()

	marketAddress := viper.GetString("BlockchainMarketAddress")
	market, err := generated.NewMarket(common.HexToAddress(marketAddress), conn)

	if err != nil {
		log.Fatalf("Failed to instantiate a Market contract: %v", err)
	}

	return market
}

// MarketPools - List all available market pools
func MarketPools() ([]common.Address, error) {
	market := ConnectMarket()

	pools, err := market.GetAllPools(nil)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

//MarketCreatePool - Create new pool
func MarketCreatePool(passphrase, publicKey string) (*types.Transaction, error) {
	market := ConnectMarket()
	auth := GetDefaultAuth(passphrase)

	transaction, err := market.CreatePool(auth, publicKey)
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}

	return transaction, nil
}
