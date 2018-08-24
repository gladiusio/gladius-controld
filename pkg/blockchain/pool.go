package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gladiusio/gladius-controld/pkg/blockchain/generated"
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

func PoolRetrieveApplicationServerUrl(poolAddress string, ga *GladiusAccountManager) (string, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	address, err := ga.GetAccountAddress()
	if err != nil {
		return "", err
	}

	url, err := pool.GetUrl(&bind.CallOpts{From: *address})
	if err != nil {
		return "", err
	}

	return url, nil
}

func PoolSetApplicationServerUrl(passphrase, poolAddress, url string) (*types.Transaction, error) {
	pool := ConnectPool(common.HexToAddress(poolAddress))
	ga := NewGladiusAccountManager()

	auth, err := ga.GetAuth(passphrase)
	if err != nil {
		return nil, err
	}

	transaction, err := pool.SetUrl(auth, url)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
