package blockchain

import (
	"encoding/json"
	"errors"
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
	ga := NewGladiusAccountManager()

	publicKey, err := pool.PublicKey(&bind.CallOpts{From: ga.GetAccountAddress()})
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
	ga := NewGladiusAccountManager()

	publicDataResponse, err := pool.PublicData(&bind.CallOpts{From: ga.GetAccountAddress()})
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
	ga := NewGladiusAccountManager()

	auth, err := ga.GetAuth(passphrase)
	if err != nil {
		return nil, err
	}

	transaction, err := pool.SetPublicData(auth, data)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func PoolNodes(poolAddress string) (*[]common.Address, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	ga := NewGladiusAccountManager()

	nodeAddressList, err := pool.GetNodeList(&bind.CallOpts{From: ga.GetAccountAddress()})
	if err != nil {
		return nil, err
	}
	return &nodeAddressList, nil
}

func PoolNodesWithData(poolAddress common.Address, nodeAddresses *[]common.Address, status int) (*string, error) {
	var filter = status >= 0

	response := "["

	for _, nodeAddress := range *nodeAddresses {
		nodeApplication, err := NodeRetrieveApplication(&nodeAddress, &poolAddress)
		if err != nil {
			return nil, err
		}

		if filter && nodeApplication.Status == status {
			response += nodeApplication.String() + ","
		} else if !filter {
			response += nodeApplication.String() + ","
		}
	}
	response = strings.TrimRight(response, ",")
	response += "]"

	return &response, nil
}

func PoolUpdateNodeStatus(passphrase, poolAddress, nodeAddress string, status int) (*types.Transaction, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	var err error
	ga := NewGladiusAccountManager()

	auth, err := ga.GetAuth(passphrase)
	if err != nil {
		return nil, err
	}

	var transaction *types.Transaction

	switch status {
	case 0:
		// Unavailable
		err = errors.New("PoolUpdateNodeStatus - Node cannot change status to `Unavailable`")
	case 1:
		// Approved
		println("aprrove")
		transaction, err = pool.AcceptNode(auth, common.HexToAddress(nodeAddress))
	case 2:
		// Rejected
		println("reject")
		transaction, err = pool.RejectNode(auth, common.HexToAddress(nodeAddress))
	case 3:
		// Pending
		err = errors.New("PoolUpdateNodeStatus - Node cannot change status to `Pending`")
	}

	if err != nil {
		return nil, err
	}

	return transaction, nil
}