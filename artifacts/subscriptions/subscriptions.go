// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subscriptions

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

// SubscriptionsABI is the input ABI used to generate the binding from.
const SubscriptionsABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_payers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"expiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_subscriptions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renewBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"renewable\",\"type\":\"bool\"}],\"name\":\"subscribe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration_\",\"type\":\"uint256\"}],\"name\":\"updateDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"renewBefore\",\"type\":\"uint256\"}],\"name\":\"updateRenewBefore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subredditPoints\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"renewBefore\",\"type\":\"uint256\"}],\"name\":\"initializecontract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_subredditPoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renewBeforeVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price_\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"durationVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"renew\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"renewable\",\"type\":\"bool\"}],\"name\":\"Subscribed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"DurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"renewBefore\",\"type\":\"uint256\"}],\"name\":\"RenewBeforeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Subscriptions is an auto generated Go binding around an Ethereum contract.
type Subscriptions struct {
	SubscriptionsCaller     // Read-only binding to the contract
	SubscriptionsTransactor // Write-only binding to the contract
	SubscriptionsFilterer   // Log filterer for contract events
}

// SubscriptionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubscriptionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscriptionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscriptionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscriptionsSession struct {
	Contract     *Subscriptions    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubscriptionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscriptionsCallerSession struct {
	Contract *SubscriptionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SubscriptionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscriptionsTransactorSession struct {
	Contract     *SubscriptionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SubscriptionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubscriptionsRaw struct {
	Contract *Subscriptions // Generic contract binding to access the raw methods on
}

// SubscriptionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscriptionsCallerRaw struct {
	Contract *SubscriptionsCaller // Generic read-only contract binding to access the raw methods on
}

// SubscriptionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscriptionsTransactorRaw struct {
	Contract *SubscriptionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscriptions creates a new instance of Subscriptions, bound to a specific deployed contract.
func NewSubscriptions(address common.Address, backend bind.ContractBackend) (*Subscriptions, error) {
	contract, err := bindSubscriptions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subscriptions{SubscriptionsCaller: SubscriptionsCaller{contract: contract}, SubscriptionsTransactor: SubscriptionsTransactor{contract: contract}, SubscriptionsFilterer: SubscriptionsFilterer{contract: contract}}, nil
}

// NewSubscriptionsCaller creates a new read-only instance of Subscriptions, bound to a specific deployed contract.
func NewSubscriptionsCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionsCaller, error) {
	contract, err := bindSubscriptions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionsCaller{contract: contract}, nil
}

// NewSubscriptionsTransactor creates a new write-only instance of Subscriptions, bound to a specific deployed contract.
func NewSubscriptionsTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionsTransactor, error) {
	contract, err := bindSubscriptions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionsTransactor{contract: contract}, nil
}

// NewSubscriptionsFilterer creates a new log filterer instance of Subscriptions, bound to a specific deployed contract.
func NewSubscriptionsFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionsFilterer, error) {
	contract, err := bindSubscriptions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionsFilterer{contract: contract}, nil
}

// bindSubscriptions binds a generic wrapper to an already deployed contract.
func bindSubscriptions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscriptions *SubscriptionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subscriptions.Contract.SubscriptionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscriptions *SubscriptionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscriptions.Contract.SubscriptionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscriptions *SubscriptionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscriptions.Contract.SubscriptionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subscriptions *SubscriptionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subscriptions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subscriptions *SubscriptionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscriptions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subscriptions *SubscriptionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subscriptions.Contract.contract.Transact(opts, method, params...)
}

// Payers is a free data retrieval call binding the contract method 0x0040de43.
//
// Solidity: function _payers(address ) constant returns(address)
func (_Subscriptions *SubscriptionsCaller) Payers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "_payers", arg0)
	return *ret0, err
}

// Payers is a free data retrieval call binding the contract method 0x0040de43.
//
// Solidity: function _payers(address ) constant returns(address)
func (_Subscriptions *SubscriptionsSession) Payers(arg0 common.Address) (common.Address, error) {
	return _Subscriptions.Contract.Payers(&_Subscriptions.CallOpts, arg0)
}

// Payers is a free data retrieval call binding the contract method 0x0040de43.
//
// Solidity: function _payers(address ) constant returns(address)
func (_Subscriptions *SubscriptionsCallerSession) Payers(arg0 common.Address) (common.Address, error) {
	return _Subscriptions.Contract.Payers(&_Subscriptions.CallOpts, arg0)
}

