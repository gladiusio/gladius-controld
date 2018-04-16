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
const PoolABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"work\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"publicData\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"owed\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint256\"},{\"name\":\"completed\",\"type\":\"uint256\"},{\"name\":\"paid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allocateFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_publicKey\",\"type\":\"string\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_publicData\",\"type\":\"string\"}],\"name\":\"setPublicData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getBalanceStructFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getTotalBalanceFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getOwedBalanceFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allocateFundsFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"},{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"logWorkFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getClientList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addClient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"acceptNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"}],\"name\":\"acceptClient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"rejectNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"}],\"name\":\"rejectClient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(owed uint256, total uint256, completed uint256, paid uint256)
func (_Pool *PoolCaller) Balance(opts *bind.CallOpts) (struct {
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
	err := _Pool.contract.Call(opts, out, "balance")
	return *ret, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(owed uint256, total uint256, completed uint256, paid uint256)
func (_Pool *PoolSession) Balance() (struct {
	Owed      *big.Int
	Total     *big.Int
	Completed *big.Int
	Paid      *big.Int
}, error) {
	return _Pool.Contract.Balance(&_Pool.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(owed uint256, total uint256, completed uint256, paid uint256)
func (_Pool *PoolCallerSession) Balance() (struct {
	Owed      *big.Int
	Total     *big.Int
	Completed *big.Int
	Paid      *big.Int
}, error) {
	return _Pool.Contract.Balance(&_Pool.CallOpts)
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

// GetBalanceStructFor is a free data retrieval call binding the contract method 0x7cc62dd7.
//
// Solidity: function getBalanceStructFor(_user address) constant returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolCaller) GetBalanceStructFor(opts *bind.CallOpts, _user common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Pool.contract.Call(opts, out, "getBalanceStructFor", _user)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetBalanceStructFor is a free data retrieval call binding the contract method 0x7cc62dd7.
//
// Solidity: function getBalanceStructFor(_user address) constant returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolSession) GetBalanceStructFor(_user common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.GetBalanceStructFor(&_Pool.CallOpts, _user)
}

// GetBalanceStructFor is a free data retrieval call binding the contract method 0x7cc62dd7.
//
// Solidity: function getBalanceStructFor(_user address) constant returns(uint256, uint256, uint256, uint256)
func (_Pool *PoolCallerSession) GetBalanceStructFor(_user common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pool.Contract.GetBalanceStructFor(&_Pool.CallOpts, _user)
}

// GetClientList is a free data retrieval call binding the contract method 0xd644f767.
//
// Solidity: function getClientList() constant returns(address[])
func (_Pool *PoolCaller) GetClientList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getClientList")
	return *ret0, err
}

// GetClientList is a free data retrieval call binding the contract method 0xd644f767.
//
// Solidity: function getClientList() constant returns(address[])
func (_Pool *PoolSession) GetClientList() ([]common.Address, error) {
	return _Pool.Contract.GetClientList(&_Pool.CallOpts)
}

// GetClientList is a free data retrieval call binding the contract method 0xd644f767.
//
// Solidity: function getClientList() constant returns(address[])
func (_Pool *PoolCallerSession) GetClientList() ([]common.Address, error) {
	return _Pool.Contract.GetClientList(&_Pool.CallOpts)
}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() constant returns(address[])
func (_Pool *PoolCaller) GetNodeList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getNodeList")
	return *ret0, err
}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() constant returns(address[])
func (_Pool *PoolSession) GetNodeList() ([]common.Address, error) {
	return _Pool.Contract.GetNodeList(&_Pool.CallOpts)
}

// GetNodeList is a free data retrieval call binding the contract method 0x53f3b713.
//
// Solidity: function getNodeList() constant returns(address[])
func (_Pool *PoolCallerSession) GetNodeList() ([]common.Address, error) {
	return _Pool.Contract.GetNodeList(&_Pool.CallOpts)
}

// GetOwedBalanceFor is a free data retrieval call binding the contract method 0xac28949c.
//
// Solidity: function getOwedBalanceFor(_user address) constant returns(uint256)
func (_Pool *PoolCaller) GetOwedBalanceFor(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getOwedBalanceFor", _user)
	return *ret0, err
}

// GetOwedBalanceFor is a free data retrieval call binding the contract method 0xac28949c.
//
// Solidity: function getOwedBalanceFor(_user address) constant returns(uint256)
func (_Pool *PoolSession) GetOwedBalanceFor(_user common.Address) (*big.Int, error) {
	return _Pool.Contract.GetOwedBalanceFor(&_Pool.CallOpts, _user)
}

// GetOwedBalanceFor is a free data retrieval call binding the contract method 0xac28949c.
//
// Solidity: function getOwedBalanceFor(_user address) constant returns(uint256)
func (_Pool *PoolCallerSession) GetOwedBalanceFor(_user common.Address) (*big.Int, error) {
	return _Pool.Contract.GetOwedBalanceFor(&_Pool.CallOpts, _user)
}

// GetTotalBalanceFor is a free data retrieval call binding the contract method 0xa8f9868e.
//
// Solidity: function getTotalBalanceFor(_user address) constant returns(uint256)
func (_Pool *PoolCaller) GetTotalBalanceFor(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "getTotalBalanceFor", _user)
	return *ret0, err
}

// GetTotalBalanceFor is a free data retrieval call binding the contract method 0xa8f9868e.
//
// Solidity: function getTotalBalanceFor(_user address) constant returns(uint256)
func (_Pool *PoolSession) GetTotalBalanceFor(_user common.Address) (*big.Int, error) {
	return _Pool.Contract.GetTotalBalanceFor(&_Pool.CallOpts, _user)
}

// GetTotalBalanceFor is a free data retrieval call binding the contract method 0xa8f9868e.
//
// Solidity: function getTotalBalanceFor(_user address) constant returns(uint256)
func (_Pool *PoolCallerSession) GetTotalBalanceFor(_user common.Address) (*big.Int, error) {
	return _Pool.Contract.GetTotalBalanceFor(&_Pool.CallOpts, _user)
}

// PublicData is a free data retrieval call binding the contract method 0x9a76717a.
//
// Solidity: function publicData() constant returns(string)
func (_Pool *PoolCaller) PublicData(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "publicData")
	return *ret0, err
}

// PublicData is a free data retrieval call binding the contract method 0x9a76717a.
//
// Solidity: function publicData() constant returns(string)
func (_Pool *PoolSession) PublicData() (string, error) {
	return _Pool.Contract.PublicData(&_Pool.CallOpts)
}

// PublicData is a free data retrieval call binding the contract method 0x9a76717a.
//
// Solidity: function publicData() constant returns(string)
func (_Pool *PoolCallerSession) PublicData() (string, error) {
	return _Pool.Contract.PublicData(&_Pool.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_Pool *PoolCaller) PublicKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "publicKey")
	return *ret0, err
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_Pool *PoolSession) PublicKey() (string, error) {
	return _Pool.Contract.PublicKey(&_Pool.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() constant returns(string)
func (_Pool *PoolCallerSession) PublicKey() (string, error) {
	return _Pool.Contract.PublicKey(&_Pool.CallOpts)
}

// AcceptClient is a paid mutator transaction binding the contract method 0xa3eb9444.
//
// Solidity: function acceptClient(_client address) returns()
func (_Pool *PoolTransactor) AcceptClient(opts *bind.TransactOpts, _client common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "acceptClient", _client)
}

// AcceptClient is a paid mutator transaction binding the contract method 0xa3eb9444.
//
// Solidity: function acceptClient(_client address) returns()
func (_Pool *PoolSession) AcceptClient(_client common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AcceptClient(&_Pool.TransactOpts, _client)
}

// AcceptClient is a paid mutator transaction binding the contract method 0xa3eb9444.
//
// Solidity: function acceptClient(_client address) returns()
func (_Pool *PoolTransactorSession) AcceptClient(_client common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AcceptClient(&_Pool.TransactOpts, _client)
}

// AcceptNode is a paid mutator transaction binding the contract method 0x400008cf.
//
// Solidity: function acceptNode(_node address) returns()
func (_Pool *PoolTransactor) AcceptNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "acceptNode", _node)
}

