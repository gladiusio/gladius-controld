package blockchain

import (

"encoding/json"
"github.com/ethereum/go-ethereum/accounts/abi/bind"
"github.com/ethereum/go-ethereum/common"
"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
"log"
"strings"
)

// ConnectNode - Connect and grab node
func ConnectPool(poolAddress common.Address) *generated.Pool {
	conn := ConnectClient()
	pool, err := generated.NewPool(poolAddress, conn)

	if err != nil {
		log.Fatalf("Failed to instantiate a Node contract: %v", err)
	}

	return pool
}

type PoolPublicData struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Rating       string `json:"rating"`
	NodeCount    string `json:"nodeCount"`
	MaxBandwidth string `json:"maxBandwidth"`
	URL          string `json:"url"`
}

func PoolRetrievePublicData(poolAddress string, ga *GladiusAccountManager) (*PoolPublicData, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	address, err := ga.GetAccountAddress()
	if err != nil {
		return nil, err
	}

	publicDataResponse, err := pool.PublicData(&bind.CallOpts{From: *address})
	if err != nil {
		return nil, err
	}

	dataReader := strings.NewReader(publicDataResponse)
	decoder := json.NewDecoder(dataReader)
	var poolPublicData PoolPublicData
	decoder.Decode(&poolPublicData)
	return &poolPublicData, nil
}