// SubredditPoints is a free data retrieval call binding the contract method 0x5120d8d1.
//
// Solidity: function _subredditPoints() constant returns(address)
func (_Subscriptions *SubscriptionsCaller) SubredditPoints(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "_subredditPoints")
	return *ret0, err
}

// SubredditPoints is a free data retrieval call binding the contract method 0x5120d8d1.
//
// Solidity: function _subredditPoints() constant returns(address)
func (_Subscriptions *SubscriptionsSession) SubredditPoints() (common.Address, error) {
	return _Subscriptions.Contract.SubredditPoints(&_Subscriptions.CallOpts)
}

// SubredditPoints is a free data retrieval call binding the contract method 0x5120d8d1.
//
// Solidity: function _subredditPoints() constant returns(address)
func (_Subscriptions *SubscriptionsCallerSession) SubredditPoints() (common.Address, error) {
	return _Subscriptions.Contract.SubredditPoints(&_Subscriptions.CallOpts)
}

// Subscriptions is a free data retrieval call binding the contract method 0x0dda029e.
//
// Solidity: function _subscriptions(address ) constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) Subscriptions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "_subscriptions", arg0)
	return *ret0, err
}

// Subscriptions is a free data retrieval call binding the contract method 0x0dda029e.
//
// Solidity: function _subscriptions(address ) constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) Subscriptions(arg0 common.Address) (*big.Int, error) {
	return _Subscriptions.Contract.Subscriptions(&_Subscriptions.CallOpts, arg0)
}

// Subscriptions is a free data retrieval call binding the contract method 0x0dda029e.
//
// Solidity: function _subscriptions(address ) constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) Subscriptions(arg0 common.Address) (*big.Int, error) {
	return _Subscriptions.Contract.Subscriptions(&_Subscriptions.CallOpts, arg0)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) Duration() (*big.Int, error) {
	return _Subscriptions.Contract.Duration(&_Subscriptions.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) Duration() (*big.Int, error) {
	return _Subscriptions.Contract.Duration(&_Subscriptions.CallOpts)
}

// DurationVal is a free data retrieval call binding the contract method 0xfc961994.
//
// Solidity: function durationVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) DurationVal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "durationVal")
	return *ret0, err
}

// DurationVal is a free data retrieval call binding the contract method 0xfc961994.
//
// Solidity: function durationVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) DurationVal() (*big.Int, error) {
	return _Subscriptions.Contract.DurationVal(&_Subscriptions.CallOpts)
}

// DurationVal is a free data retrieval call binding the contract method 0xfc961994.
//
// Solidity: function durationVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) DurationVal() (*big.Int, error) {
	return _Subscriptions.Contract.DurationVal(&_Subscriptions.CallOpts)
}

// Expiration is a free data retrieval call binding the contract method 0x04dbe3fe.
//
// Solidity: function expiration(address account) constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) Expiration(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "expiration", account)
	return *ret0, err
}

// Expiration is a free data retrieval call binding the contract method 0x04dbe3fe.
//
// Solidity: function expiration(address account) constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) Expiration(account common.Address) (*big.Int, error) {
	return _Subscriptions.Contract.Expiration(&_Subscriptions.CallOpts, account)
}

// Expiration is a free data retrieval call binding the contract method 0x04dbe3fe.
//
// Solidity: function expiration(address account) constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) Expiration(account common.Address) (*big.Int, error) {
	return _Subscriptions.Contract.Expiration(&_Subscriptions.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Subscriptions *SubscriptionsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Subscriptions *SubscriptionsSession) IsOwner() (bool, error) {
	return _Subscriptions.Contract.IsOwner(&_Subscriptions.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Subscriptions *SubscriptionsCallerSession) IsOwner() (bool, error) {
	return _Subscriptions.Contract.IsOwner(&_Subscriptions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subscriptions *SubscriptionsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subscriptions *SubscriptionsSession) Owner() (common.Address, error) {
	return _Subscriptions.Contract.Owner(&_Subscriptions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subscriptions *SubscriptionsCallerSession) Owner() (common.Address, error) {
	return _Subscriptions.Contract.Owner(&_Subscriptions.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) Price() (*big.Int, error) {
	return _Subscriptions.Contract.Price(&_Subscriptions.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) Price() (*big.Int, error) {
	return _Subscriptions.Contract.Price(&_Subscriptions.CallOpts)
}

// PriceVal is a free data retrieval call binding the contract method 0x37d7839e.
//
// Solidity: function priceVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) PriceVal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "priceVal")
	return *ret0, err
}

// PriceVal is a free data retrieval call binding the contract method 0x37d7839e.
//
// Solidity: function priceVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) PriceVal() (*big.Int, error) {
	return _Subscriptions.Contract.PriceVal(&_Subscriptions.CallOpts)
}

// PriceVal is a free data retrieval call binding the contract method 0x37d7839e.
//
// Solidity: function priceVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) PriceVal() (*big.Int, error) {
	return _Subscriptions.Contract.PriceVal(&_Subscriptions.CallOpts)
}

// RenewBefore is a free data retrieval call binding the contract method 0x1085c9ba.
//
// Solidity: function renewBefore() constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) RenewBefore(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "renewBefore")
	return *ret0, err
}

// RenewBefore is a free data retrieval call binding the contract method 0x1085c9ba.
//
// Solidity: function renewBefore() constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) RenewBefore() (*big.Int, error) {
	return _Subscriptions.Contract.RenewBefore(&_Subscriptions.CallOpts)
}

