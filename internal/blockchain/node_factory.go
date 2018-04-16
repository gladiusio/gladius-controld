package blockchain

import (
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

// NodeForAddress - returns node address for wallet
//func NodeForAddress(hash common.Hash) *generated.Node {
//factory := ConnectNodeFactory()

//}
