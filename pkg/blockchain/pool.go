package blockchain

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
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

func PoolRetrievePublicKey(poolAddress string) (string, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	publicKey, err := pool.PublicKey(&bind.CallOpts{From: GetDefaultAccountAddress()})
	if err != nil {
		return "null", nil
	}

	return publicKey, nil
}

type PoolPublicData struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Rating       string `json:"rating"`
	NodeCount    string `json:"nodeCount"`
	MaxBandwidth string `json:"maxBandwidth"`
}

func (d *PoolPublicData) String() string {
	json, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(json)
}

func PoolRetrievePublicData(poolAddress string) (*PoolPublicData, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	publicDataResponse, err := pool.PublicData(&bind.CallOpts{From: GetDefaultAccountAddress()})
	if err != nil {
		return nil, err
	}

	dataReader := strings.NewReader(publicDataResponse)
	decoder := json.NewDecoder(dataReader)
	var poolPublicData PoolPublicData
	decoder.Decode(&poolPublicData)
	return &poolPublicData, nil
}

func PoolSetPublicData(passphrase, poolAddress, data string) (*types.Transaction, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))

	auth, err := GetDefaultAuth(passphrase)
	if err != nil {
		return nil, err
	}

	transaction, err := pool.SetPublicData(auth, data)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func Nodes(poolAddress string) (*[]common.Address, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	nodeAddressList, err := pool.GetNodeList(&bind.CallOpts{From: GetDefaultAccountAddress()})
	if err != nil {
		return nil, err
	}
	return &nodeAddressList, nil
}

//
// func NodesByStatus(poolAddress string, status int) (*[]generated.Node, error) {}