// AcceptNode is a paid mutator transaction binding the contract method 0x400008cf.
//
// Solidity: function acceptNode(_node address) returns()
func (_Pool *PoolSession) AcceptNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AcceptNode(&_Pool.TransactOpts, _node)
}

// AcceptNode is a paid mutator transaction binding the contract method 0x400008cf.
//
// Solidity: function acceptNode(_node address) returns()
func (_Pool *PoolTransactorSession) AcceptNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AcceptNode(&_Pool.TransactOpts, _node)
}

// AddClient is a paid mutator transaction binding the contract method 0x533aa474.
//
// Solidity: function addClient() returns()
func (_Pool *PoolTransactor) AddClient(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addClient")
}

// AddClient is a paid mutator transaction binding the contract method 0x533aa474.
//
// Solidity: function addClient() returns()
func (_Pool *PoolSession) AddClient() (*types.Transaction, error) {
	return _Pool.Contract.AddClient(&_Pool.TransactOpts)
}

// AddClient is a paid mutator transaction binding the contract method 0x533aa474.
//
// Solidity: function addClient() returns()
func (_Pool *PoolTransactorSession) AddClient() (*types.Transaction, error) {
	return _Pool.Contract.AddClient(&_Pool.TransactOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0xe07c60e1.
//
// Solidity: function addNode() returns()
func (_Pool *PoolTransactor) AddNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addNode")
}

// AddNode is a paid mutator transaction binding the contract method 0xe07c60e1.
//
// Solidity: function addNode() returns()
func (_Pool *PoolSession) AddNode() (*types.Transaction, error) {
	return _Pool.Contract.AddNode(&_Pool.TransactOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0xe07c60e1.
//
// Solidity: function addNode() returns()
func (_Pool *PoolTransactorSession) AddNode() (*types.Transaction, error) {
	return _Pool.Contract.AddNode(&_Pool.TransactOpts)
}

// AllocateFunds is a paid mutator transaction binding the contract method 0xc9f8cc67.
//
// Solidity: function allocateFunds(_amount uint256) returns(bool)
func (_Pool *PoolTransactor) AllocateFunds(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "allocateFunds", _amount)
}

// AllocateFunds is a paid mutator transaction binding the contract method 0xc9f8cc67.
//
// Solidity: function allocateFunds(_amount uint256) returns(bool)
func (_Pool *PoolSession) AllocateFunds(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AllocateFunds(&_Pool.TransactOpts, _amount)
}

// AllocateFunds is a paid mutator transaction binding the contract method 0xc9f8cc67.
//
// Solidity: function allocateFunds(_amount uint256) returns(bool)
func (_Pool *PoolTransactorSession) AllocateFunds(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AllocateFunds(&_Pool.TransactOpts, _amount)
}

// AllocateFundsFrom is a paid mutator transaction binding the contract method 0xa1f2e9bf.
//
// Solidity: function allocateFundsFrom(_client address, _amount uint256) returns(bool)
func (_Pool *PoolTransactor) AllocateFundsFrom(opts *bind.TransactOpts, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "allocateFundsFrom", _client, _amount)
}

// AllocateFundsFrom is a paid mutator transaction binding the contract method 0xa1f2e9bf.
//
// Solidity: function allocateFundsFrom(_client address, _amount uint256) returns(bool)
func (_Pool *PoolSession) AllocateFundsFrom(_client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AllocateFundsFrom(&_Pool.TransactOpts, _client, _amount)
}

// AllocateFundsFrom is a paid mutator transaction binding the contract method 0xa1f2e9bf.
//
// Solidity: function allocateFundsFrom(_client address, _amount uint256) returns(bool)
func (_Pool *PoolTransactorSession) AllocateFundsFrom(_client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AllocateFundsFrom(&_Pool.TransactOpts, _client, _amount)
}

// LogWorkFrom is a paid mutator transaction binding the contract method 0x32a6a376.
//
// Solidity: function logWorkFrom(_node address, _client address, _amount uint256) returns(bool)
func (_Pool *PoolTransactor) LogWorkFrom(opts *bind.TransactOpts, _node common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "logWorkFrom", _node, _client, _amount)
}

