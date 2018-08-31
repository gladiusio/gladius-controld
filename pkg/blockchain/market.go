package blockchain

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-application-server/pkg/db/models"
	"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
	"github.com/gladiusio/gladius-controld/pkg/routing/response"
	"github.com/gladiusio/gladius-controld/pkg/utils"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// ConnectMarket - Connect and return configured market
func ConnectMarket() *generated.Market {

	conn := ConnectClient()

	marketAddress := viper.GetString("blockchain.marketAddress")
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

	pools, err := market.GetOwnerAllPools(&bind.CallOpts{From: *address}, *address)
	if err != nil {
		return PoolArrayResponse{}, err
	}

	return MarketPoolAddressesToArrayResponse(pools, includeData, ga)
}

type PoolArrayResponse struct {
	Pools []PoolResponse `json:"pools"`
}

type PoolResponse struct {
	Address string                 `json:"address"`
	Url     string                 `json:"url,omitempty"`
	Data    models.PoolInformation `json:"data,omitempty"`
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
			poolUrl, err := PoolRetrieveApplicationServerUrl(poolAddress.String(), ga)

			poolInformationResponse, err := utils.SendRequest(http.MethodGet, poolUrl+"server/info", nil)
			var defaultResponse response.DefaultResponse
			json.Unmarshal([]byte(poolInformationResponse), &defaultResponse)

			var poolInformation models.PoolInformation
			poolInfoByteArray, err := json.Marshal(defaultResponse.Response)
			json.Unmarshal(poolInfoByteArray, &poolInformation)
			poolInformation.Url = poolUrl

			poolResponse = PoolResponse{poolAddress.String(), poolUrl, poolInformation}
			if err != nil {
				return PoolArrayResponse{}, err
			}
		} else {
			poolResponse = PoolResponse{poolAddress.String(), "", models.PoolInformation{}}
		}
		pools.Pools = append(pools.Pools, poolResponse)
	}

	return pools, nil
}

//MarketCreatePool - Create new pool
func MarketCreatePool(passphrase string, ga *GladiusAccountManager) (*types.Transaction, error) {
	market := ConnectMarket()

	auth, err := ga.GetAuth(passphrase)
	if err != nil {
		return nil, err
	}

	transaction, err := market.CreatePool(auth)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}