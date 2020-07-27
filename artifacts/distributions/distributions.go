// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package distributions

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

// DistributionsABI is the input ABI used to generate the binding from.
const DistributionsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"karmaSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"subredditPoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"karma\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_roundPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalKarma\",\"type\":\"uint256\"}],\"name\":\"advanceToRound\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"subreddit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"karma\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subredditPoints_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialSupply_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialKarma_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"karmaSource_\",\"type\":\"address\"}],\"name\":\"initializecontract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPoints\",\"type\":\"uint256\"}],\"name\":\"AdvanceRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"points\",\"type\":\"uint256\"}],\"name\":\"ClaimPoints\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Distributions is an auto generated Go binding around an Ethereum contract.
type Distributions struct {
	DistributionsCaller     // Read-only binding to the contract
	DistributionsTransactor // Write-only binding to the contract
	DistributionsFilterer   // Log filterer for contract events
}

// DistributionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributionsSession struct {
	Contract     *Distributions    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributionsCallerSession struct {
	Contract *DistributionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DistributionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributionsTransactorSession struct {
	Contract     *DistributionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DistributionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributionsRaw struct {
	Contract *Distributions // Generic contract binding to access the raw methods on
}

// DistributionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributionsCallerRaw struct {
	Contract *DistributionsCaller // Generic read-only contract binding to access the raw methods on
}

// DistributionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributionsTransactorRaw struct {
	Contract *DistributionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributions creates a new instance of Distributions, bound to a specific deployed contract.
func NewDistributions(address common.Address, backend bind.ContractBackend) (*Distributions, error) {
	contract, err := bindDistributions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distributions{DistributionsCaller: DistributionsCaller{contract: contract}, DistributionsTransactor: DistributionsTransactor{contract: contract}, DistributionsFilterer: DistributionsFilterer{contract: contract}}, nil
}

// NewDistributionsCaller creates a new read-only instance of Distributions, bound to a specific deployed contract.
func NewDistributionsCaller(address common.Address, caller bind.ContractCaller) (*DistributionsCaller, error) {
	contract, err := bindDistributions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionsCaller{contract: contract}, nil
}

// NewDistributionsTransactor creates a new write-only instance of Distributions, bound to a specific deployed contract.
func NewDistributionsTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributionsTransactor, error) {
	contract, err := bindDistributions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionsTransactor{contract: contract}, nil
}

// NewDistributionsFilterer creates a new log filterer instance of Distributions, bound to a specific deployed contract.
func NewDistributionsFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributionsFilterer, error) {
	contract, err := bindDistributions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributionsFilterer{contract: contract}, nil
}

