package blockchain

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// GladiusAccountManager is a type that allows the user to create a keystore file,
// create an in it, and preform actions on the first account stored.
type GladiusAccountManager struct {
	keystore *keystore.KeyStore
}

func NewGladiusAccountManager() *GladiusAccountManager {
	var pathTemp string = viper.GetString("DirWallet")
	ks := keystore.NewKeyStore(
		pathTemp,
		keystore.LightScryptN,
		keystore.LightScryptP)

	return &GladiusAccountManager{keystore: ks}
}

func (ga GladiusAccountManager) Keystore() *keystore.KeyStore {
	return ga.keystore
}

func (ga GladiusAccountManager) UnlockAccount(passphrase string) error {
	return ga.Keystore().Unlock(ga.GetAccount(), passphrase)
}

func (ga GladiusAccountManager) AccountResponseFormatter() string {
	address := ga.GetAccountAddress()
	accountAddress := fmt.Sprintf("0x%x", address)

	return "{ \"address\": \"" + accountAddress + "\"}"
}

func (ga GladiusAccountManager) CreateAccount(passphrase string) (accounts.Account, error) {
	ks := ga.Keystore()
	if len(ga.Keystore().Accounts()) < 1 {
		return ks.NewAccount(passphrase)
	}
	return accounts.Account{}, errors.New("gladius account already exists")

}

func (ga GladiusAccountManager) GetAccountAddress() common.Address {
	return ga.GetAccount().Address
}

func (ga GladiusAccountManager) GetAccount() accounts.Account {
	keystore := ga.Keystore()

	return keystore.Accounts()[0]
}

func (ga GladiusAccountManager) GetAuth(passphrase string) (*bind.TransactOpts, error) {
	// Create a JSON blob with the same passphrase used to decrypt it
	key, err := ga.Keystore().Export(ga.GetAccount(), passphrase, passphrase)
	if err != nil {
		return nil, err
	}

	// Create a transactor from the key file
	auth, err := bind.NewTransactor(strings.NewReader(string(key)), passphrase)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

// TODO: Move somewhere more logical...
func GetPGPPublicKey() (string, error) {
	var pathTemp string = viper.GetString("DirKeys")
	keyringFileBuffer, err := ioutil.ReadFile(pathTemp + "/public.asc")
	if err != nil {
		return "", err
	}

	publicKey := string(keyringFileBuffer)

	return publicKey, nil
}