// RenewBefore is a free data retrieval call binding the contract method 0x1085c9ba.
//
// Solidity: function renewBefore() constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) RenewBefore() (*big.Int, error) {
	return _Subscriptions.Contract.RenewBefore(&_Subscriptions.CallOpts)
}

// RenewBeforeVal is a free data retrieval call binding the contract method 0x6130e66b.
//
// Solidity: function renewBeforeVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsCaller) RenewBeforeVal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subscriptions.contract.Call(opts, out, "renewBeforeVal")
	return *ret0, err
}

// RenewBeforeVal is a free data retrieval call binding the contract method 0x6130e66b.
//
// Solidity: function renewBeforeVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsSession) RenewBeforeVal() (*big.Int, error) {
	return _Subscriptions.Contract.RenewBeforeVal(&_Subscriptions.CallOpts)
}

// RenewBeforeVal is a free data retrieval call binding the contract method 0x6130e66b.
//
// Solidity: function renewBeforeVal() constant returns(uint256)
func (_Subscriptions *SubscriptionsCallerSession) RenewBeforeVal() (*big.Int, error) {
	return _Subscriptions.Contract.RenewBeforeVal(&_Subscriptions.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x4c33fe94.
//
// Solidity: function cancel(address recipient) returns()
func (_Subscriptions *SubscriptionsTransactor) Cancel(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "cancel", recipient)
}

// Cancel is a paid mutator transaction binding the contract method 0x4c33fe94.
//
// Solidity: function cancel(address recipient) returns()
func (_Subscriptions *SubscriptionsSession) Cancel(recipient common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.Cancel(&_Subscriptions.TransactOpts, recipient)
}

// Cancel is a paid mutator transaction binding the contract method 0x4c33fe94.
//
// Solidity: function cancel(address recipient) returns()
func (_Subscriptions *SubscriptionsTransactorSession) Cancel(recipient common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.Cancel(&_Subscriptions.TransactOpts, recipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Subscriptions *SubscriptionsTransactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Subscriptions *SubscriptionsSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.Initialize(&_Subscriptions.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Subscriptions *SubscriptionsTransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.Initialize(&_Subscriptions.TransactOpts, sender)
}

// Initializecontract is a paid mutator transaction binding the contract method 0x4c7dbef1.
//
// Solidity: function initializecontract(address owner_, address subredditPoints, uint256 price_, uint256 duration_, uint256 renewBefore) returns()
func (_Subscriptions *SubscriptionsTransactor) Initializecontract(opts *bind.TransactOpts, owner_ common.Address, subredditPoints common.Address, price_ *big.Int, duration_ *big.Int, renewBefore *big.Int) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "initializecontract", owner_, subredditPoints, price_, duration_, renewBefore)
}

// Initializecontract is a paid mutator transaction binding the contract method 0x4c7dbef1.
//
// Solidity: function initializecontract(address owner_, address subredditPoints, uint256 price_, uint256 duration_, uint256 renewBefore) returns()
func (_Subscriptions *SubscriptionsSession) Initializecontract(owner_ common.Address, subredditPoints common.Address, price_ *big.Int, duration_ *big.Int, renewBefore *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.Initializecontract(&_Subscriptions.TransactOpts, owner_, subredditPoints, price_, duration_, renewBefore)
}

// Initializecontract is a paid mutator transaction binding the contract method 0x4c7dbef1.
//
// Solidity: function initializecontract(address owner_, address subredditPoints, uint256 price_, uint256 duration_, uint256 renewBefore) returns()
func (_Subscriptions *SubscriptionsTransactorSession) Initializecontract(owner_ common.Address, subredditPoints common.Address, price_ *big.Int, duration_ *big.Int, renewBefore *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.Initializecontract(&_Subscriptions.TransactOpts, owner_, subredditPoints, price_, duration_, renewBefore)
}

// Renew is a paid mutator transaction binding the contract method 0xfdc28c8c.
//
// Solidity: function renew(address recipient) returns()
func (_Subscriptions *SubscriptionsTransactor) Renew(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "renew", recipient)
}