// LogWorkFrom is a paid mutator transaction binding the contract method 0x32a6a376.
//
// Solidity: function logWorkFrom(_node address, _client address, _amount uint256) returns(bool)
func (_Pool *PoolSession) LogWorkFrom(_node common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.LogWorkFrom(&_Pool.TransactOpts, _node, _client, _amount)
}

// LogWorkFrom is a paid mutator transaction binding the contract method 0x32a6a376.
//
// Solidity: function logWorkFrom(_node address, _client address, _amount uint256) returns(bool)
func (_Pool *PoolTransactorSession) LogWorkFrom(_node common.Address, _client common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.LogWorkFrom(&_Pool.TransactOpts, _node, _client, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(_amount uint256) returns(bool)
func (_Pool *PoolTransactor) Pay(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "pay", _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(_amount uint256) returns(bool)
func (_Pool *PoolSession) Pay(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Pay(&_Pool.TransactOpts, _amount)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(_amount uint256) returns(bool)
func (_Pool *PoolTransactorSession) Pay(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Pay(&_Pool.TransactOpts, _amount)
}

// Payout is a paid mutator transaction binding the contract method 0x117de2fd.
//
// Solidity: function payout(_node address, _amount uint256) returns(bool)
func (_Pool *PoolTransactor) Payout(opts *bind.TransactOpts, _node common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "payout", _node, _amount)
}

// Payout is a paid mutator transaction binding the contract method 0x117de2fd.
//
// Solidity: function payout(_node address, _amount uint256) returns(bool)
func (_Pool *PoolSession) Payout(_node common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Payout(&_Pool.TransactOpts, _node, _amount)
}

// Payout is a paid mutator transaction binding the contract method 0x117de2fd.
//
// Solidity: function payout(_node address, _amount uint256) returns(bool)
func (_Pool *PoolTransactorSession) Payout(_node common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Payout(&_Pool.TransactOpts, _node, _amount)
}

// RejectClient is a paid mutator transaction binding the contract method 0xe086f581.
//
// Solidity: function rejectClient(_client address) returns()
func (_Pool *PoolTransactor) RejectClient(opts *bind.TransactOpts, _client common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "rejectClient", _client)
}

// RejectClient is a paid mutator transaction binding the contract method 0xe086f581.
//
// Solidity: function rejectClient(_client address) returns()
func (_Pool *PoolSession) RejectClient(_client common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RejectClient(&_Pool.TransactOpts, _client)
}

// RejectClient is a paid mutator transaction binding the contract method 0xe086f581.
//
// Solidity: function rejectClient(_client address) returns()
func (_Pool *PoolTransactorSession) RejectClient(_client common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RejectClient(&_Pool.TransactOpts, _client)
}

// RejectNode is a paid mutator transaction binding the contract method 0xaef2a618.
//
// Solidity: function rejectNode(_node address) returns()
func (_Pool *PoolTransactor) RejectNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "rejectNode", _node)
}

