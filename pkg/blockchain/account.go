package blockchain

import (
	"errors"
	"io/ioutil"
	"strings"

	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

// GladiusAccountManager is a type that allows the user to create a keystore file,
// create an in it, and preform actions on the first account stored.
type GladiusAccountManager struct {
	keystore *keystore.KeyStore
}

// NewGladiusAccountManager creates a new gladius account manager
func NewGladiusAccountManager() *GladiusAccountManager {
	var pathTemp = viper.GetString("DirWallet")

	ks := keystore.NewKeyStore(
		pathTemp,
		keystore.LightScryptN,
		keystore.LightScryptP)

	return &GladiusAccountManager{keystore: ks}
}

// Keystore gets the keystore associated with the account manager
func (ga GladiusAccountManager) Keystore() *keystore.KeyStore {
	return ga.keystore
}

//UnlockAccount Unlocks the account
func (ga GladiusAccountManager) UnlockAccount(passphrase string) (bool, error) {
	account, err := ga.GetAccount()
	if err != nil {
		return false, err
	}

	err = ga.Keystore().Unlock(*account, passphrase)
	if err == nil {
		return true, nil
	}

	return false, err
}

// CreateAccount will create an account if there isn't one already
func (ga GladiusAccountManager) CreateAccount(passphrase string) (accounts.Account, error) {
	ks := ga.Keystore()
	if len(ga.Keystore().Accounts()) < 1 {
		return ks.NewAccount(passphrase)
	}
	return accounts.Account{}, errors.New("gladius account already exists")

}

// GetAccountAddress gets the account address
func (ga GladiusAccountManager) GetAccountAddress() (*common.Address, error) {
	account, err := ga.GetAccount()
	if err != nil {
		return nil, err
	}

	return &account.Address, nil
}

// GetAccount gets the actual account type
func (ga GladiusAccountManager) GetAccount() (*accounts.Account, error) {
	store := ga.Keystore()
	if len(store.Accounts()) < 1 {
		return nil, errors.New("account retrieval error, no existing accounts found")
	}

	account := store.Accounts()[0]

	return &account, nil
}

type BalanceType int32

const (
	ETH BalanceType = 0
	GLA BalanceType = 1
)

type Balance struct {
	Value  uint64 `json:"value"`
	Symbol string `json:"symbol"`
}

func GetAccountBalance(address common.Address, symbol BalanceType) (Balance, error) {
	var resp *http.Response
	var err error
	var symbolString string

	glaTokenAddress := "0x972c1e9698b218acc3e7869c1ccfefd3808bdbb2"

	switch symbol {
	case ETH:
		resp, err = http.Get("https://api-ropsten.etherscan.io/api?module=account&action=balance&address=" + address.String() + "&tag=latest&apikey=3VRW685YYESSYIFVND3DVN9ZNF4BTT1GB8")
		symbolString = "ETH"
		break
	case GLA:
		// https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=0x57d90b64a1a57749b0f932f1a3395792e12e7055&address=0xe04f27eb70e025b78871a2ad7eabe85e61212761&tag=latest&apikey=YourApiKeyToken
		resp, err = http.Get("https://api-ropsten.etherscan.io/api?module=account&action=tokenbalance&contractaddress=" + glaTokenAddress + "&address=" + address.String() + "&apikey=3VRW685YYESSYIFVND3DVN9ZNF4BTT1GB8")
		symbolString = "GLA"
		break
	}

	if err != nil {
		return Balance{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	type etherscanResult struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}

	result := etherscanResult{}
	json.Unmarshal(body, &result)

	balanceInt, err := strconv.ParseUint(result.Result, 10, 64)
	if err != nil {
		return Balance{}, err
	}

	balance := Balance{Value: balanceInt, Symbol: symbolString}

	return balance, nil
}

type TransactionOptions struct {
	Filters *TransactionFilter `json:"filters"`
}

type TransactionFilter struct {
	EthTransfer bool `json:"eth_transfer"`
}

type EtherscanTransactionsResponse struct {
	Status       string                 `json:"status"`
	Message      string                 `json:"message"`
	Transactions []EtherscanTransaction `json:"result"`
}

type EtherscanTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}

func GetAccountTransactions(address common.Address, options TransactionOptions) (EtherscanTransactionsResponse, error) {
	resp, err := http.Get("https://api-ropsten.etherscan.io/api?module=account&action=txlist&address=" + address.String() + "&startblock=0&endblock=latest&sort=asc&apikey=3VRW685YYESSYIFVND3DVN9ZNF4BTT1GB8")

	if err != nil {
		return EtherscanTransactionsResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	result := EtherscanTransactionsResponse{}
	json.Unmarshal(body, &result)
	var filteredTransactions []EtherscanTransaction

	if options.Filters != nil {
		for _, transaction := range result.Transactions {
			if options.Filters.EthTransfer && transaction.Value != "0" {
				filteredTransactions = append(filteredTransactions, transaction)
			} else if !options.Filters.EthTransfer && transaction.Value == "0" {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}

		result.Transactions = filteredTransactions
	}

	return result, nil
}

// GetAuth gets the authenticator for the go bindings of our smart contracts
func (ga GladiusAccountManager) GetAuth(passphrase string) (*bind.TransactOpts, error) {
	account, err := ga.GetAccount()
	if err != nil {
		return nil, err
	}
	// Create a JSON blob with the same passphrase used to decrypt it
	key, err := ga.Keystore().Export(*account, passphrase, passphrase)
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
	var pathTemp = viper.GetString("DirKeys")
	keyringFileBuffer, err := ioutil.ReadFile(pathTemp + "/public.asc")
	if err != nil {
		return "", err
	}

	publicKey := string(keyringFileBuffer)

	return publicKey, nil
}