// bindDistributions binds a generic wrapper to an already deployed contract.
func bindDistributions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistributionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributions *DistributionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Distributions.Contract.DistributionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributions *DistributionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributions.Contract.DistributionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributions *DistributionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributions.Contract.DistributionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributions *DistributionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Distributions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributions *DistributionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributions *DistributionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributions.Contract.contract.Transact(opts, method, params...)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_Distributions *DistributionsCaller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Distributions.contract.Call(opts, out, "initialSupply")
	return *ret0, err
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_Distributions *DistributionsSession) InitialSupply() (*big.Int, error) {
	return _Distributions.Contract.InitialSupply(&_Distributions.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_Distributions *DistributionsCallerSession) InitialSupply() (*big.Int, error) {
	return _Distributions.Contract.InitialSupply(&_Distributions.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Distributions *DistributionsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Distributions.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Distributions *DistributionsSession) IsOwner() (bool, error) {
	return _Distributions.Contract.IsOwner(&_Distributions.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Distributions *DistributionsCallerSession) IsOwner() (bool, error) {
	return _Distributions.Contract.IsOwner(&_Distributions.CallOpts)
}

// KarmaSource is a free data retrieval call binding the contract method 0x0015d5b1.
//
// Solidity: function karmaSource() constant returns(address)
func (_Distributions *DistributionsCaller) KarmaSource(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Distributions.contract.Call(opts, out, "karmaSource")
	return *ret0, err
}

// KarmaSource is a free data retrieval call binding the contract method 0x0015d5b1.
//
// Solidity: function karmaSource() constant returns(address)
func (_Distributions *DistributionsSession) KarmaSource() (common.Address, error) {
	return _Distributions.Contract.KarmaSource(&_Distributions.CallOpts)
}

// KarmaSource is a free data retrieval call binding the contract method 0x0015d5b1.
//
// Solidity: function karmaSource() constant returns(address)
func (_Distributions *DistributionsCallerSession) KarmaSource() (common.Address, error) {
	return _Distributions.Contract.KarmaSource(&_Distributions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Distributions *DistributionsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Distributions.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Distributions *DistributionsSession) Owner() (common.Address, error) {
	return _Distributions.Contract.Owner(&_Distributions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Distributions *DistributionsCallerSession) Owner() (common.Address, error) {
	return _Distributions.Contract.Owner(&_Distributions.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() constant returns(string)
func (_Distributions *DistributionsCaller) Subreddit(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Distributions.contract.Call(opts, out, "subreddit")
	return *ret0, err
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() constant returns(string)
func (_Distributions *DistributionsSession) Subreddit() (string, error) {
	return _Distributions.Contract.Subreddit(&_Distributions.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() constant returns(string)
func (_Distributions *DistributionsCallerSession) Subreddit() (string, error) {
	return _Distributions.Contract.Subreddit(&_Distributions.CallOpts)
}

// SubredditPoints is a free data retrieval call binding the contract method 0x89304061.
//
// Solidity: function subredditPoints() constant returns(address)
func (_Distributions *DistributionsCaller) SubredditPoints(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Distributions.contract.Call(opts, out, "subredditPoints")
	return *ret0, err
}

// SubredditPoints is a free data retrieval call binding the contract method 0x89304061.
//
// Solidity: function subredditPoints() constant returns(address)
func (_Distributions *DistributionsSession) SubredditPoints() (common.Address, error) {
	return _Distributions.Contract.SubredditPoints(&_Distributions.CallOpts)
}

// SubredditPoints is a free data retrieval call binding the contract method 0x89304061.
//
// Solidity: function subredditPoints() constant returns(address)
func (_Distributions *DistributionsCallerSession) SubredditPoints() (common.Address, error) {
	return _Distributions.Contract.SubredditPoints(&_Distributions.CallOpts)
}

// AdvanceToRound is a paid mutator transaction binding the contract method 0xba1454ae.
//
// Solidity: function advanceToRound(uint256 round, uint256 _roundPoints, uint256 _totalKarma) returns()
func (_Distributions *DistributionsTransactor) AdvanceToRound(opts *bind.TransactOpts, round *big.Int, _roundPoints *big.Int, _totalKarma *big.Int) (*types.Transaction, error) {
	return _Distributions.contract.Transact(opts, "advanceToRound", round, _roundPoints, _totalKarma)
}

// AdvanceToRound is a paid mutator transaction binding the contract method 0xba1454ae.
//
// Solidity: function advanceToRound(uint256 round, uint256 _roundPoints, uint256 _totalKarma) returns()
func (_Distributions *DistributionsSession) AdvanceToRound(round *big.Int, _roundPoints *big.Int, _totalKarma *big.Int) (*types.Transaction, error) {
	return _Distributions.Contract.AdvanceToRound(&_Distributions.TransactOpts, round, _roundPoints, _totalKarma)
}

// AdvanceToRound is a paid mutator transaction binding the contract method 0xba1454ae.
//
// Solidity: function advanceToRound(uint256 round, uint256 _roundPoints, uint256 _totalKarma) returns()
func (_Distributions *DistributionsTransactorSession) AdvanceToRound(round *big.Int, _roundPoints *big.Int, _totalKarma *big.Int) (*types.Transaction, error) {
	return _Distributions.Contract.AdvanceToRound(&_Distributions.TransactOpts, round, _roundPoints, _totalKarma)
}

// Claim is a paid mutator transaction binding the contract method 0x99016142.
//
// Solidity: function claim(uint256 round, address account, uint256 karma, bytes signature) returns()
func (_Distributions *DistributionsTransactor) Claim(opts *bind.TransactOpts, round *big.Int, account common.Address, karma *big.Int, signature []byte) (*types.Transaction, error) {
	return _Distributions.contract.Transact(opts, "claim", round, account, karma, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x99016142.
//
// Solidity: function claim(uint256 round, address account, uint256 karma, bytes signature) returns()
func (_Distributions *DistributionsSession) Claim(round *big.Int, account common.Address, karma *big.Int, signature []byte) (*types.Transaction, error) {
	return _Distributions.Contract.Claim(&_Distributions.TransactOpts, round, account, karma, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x99016142.
//
// Solidity: function claim(uint256 round, address account, uint256 karma, bytes signature) returns()
func (_Distributions *DistributionsTransactorSession) Claim(round *big.Int, account common.Address, karma *big.Int, signature []byte) (*types.Transaction, error) {
	return _Distributions.Contract.Claim(&_Distributions.TransactOpts, round, account, karma, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Distributions *DistributionsTransactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Distributions.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Distributions *DistributionsSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Distributions.Contract.Initialize(&_Distributions.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Distributions *DistributionsTransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Distributions.Contract.Initialize(&_Distributions.TransactOpts, sender)
}

// Initializecontract is a paid mutator transaction binding the contract method 0xfb687d0b.
//
// Solidity: function initializecontract(address owner_, address subredditPoints_, uint256 initialSupply_, uint256 initialKarma_, address karmaSource_) returns()
func (_Distributions *DistributionsTransactor) Initializecontract(opts *bind.TransactOpts, owner_ common.Address, subredditPoints_ common.Address, initialSupply_ *big.Int, initialKarma_ *big.Int, karmaSource_ common.Address) (*types.Transaction, error) {
	return _Distributions.contract.Transact(opts, "initializecontract", owner_, subredditPoints_, initialSupply_, initialKarma_, karmaSource_)
}

// Initializecontract is a paid mutator transaction binding the contract method 0xfb687d0b.
//
// Solidity: function initializecontract(address owner_, address subredditPoints_, uint256 initialSupply_, uint256 initialKarma_, address karmaSource_) returns()
func (_Distributions *DistributionsSession) Initializecontract(owner_ common.Address, subredditPoints_ common.Address, initialSupply_ *big.Int, initialKarma_ *big.Int, karmaSource_ common.Address) (*types.Transaction, error) {
	return _Distributions.Contract.Initializecontract(&_Distributions.TransactOpts, owner_, subredditPoints_, initialSupply_, initialKarma_, karmaSource_)
}

// Initializecontract is a paid mutator transaction binding the contract method 0xfb687d0b.
//
// Solidity: function initializecontract(address owner_, address subredditPoints_, uint256 initialSupply_, uint256 initialKarma_, address karmaSource_) returns()
func (_Distributions *DistributionsTransactorSession) Initializecontract(owner_ common.Address, subredditPoints_ common.Address, initialSupply_ *big.Int, initialKarma_ *big.Int, karmaSource_ common.Address) (*types.Transaction, error) {
	return _Distributions.Contract.Initializecontract(&_Distributions.TransactOpts, owner_, subredditPoints_, initialSupply_, initialKarma_, karmaSource_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distributions *DistributionsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributions.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distributions *DistributionsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distributions.Contract.RenounceOwnership(&_Distributions.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distributions *DistributionsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distributions.Contract.RenounceOwnership(&_Distributions.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distributions *DistributionsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Distributions.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distributions *DistributionsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distributions.Contract.TransferOwnership(&_Distributions.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distributions *DistributionsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distributions.Contract.TransferOwnership(&_Distributions.TransactOpts, newOwner)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xec0e642f.
//
// Solidity: function verifySignature(address account, uint256 round, uint256 karma, bytes signature) returns(address)
func (_Distributions *DistributionsTransactor) VerifySignature(opts *bind.TransactOpts, account common.Address, round *big.Int, karma *big.Int, signature []byte) (*types.Transaction, error) {
	return _Distributions.contract.Transact(opts, "verifySignature", account, round, karma, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xec0e642f.
//
// Solidity: function verifySignature(address account, uint256 round, uint256 karma, bytes signature) returns(address)
func (_Distributions *DistributionsSession) VerifySignature(account common.Address, round *big.Int, karma *big.Int, signature []byte) (*types.Transaction, error) {
	return _Distributions.Contract.VerifySignature(&_Distributions.TransactOpts, account, round, karma, signature)
}

// VerifySignature is a paid mutator transaction binding the contract method 0xec0e642f.
//
// Solidity: function verifySignature(address account, uint256 round, uint256 karma, bytes signature) returns(address)
func (_Distributions *DistributionsTransactorSession) VerifySignature(account common.Address, round *big.Int, karma *big.Int, signature []byte) (*types.Transaction, error) {
	return _Distributions.Contract.VerifySignature(&_Distributions.TransactOpts, account, round, karma, signature)
}

// DistributionsAdvanceRoundIterator is returned from FilterAdvanceRound and is used to iterate over the raw logs and unpacked data for AdvanceRound events raised by the Distributions contract.
type DistributionsAdvanceRoundIterator struct {
	Event *DistributionsAdvanceRound // Event containing the contract specifics and raw log

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
func (it *DistributionsAdvanceRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionsAdvanceRound)
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
		it.Event = new(DistributionsAdvanceRound)
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
func (it *DistributionsAdvanceRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionsAdvanceRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionsAdvanceRound represents a AdvanceRound event raised by the Distributions contract.
type DistributionsAdvanceRound struct {
	Round       *big.Int
	TotalPoints *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAdvanceRound is a free log retrieval operation binding the contract event 0x89b058ec860e65024050f0e6e93c04b4575b854e5e9938f34214c4562130a484.
//
// Solidity: event AdvanceRound(uint256 round, uint256 totalPoints)
func (_Distributions *DistributionsFilterer) FilterAdvanceRound(opts *bind.FilterOpts) (*DistributionsAdvanceRoundIterator, error) {

	logs, sub, err := _Distributions.contract.FilterLogs(opts, "AdvanceRound")
	if err != nil {
		return nil, err
	}
	return &DistributionsAdvanceRoundIterator{contract: _Distributions.contract, event: "AdvanceRound", logs: logs, sub: sub}, nil
}

// WatchAdvanceRound is a free log subscription operation binding the contract event 0x89b058ec860e65024050f0e6e93c04b4575b854e5e9938f34214c4562130a484.
//
// Solidity: event AdvanceRound(uint256 round, uint256 totalPoints)
func (_Distributions *DistributionsFilterer) WatchAdvanceRound(opts *bind.WatchOpts, sink chan<- *DistributionsAdvanceRound) (event.Subscription, error) {

	logs, sub, err := _Distributions.contract.WatchLogs(opts, "AdvanceRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionsAdvanceRound)
				if err := _Distributions.contract.UnpackLog(event, "AdvanceRound", log); err != nil {
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

// ParseAdvanceRound is a log parse operation binding the contract event 0x89b058ec860e65024050f0e6e93c04b4575b854e5e9938f34214c4562130a484.
//
// Solidity: event AdvanceRound(uint256 round, uint256 totalPoints)
func (_Distributions *DistributionsFilterer) ParseAdvanceRound(log types.Log) (*DistributionsAdvanceRound, error) {
	event := new(DistributionsAdvanceRound)
	if err := _Distributions.contract.UnpackLog(event, "AdvanceRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DistributionsClaimPointsIterator is returned from FilterClaimPoints and is used to iterate over the raw logs and unpacked data for ClaimPoints events raised by the Distributions contract.
type DistributionsClaimPointsIterator struct {
	Event *DistributionsClaimPoints // Event containing the contract specifics and raw log

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
func (it *DistributionsClaimPointsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionsClaimPoints)
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
		it.Event = new(DistributionsClaimPoints)
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
func (it *DistributionsClaimPointsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionsClaimPointsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionsClaimPoints represents a ClaimPoints event raised by the Distributions contract.
type DistributionsClaimPoints struct {
	Round  *big.Int
	User   common.Address
	Points *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimPoints is a free log retrieval operation binding the contract event 0x25a2e12a9fce51ee2c4bc04d844e80c24a5f62ce66bf1f25e9da0ed7694d1810.
//
// Solidity: event ClaimPoints(uint256 round, address indexed user, uint256 points)
func (_Distributions *DistributionsFilterer) FilterClaimPoints(opts *bind.FilterOpts, user []common.Address) (*DistributionsClaimPointsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Distributions.contract.FilterLogs(opts, "ClaimPoints", userRule)
	if err != nil {
		return nil, err
	}
	return &DistributionsClaimPointsIterator{contract: _Distributions.contract, event: "ClaimPoints", logs: logs, sub: sub}, nil
}

// WatchClaimPoints is a free log subscription operation binding the contract event 0x25a2e12a9fce51ee2c4bc04d844e80c24a5f62ce66bf1f25e9da0ed7694d1810.
//
// Solidity: event ClaimPoints(uint256 round, address indexed user, uint256 points)
func (_Distributions *DistributionsFilterer) WatchClaimPoints(opts *bind.WatchOpts, sink chan<- *DistributionsClaimPoints, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Distributions.contract.WatchLogs(opts, "ClaimPoints", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionsClaimPoints)
				if err := _Distributions.contract.UnpackLog(event, "ClaimPoints", log); err != nil {
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

// ParseClaimPoints is a log parse operation binding the contract event 0x25a2e12a9fce51ee2c4bc04d844e80c24a5f62ce66bf1f25e9da0ed7694d1810.
//
// Solidity: event ClaimPoints(uint256 round, address indexed user, uint256 points)
func (_Distributions *DistributionsFilterer) ParseClaimPoints(log types.Log) (*DistributionsClaimPoints, error) {
	event := new(DistributionsClaimPoints)
	if err := _Distributions.contract.UnpackLog(event, "ClaimPoints", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DistributionsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Distributions contract.
type DistributionsOwnershipTransferredIterator struct {
	Event *DistributionsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DistributionsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionsOwnershipTransferred)
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
		it.Event = new(DistributionsOwnershipTransferred)
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
func (it *DistributionsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionsOwnershipTransferred represents a OwnershipTransferred event raised by the Distributions contract.
type DistributionsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distributions *DistributionsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DistributionsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distributions.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DistributionsOwnershipTransferredIterator{contract: _Distributions.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distributions *DistributionsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DistributionsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distributions.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionsOwnershipTransferred)
				if err := _Distributions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Distributions *DistributionsFilterer) ParseOwnershipTransferred(log types.Log) (*DistributionsOwnershipTransferred, error) {
	event := new(DistributionsOwnershipTransferred)
	if err := _Distributions.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
