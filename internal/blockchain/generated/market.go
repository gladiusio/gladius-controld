// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MarketABI is the input ABI used to generate the binding from.
const MarketABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"owed\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint256\"},{\"name\":\"completed\",\"type\":\"uint256\"},{\"name\":\"paid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allocateFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"work\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"work\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPoolsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payouts\",\"outputs\":[{\"name\":\"pool\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"owed\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint256\"},{\"name\":\"completed\",\"type\":\"uint256\"},{\"name\":\"paid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allocateFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"clientToPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownedPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketPoolsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_gladiusToken\",\"type\":\"address\"},{\"name\":\"_joinCost\",\"type\":\"uint256\"},{\"name\":\"_maxPayout\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicKey\",\"type\":\"string\"}],\"name\":\"createPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"getOwnedPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"}],\"name\":\"getClientPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_node\",\"type\":\"address\"},{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logWorkFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_node\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"joinMarketplace\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allocateClientFundsTo\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarketPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// AllPoolsList is a free data retrieval call binding the contract method 0x894958aa.
//
// Solidity: function allPoolsList( uint256) constant returns(address)
func (_Market *MarketCaller) AllPoolsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "allPoolsList", arg0)
	return *ret0, err
}

// AllPoolsList is a free data retrieval call binding the contract method 0x894958aa.
//
// Solidity: function allPoolsList( uint256) constant returns(address)
func (_Market *MarketSession) AllPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.AllPoolsList(&_Market.CallOpts, arg0)
}

// AllPoolsList is a free data retrieval call binding the contract method 0x894958aa.
//
// Solidity: function allPoolsList( uint256) constant returns(address)
func (_Market *MarketCallerSession) AllPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.AllPoolsList(&_Market.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(owed uint256, total uint256, completed uint256, paid uint256)
func (_Market *MarketCaller) Balance(opts *bind.CallOpts) (struct {
	Owed      *big.Int
	Total     *big.Int
	Completed *big.Int
	Paid      *big.Int
}, error) {
	ret := new(struct {
		Owed      *big.Int
		Total     *big.Int
		Completed *big.Int
		Paid      *big.Int
	})
	out := ret
	err := _Market.contract.Call(opts, out, "balance")
	return *ret, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(owed uint256, total uint256, completed uint256, paid uint256)
func (_Market *MarketSession) Balance() (struct {
	Owed      *big.Int
	Total     *big.Int
	Completed *big.Int
	Paid      *big.Int
}, error) {
	return _Market.Contract.Balance(&_Market.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(owed uint256, total uint256, completed uint256, paid uint256)
func (_Market *MarketCallerSession) Balance() (struct {
	Owed      *big.Int
	Total     *big.Int
	Completed *big.Int
	Paid      *big.Int
}, error) {
	return _Market.Contract.Balance(&_Market.CallOpts)
}

// ClientToPool is a free data retrieval call binding the contract method 0xcfa4cc4f.
//
// Solidity: function clientToPool( address) constant returns(address)
func (_Market *MarketCaller) ClientToPool(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "clientToPool", arg0)
	return *ret0, err
}

// ClientToPool is a free data retrieval call binding the contract method 0xcfa4cc4f.
//
// Solidity: function clientToPool( address) constant returns(address)
func (_Market *MarketSession) ClientToPool(arg0 common.Address) (common.Address, error) {
	return _Market.Contract.ClientToPool(&_Market.CallOpts, arg0)
}

// ClientToPool is a free data retrieval call binding the contract method 0xcfa4cc4f.
//
// Solidity: function clientToPool( address) constant returns(address)
func (_Market *MarketCallerSession) ClientToPool(arg0 common.Address) (common.Address, error) {
	return _Market.Contract.ClientToPool(&_Market.CallOpts, arg0)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_Market *MarketCaller) GetAllPools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getAllPools")
	return *ret0, err
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_Market *MarketSession) GetAllPools() ([]common.Address, error) {
	return _Market.Contract.GetAllPools(&_Market.CallOpts)
}

// GetAllPools is a free data retrieval call binding the contract method 0xd88ff1f4.
//
// Solidity: function getAllPools() constant returns(address[])
func (_Market *MarketCallerSession) GetAllPools() ([]common.Address, error) {
	return _Market.Contract.GetAllPools(&_Market.CallOpts)
}

// GetClientPool is a free data retrieval call binding the contract method 0x2db5b54e.
//
// Solidity: function getClientPool(_client address) constant returns(address)
func (_Market *MarketCaller) GetClientPool(opts *bind.CallOpts, _client common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getClientPool", _client)
	return *ret0, err
}

// GetClientPool is a free data retrieval call binding the contract method 0x2db5b54e.
//
// Solidity: function getClientPool(_client address) constant returns(address)
func (_Market *MarketSession) GetClientPool(_client common.Address) (common.Address, error) {
	return _Market.Contract.GetClientPool(&_Market.CallOpts, _client)
}

// GetClientPool is a free data retrieval call binding the contract method 0x2db5b54e.
//
// Solidity: function getClientPool(_client address) constant returns(address)
func (_Market *MarketCallerSession) GetClientPool(_client common.Address) (common.Address, error) {
	return _Market.Contract.GetClientPool(&_Market.CallOpts, _client)
}

// GetMarketPools is a free data retrieval call binding the contract method 0xb42cb6f5.
//
// Solidity: function getMarketPools() constant returns(address[])
func (_Market *MarketCaller) GetMarketPools(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getMarketPools")
	return *ret0, err
}

// GetMarketPools is a free data retrieval call binding the contract method 0xb42cb6f5.
//
// Solidity: function getMarketPools() constant returns(address[])
func (_Market *MarketSession) GetMarketPools() ([]common.Address, error) {
	return _Market.Contract.GetMarketPools(&_Market.CallOpts)
}

// GetMarketPools is a free data retrieval call binding the contract method 0xb42cb6f5.
//
// Solidity: function getMarketPools() constant returns(address[])
func (_Market *MarketCallerSession) GetMarketPools() ([]common.Address, error) {
	return _Market.Contract.GetMarketPools(&_Market.CallOpts)
}

// GetOwnedPools is a free data retrieval call binding the contract method 0x33e7f452.
//
// Solidity: function getOwnedPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketCaller) GetOwnedPools(opts *bind.CallOpts, _ownerAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getOwnedPools", _ownerAddress)
	return *ret0, err
}

// GetOwnedPools is a free data retrieval call binding the contract method 0x33e7f452.
//
// Solidity: function getOwnedPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketSession) GetOwnedPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnedPools(&_Market.CallOpts, _ownerAddress)
}

