package blockchain

import (
	"encoding/json"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
	"github.com/spf13/viper"
)

// ConnectMarket - Connect and return configured market
func ConnectMarket() *generated.Market {

	conn := ConnectClient()

	marketAddress := viper.GetString("blockchain.market")
	market, err := generated.NewMarket(common.HexToAddress(marketAddress), conn)

	if err != nil {
		log.Fatalf("Failed to instantiate a Market contract: %v", err)
	}

	return market
}

func MarketPoolsOwnedByUser(includeData bool, ga *GladiusAccountManager) (PoolArrayResponse, error) {
	market := ConnectMarket()

	address, err := ga.GetAccountAddress()
	if err != nil {
		return PoolArrayResponse{}, err
	}

	pools, err := market.GetOwnedPools(&bind.CallOpts{From: *address}, *address)
	if err != nil {
		return PoolArrayResponse{}, err
	}

	return MarketPoolAddressesToArrayResponse(pools, includeData, ga)
}

type PoolArrayResponse struct {
	Pools []PoolResponse `json:"pools"`
}

type PoolResponse struct {
	Address string          `json:"address"`
	Data    *PoolPublicData `json:"data,omitempty"`
}

func (d *PoolResponse) String() string {
	jsonResponse, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(jsonResponse)
}

func MarketPools(includeData bool, ga *GladiusAccountManager) (PoolArrayResponse, error) {
	market := ConnectMarket()

	poolAddresses, err := market.GetAllPools(nil)
	if err != nil {
		return PoolArrayResponse{}, err
	}

	return MarketPoolAddressesToArrayResponse(poolAddresses, includeData, ga)
}

func MarketPoolAddressesToArrayResponse(poolAddresses []common.Address, includeData bool, ga *GladiusAccountManager) (PoolArrayResponse, error) {
	var pools PoolArrayResponse

	for _, poolAddress := range poolAddresses {
		var poolResponse PoolResponse
		if includeData {
			poolData, err := PoolRetrievePublicData(poolAddress.String(), ga)
			poolResponse = PoolResponse{poolAddress.String(), poolData}
			if err != nil {
				return PoolArrayResponse{}, err
			}
		} else {
			poolResponse = PoolResponse{poolAddress.String(), nil}
		}
		pools.Pools = append(pools.Pools, poolResponse)
	}

	return pools, nil
}

//MarketCreatePool - Create new pool
func MarketCreatePool(passphrase, publicKey string, ga *GladiusAccountManager) (*types.Transaction, error) {
	market := ConnectMarket()

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
