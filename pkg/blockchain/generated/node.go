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

// NodeABI is the input ABI used to generate the binding from.
const NodeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicData\",\"type\":\"string\"}],\"name\":\"setPublicData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"int256\"}],\"name\":\"setStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"changePoolData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"applyToPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Node *NodeCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Node.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Node *NodeSession) Data() (string, error) {
	return _Node.Contract.Data(&_Node.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Node *NodeCallerSession) Data() (string, error) {
	return _Node.Contract.Data(&_Node.CallOpts)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(_pool address) constant returns(string)
func (_Node *NodeCaller) GetPoolData(opts *bind.CallOpts, _pool common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Node.contract.Call(opts, out, "getPoolData", _pool)
	return *ret0, err
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(_pool address) constant returns(string)
func (_Node *NodeSession) GetPoolData(_pool common.Address) (string, error) {
	return _Node.Contract.GetPoolData(&_Node.CallOpts, _pool)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(_pool address) constant returns(string)
func (_Node *NodeCallerSession) GetPoolData(_pool common.Address) (string, error) {
	return _Node.Contract.GetPoolData(&_Node.CallOpts, _pool)
}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() constant returns(address[])
func (_Node *NodeCaller) GetPoolList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Node.contract.Call(opts, out, "getPoolList")
	return *ret0, err
}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() constant returns(address[])
func (_Node *NodeSession) GetPoolList() ([]common.Address, error) {
	return _Node.Contract.GetPoolList(&_Node.CallOpts)
}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() constant returns(address[])
func (_Node *NodeCallerSession) GetPoolList() ([]common.Address, error) {
	return _Node.Contract.GetPoolList(&_Node.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(_pool address) constant returns(int256)
func (_Node *NodeCaller) GetStatus(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Node.contract.Call(opts, out, "getStatus", _pool)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(_pool address) constant returns(int256)
func (_Node *NodeSession) GetStatus(_pool common.Address) (*big.Int, error) {
	return _Node.Contract.GetStatus(&_Node.CallOpts, _pool)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(_pool address) constant returns(int256)
func (_Node *NodeCallerSession) GetStatus(_pool common.Address) (*big.Int, error) {
	return _Node.Contract.GetStatus(&_Node.CallOpts, _pool)
}

// PublicData is a free data retrieval call binding the contract method 0x9a76717a.
//
// Solidity: function publicData() constant returns(string)
func (_Node *NodeCaller) PublicData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Node.contract.Call(opts, out, "publicData")
	return *ret0, err
}

// PublicData is a free data retrieval call binding the contract method 0x9a76717a.
//
// Solidity: function publicData() constant returns(string)
func (_Node *NodeSession) PublicData() (string, error) {
	return _Node.Contract.PublicData(&_Node.CallOpts)
}

// PublicData is a free data retrieval call binding the contract method 0x9a76717a.
//
// Solidity: function publicData() constant returns(string)
func (_Node *NodeCallerSession) PublicData() (string, error) {
	return _Node.Contract.PublicData(&_Node.CallOpts)
}

// ApplyToPool is a paid mutator transaction binding the contract method 0xb5eca7ec.
//
// Solidity: function applyToPool(_pool address, _data string) returns()
func (_Node *NodeTransactor) ApplyToPool(opts *bind.TransactOpts, _pool common.Address, _data string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "applyToPool", _pool, _data)
}

// ApplyToPool is a paid mutator transaction binding the contract method 0xb5eca7ec.
//
// Solidity: function applyToPool(_pool address, _data string) returns()
func (_Node *NodeSession) ApplyToPool(_pool common.Address, _data string) (*types.Transaction, error) {
	return _Node.Contract.ApplyToPool(&_Node.TransactOpts, _pool, _data)
}

// ApplyToPool is a paid mutator transaction binding the contract method 0xb5eca7ec.
//
// Solidity: function applyToPool(_pool address, _data string) returns()
func (_Node *NodeTransactorSession) ApplyToPool(_pool common.Address, _data string) (*types.Transaction, error) {
	return _Node.Contract.ApplyToPool(&_Node.TransactOpts, _pool, _data)
}

// ChangePoolData is a paid mutator transaction binding the contract method 0x4b8aad60.
//
// Solidity: function changePoolData(_pool address, _data string) returns()
func (_Node *NodeTransactor) ChangePoolData(opts *bind.TransactOpts, _pool common.Address, _data string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "changePoolData", _pool, _data)
}

// ChangePoolData is a paid mutator transaction binding the contract method 0x4b8aad60.
//
// Solidity: function changePoolData(_pool address, _data string) returns()
func (_Node *NodeSession) ChangePoolData(_pool common.Address, _data string) (*types.Transaction, error) {
	return _Node.Contract.ChangePoolData(&_Node.TransactOpts, _pool, _data)
}

// ChangePoolData is a paid mutator transaction binding the contract method 0x4b8aad60.
//
// Solidity: function changePoolData(_pool address, _data string) returns()
func (_Node *NodeTransactorSession) ChangePoolData(_pool common.Address, _data string) (*types.Transaction, error) {
	return _Node.Contract.ChangePoolData(&_Node.TransactOpts, _pool, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Node *NodeTransactor) SetData(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Node *NodeSession) SetData(_data string) (*types.Transaction, error) {
	return _Node.Contract.SetData(&_Node.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x47064d6a.
//
// Solidity: function setData(_data string) returns()
func (_Node *NodeTransactorSession) SetData(_data string) (*types.Transaction, error) {
	return _Node.Contract.SetData(&_Node.TransactOpts, _data)
}

// SetPublicData is a paid mutator transaction binding the contract method 0xfddbd275.
//
// Solidity: function setPublicData(_publicData string) returns()
func (_Node *NodeTransactor) SetPublicData(opts *bind.TransactOpts, _publicData string) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setPublicData", _publicData)
}

// SetPublicData is a paid mutator transaction binding the contract method 0xfddbd275.
//
// Solidity: function setPublicData(_publicData string) returns()
func (_Node *NodeSession) SetPublicData(_publicData string) (*types.Transaction, error) {
	return _Node.Contract.SetPublicData(&_Node.TransactOpts, _publicData)
}

// SetPublicData is a paid mutator transaction binding the contract method 0xfddbd275.
//
// Solidity: function setPublicData(_publicData string) returns()
func (_Node *NodeTransactorSession) SetPublicData(_publicData string) (*types.Transaction, error) {
	return _Node.Contract.SetPublicData(&_Node.TransactOpts, _publicData)
}

// SetStatus is a paid mutator transaction binding the contract method 0x17bc269b.
//
// Solidity: function setStatus(_status int256) returns()
func (_Node *NodeTransactor) SetStatus(opts *bind.TransactOpts, _status *big.Int) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "setStatus", _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x17bc269b.
//
// Solidity: function setStatus(_status int256) returns()
func (_Node *NodeSession) SetStatus(_status *big.Int) (*types.Transaction, error) {
	return _Node.Contract.SetStatus(&_Node.TransactOpts, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x17bc269b.
//
// Solidity: function setStatus(_status int256) returns()
func (_Node *NodeTransactorSession) SetStatus(_status *big.Int) (*types.Transaction, error) {
	return _Node.Contract.SetStatus(&_Node.TransactOpts, _status)
}