// GetOwnedPools is a free data retrieval call binding the contract method 0x33e7f452.
//
// Solidity: function getOwnedPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketCallerSession) GetOwnedPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnedPools(&_Market.CallOpts, _ownerAddress)
}

// MarketPools is a free data retrieval call binding the contract method 0x86bf8056.
//
// Solidity: function marketPools( address,  uint256) constant returns(address)
func (_Market *MarketCaller) MarketPools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "marketPools", arg0, arg1)
	return *ret0, err
}

// MarketPools is a free data retrieval call binding the contract method 0x86bf8056.
//
// Solidity: function marketPools( address,  uint256) constant returns(address)
func (_Market *MarketSession) MarketPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPools(&_Market.CallOpts, arg0, arg1)
}

// MarketPools is a free data retrieval call binding the contract method 0x86bf8056.
//
// Solidity: function marketPools( address,  uint256) constant returns(address)
func (_Market *MarketCallerSession) MarketPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPools(&_Market.CallOpts, arg0, arg1)
}

// MarketPoolsList is a free data retrieval call binding the contract method 0xe9e24234.
//
// Solidity: function marketPoolsList( uint256) constant returns(address)
func (_Market *MarketCaller) MarketPoolsList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "marketPoolsList", arg0)
	return *ret0, err
}

// MarketPoolsList is a free data retrieval call binding the contract method 0xe9e24234.
//
// Solidity: function marketPoolsList( uint256) constant returns(address)
func (_Market *MarketSession) MarketPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolsList(&_Market.CallOpts, arg0)
}

// MarketPoolsList is a free data retrieval call binding the contract method 0xe9e24234.
//
// Solidity: function marketPoolsList( uint256) constant returns(address)
func (_Market *MarketCallerSession) MarketPoolsList(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolsList(&_Market.CallOpts, arg0)
}

// OwnedPools is a free data retrieval call binding the contract method 0xdab66fb6.
//
// Solidity: function ownedPools( address,  uint256) constant returns(address)
func (_Market *MarketCaller) OwnedPools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "ownedPools", arg0, arg1)
	return *ret0, err
}

// OwnedPools is a free data retrieval call binding the contract method 0xdab66fb6.
//
// Solidity: function ownedPools( address,  uint256) constant returns(address)
func (_Market *MarketSession) OwnedPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.OwnedPools(&_Market.CallOpts, arg0, arg1)
}

