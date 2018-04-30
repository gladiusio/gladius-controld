package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateKeyStore(passphrase string) (string, error) {
	var path string = os.Getenv("HOME") + "/.config/gladius/wallet"
	ks := keystore.NewKeyStore(
		path,
		keystore.LightScryptN,
		keystore.LightScryptP)

	ks.NewAccount(passphrase)

	address := ks.Accounts()[0].Address
	accountAddress := fmt.Sprintf("0x%x", address)

	accountPath := ks.Accounts()[0].URL.Path

	response := "{ \"address\": \"" + accountAddress + "\", \"path\": \"" + accountPath + "\"}"

	return response, nil
}

func UnlockAccount(passphrase string) {

}

// GetAuth - Temporary auth retrieval
func GetAuth(passphrase string) *bind.TransactOpts {

	key, _ := ioutil.ReadFile("/Users/nate/.config/gladius/wallet/UTC--2018-04-30T19-38-05.856804037Z--759eef8b0c929c452f710e07ae5d92988e8698f0")

	auth, err := bind.NewTransactor(strings.NewReader(string(key)), passphrase)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	return auth
}
