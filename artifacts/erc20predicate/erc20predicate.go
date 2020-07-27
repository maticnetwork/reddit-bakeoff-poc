// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20predicate

import (
	"math/big"
	"strings"

	ethereum "github.com/maticnetwork/bor"
	"github.com/maticnetwork/bor/accounts/abi"
	"github.com/maticnetwork/bor/accounts/abi/bind"
	"github.com/maticnetwork/bor/common"
	"github.com/maticnetwork/bor/core/types"
	"github.com/maticnetwork/bor/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Erc20predicateABI is the input ABI used to generate the binding from.
const Erc20predicateABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFinalizeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAINID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"startExitWithBurntTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"exitTx\",\"type\":\"bytes\"}],\"name\":\"startExitForOutgoingErc20Transfer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"exitTx\",\"type\":\"bytes\"}],\"name\":\"startExitForIncomingErc20Transfer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"exit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"inputUtxo\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"challengeData\",\"type\":\"bytes\"}],\"name\":\"verifyDeprecation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"interpretStateUpdate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Erc20predicate is an auto generated Go binding around an Ethereum contract.
type Erc20predicate struct {
	Erc20predicateCaller     // Read-only binding to the contract
	Erc20predicateTransactor // Write-only binding to the contract
	Erc20predicateFilterer   // Log filterer for contract events
}

// Erc20predicateCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20predicateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20predicateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20predicateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20predicateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20predicateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20predicateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20predicateSession struct {
	Contract     *Erc20predicate   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20predicateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20predicateCallerSession struct {
	Contract *Erc20predicateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Erc20predicateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20predicateTransactorSession struct {
	Contract     *Erc20predicateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Erc20predicateRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20predicateRaw struct {
	Contract *Erc20predicate // Generic contract binding to access the raw methods on
}

// Erc20predicateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20predicateCallerRaw struct {
	Contract *Erc20predicateCaller // Generic read-only contract binding to access the raw methods on
}

// Erc20predicateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20predicateTransactorRaw struct {
	Contract *Erc20predicateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20predicate creates a new instance of Erc20predicate, bound to a specific deployed contract.
func NewErc20predicate(address common.Address, backend bind.ContractBackend) (*Erc20predicate, error) {
	contract, err := bindErc20predicate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20predicate{Erc20predicateCaller: Erc20predicateCaller{contract: contract}, Erc20predicateTransactor: Erc20predicateTransactor{contract: contract}, Erc20predicateFilterer: Erc20predicateFilterer{contract: contract}}, nil
}

// NewErc20predicateCaller creates a new read-only instance of Erc20predicate, bound to a specific deployed contract.
func NewErc20predicateCaller(address common.Address, caller bind.ContractCaller) (*Erc20predicateCaller, error) {
	contract, err := bindErc20predicate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20predicateCaller{contract: contract}, nil
}

// NewErc20predicateTransactor creates a new write-only instance of Erc20predicate, bound to a specific deployed contract.
func NewErc20predicateTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc20predicateTransactor, error) {
	contract, err := bindErc20predicate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20predicateTransactor{contract: contract}, nil
}

// NewErc20predicateFilterer creates a new log filterer instance of Erc20predicate, bound to a specific deployed contract.
func NewErc20predicateFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc20predicateFilterer, error) {
	contract, err := bindErc20predicate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20predicateFilterer{contract: contract}, nil
}