// OwnedPools is a free data retrieval call binding the contract method 0xdab66fb6.
//
// Solidity: function ownedPools( address,  uint256) constant returns(address)
func (_Market *MarketCallerSession) OwnedPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.OwnedPools(&_Market.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Market *MarketCallerSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Payouts is a free data retrieval call binding the contract method 0x9d484693.
//
// Solidity: function payouts( address,  uint256) constant returns(pool address, user address, amount uint256, timestamp uint256)
func (_Market *MarketCaller) Payouts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Pool      common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Pool      common.Address
		User      common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	out := ret
	err := _Market.contract.Call(opts, out, "payouts", arg0, arg1)
	return *ret, err
}

// Payouts is a free data retrieval call binding the contract method 0x9d484693.
//
// Solidity: function payouts( address,  uint256) constant returns(pool address, user address, amount uint256, timestamp uint256)
func (_Market *MarketSession) Payouts(arg0 common.Address, arg1 *big.Int) (struct {
	Pool      common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Market.Contract.Payouts(&_Market.CallOpts, arg0, arg1)
}

// Payouts is a free data retrieval call binding the contract method 0x9d484693.
//
// Solidity: function payouts( address,  uint256) constant returns(pool address, user address, amount uint256, timestamp uint256)
func (_Market *MarketCallerSession) Payouts(arg0 common.Address, arg1 *big.Int) (struct {
	Pool      common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Market.Contract.Payouts(&_Market.CallOpts, arg0, arg1)
}

// PoolToOwner is a free data retrieval call binding the contract method 0xd4cced9a.
//
// Solidity: function poolToOwner( address) constant returns(address)
func (_Market *MarketCaller) PoolToOwner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "poolToOwner", arg0)
	return *ret0, err
}

// PoolToOwner is a free data retrieval call binding the contract method 0xd4cced9a.
//
// Solidity: function poolToOwner( address) constant returns(address)
func (_Market *MarketSession) PoolToOwner(arg0 common.Address) (common.Address, error) {
	return _Market.Contract.PoolToOwner(&_Market.CallOpts, arg0)
}

// PoolToOwner is a free data retrieval call binding the contract method 0xd4cced9a.
//
// Solidity: function poolToOwner( address) constant returns(address)
func (_Market *MarketCallerSession) PoolToOwner(arg0 common.Address) (common.Address, error) {
	return _Market.Contract.PoolToOwner(&_Market.CallOpts, arg0)
}

// AllocateClientFundsTo is a paid mutator transaction binding the contract method 0xc93d248c.
//
// Solidity: function allocateClientFundsTo(_pool address, _client address, _amount uint256) returns(bool)
func (_Market *MarketTransactor) AllocateClientFundsTo(opts *bind.TransactOpts, _pool common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "allocateClientFundsTo", _pool, _client, _amount)
}

// AllocateClientFundsTo is a paid mutator transaction binding the contract method 0xc93d248c.
//
// Solidity: function allocateClientFundsTo(_pool address, _client address, _amount uint256) returns(bool)
func (_Market *MarketSession) AllocateClientFundsTo(_pool common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AllocateClientFundsTo(&_Market.TransactOpts, _pool, _client, _amount)
}

// AllocateClientFundsTo is a paid mutator transaction binding the contract method 0xc93d248c.
//
// Solidity: function allocateClientFundsTo(_pool address, _client address, _amount uint256) returns(bool)
func (_Market *MarketTransactorSession) AllocateClientFundsTo(_pool common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AllocateClientFundsTo(&_Market.TransactOpts, _pool, _client, _amount)
}

// AllocateFunds is a paid mutator transaction binding the contract method 0xc9f8cc67.
//
// Solidity: function allocateFunds(_amount uint256) returns(bool)
func (_Market *MarketTransactor) AllocateFunds(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "allocateFunds", _amount)
}

// AllocateFunds is a paid mutator transaction binding the contract method 0xc9f8cc67.
//
// Solidity: function allocateFunds(_amount uint256) returns(bool)
func (_Market *MarketSession) AllocateFunds(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AllocateFunds(&_Market.TransactOpts, _amount)
}

// AllocateFunds is a paid mutator transaction binding the contract method 0xc9f8cc67.
//
// Solidity: function allocateFunds(_amount uint256) returns(bool)
func (_Market *MarketTransactorSession) AllocateFunds(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AllocateFunds(&_Market.TransactOpts, _amount)
}

// CreatePool is a paid mutator transaction binding the contract method 0xd0d13036.
//
// Solidity: function createPool(_publicKey string) returns(address)
func (_Market *MarketTransactor) CreatePool(opts *bind.TransactOpts, _publicKey string) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "createPool", _publicKey)
}