// RejectNode is a paid mutator transaction binding the contract method 0xaef2a618.
//
// Solidity: function rejectNode(_node address) returns()
func (_Pool *PoolSession) RejectNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RejectNode(&_Pool.TransactOpts, _node)
}

// RejectNode is a paid mutator transaction binding the contract method 0xaef2a618.
//
// Solidity: function rejectNode(_node address) returns()
func (_Pool *PoolTransactorSession) RejectNode(_node common.Address) (*types.Transaction, error) {
	return _Pool.Contract.RejectNode(&_Pool.TransactOpts, _node)
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

// SetPublicData is a paid mutator transaction binding the contract method 0xfddbd275.
//
// Solidity: function setPublicData(_publicData string) returns()
func (_Pool *PoolTransactor) SetPublicData(opts *bind.TransactOpts, _publicData string) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setPublicData", _publicData)
}

// SetPublicData is a paid mutator transaction binding the contract method 0xfddbd275.
//
// Solidity: function setPublicData(_publicData string) returns()
func (_Pool *PoolSession) SetPublicData(_publicData string) (*types.Transaction, error) {
	return _Pool.Contract.SetPublicData(&_Pool.TransactOpts, _publicData)
}

// SetPublicData is a paid mutator transaction binding the contract method 0xfddbd275.
//
// Solidity: function setPublicData(_publicData string) returns()
func (_Pool *PoolTransactorSession) SetPublicData(_publicData string) (*types.Transaction, error) {
	return _Pool.Contract.SetPublicData(&_Pool.TransactOpts, _publicData)
}

// Work is a paid mutator transaction binding the contract method 0x5858d161.
//
// Solidity: function work(_amount uint256) returns(bool)
func (_Pool *PoolTransactor) Work(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "work", _amount)
}

// Work is a paid mutator transaction binding the contract method 0x5858d161.
//
// Solidity: function work(_amount uint256) returns(bool)
func (_Pool *PoolSession) Work(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Work(&_Pool.TransactOpts, _amount)
}

// Work is a paid mutator transaction binding the contract method 0x5858d161.
//
// Solidity: function work(_amount uint256) returns(bool)
func (_Pool *PoolTransactorSession) Work(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Work(&_Pool.TransactOpts, _amount)
}