// bindErc20predicate binds a generic wrapper to an already deployed contract.
func bindErc20predicate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20predicateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20predicate *Erc20predicateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc20predicate.Contract.Erc20predicateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20predicate *Erc20predicateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20predicate.Contract.Erc20predicateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20predicate *Erc20predicateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20predicate.Contract.Erc20predicateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20predicate *Erc20predicateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Erc20predicate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20predicate *Erc20predicateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20predicate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20predicate *Erc20predicateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20predicate.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() constant returns(uint256)
func (_Erc20predicate *Erc20predicateCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Erc20predicate.contract.Call(opts, out, "CHAINID")
	return *ret0, err
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() constant returns(uint256)
func (_Erc20predicate *Erc20predicateSession) CHAINID() (*big.Int, error) {
	return _Erc20predicate.Contract.CHAINID(&_Erc20predicate.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() constant returns(uint256)
func (_Erc20predicate *Erc20predicateCallerSession) CHAINID() (*big.Int, error) {
	return _Erc20predicate.Contract.CHAINID(&_Erc20predicate.CallOpts)
}

// InterpretStateUpdate is a free data retrieval call binding the contract method 0x82e3464c.
//
// Solidity: function interpretStateUpdate(bytes state) constant returns(bytes)
func (_Erc20predicate *Erc20predicateCaller) InterpretStateUpdate(opts *bind.CallOpts, state []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Erc20predicate.contract.Call(opts, out, "interpretStateUpdate", state)
	return *ret0, err
}

// InterpretStateUpdate is a free data retrieval call binding the contract method 0x82e3464c.
//
// Solidity: function interpretStateUpdate(bytes state) constant returns(bytes)
func (_Erc20predicate *Erc20predicateSession) InterpretStateUpdate(state []byte) ([]byte, error) {
	return _Erc20predicate.Contract.InterpretStateUpdate(&_Erc20predicate.CallOpts, state)
}

// InterpretStateUpdate is a free data retrieval call binding the contract method 0x82e3464c.
//
// Solidity: function interpretStateUpdate(bytes state) constant returns(bytes)
func (_Erc20predicate *Erc20predicateCallerSession) InterpretStateUpdate(state []byte) ([]byte, error) {
	return _Erc20predicate.Contract.InterpretStateUpdate(&_Erc20predicate.CallOpts, state)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() constant returns(bytes)
func (_Erc20predicate *Erc20predicateCaller) NetworkId(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Erc20predicate.contract.Call(opts, out, "networkId")
	return *ret0, err
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() constant returns(bytes)
func (_Erc20predicate *Erc20predicateSession) NetworkId() ([]byte, error) {
	return _Erc20predicate.Contract.NetworkId(&_Erc20predicate.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() constant returns(bytes)
func (_Erc20predicate *Erc20predicateCallerSession) NetworkId() ([]byte, error) {
	return _Erc20predicate.Contract.NetworkId(&_Erc20predicate.CallOpts)
}

// OnFinalizeExit is a paid mutator transaction binding the contract method 0x7bd94e03.
//
// Solidity: function onFinalizeExit(bytes data) returns()
func (_Erc20predicate *Erc20predicateTransactor) OnFinalizeExit(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Erc20predicate.contract.Transact(opts, "onFinalizeExit", data)
}

// OnFinalizeExit is a paid mutator transaction binding the contract method 0x7bd94e03.
//
// Solidity: function onFinalizeExit(bytes data) returns()
func (_Erc20predicate *Erc20predicateSession) OnFinalizeExit(data []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.OnFinalizeExit(&_Erc20predicate.TransactOpts, data)
}

// OnFinalizeExit is a paid mutator transaction binding the contract method 0x7bd94e03.
//
// Solidity: function onFinalizeExit(bytes data) returns()
func (_Erc20predicate *Erc20predicateTransactorSession) OnFinalizeExit(data []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.OnFinalizeExit(&_Erc20predicate.TransactOpts, data)
}

// StartExitForIncomingErc20Transfer is a paid mutator transaction binding the contract method 0x378383b3.
//
// Solidity: function startExitForIncomingErc20Transfer(bytes data, bytes exitTx) returns(address, uint256)
func (_Erc20predicate *Erc20predicateTransactor) StartExitForIncomingErc20Transfer(opts *bind.TransactOpts, data []byte, exitTx []byte) (*types.Transaction, error) {
	return _Erc20predicate.contract.Transact(opts, "startExitForIncomingErc20Transfer", data, exitTx)
}

// StartExitForIncomingErc20Transfer is a paid mutator transaction binding the contract method 0x378383b3.
//
// Solidity: function startExitForIncomingErc20Transfer(bytes data, bytes exitTx) returns(address, uint256)
func (_Erc20predicate *Erc20predicateSession) StartExitForIncomingErc20Transfer(data []byte, exitTx []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.StartExitForIncomingErc20Transfer(&_Erc20predicate.TransactOpts, data, exitTx)
}

// StartExitForIncomingErc20Transfer is a paid mutator transaction binding the contract method 0x378383b3.
//
// Solidity: function startExitForIncomingErc20Transfer(bytes data, bytes exitTx) returns(address, uint256)
func (_Erc20predicate *Erc20predicateTransactorSession) StartExitForIncomingErc20Transfer(data []byte, exitTx []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.StartExitForIncomingErc20Transfer(&_Erc20predicate.TransactOpts, data, exitTx)
}

// StartExitForOutgoingErc20Transfer is a paid mutator transaction binding the contract method 0x3cccdaaf.
//
// Solidity: function startExitForOutgoingErc20Transfer(bytes data, bytes exitTx) returns(address, uint256)
func (_Erc20predicate *Erc20predicateTransactor) StartExitForOutgoingErc20Transfer(opts *bind.TransactOpts, data []byte, exitTx []byte) (*types.Transaction, error) {
	return _Erc20predicate.contract.Transact(opts, "startExitForOutgoingErc20Transfer", data, exitTx)
}

// StartExitForOutgoingErc20Transfer is a paid mutator transaction binding the contract method 0x3cccdaaf.
//
// Solidity: function startExitForOutgoingErc20Transfer(bytes data, bytes exitTx) returns(address, uint256)
func (_Erc20predicate *Erc20predicateSession) StartExitForOutgoingErc20Transfer(data []byte, exitTx []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.StartExitForOutgoingErc20Transfer(&_Erc20predicate.TransactOpts, data, exitTx)
}

// StartExitForOutgoingErc20Transfer is a paid mutator transaction binding the contract method 0x3cccdaaf.
//
// Solidity: function startExitForOutgoingErc20Transfer(bytes data, bytes exitTx) returns(address, uint256)
func (_Erc20predicate *Erc20predicateTransactorSession) StartExitForOutgoingErc20Transfer(data []byte, exitTx []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.StartExitForOutgoingErc20Transfer(&_Erc20predicate.TransactOpts, data, exitTx)
}

// StartExitWithBurntTokens is a paid mutator transaction binding the contract method 0x7c5264b4.
//
// Solidity: function startExitWithBurntTokens(bytes data) returns()
func (_Erc20predicate *Erc20predicateTransactor) StartExitWithBurntTokens(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Erc20predicate.contract.Transact(opts, "startExitWithBurntTokens", data)
}

// StartExitWithBurntTokens is a paid mutator transaction binding the contract method 0x7c5264b4.
//
// Solidity: function startExitWithBurntTokens(bytes data) returns()
func (_Erc20predicate *Erc20predicateSession) StartExitWithBurntTokens(data []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.StartExitWithBurntTokens(&_Erc20predicate.TransactOpts, data)
}

// StartExitWithBurntTokens is a paid mutator transaction binding the contract method 0x7c5264b4.
//
// Solidity: function startExitWithBurntTokens(bytes data) returns()
func (_Erc20predicate *Erc20predicateTransactorSession) StartExitWithBurntTokens(data []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.StartExitWithBurntTokens(&_Erc20predicate.TransactOpts, data)
}

// VerifyDeprecation is a paid mutator transaction binding the contract method 0xec58410c.
//
// Solidity: function verifyDeprecation(bytes exit, bytes inputUtxo, bytes challengeData) returns(bool)
func (_Erc20predicate *Erc20predicateTransactor) VerifyDeprecation(opts *bind.TransactOpts, exit []byte, inputUtxo []byte, challengeData []byte) (*types.Transaction, error) {
	return _Erc20predicate.contract.Transact(opts, "verifyDeprecation", exit, inputUtxo, challengeData)
}

// VerifyDeprecation is a paid mutator transaction binding the contract method 0xec58410c.
//
// Solidity: function verifyDeprecation(bytes exit, bytes inputUtxo, bytes challengeData) returns(bool)
func (_Erc20predicate *Erc20predicateSession) VerifyDeprecation(exit []byte, inputUtxo []byte, challengeData []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.VerifyDeprecation(&_Erc20predicate.TransactOpts, exit, inputUtxo, challengeData)
}

// VerifyDeprecation is a paid mutator transaction binding the contract method 0xec58410c.
//
// Solidity: function verifyDeprecation(bytes exit, bytes inputUtxo, bytes challengeData) returns(bool)
func (_Erc20predicate *Erc20predicateTransactorSession) VerifyDeprecation(exit []byte, inputUtxo []byte, challengeData []byte) (*types.Transaction, error) {
	return _Erc20predicate.Contract.VerifyDeprecation(&_Erc20predicate.TransactOpts, exit, inputUtxo, challengeData)
}
