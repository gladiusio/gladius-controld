package blockchain

import (
	"encoding/json"
	"log"
	"math/big"
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

type Node struct {
	Address string   `json:"address"`
	Status  int      `json:"status"`
	Data    NodeData `json:"data"`
}

func (d *Node) String() string {
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

func Node(nodeAddress string) (*Node, error) {
	node := ConnectNode(*nodeAddress)
}

func NodeRetrieveData() (*NodeData, error) {
	nodeAddress, _ := NodeOwnedByUser()
	node := ConnectNode(*nodeAddress)

	encData, _ := node.Data(&bind.CallOpts{From: GetDefaultAccountAddress()})
	data, _ := crypto.DecryptData(encData)

	dataReader := strings.NewReader(data)
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

	auth, err := GetDefaultAuth(passphrase)
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
	println(data.String())
	if err != nil {
		return nil, err
	}

	poolPubKey, err := PoolRetrievePublicKey(poolAddress)
	println(poolPubKey)
	if err != nil {
		return nil, err
	}

	encData, err := crypto.EncryptMessage(data.String(), poolPubKey)
	if err != nil {
		return nil, err
	}

	auth, err := GetDefaultAuth(passphrase)
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
	node := ConnectNode(common.HexToAddress(nodeAddress))
	statusCode, err := node.GetStatus(&bind.CallOpts{From: GetDefaultAccountAddress()}, common.HexToAddress(poolAddress))

	if err != nil {
		return nil, err
	}

	return statusCode, nil
}
