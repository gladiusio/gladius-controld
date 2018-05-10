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

// NodeFactoryABI is the input ABI used to generate the binding from.
const NodeFactoryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"createNode\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeFactory is an auto generated Go binding around an Ethereum contract.
type NodeFactory struct {
	NodeFactoryCaller     // Read-only binding to the contract
	NodeFactoryTransactor // Write-only binding to the contract
	NodeFactoryFilterer   // Log filterer for contract events
}

// NodeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeFactorySession struct {
	Contract     *NodeFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeFactoryCallerSession struct {
	Contract *NodeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeFactoryTransactorSession struct {
	Contract     *NodeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeFactoryRaw struct {
	Contract *NodeFactory // Generic contract binding to access the raw methods on
}

// NodeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeFactoryCallerRaw struct {
	Contract *NodeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// NodeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeFactoryTransactorRaw struct {
	Contract *NodeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeFactory creates a new instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactory(address common.Address, backend bind.ContractBackend) (*NodeFactory, error) {
	contract, err := bindNodeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeFactory{NodeFactoryCaller: NodeFactoryCaller{contract: contract}, NodeFactoryTransactor: NodeFactoryTransactor{contract: contract}, NodeFactoryFilterer: NodeFactoryFilterer{contract: contract}}, nil
}

// NewNodeFactoryCaller creates a new read-only instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryCaller(address common.Address, caller bind.ContractCaller) (*NodeFactoryCaller, error) {
	contract, err := bindNodeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryCaller{contract: contract}, nil
}

// NewNodeFactoryTransactor creates a new write-only instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeFactoryTransactor, error) {
	contract, err := bindNodeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryTransactor{contract: contract}, nil
}

// NewNodeFactoryFilterer creates a new log filterer instance of NodeFactory, bound to a specific deployed contract.
func NewNodeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFactoryFilterer, error) {
	contract, err := bindNodeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFactoryFilterer{contract: contract}, nil
}

// bindNodeFactory binds a generic wrapper to an already deployed contract.
func bindNodeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeFactory *NodeFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeFactory.Contract.NodeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeFactory *NodeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeFactory.Contract.NodeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeFactory *NodeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeFactory.Contract.NodeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeFactory *NodeFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeFactory *NodeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeFactory *NodeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeFactory.Contract.contract.Transact(opts, method, params...)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() constant returns(address)
func (_NodeFactory *NodeFactoryCaller) GetNodeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodeFactory.contract.Call(opts, out, "getNodeAddress")
	return *ret0, err
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() constant returns(address)
func (_NodeFactory *NodeFactorySession) GetNodeAddress() (common.Address, error) {
	return _NodeFactory.Contract.GetNodeAddress(&_NodeFactory.CallOpts)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() constant returns(address)
func (_NodeFactory *NodeFactoryCallerSession) GetNodeAddress() (common.Address, error) {
	return _NodeFactory.Contract.GetNodeAddress(&_NodeFactory.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(uint256)
func (_NodeFactory *NodeFactoryCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeFactory.contract.Call(opts, out, "getNodeCount")
	return *ret0, err
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(uint256)
func (_NodeFactory *NodeFactorySession) GetNodeCount() (*big.Int, error) {
	return _NodeFactory.Contract.GetNodeCount(&_NodeFactory.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() constant returns(uint256)
func (_NodeFactory *NodeFactoryCallerSession) GetNodeCount() (*big.Int, error) {
	return _NodeFactory.Contract.GetNodeCount(&_NodeFactory.CallOpts)
}

// CreateNode is a paid mutator transaction binding the contract method 0x1c3917b8.
//
// Solidity: function createNode() returns(address)
func (_NodeFactory *NodeFactoryTransactor) CreateNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeFactory.contract.Transact(opts, "createNode")
}

// CreateNode is a paid mutator transaction binding the contract method 0x1c3917b8.
//
// Solidity: function createNode() returns(address)
func (_NodeFactory *NodeFactorySession) CreateNode() (*types.Transaction, error) {
	return _NodeFactory.Contract.CreateNode(&_NodeFactory.TransactOpts)
}

// CreateNode is a paid mutator transaction binding the contract method 0x1c3917b8.
//
// Solidity: function createNode() returns(address)
func (_NodeFactory *NodeFactoryTransactorSession) CreateNode() (*types.Transaction, error) {
	return _NodeFactory.Contract.CreateNode(&_NodeFactory.TransactOpts)
}
