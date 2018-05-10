package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
	"github.com/spf13/viper"
)

// ConnectNodeFactory - Establish Connection to deployed NodeFactory contract
func ConnectNodeFactory() (*generated.NodeFactory, error) {
	conn := ConnectClient()

	nodeFactoryAddress := viper.GetString("BlockchainNodeFactoryAddress")
	nodeFactory, err := generated.NewNodeFactory(common.HexToAddress(nodeFactoryAddress), conn)

	if err != nil {
		return nil, err
	}

	return nodeFactory, nil
}

// NodeForAccount - returns node address for wallet
func NodeOwnedByUser() (*common.Address, error) {
	address := GetDefaultAccountAddress()
	return NodeForAccount(address)
}

func NodeForAccount(ownerAddress common.Address) (*common.Address, error) {
	factory, err := ConnectNodeFactory()

	if err != nil {
		return nil, err
	}

	address, err := factory.GetNodeAddress(&bind.CallOpts{From: ownerAddress})

	if err != nil {
		return nil, err
	}

	return &address, nil
}

func CreateNode(passphrase string) (*types.Transaction, error) {
	factory, err := ConnectNodeFactory()

	if err != nil {
		return nil, err
	}

	auth := GetDefaultAuth(passphrase)

	transaction, err := factory.CreateNode(auth)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
