package blockchain

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Tx - Get Transaction
func Tx(txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	conn := ConnectClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return conn.TransactionByHash(ctx, txHash)
}

// TxReceipt - Get Receipt of Transaction
func TxReceipt(txHash common.Hash) (*types.Receipt, error) {
	conn := ConnectClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return conn.TransactionReceipt(ctx, txHash)
}

func ApplicationStatusFromString(status string) (int, error) {
	switch status {
	case "unavailable":
		// Unavailable
		return 0, nil
	case "approved", "approve":
		// Approved
		return 1, nil
	case "rejected", "reject":
		// Rejected
		return 2, nil
	case "pending":
		// Pending
		return 3, nil
	}
	return -1, errors.New("Status Not Found for: " + status)
}
