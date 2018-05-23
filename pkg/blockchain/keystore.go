package blockchain

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

func keystoreForPath(path string) (*keystore.KeyStore, error) {
	var pathTemp string = viper.GetString("DirWallet")
	ks := keystore.NewKeyStore(
		pathTemp,
		keystore.LightScryptN,
		keystore.LightScryptP)

	return ks, nil
}

func AccountResponseFormatter(account *accounts.Account) string {
	address := account.Address
	accountAddress := fmt.Sprintf("0x%x", address)

	return "{ \"address\": \"" + accountAddress + "\"}"
}

func WalletResponseFormatter(wallet accounts.Wallet) string {
	walletStatus, _ := wallet.Status()
	walletAddress := wallet.Accounts()[0].Address

	return "{ \"status\": \"" + walletStatus + "\", \"address\": \"" + walletAddress.String() + "\"}"
}

func CreateAccount(passphrase string) (accounts.Account, error) {
	ks, _ := keystoreForPath("")
	return ks.NewAccount(passphrase)
}

func Wallets() []accounts.Wallet {
	ks, _ := keystoreForPath("")

	return ks.Wallets()
}

func OpenWallet(accountIndex int, passphrase string) accounts.Wallet {
	ks, _ := keystoreForPath("")
	account := ks.Accounts()[accountIndex]
	err := ks.Unlock(account, passphrase)

	wallet := ks.Wallets()[accountIndex]
	wallet.Open(passphrase)

	if err != nil {
		log.Fatal(err)
	}

	return wallet
}

func CloseWallet(accountIndex int) {
	ks, _ := keystoreForPath("")
	wallet := ks.Wallets()[accountIndex]
	wallet.Close()
}

func GetPGPPublicKey() (string, error) {
	var pathTemp string = viper.GetString("DirKeys")
	keyringFileBuffer, err := ioutil.ReadFile(pathTemp + "/public.asc")
	if err != nil {
		return "", err
	}

	publicKey := string(keyringFileBuffer)

	return publicKey, nil
}

func GetDefaultAccountAddress() common.Address {
	return GetAccountAddress(0)
}

func GetAccountAddress(index int) common.Address {
	ks, _ := keystoreForPath("")
	wallet := ks.Wallets()[index]
	return wallet.Accounts()[index].Address
}

func GetDefaultAuth(passphrase string) (*bind.TransactOpts, error) {
	return GetAuth(passphrase, 0)
}

// GetAuth - Temporary auth retrieval
func GetAuth(passphrase string, index int) (*bind.TransactOpts, error) {
	wallet := Wallets()[index]

	key, err := ioutil.ReadFile(wallet.URL().Path)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewTransactor(strings.NewReader(string(key)), passphrase)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