// Renew is a paid mutator transaction binding the contract method 0xfdc28c8c.
//
// Solidity: function renew(address recipient) returns()
func (_Subscriptions *SubscriptionsSession) Renew(recipient common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.Renew(&_Subscriptions.TransactOpts, recipient)
}

// Renew is a paid mutator transaction binding the contract method 0xfdc28c8c.
//
// Solidity: function renew(address recipient) returns()
func (_Subscriptions *SubscriptionsTransactorSession) Renew(recipient common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.Renew(&_Subscriptions.TransactOpts, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscriptions *SubscriptionsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscriptions *SubscriptionsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subscriptions.Contract.RenounceOwnership(&_Subscriptions.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subscriptions *SubscriptionsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subscriptions.Contract.RenounceOwnership(&_Subscriptions.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x19bc469e.
//
// Solidity: function subscribe(address recipient, bool renewable) returns()
func (_Subscriptions *SubscriptionsTransactor) Subscribe(opts *bind.TransactOpts, recipient common.Address, renewable bool) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "subscribe", recipient, renewable)
}

// Subscribe is a paid mutator transaction binding the contract method 0x19bc469e.
//
// Solidity: function subscribe(address recipient, bool renewable) returns()
func (_Subscriptions *SubscriptionsSession) Subscribe(recipient common.Address, renewable bool) (*types.Transaction, error) {
	return _Subscriptions.Contract.Subscribe(&_Subscriptions.TransactOpts, recipient, renewable)
}

// Subscribe is a paid mutator transaction binding the contract method 0x19bc469e.
//
// Solidity: function subscribe(address recipient, bool renewable) returns()
func (_Subscriptions *SubscriptionsTransactorSession) Subscribe(recipient common.Address, renewable bool) (*types.Transaction, error) {
	return _Subscriptions.Contract.Subscribe(&_Subscriptions.TransactOpts, recipient, renewable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscriptions *SubscriptionsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscriptions *SubscriptionsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.TransferOwnership(&_Subscriptions.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subscriptions *SubscriptionsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subscriptions.Contract.TransferOwnership(&_Subscriptions.TransactOpts, newOwner)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x1b50ad09.
//
// Solidity: function updateDuration(uint256 duration_) returns()
func (_Subscriptions *SubscriptionsTransactor) UpdateDuration(opts *bind.TransactOpts, duration_ *big.Int) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "updateDuration", duration_)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x1b50ad09.
//
// Solidity: function updateDuration(uint256 duration_) returns()
func (_Subscriptions *SubscriptionsSession) UpdateDuration(duration_ *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.UpdateDuration(&_Subscriptions.TransactOpts, duration_)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x1b50ad09.
//
// Solidity: function updateDuration(uint256 duration_) returns()
func (_Subscriptions *SubscriptionsTransactorSession) UpdateDuration(duration_ *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.UpdateDuration(&_Subscriptions.TransactOpts, duration_)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 price_) returns()
func (_Subscriptions *SubscriptionsTransactor) UpdatePrice(opts *bind.TransactOpts, price_ *big.Int) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "updatePrice", price_)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 price_) returns()
func (_Subscriptions *SubscriptionsSession) UpdatePrice(price_ *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.UpdatePrice(&_Subscriptions.TransactOpts, price_)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 price_) returns()
func (_Subscriptions *SubscriptionsTransactorSession) UpdatePrice(price_ *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.UpdatePrice(&_Subscriptions.TransactOpts, price_)
}

// UpdateRenewBefore is a paid mutator transaction binding the contract method 0x474738b7.
//
// Solidity: function updateRenewBefore(uint256 renewBefore) returns()
func (_Subscriptions *SubscriptionsTransactor) UpdateRenewBefore(opts *bind.TransactOpts, renewBefore *big.Int) (*types.Transaction, error) {
	return _Subscriptions.contract.Transact(opts, "updateRenewBefore", renewBefore)
}

// UpdateRenewBefore is a paid mutator transaction binding the contract method 0x474738b7.
//
// Solidity: function updateRenewBefore(uint256 renewBefore) returns()
func (_Subscriptions *SubscriptionsSession) UpdateRenewBefore(renewBefore *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.UpdateRenewBefore(&_Subscriptions.TransactOpts, renewBefore)
}

// UpdateRenewBefore is a paid mutator transaction binding the contract method 0x474738b7.
//
// Solidity: function updateRenewBefore(uint256 renewBefore) returns()
func (_Subscriptions *SubscriptionsTransactorSession) UpdateRenewBefore(renewBefore *big.Int) (*types.Transaction, error) {
	return _Subscriptions.Contract.UpdateRenewBefore(&_Subscriptions.TransactOpts, renewBefore)
}

// SubscriptionsCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the Subscriptions contract.
type SubscriptionsCanceledIterator struct {
	Event *SubscriptionsCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubscriptionsCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionsCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubscriptionsCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubscriptionsCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionsCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionsCanceled represents a Canceled event raised by the Subscriptions contract.
type SubscriptionsCanceled struct {
	Recipient common.Address
	ExpiresAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0xf3a6ef5718c05d9183af076f5753197b68b04552a763c34796637d6134bdd0f2.
//
// Solidity: event Canceled(address indexed recipient, uint256 expiresAt)
func (_Subscriptions *SubscriptionsFilterer) FilterCanceled(opts *bind.FilterOpts, recipient []common.Address) (*SubscriptionsCanceledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Subscriptions.contract.FilterLogs(opts, "Canceled", recipientRule)
	if err != nil {
		return nil, err
	}
	return &SubscriptionsCanceledIterator{contract: _Subscriptions.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0xf3a6ef5718c05d9183af076f5753197b68b04552a763c34796637d6134bdd0f2.
//
// Solidity: event Canceled(address indexed recipient, uint256 expiresAt)
func (_Subscriptions *SubscriptionsFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *SubscriptionsCanceled, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Subscriptions.contract.WatchLogs(opts, "Canceled", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionsCanceled)
				if err := _Subscriptions.contract.UnpackLog(event, "Canceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCanceled is a log parse operation binding the contract event 0xf3a6ef5718c05d9183af076f5753197b68b04552a763c34796637d6134bdd0f2.
//
// Solidity: event Canceled(address indexed recipient, uint256 expiresAt)
func (_Subscriptions *SubscriptionsFilterer) ParseCanceled(log types.Log) (*SubscriptionsCanceled, error) {
	event := new(SubscriptionsCanceled)
	if err := _Subscriptions.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubscriptionsDurationUpdatedIterator is returned from FilterDurationUpdated and is used to iterate over the raw logs and unpacked data for DurationUpdated events raised by the Subscriptions contract.
type SubscriptionsDurationUpdatedIterator struct {
	Event *SubscriptionsDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubscriptionsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionsDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubscriptionsDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubscriptionsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionsDurationUpdated represents a DurationUpdated event raised by the Subscriptions contract.
type SubscriptionsDurationUpdated struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDurationUpdated is a free log retrieval operation binding the contract event 0x91abcc2d6823e3a3f11d31b208dd3065d2c6a791f1c7c9fe96a42ce12897eac5.
//
// Solidity: event DurationUpdated(uint256 duration)
func (_Subscriptions *SubscriptionsFilterer) FilterDurationUpdated(opts *bind.FilterOpts) (*SubscriptionsDurationUpdatedIterator, error) {

	logs, sub, err := _Subscriptions.contract.FilterLogs(opts, "DurationUpdated")
	if err != nil {
		return nil, err
	}
	return &SubscriptionsDurationUpdatedIterator{contract: _Subscriptions.contract, event: "DurationUpdated", logs: logs, sub: sub}, nil
}

// WatchDurationUpdated is a free log subscription operation binding the contract event 0x91abcc2d6823e3a3f11d31b208dd3065d2c6a791f1c7c9fe96a42ce12897eac5.
//
// Solidity: event DurationUpdated(uint256 duration)
func (_Subscriptions *SubscriptionsFilterer) WatchDurationUpdated(opts *bind.WatchOpts, sink chan<- *SubscriptionsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Subscriptions.contract.WatchLogs(opts, "DurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionsDurationUpdated)
				if err := _Subscriptions.contract.UnpackLog(event, "DurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDurationUpdated is a log parse operation binding the contract event 0x91abcc2d6823e3a3f11d31b208dd3065d2c6a791f1c7c9fe96a42ce12897eac5.
//
// Solidity: event DurationUpdated(uint256 duration)
func (_Subscriptions *SubscriptionsFilterer) ParseDurationUpdated(log types.Log) (*SubscriptionsDurationUpdated, error) {
	event := new(SubscriptionsDurationUpdated)
	if err := _Subscriptions.contract.UnpackLog(event, "DurationUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubscriptionsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Subscriptions contract.
type SubscriptionsOwnershipTransferredIterator struct {
	Event *SubscriptionsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubscriptionsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubscriptionsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubscriptionsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionsOwnershipTransferred represents a OwnershipTransferred event raised by the Subscriptions contract.
type SubscriptionsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subscriptions *SubscriptionsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubscriptionsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subscriptions.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubscriptionsOwnershipTransferredIterator{contract: _Subscriptions.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subscriptions *SubscriptionsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubscriptionsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subscriptions.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionsOwnershipTransferred)
				if err := _Subscriptions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subscriptions *SubscriptionsFilterer) ParseOwnershipTransferred(log types.Log) (*SubscriptionsOwnershipTransferred, error) {
	event := new(SubscriptionsOwnershipTransferred)
	if err := _Subscriptions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubscriptionsPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the Subscriptions contract.
type SubscriptionsPriceUpdatedIterator struct {
	Event *SubscriptionsPriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubscriptionsPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionsPriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubscriptionsPriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubscriptionsPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionsPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionsPriceUpdated represents a PriceUpdated event raised by the Subscriptions contract.
type SubscriptionsPriceUpdated struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x66cbca4f3c64fecf1dcb9ce094abcf7f68c3450a1d4e3a8e917dd621edb4ebe0.
//
// Solidity: event PriceUpdated(uint256 price)
func (_Subscriptions *SubscriptionsFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*SubscriptionsPriceUpdatedIterator, error) {

	logs, sub, err := _Subscriptions.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &SubscriptionsPriceUpdatedIterator{contract: _Subscriptions.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x66cbca4f3c64fecf1dcb9ce094abcf7f68c3450a1d4e3a8e917dd621edb4ebe0.
//
// Solidity: event PriceUpdated(uint256 price)
func (_Subscriptions *SubscriptionsFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *SubscriptionsPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _Subscriptions.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionsPriceUpdated)
				if err := _Subscriptions.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceUpdated is a log parse operation binding the contract event 0x66cbca4f3c64fecf1dcb9ce094abcf7f68c3450a1d4e3a8e917dd621edb4ebe0.
//
// Solidity: event PriceUpdated(uint256 price)
func (_Subscriptions *SubscriptionsFilterer) ParsePriceUpdated(log types.Log) (*SubscriptionsPriceUpdated, error) {
	event := new(SubscriptionsPriceUpdated)
	if err := _Subscriptions.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubscriptionsRenewBeforeUpdatedIterator is returned from FilterRenewBeforeUpdated and is used to iterate over the raw logs and unpacked data for RenewBeforeUpdated events raised by the Subscriptions contract.
type SubscriptionsRenewBeforeUpdatedIterator struct {
	Event *SubscriptionsRenewBeforeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubscriptionsRenewBeforeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionsRenewBeforeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubscriptionsRenewBeforeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubscriptionsRenewBeforeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionsRenewBeforeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionsRenewBeforeUpdated represents a RenewBeforeUpdated event raised by the Subscriptions contract.
type SubscriptionsRenewBeforeUpdated struct {
	RenewBefore *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRenewBeforeUpdated is a free log retrieval operation binding the contract event 0x910b02595d6b5ce3e26daa3fa59d878d520b015b80981aaf497e918a841b4e51.
//
// Solidity: event RenewBeforeUpdated(uint256 renewBefore)
func (_Subscriptions *SubscriptionsFilterer) FilterRenewBeforeUpdated(opts *bind.FilterOpts) (*SubscriptionsRenewBeforeUpdatedIterator, error) {

	logs, sub, err := _Subscriptions.contract.FilterLogs(opts, "RenewBeforeUpdated")
	if err != nil {
		return nil, err
	}
	return &SubscriptionsRenewBeforeUpdatedIterator{contract: _Subscriptions.contract, event: "RenewBeforeUpdated", logs: logs, sub: sub}, nil
}

// WatchRenewBeforeUpdated is a free log subscription operation binding the contract event 0x910b02595d6b5ce3e26daa3fa59d878d520b015b80981aaf497e918a841b4e51.
//
// Solidity: event RenewBeforeUpdated(uint256 renewBefore)
func (_Subscriptions *SubscriptionsFilterer) WatchRenewBeforeUpdated(opts *bind.WatchOpts, sink chan<- *SubscriptionsRenewBeforeUpdated) (event.Subscription, error) {

	logs, sub, err := _Subscriptions.contract.WatchLogs(opts, "RenewBeforeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionsRenewBeforeUpdated)
				if err := _Subscriptions.contract.UnpackLog(event, "RenewBeforeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRenewBeforeUpdated is a log parse operation binding the contract event 0x910b02595d6b5ce3e26daa3fa59d878d520b015b80981aaf497e918a841b4e51.
//
// Solidity: event RenewBeforeUpdated(uint256 renewBefore)
func (_Subscriptions *SubscriptionsFilterer) ParseRenewBeforeUpdated(log types.Log) (*SubscriptionsRenewBeforeUpdated, error) {
	event := new(SubscriptionsRenewBeforeUpdated)
	if err := _Subscriptions.contract.UnpackLog(event, "RenewBeforeUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubscriptionsSubscribedIterator is returned from FilterSubscribed and is used to iterate over the raw logs and unpacked data for Subscribed events raised by the Subscriptions contract.
type SubscriptionsSubscribedIterator struct {
	Event *SubscriptionsSubscribed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SubscriptionsSubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionsSubscribed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SubscriptionsSubscribed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SubscriptionsSubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionsSubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionsSubscribed represents a Subscribed event raised by the Subscriptions contract.
type SubscriptionsSubscribed struct {
	Recipient    common.Address
	Payer        common.Address
	BurnedPoints *big.Int
	ExpiresAt    *big.Int
	Renewable    bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubscribed is a free log retrieval operation binding the contract event 0xd38b19993495aaecd8b724dae5aba66885a681f4cbe90b11feb232ef36ed4e35.
//
// Solidity: event Subscribed(address indexed recipient, address indexed payer, uint256 burnedPoints, uint256 expiresAt, bool renewable)
func (_Subscriptions *SubscriptionsFilterer) FilterSubscribed(opts *bind.FilterOpts, recipient []common.Address, payer []common.Address) (*SubscriptionsSubscribedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Subscriptions.contract.FilterLogs(opts, "Subscribed", recipientRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &SubscriptionsSubscribedIterator{contract: _Subscriptions.contract, event: "Subscribed", logs: logs, sub: sub}, nil
}

// WatchSubscribed is a free log subscription operation binding the contract event 0xd38b19993495aaecd8b724dae5aba66885a681f4cbe90b11feb232ef36ed4e35.
//
// Solidity: event Subscribed(address indexed recipient, address indexed payer, uint256 burnedPoints, uint256 expiresAt, bool renewable)
func (_Subscriptions *SubscriptionsFilterer) WatchSubscribed(opts *bind.WatchOpts, sink chan<- *SubscriptionsSubscribed, recipient []common.Address, payer []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Subscriptions.contract.WatchLogs(opts, "Subscribed", recipientRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionsSubscribed)
				if err := _Subscriptions.contract.UnpackLog(event, "Subscribed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscribed is a log parse operation binding the contract event 0xd38b19993495aaecd8b724dae5aba66885a681f4cbe90b11feb232ef36ed4e35.
//
// Solidity: event Subscribed(address indexed recipient, address indexed payer, uint256 burnedPoints, uint256 expiresAt, bool renewable)
func (_Subscriptions *SubscriptionsFilterer) ParseSubscribed(log types.Log) (*SubscriptionsSubscribed, error) {
	event := new(SubscriptionsSubscribed)
	if err := _Subscriptions.contract.UnpackLog(event, "Subscribed", log); err != nil {
		return nil, err
	}
	return event, nil
}
