package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
