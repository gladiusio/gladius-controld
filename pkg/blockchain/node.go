package blockchain

import (
	"encoding/json"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
	"github.com/gladiusio/gladius-controld/pkg/crypto"
)

// ConnectNode - Connect and grab node
func ConnectNode(nodeAddress common.Address) *generated.Node {

	conn := ConnectClient()

	node, err := generated.NewNode(nodeAddress, conn)

	if err != nil {
		log.Fatalf("Failed to instantiate a Node contract: %v", err)
	}

	return node
}

type NodeApplication struct {
	Address string   `json:"address"`
	Status  int      `json:"status"`
	Data    NodeData `json:"data"`
}

func (d *NodeApplication) String() string {
	json, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(json)
}

type NodeData struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	IP     string `json:"ip"`
	Status string `json:"status"`
}

func (d *NodeData) String() string {
	json, err := json.Marshal(d)
	if err != nil {
		return "{}"
	}

	return string(json)
}

func NodeRetrieveApplication(nodeAddress, poolAddress *common.Address) (*NodeApplication, error) {
	nodeData, err := NodeRetrievePoolData(nodeAddress, poolAddress)
	if err != nil {
		return nil, err
	}

	nodeStatus, err := NodeApplicationStatus(nodeAddress.String(), poolAddress.String())
	if err != nil {
		return nil, err
	}

	statusInt, err := strconv.Atoi(nodeStatus.String())
	if err != nil {
		return nil, err
	}

	nodeStruct := &NodeApplication{
		Address: nodeAddress.String(),
		Status:  statusInt,
		Data:    *nodeData,
	}

	return nodeStruct, nil
}

func NodeRetrieveData() (*NodeData, error) {
	nodeAddress, _ := NodeOwnedByUser()
	node := ConnectNode(*nodeAddress)
	ga := NewGladiusAccountManager()

	encData, err := node.Data(&bind.CallOpts{From: ga.GetAccountAddress()})
	if err != nil {
		return nil, err
	}

	data, err := crypto.DecryptData(encData)
	if err != nil {
		return nil, err
	}

	dataReader := strings.NewReader(data)
	decoder := json.NewDecoder(dataReader)
	var nodeData NodeData
	decoder.Decode(&nodeData)
	return &nodeData, nil
}

func NodeRetrievePoolData(nodeAddress, poolAddress *common.Address) (*NodeData, error) {
	node := ConnectNode(*nodeAddress)
	ga := NewGladiusAccountManager()

	encPoolData, err := node.GetPoolData(&bind.CallOpts{From: ga.GetAccountAddress()}, *poolAddress)
	if err != nil {
		return nil, err
	}

	poolApplication, err := crypto.DecryptData(encPoolData)
	if err != nil {
		return nil, err
	}

	dataReader := strings.NewReader(poolApplication)
	decoder := json.NewDecoder(dataReader)

	var nodeData NodeData
	decoder.Decode(&nodeData)

	return &nodeData, nil
}

func NodeSetData(passphrase string, data *NodeData) (*types.Transaction, error) {
	nodeAddress, _ := NodeOwnedByUser()
	node := ConnectNode(*nodeAddress)

	encData, err := crypto.EncryptData(data.String())
	if err != nil {
		return nil, err
	}

	ga := NewGladiusAccountManager()

	auth, err := ga.GetAuth(passphrase)
	if err != nil {
		return nil, err
	}

	transaction, err := node.SetData(auth, encData)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func NodeApplyToPool(passphrase, nodeAddress, poolAddress string) (*types.Transaction, error) {
	node := ConnectNode(common.HexToAddress(nodeAddress))

	data, err := NodeRetrieveData()

	if err != nil {
		return nil, err
	}

	poolPubKey, err := PoolRetrievePublicKey(poolAddress)

	if err != nil {
		return nil, err
	}

	encData, err := crypto.EncryptMessage(data.String(), poolPubKey)
	if err != nil {
		return nil, err
	}

	ga := NewGladiusAccountManager()

	auth, err := ga.GetAuth(passphrase)
	if err != nil {
		return nil, err
	}

	transaction, err := node.ApplyToPool(auth, common.HexToAddress(poolAddress), encData)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func NodeApplicationStatus(nodeAddress, poolAddress string) (*big.Int, error) {
	parsedNodeAddress := common.HexToAddress(nodeAddress)
	node := ConnectNode(parsedNodeAddress)
	statusCode, err := node.GetStatus(&bind.CallOpts{From: common.HexToAddress(poolAddress)}, common.HexToAddress(poolAddress))
	if err != nil {
		return nil, err
	}

	return statusCode, nil
}
