package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
