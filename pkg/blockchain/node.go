package blockchain

import (
	"encoding/json"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

type NodeResponse struct {
	Address string    `json:"address"`
	Data    *NodeData `json:"data"`
}

type NodeData struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	IP     string `json:"ip"`
	Status string `json:"status"`
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

func NodeRetrievePoolData(nodeAddress, poolAddress *common.Address) (*NodeData, error) {
	node := ConnectNode(*nodeAddress)
	ga := NewGladiusAccountManager()
	address, err := ga.GetAccountAddress()
	if err != nil {
		return nil, err
	}

	encPoolData, err := node.GetPoolData(&bind.CallOpts{From: *address}, *poolAddress)
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

func NodeApplicationStatus(nodeAddress, poolAddress string) (*big.Int, error) {
	parsedNodeAddress := common.HexToAddress(nodeAddress)
	node := ConnectNode(parsedNodeAddress)
	statusCode, err := node.GetStatus(&bind.CallOpts{From: common.HexToAddress(poolAddress)}, common.HexToAddress(poolAddress))
	if err != nil {
		return nil, err
	}

	return statusCode, nil
}