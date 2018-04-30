package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nfeld9807/rest-api/internal/blockchain/generated"
)

// ConnectNodeFactory - Establish Connection to deployed NodeFactory contract
func ConnectNodeFactory() (*generated.NodeFactory, error) {
	conn := ConnectClient()

	nodeFactory, err := generated.NewNodeFactory(common.HexToAddress("0x85f0129d0b40b0ed15d97b657872b55cf91ae7de"), conn)

	if err != nil {
		return nil, err
	}

	return nodeFactory, nil
}

// NodeForAccount - returns node address for wallet
func NodeForAccount(ownerAddress string) (*common.Address, error) {
	factory, err := ConnectNodeFactory()

	if err != nil {
		return nil, err
	}

	address, err := factory.GetNodeAddress(&bind.CallOpts{From: common.HexToAddress(ownerAddress)})

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func CreateNode() (string, error) {
	factory, err := ConnectNodeFactory()

	if err != nil {
		return "null", err
	}

	auth := GetAuth("password")

	transaction, err := factory.CreateNode(auth)

	if err != nil {
		return "null", err
	}

	txHash := fmt.Sprintf("0x%x", transaction.Hash())

	return txHash, nil
}
