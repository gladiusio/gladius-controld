package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nfeld9807/rest-api/internal/blockchain/generated"
	"log"
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

//func NodeRetrieveData() (*NodeData, error) {
//nodeAddress, _ := NodeOwnedByUser()
//node := ConnectNode(*nodeAddress)

//encData, _ := node.Data(&bind.CallOpts{From: GetDefaultAccountAddress()})
//data, _ := crypto.DecryptData(encData)

//dataReader := strings.NewReader(data)
//decoder := json.NewDecoder(dataReader)
//var nodeData NodeData
//decoder.Decode(&nodeData)
//return &nodeData, nil
//}

//func NodeSetData(passphrase string, data *NodeData) (*types.Transaction, error) {
//nodeAddress, _ := NodeOwnedByUser()
//node := ConnectNode(*nodeAddress)

//encData, err := crypto.EncryptData(data.String())

//transaction, err := node.SetData(GetDefaultAuth(passphrase), encData)

//if err != nil {
//return nil, err
//}

//return transaction, nil
//}
