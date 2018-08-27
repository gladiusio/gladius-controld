// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MarketABI is the input ABI used to generate the binding from.
const MarketABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPoolsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketPoolOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketPoolsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[],\"name\":\"createPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMarketPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"getOwnerAllPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"getOwnerMarketPools\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"addToMarketplace\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Market *MarketCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Market *MarketSession) Data() (string, error) {
	return _Market.Contract.Data(&_Market.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Market *MarketCallerSession) Data() (string, error) {
	return _Market.Contract.Data(&_Market.CallOpts)
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

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Market *MarketCaller) GetData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getData")
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Market *MarketSession) GetData() (string, error) {
	return _Market.Contract.GetData(&_Market.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Market *MarketCallerSession) GetData() (string, error) {
	return _Market.Contract.GetData(&_Market.CallOpts)
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

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Market *MarketCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Market *MarketSession) GetOwner() (common.Address, error) {
	return _Market.Contract.GetOwner(&_Market.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Market *MarketCallerSession) GetOwner() (common.Address, error) {
	return _Market.Contract.GetOwner(&_Market.CallOpts)
}

// GetOwnerAllPools is a free data retrieval call binding the contract method 0xaebce2d4.
//
// Solidity: function getOwnerAllPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketCaller) GetOwnerAllPools(opts *bind.CallOpts, _ownerAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getOwnerAllPools", _ownerAddress)
	return *ret0, err
}

// GetOwnerAllPools is a free data retrieval call binding the contract method 0xaebce2d4.
//
// Solidity: function getOwnerAllPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketSession) GetOwnerAllPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerAllPools(&_Market.CallOpts, _ownerAddress)
}

// GetOwnerAllPools is a free data retrieval call binding the contract method 0xaebce2d4.
//
// Solidity: function getOwnerAllPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketCallerSession) GetOwnerAllPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerAllPools(&_Market.CallOpts, _ownerAddress)
}

// GetOwnerMarketPools is a free data retrieval call binding the contract method 0x17e421c8.
//
// Solidity: function getOwnerMarketPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketCaller) GetOwnerMarketPools(opts *bind.CallOpts, _ownerAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getOwnerMarketPools", _ownerAddress)
	return *ret0, err
}

// GetOwnerMarketPools is a free data retrieval call binding the contract method 0x17e421c8.
//
// Solidity: function getOwnerMarketPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketSession) GetOwnerMarketPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerMarketPools(&_Market.CallOpts, _ownerAddress)
}

// GetOwnerMarketPools is a free data retrieval call binding the contract method 0x17e421c8.
//
// Solidity: function getOwnerMarketPools(_ownerAddress address) constant returns(address[])
func (_Market *MarketCallerSession) GetOwnerMarketPools(_ownerAddress common.Address) ([]common.Address, error) {
	return _Market.Contract.GetOwnerMarketPools(&_Market.CallOpts, _ownerAddress)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(_pool address) constant returns(address)
func (_Market *MarketCaller) GetPoolOwner(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getPoolOwner", _pool)
	return *ret0, err
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(_pool address) constant returns(address)
func (_Market *MarketSession) GetPoolOwner(_pool common.Address) (common.Address, error) {
	return _Market.Contract.GetPoolOwner(&_Market.CallOpts, _pool)
}

// GetPoolOwner is a free data retrieval call binding the contract method 0x7cf25517.
//
// Solidity: function getPoolOwner(_pool address) constant returns(address)
func (_Market *MarketCallerSession) GetPoolOwner(_pool common.Address) (common.Address, error) {
	return _Market.Contract.GetPoolOwner(&_Market.CallOpts, _pool)
}

// MarketPoolOwners is a free data retrieval call binding the contract method 0xdb20ee33.
//
// Solidity: function marketPoolOwners( address,  uint256) constant returns(address)
func (_Market *MarketCaller) MarketPoolOwners(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "marketPoolOwners", arg0, arg1)
	return *ret0, err
}

// MarketPoolOwners is a free data retrieval call binding the contract method 0xdb20ee33.
//
// Solidity: function marketPoolOwners( address,  uint256) constant returns(address)
func (_Market *MarketSession) MarketPoolOwners(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolOwners(&_Market.CallOpts, arg0, arg1)
}

// MarketPoolOwners is a free data retrieval call binding the contract method 0xdb20ee33.
//
// Solidity: function marketPoolOwners( address,  uint256) constant returns(address)
func (_Market *MarketCallerSession) MarketPoolOwners(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.MarketPoolOwners(&_Market.CallOpts, arg0, arg1)
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

// PoolOwners is a free data retrieval call binding the contract method 0x6a05e5fe.
//
// Solidity: function poolOwners( address,  uint256) constant returns(address)
func (_Market *MarketCaller) PoolOwners(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "poolOwners", arg0, arg1)
	return *ret0, err
}

// PoolOwners is a free data retrieval call binding the contract method 0x6a05e5fe.
//
// Solidity: function poolOwners( address,  uint256) constant returns(address)
func (_Market *MarketSession) PoolOwners(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.PoolOwners(&_Market.CallOpts, arg0, arg1)
}

// PoolOwners is a free data retrieval call binding the contract method 0x6a05e5fe.
//
// Solidity: function poolOwners( address,  uint256) constant returns(address)
func (_Market *MarketCallerSession) PoolOwners(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Market.Contract.PoolOwners(&_Market.CallOpts, arg0, arg1)
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

// AddToMarketplace is a paid mutator transaction binding the contract method 0xd00e1e72.
//
// Solidity: function addToMarketplace(_poolAddress address) returns(bool)
func (_Market *MarketTransactor) AddToMarketplace(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "addToMarketplace", _poolAddress)
}

// AddToMarketplace is a paid mutator transaction binding the contract method 0xd00e1e72.
//
// Solidity: function addToMarketplace(_poolAddress address) returns(bool)
func (_Market *MarketSession) AddToMarketplace(_poolAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddToMarketplace(&_Market.TransactOpts, _poolAddress)
}

// AddToMarketplace is a paid mutator transaction binding the contract method 0xd00e1e72.
//
// Solidity: function addToMarketplace(_poolAddress address) returns(bool)
func (_Market *MarketTransactorSession) AddToMarketplace(_poolAddress common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddToMarketplace(&_Market.TransactOpts, _poolAddress)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Market *MarketTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Market *MarketSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Market.Contract.ChangeOwner(&_Market.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Market *MarketTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Market.Contract.ChangeOwner(&_Market.TransactOpts, _owner)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9a06b113.
//
// Solidity: function createPool() returns(address)
func (_Market *MarketTransactor) CreatePool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "createPool")
}

// CreatePool is a paid mutator transaction binding the contract method 0x9a06b113.
//
// Solidity: function createPool() returns(address)
func (_Market *MarketSession) CreatePool() (*types.Transaction, error) {
	return _Market.Contract.CreatePool(&_Market.TransactOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9a06b113.
//
// Solidity: function createPool() returns(address)
func (_Market *MarketTransactorSession) CreatePool() (*types.Transaction, error) {
	return _Market.Contract.CreatePool(&_Market.TransactOpts)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Market *MarketTransactor) SetData(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Market *MarketSession) SetData(_data string) (*types.Transaction, error) {
	return _Market.Contract.SetData(&_Market.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Market *MarketTransactorSession) SetData(_data string) (*types.Transaction, error) {
	return _Market.Contract.SetData(&_Market.TransactOpts, _data)
}