// CreatePool is a paid mutator transaction binding the contract method 0xd0d13036.
//
// Solidity: function createPool(_publicKey string) returns(address)
func (_Market *MarketSession) CreatePool(_publicKey string) (*types.Transaction, error) {
	return _Market.Contract.CreatePool(&_Market.TransactOpts, _publicKey)
}

// CreatePool is a paid mutator transaction binding the contract method 0xd0d13036.
//
// Solidity: function createPool(_publicKey string) returns(address)
func (_Market *MarketTransactorSession) CreatePool(_publicKey string) (*types.Transaction, error) {
	return _Market.Contract.CreatePool(&_Market.TransactOpts, _publicKey)
}

// JoinMarketplace is a paid mutator transaction binding the contract method 0xac0b0445.
//
// Solidity: function joinMarketplace(_poolAddress address) returns(bool)
func (_Market *MarketTransactor) JoinMarketplace(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "joinMarketplace", _poolAddress)
}

// JoinMarketplace is a paid mutator transaction binding the contract method 0xac0b0445.
//
// Solidity: function joinMarketplace(_poolAddress address) returns(bool)
func (_Market *MarketSession) JoinMarketplace(_poolAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.JoinMarketplace(&_Market.TransactOpts, _poolAddress)
}

// JoinMarketplace is a paid mutator transaction binding the contract method 0xac0b0445.
//
// Solidity: function joinMarketplace(_poolAddress address) returns(bool)
func (_Market *MarketTransactorSession) JoinMarketplace(_poolAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.JoinMarketplace(&_Market.TransactOpts, _poolAddress)
}

// LogWorkFrom is a paid mutator transaction binding the contract method 0x47ab23dc.
//
// Solidity: function logWorkFrom(_pool address, _node address, _client address, _amount uint256) returns()
func (_Market *MarketTransactor) LogWorkFrom(opts *bind.TransactOpts, _pool common.Address, _node common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "logWorkFrom", _pool, _node, _client, _amount)
}

// LogWorkFrom is a paid mutator transaction binding the contract method 0x47ab23dc.
//
// Solidity: function logWorkFrom(_pool address, _node address, _client address, _amount uint256) returns()
func (_Market *MarketSession) LogWorkFrom(_pool common.Address, _node common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.LogWorkFrom(&_Market.TransactOpts, _pool, _node, _client, _amount)
}

// LogWorkFrom is a paid mutator transaction binding the contract method 0x47ab23dc.
//
// Solidity: function logWorkFrom(_pool address, _node address, _client address, _amount uint256) returns()
func (_Market *MarketTransactorSession) LogWorkFrom(_pool common.Address, _node common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.LogWorkFrom(&_Market.TransactOpts, _pool, _node, _client, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(_amount uint256) returns(bool)
func (_Market *MarketTransactor) Pay(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "pay", _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(_amount uint256) returns(bool)
func (_Market *MarketSession) Pay(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Pay(&_Market.TransactOpts, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(_amount uint256) returns(bool)
func (_Market *MarketTransactorSession) Pay(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Pay(&_Market.TransactOpts, _amount)
}

// Payout is a paid mutator transaction binding the contract method 0x20f801d4.
//
// Solidity: function payout(_pool address, _node address, _amount uint256) returns(bool)
func (_Market *MarketTransactor) Payout(opts *bind.TransactOpts, _pool common.Address, _node common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "payout", _pool, _node, _amount)
}

// Payout is a paid mutator transaction binding the contract method 0x20f801d4.
//
// Solidity: function payout(_pool address, _node address, _amount uint256) returns(bool)
func (_Market *MarketSession) Payout(_pool common.Address, _node common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Payout(&_Market.TransactOpts, _pool, _node, _amount)
}

// Payout is a paid mutator transaction binding the contract method 0x20f801d4.
//
// Solidity: function payout(_pool address, _node address, _amount uint256) returns(bool)
func (_Market *MarketTransactorSession) Payout(_pool common.Address, _node common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Payout(&_Market.TransactOpts, _pool, _node, _amount)
}

// Work is a paid mutator transaction binding the contract method 0x5858d161.
//
// Solidity: function work(_amount uint256) returns(bool)
func (_Market *MarketTransactor) Work(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "work", _amount)
}

// Work is a paid mutator transaction binding the contract method 0x5858d161.
//
// Solidity: function work(_amount uint256) returns(bool)
func (_Market *MarketSession) Work(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Work(&_Market.TransactOpts, _amount)
}

// Work is a paid mutator transaction binding the contract method 0x5858d161.
//
// Solidity: function work(_amount uint256) returns(bool)
func (_Market *MarketTransactorSession) Work(_amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Work(&_Market.TransactOpts, _amount)
}
