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

// PoolABI is the input ABI used to generate the binding from.
const PoolABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"url\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"masterNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMasterNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSeedNode\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUrl\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"setSeedNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"addMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Pool *PoolCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Pool *PoolSession) Data() (string, error) {
	return _Pool.Contract.Data(&_Pool.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Pool *PoolCallerSession) Data() (string, error) {
	return _Pool.Contract.Data(&_Pool.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Pool *PoolCaller) GetData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getData")
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Pool *PoolSession) GetData() (string, error) {
	return _Pool.Contract.GetData(&_Pool.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(string)
func (_Pool *PoolCallerSession) GetData() (string, error) {
	return _Pool.Contract.GetData(&_Pool.CallOpts)
}

// GetMasterNodes is a free data retrieval call binding the contract method 0x67eeae72.
//
// Solidity: function getMasterNodes() constant returns(address[])
func (_Pool *PoolCaller) GetMasterNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getMasterNodes")
	return *ret0, err
}

// GetMasterNodes is a free data retrieval call binding the contract method 0x67eeae72.
//
// Solidity: function getMasterNodes() constant returns(address[])
func (_Pool *PoolSession) GetMasterNodes() ([]common.Address, error) {
	return _Pool.Contract.GetMasterNodes(&_Pool.CallOpts)
}

// GetMasterNodes is a free data retrieval call binding the contract method 0x67eeae72.
//
// Solidity: function getMasterNodes() constant returns(address[])
func (_Pool *PoolCallerSession) GetMasterNodes() ([]common.Address, error) {
	return _Pool.Contract.GetMasterNodes(&_Pool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Pool *PoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Pool *PoolSession) GetOwner() (common.Address, error) {
	return _Pool.Contract.GetOwner(&_Pool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Pool *PoolCallerSession) GetOwner() (common.Address, error) {
	return _Pool.Contract.GetOwner(&_Pool.CallOpts)
}

// GetSeedNode is a free data retrieval call binding the contract method 0x3c714980.
//
// Solidity: function getSeedNode() constant returns(string)
func (_Pool *PoolCaller) GetSeedNode(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getSeedNode")
	return *ret0, err
}

// GetSeedNode is a free data retrieval call binding the contract method 0x3c714980.
//
// Solidity: function getSeedNode() constant returns(string)
func (_Pool *PoolSession) GetSeedNode() (string, error) {
	return _Pool.Contract.GetSeedNode(&_Pool.CallOpts)
}

// GetSeedNode is a free data retrieval call binding the contract method 0x3c714980.
//
// Solidity: function getSeedNode() constant returns(string)
func (_Pool *PoolCallerSession) GetSeedNode() (string, error) {
	return _Pool.Contract.GetSeedNode(&_Pool.CallOpts)
}

// GetUrl is a free data retrieval call binding the contract method 0xd6bd8727.
//
// Solidity: function getUrl() constant returns(string)
func (_Pool *PoolCaller) GetUrl(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getUrl")
	return *ret0, err
}

// GetUrl is a free data retrieval call binding the contract method 0xd6bd8727.
//
// Solidity: function getUrl() constant returns(string)
func (_Pool *PoolSession) GetUrl() (string, error) {
	return _Pool.Contract.GetUrl(&_Pool.CallOpts)
}

// GetUrl is a free data retrieval call binding the contract method 0xd6bd8727.
//
// Solidity: function getUrl() constant returns(string)
func (_Pool *PoolCallerSession) GetUrl() (string, error) {
	return _Pool.Contract.GetUrl(&_Pool.CallOpts)
}

// MasterNodes is a free data retrieval call binding the contract method 0xca2e6c77.
//
// Solidity: function masterNodes( uint256) constant returns(address)
func (_Pool *PoolCaller) MasterNodes(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "masterNodes", arg0)
	return *ret0, err
}

// MasterNodes is a free data retrieval call binding the contract method 0xca2e6c77.
//
// Solidity: function masterNodes( uint256) constant returns(address)
func (_Pool *PoolSession) MasterNodes(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.MasterNodes(&_Pool.CallOpts, arg0)
}

// MasterNodes is a free data retrieval call binding the contract method 0xca2e6c77.
//
// Solidity: function masterNodes( uint256) constant returns(address)
func (_Pool *PoolCallerSession) MasterNodes(arg0 *big.Int) (common.Address, error) {
	return _Pool.Contract.MasterNodes(&_Pool.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// SeedNode is a free data retrieval call binding the contract method 0xe0a18ee7.
//
// Solidity: function seedNode() constant returns(string)
func (_Pool *PoolCaller) SeedNode(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "seedNode")
	return *ret0, err
}

// SeedNode is a free data retrieval call binding the contract method 0xe0a18ee7.
//
// Solidity: function seedNode() constant returns(string)
func (_Pool *PoolSession) SeedNode() (string, error) {
	return _Pool.Contract.SeedNode(&_Pool.CallOpts)
}

// SeedNode is a free data retrieval call binding the contract method 0xe0a18ee7.
//
// Solidity: function seedNode() constant returns(string)
func (_Pool *PoolCallerSession) SeedNode() (string, error) {
	return _Pool.Contract.SeedNode(&_Pool.CallOpts)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() constant returns(string)
func (_Pool *PoolCaller) Url(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "url")
	return *ret0, err
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() constant returns(string)
func (_Pool *PoolSession) Url() (string, error) {
	return _Pool.Contract.Url(&_Pool.CallOpts)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() constant returns(string)
func (_Pool *PoolCallerSession) Url() (string, error) {
	return _Pool.Contract.Url(&_Pool.CallOpts)
}

// AddMasterNode is a paid mutator transaction binding the contract method 0x6b8a98a9.
//
// Solidity: function addMasterNode(_node address) returns()
func (_Pool *PoolTransactor) AddMasterNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addMasterNode", _node)
}

// AddMasterNode is a paid mutator transaction binding the contract method 0x6b8a98a9.
//
// Solidity: function addMasterNode(_node address) returns()
func (_Pool *PoolSession) AddMasterNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddMasterNode(&_Pool.TransactOpts, _node)
}

// AddMasterNode is a paid mutator transaction binding the contract method 0x6b8a98a9.
//
// Solidity: function addMasterNode(_node address) returns()
func (_Pool *PoolTransactorSession) AddMasterNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddMasterNode(&_Pool.TransactOpts, _node)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Pool *PoolTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Pool *PoolSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ChangeOwner(&_Pool.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Pool *PoolTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.ChangeOwner(&_Pool.TransactOpts, _owner)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Pool *PoolTransactor) SetData(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Pool *PoolSession) SetData(_data string) (*types.Transaction, error) {
	return _Pool.Contract.SetData(&_Pool.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Pool *PoolTransactorSession) SetData(_data string) (*types.Transaction, error) {
	return _Pool.Contract.SetData(&_Pool.TransactOpts, _data)
}

// SetSeedNode is a paid mutator transaction binding the contract method 0x6bc3af27.
//
// Solidity: function setSeedNode(_node string) returns()
func (_Pool *PoolTransactor) SetSeedNode(opts *bind.TransactOpts, _node string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setSeedNode", _node)
}

// SetSeedNode is a paid mutator transaction binding the contract method 0x6bc3af27.
//
// Solidity: function setSeedNode(_node string) returns()
func (_Pool *PoolSession) SetSeedNode(_node string) (*types.Transaction, error) {
	return _Pool.Contract.SetSeedNode(&_Pool.TransactOpts, _node)
}

// SetSeedNode is a paid mutator transaction binding the contract method 0x6bc3af27.
//
// Solidity: function setSeedNode(_node string) returns()
func (_Pool *PoolTransactorSession) SetSeedNode(_node string) (*types.Transaction, error) {
	return _Pool.Contract.SetSeedNode(&_Pool.TransactOpts, _node)
}

// SetUrl is a paid mutator transaction binding the contract method 0x252498a2.
//
// Solidity: function setUrl(_url string) returns()
func (_Pool *PoolTransactor) SetUrl(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUrl", _url)
}

// SetUrl is a paid mutator transaction binding the contract method 0x252498a2.
//
// Solidity: function setUrl(_url string) returns()
func (_Pool *PoolSession) SetUrl(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetUrl(&_Pool.TransactOpts, _url)
}

// SetUrl is a paid mutator transaction binding the contract method 0x252498a2.
//
// Solidity: function setUrl(_url string) returns()
func (_Pool *PoolTransactorSession) SetUrl(_url string) (*types.Transaction, error) {
	return _Pool.Contract.SetUrl(&_Pool.TransactOpts, _url)
}
