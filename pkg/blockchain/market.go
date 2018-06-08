package blockchain

import (
	"encoding/json"
	"log"
	"strings"

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

type PoolResponse struct {
	Address string         `json:"address"`
	Data    PoolPublicData `json:"data"`
}

func (d *PoolResponse) String() string {
	json, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(json)
}

func MarketPoolsWithData() (string, error) {
	poolAddresses, err := MarketPools()
	if err != nil {
		return "[]", err
	}

	response := "["

	for _, poolAddress := range poolAddresses {
		poolData, err := PoolRetrievePublicData(poolAddress.String())
		poolResponse := PoolResponse{poolAddress.String(), *poolData}
		if err != nil {
			return "[]", err
		}

		response += poolResponse.String() + ","
	}

	response = strings.TrimRight(response, ",")
	response += "]"

	return response, nil
}

//MarketCreatePool - Create new pool
func MarketCreatePool(passphrase, publicKey string) (*types.Transaction, error) {
	market := ConnectMarket()
	ga := NewGladiusAccountManager()
	auth, err := ga.GetAuth(passphrase)
	if err != nil {
		return nil, err
	}

	transaction, err := market.CreatePool(auth, publicKey)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
