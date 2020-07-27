// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mrc20

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

// Mrc20ABI is the input ABI used to generate the binding from.
const Mrc20ABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferWithSig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"ecrecovery\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disabledHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenIdOrAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"getTokenTransferOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAINID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output2\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output2\",\"type\":\"uint256\"}],\"name\":\"LogFeeTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_childChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setParent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Mrc20 is an auto generated Go binding around an Ethereum contract.
type Mrc20 struct {
	Mrc20Caller     // Read-only binding to the contract
	Mrc20Transactor // Write-only binding to the contract
	Mrc20Filterer   // Log filterer for contract events
}

// Mrc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Mrc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mrc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Mrc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mrc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Mrc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mrc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Mrc20Session struct {
	Contract     *Mrc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mrc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Mrc20CallerSession struct {
	Contract *Mrc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Mrc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Mrc20TransactorSession struct {
	Contract     *Mrc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mrc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Mrc20Raw struct {
	Contract *Mrc20 // Generic contract binding to access the raw methods on
}

// Mrc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Mrc20CallerRaw struct {
	Contract *Mrc20Caller // Generic read-only contract binding to access the raw methods on
}

// Mrc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Mrc20TransactorRaw struct {
	Contract *Mrc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMrc20 creates a new instance of Mrc20, bound to a specific deployed contract.
func NewMrc20(address common.Address, backend bind.ContractBackend) (*Mrc20, error) {
	contract, err := bindMrc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mrc20{Mrc20Caller: Mrc20Caller{contract: contract}, Mrc20Transactor: Mrc20Transactor{contract: contract}, Mrc20Filterer: Mrc20Filterer{contract: contract}}, nil
}

// NewMrc20Caller creates a new read-only instance of Mrc20, bound to a specific deployed contract.
func NewMrc20Caller(address common.Address, caller bind.ContractCaller) (*Mrc20Caller, error) {
	contract, err := bindMrc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Mrc20Caller{contract: contract}, nil
}

// NewMrc20Transactor creates a new write-only instance of Mrc20, bound to a specific deployed contract.
func NewMrc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Mrc20Transactor, error) {
	contract, err := bindMrc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Mrc20Transactor{contract: contract}, nil
}

// NewMrc20Filterer creates a new log filterer instance of Mrc20, bound to a specific deployed contract.
func NewMrc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Mrc20Filterer, error) {
	contract, err := bindMrc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Mrc20Filterer{contract: contract}, nil
}

// bindMrc20 binds a generic wrapper to an already deployed contract.
func bindMrc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Mrc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mrc20 *Mrc20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mrc20.Contract.Mrc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mrc20 *Mrc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mrc20.Contract.Mrc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mrc20 *Mrc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mrc20.Contract.Mrc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mrc20 *Mrc20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mrc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mrc20 *Mrc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mrc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mrc20 *Mrc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mrc20.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() constant returns(uint256)
func (_Mrc20 *Mrc20Caller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "CHAINID")
	return *ret0, err
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() constant returns(uint256)
func (_Mrc20 *Mrc20Session) CHAINID() (*big.Int, error) {
	return _Mrc20.Contract.CHAINID(&_Mrc20.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() constant returns(uint256)
func (_Mrc20 *Mrc20CallerSession) CHAINID() (*big.Int, error) {
	return _Mrc20.Contract.CHAINID(&_Mrc20.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20Caller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "EIP712_DOMAIN_HASH")
	return *ret0, err
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20Session) EIP712DOMAINHASH() ([32]byte, error) {
	return _Mrc20.Contract.EIP712DOMAINHASH(&_Mrc20.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20CallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Mrc20.Contract.EIP712DOMAINHASH(&_Mrc20.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20Caller) EIP712DOMAINSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "EIP712_DOMAIN_SCHEMA_HASH")
	return *ret0, err
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20Session) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _Mrc20.Contract.EIP712DOMAINSCHEMAHASH(&_Mrc20.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20CallerSession) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _Mrc20.Contract.EIP712DOMAINSCHEMAHASH(&_Mrc20.CallOpts)
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20Caller) EIP712TOKENTRANSFERORDERSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH")
	return *ret0, err
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20Session) EIP712TOKENTRANSFERORDERSCHEMAHASH() ([32]byte, error) {
	return _Mrc20.Contract.EIP712TOKENTRANSFERORDERSCHEMAHASH(&_Mrc20.CallOpts)
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() constant returns(bytes32)
func (_Mrc20 *Mrc20CallerSession) EIP712TOKENTRANSFERORDERSCHEMAHASH() ([32]byte, error) {
	return _Mrc20.Contract.EIP712TOKENTRANSFERORDERSCHEMAHASH(&_Mrc20.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Mrc20 *Mrc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Mrc20 *Mrc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mrc20.Contract.BalanceOf(&_Mrc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Mrc20 *Mrc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mrc20.Contract.BalanceOf(&_Mrc20.CallOpts, account)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_Mrc20 *Mrc20Caller) CurrentSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "currentSupply")
	return *ret0, err
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_Mrc20 *Mrc20Session) CurrentSupply() (*big.Int, error) {
	return _Mrc20.Contract.CurrentSupply(&_Mrc20.CallOpts)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() constant returns(uint256)
func (_Mrc20 *Mrc20CallerSession) CurrentSupply() (*big.Int, error) {
	return _Mrc20.Contract.CurrentSupply(&_Mrc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Mrc20 *Mrc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Mrc20 *Mrc20Session) Decimals() (uint8, error) {
	return _Mrc20.Contract.Decimals(&_Mrc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Mrc20 *Mrc20CallerSession) Decimals() (uint8, error) {
	return _Mrc20.Contract.Decimals(&_Mrc20.CallOpts)
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) constant returns(bool)
func (_Mrc20 *Mrc20Caller) DisabledHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "disabledHashes", arg0)
	return *ret0, err
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) constant returns(bool)
func (_Mrc20 *Mrc20Session) DisabledHashes(arg0 [32]byte) (bool, error) {
	return _Mrc20.Contract.DisabledHashes(&_Mrc20.CallOpts, arg0)
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) constant returns(bool)
func (_Mrc20 *Mrc20CallerSession) DisabledHashes(arg0 [32]byte) (bool, error) {
	return _Mrc20.Contract.DisabledHashes(&_Mrc20.CallOpts, arg0)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) constant returns(address result)
func (_Mrc20 *Mrc20Caller) Ecrecovery(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "ecrecovery", hash, sig)
	return *ret0, err
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) constant returns(address result)
func (_Mrc20 *Mrc20Session) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _Mrc20.Contract.Ecrecovery(&_Mrc20.CallOpts, hash, sig)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) constant returns(address result)
func (_Mrc20 *Mrc20CallerSession) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _Mrc20.Contract.Ecrecovery(&_Mrc20.CallOpts, hash, sig)
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) constant returns(bytes32 orderHash)
func (_Mrc20 *Mrc20Caller) GetTokenTransferOrderHash(opts *bind.CallOpts, spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "getTokenTransferOrderHash", spender, tokenIdOrAmount, data, expiration)
	return *ret0, err
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) constant returns(bytes32 orderHash)
func (_Mrc20 *Mrc20Session) GetTokenTransferOrderHash(spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	return _Mrc20.Contract.GetTokenTransferOrderHash(&_Mrc20.CallOpts, spender, tokenIdOrAmount, data, expiration)
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) constant returns(bytes32 orderHash)
func (_Mrc20 *Mrc20CallerSession) GetTokenTransferOrderHash(spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	return _Mrc20.Contract.GetTokenTransferOrderHash(&_Mrc20.CallOpts, spender, tokenIdOrAmount, data, expiration)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Mrc20 *Mrc20Caller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Mrc20 *Mrc20Session) IsOwner() (bool, error) {
	return _Mrc20.Contract.IsOwner(&_Mrc20.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Mrc20 *Mrc20CallerSession) IsOwner() (bool, error) {
	return _Mrc20.Contract.IsOwner(&_Mrc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mrc20 *Mrc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mrc20 *Mrc20Session) Name() (string, error) {
	return _Mrc20.Contract.Name(&_Mrc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mrc20 *Mrc20CallerSession) Name() (string, error) {
	return _Mrc20.Contract.Name(&_Mrc20.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() constant returns(bytes)
func (_Mrc20 *Mrc20Caller) NetworkId(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "networkId")
	return *ret0, err
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() constant returns(bytes)
func (_Mrc20 *Mrc20Session) NetworkId() ([]byte, error) {
	return _Mrc20.Contract.NetworkId(&_Mrc20.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() constant returns(bytes)
func (_Mrc20 *Mrc20CallerSession) NetworkId() ([]byte, error) {
	return _Mrc20.Contract.NetworkId(&_Mrc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Mrc20 *Mrc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Mrc20 *Mrc20Session) Owner() (common.Address, error) {
	return _Mrc20.Contract.Owner(&_Mrc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Mrc20 *Mrc20CallerSession) Owner() (common.Address, error) {
	return _Mrc20.Contract.Owner(&_Mrc20.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_Mrc20 *Mrc20Caller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "parent")
	return *ret0, err
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_Mrc20 *Mrc20Session) Parent() (common.Address, error) {
	return _Mrc20.Contract.Parent(&_Mrc20.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() constant returns(address)
func (_Mrc20 *Mrc20CallerSession) Parent() (common.Address, error) {
	return _Mrc20.Contract.Parent(&_Mrc20.CallOpts)
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() constant returns(address)
func (_Mrc20 *Mrc20Caller) ParentOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "parentOwner")
	return *ret0, err
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() constant returns(address)
func (_Mrc20 *Mrc20Session) ParentOwner() (common.Address, error) {
	return _Mrc20.Contract.ParentOwner(&_Mrc20.CallOpts)
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() constant returns(address)
func (_Mrc20 *Mrc20CallerSession) ParentOwner() (common.Address, error) {
	return _Mrc20.Contract.ParentOwner(&_Mrc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Mrc20 *Mrc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Mrc20 *Mrc20Session) Symbol() (string, error) {
	return _Mrc20.Contract.Symbol(&_Mrc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Mrc20 *Mrc20CallerSession) Symbol() (string, error) {
	return _Mrc20.Contract.Symbol(&_Mrc20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Mrc20 *Mrc20Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Mrc20 *Mrc20Session) Token() (common.Address, error) {
	return _Mrc20.Contract.Token(&_Mrc20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Mrc20 *Mrc20CallerSession) Token() (common.Address, error) {
	return _Mrc20.Contract.Token(&_Mrc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Mrc20 *Mrc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mrc20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Mrc20 *Mrc20Session) TotalSupply() (*big.Int, error) {
	return _Mrc20.Contract.TotalSupply(&_Mrc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Mrc20 *Mrc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Mrc20.Contract.TotalSupply(&_Mrc20.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Mrc20 *Mrc20Transactor) Deposit(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "deposit", user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Mrc20 *Mrc20Session) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mrc20.Contract.Deposit(&_Mrc20.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Mrc20 *Mrc20TransactorSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mrc20.Contract.Deposit(&_Mrc20.TransactOpts, user, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_Mrc20 *Mrc20Transactor) Initialize(opts *bind.TransactOpts, _childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "initialize", _childChain, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_Mrc20 *Mrc20Session) Initialize(_childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.Initialize(&_Mrc20.TransactOpts, _childChain, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_Mrc20 *Mrc20TransactorSession) Initialize(_childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.Initialize(&_Mrc20.TransactOpts, _childChain, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mrc20 *Mrc20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mrc20 *Mrc20Session) RenounceOwnership() (*types.Transaction, error) {
	return _Mrc20.Contract.RenounceOwnership(&_Mrc20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mrc20 *Mrc20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mrc20.Contract.RenounceOwnership(&_Mrc20.TransactOpts)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_Mrc20 *Mrc20Transactor) SetParent(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "setParent", arg0)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_Mrc20 *Mrc20Session) SetParent(arg0 common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.SetParent(&_Mrc20.TransactOpts, arg0)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_Mrc20 *Mrc20TransactorSession) SetParent(arg0 common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.SetParent(&_Mrc20.TransactOpts, arg0)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Mrc20 *Mrc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Mrc20 *Mrc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mrc20.Contract.Transfer(&_Mrc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Mrc20 *Mrc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mrc20.Contract.Transfer(&_Mrc20.TransactOpts, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mrc20 *Mrc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mrc20 *Mrc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.TransferOwnership(&_Mrc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mrc20 *Mrc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.TransferOwnership(&_Mrc20.TransactOpts, newOwner)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_Mrc20 *Mrc20Transactor) TransferWithSig(opts *bind.TransactOpts, sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "transferWithSig", sig, amount, data, expiration, to)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_Mrc20 *Mrc20Session) TransferWithSig(sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.TransferWithSig(&_Mrc20.TransactOpts, sig, amount, data, expiration, to)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_Mrc20 *Mrc20TransactorSession) TransferWithSig(sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _Mrc20.Contract.TransferWithSig(&_Mrc20.TransactOpts, sig, amount, data, expiration, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Mrc20 *Mrc20Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mrc20.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Mrc20 *Mrc20Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Mrc20.Contract.Withdraw(&_Mrc20.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Mrc20 *Mrc20TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Mrc20.Contract.Withdraw(&_Mrc20.TransactOpts, amount)
}

// Mrc20DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Mrc20 contract.
type Mrc20DepositIterator struct {
	Event *Mrc20Deposit // Event containing the contract specifics and raw log

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
func (it *Mrc20DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mrc20Deposit)
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
		it.Event = new(Mrc20Deposit)
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
func (it *Mrc20DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mrc20DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mrc20Deposit represents a Deposit event raised by the Mrc20 contract.
type Mrc20Deposit struct {
	Token   common.Address
	From    common.Address
	Amount  *big.Int
	Input1  *big.Int
	Output1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Mrc20 *Mrc20Filterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*Mrc20DepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Mrc20.contract.FilterLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Mrc20DepositIterator{contract: _Mrc20.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Mrc20 *Mrc20Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Mrc20Deposit, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Mrc20.contract.WatchLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mrc20Deposit)
				if err := _Mrc20.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Mrc20 *Mrc20Filterer) ParseDeposit(log types.Log) (*Mrc20Deposit, error) {
	event := new(Mrc20Deposit)
	if err := _Mrc20.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Mrc20LogFeeTransferIterator is returned from FilterLogFeeTransfer and is used to iterate over the raw logs and unpacked data for LogFeeTransfer events raised by the Mrc20 contract.
type Mrc20LogFeeTransferIterator struct {
	Event *Mrc20LogFeeTransfer // Event containing the contract specifics and raw log

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
func (it *Mrc20LogFeeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mrc20LogFeeTransfer)
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
		it.Event = new(Mrc20LogFeeTransfer)
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
func (it *Mrc20LogFeeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mrc20LogFeeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mrc20LogFeeTransfer represents a LogFeeTransfer event raised by the Mrc20 contract.
type Mrc20LogFeeTransfer struct {
	Token   common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Input1  *big.Int
	Input2  *big.Int
	Output1 *big.Int
	Output2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogFeeTransfer is a free log retrieval operation binding the contract event 0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63.
//
// Solidity: event LogFeeTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Mrc20 *Mrc20Filterer) FilterLogFeeTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*Mrc20LogFeeTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mrc20.contract.FilterLogs(opts, "LogFeeTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Mrc20LogFeeTransferIterator{contract: _Mrc20.contract, event: "LogFeeTransfer", logs: logs, sub: sub}, nil
}

// WatchLogFeeTransfer is a free log subscription operation binding the contract event 0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63.
//
// Solidity: event LogFeeTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Mrc20 *Mrc20Filterer) WatchLogFeeTransfer(opts *bind.WatchOpts, sink chan<- *Mrc20LogFeeTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mrc20.contract.WatchLogs(opts, "LogFeeTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mrc20LogFeeTransfer)
				if err := _Mrc20.contract.UnpackLog(event, "LogFeeTransfer", log); err != nil {
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

// ParseLogFeeTransfer is a log parse operation binding the contract event 0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63.
//
// Solidity: event LogFeeTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Mrc20 *Mrc20Filterer) ParseLogFeeTransfer(log types.Log) (*Mrc20LogFeeTransfer, error) {
	event := new(Mrc20LogFeeTransfer)
	if err := _Mrc20.contract.UnpackLog(event, "LogFeeTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Mrc20LogTransferIterator is returned from FilterLogTransfer and is used to iterate over the raw logs and unpacked data for LogTransfer events raised by the Mrc20 contract.
type Mrc20LogTransferIterator struct {
	Event *Mrc20LogTransfer // Event containing the contract specifics and raw log

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
func (it *Mrc20LogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mrc20LogTransfer)
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
		it.Event = new(Mrc20LogTransfer)
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
func (it *Mrc20LogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mrc20LogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mrc20LogTransfer represents a LogTransfer event raised by the Mrc20 contract.
type Mrc20LogTransfer struct {
	Token   common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Input1  *big.Int
	Input2  *big.Int
	Output1 *big.Int
	Output2 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogTransfer is a free log retrieval operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Mrc20 *Mrc20Filterer) FilterLogTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*Mrc20LogTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mrc20.contract.FilterLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Mrc20LogTransferIterator{contract: _Mrc20.contract, event: "LogTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTransfer is a free log subscription operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Mrc20 *Mrc20Filterer) WatchLogTransfer(opts *bind.WatchOpts, sink chan<- *Mrc20LogTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mrc20.contract.WatchLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mrc20LogTransfer)
				if err := _Mrc20.contract.UnpackLog(event, "LogTransfer", log); err != nil {
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

// ParseLogTransfer is a log parse operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Mrc20 *Mrc20Filterer) ParseLogTransfer(log types.Log) (*Mrc20LogTransfer, error) {
	event := new(Mrc20LogTransfer)
	if err := _Mrc20.contract.UnpackLog(event, "LogTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Mrc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mrc20 contract.
type Mrc20OwnershipTransferredIterator struct {
	Event *Mrc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Mrc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mrc20OwnershipTransferred)
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
		it.Event = new(Mrc20OwnershipTransferred)
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
func (it *Mrc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mrc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mrc20OwnershipTransferred represents a OwnershipTransferred event raised by the Mrc20 contract.
type Mrc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mrc20 *Mrc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Mrc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mrc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Mrc20OwnershipTransferredIterator{contract: _Mrc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mrc20 *Mrc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Mrc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mrc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mrc20OwnershipTransferred)
				if err := _Mrc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mrc20 *Mrc20Filterer) ParseOwnershipTransferred(log types.Log) (*Mrc20OwnershipTransferred, error) {
	event := new(Mrc20OwnershipTransferred)
	if err := _Mrc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Mrc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mrc20 contract.
type Mrc20TransferIterator struct {
	Event *Mrc20Transfer // Event containing the contract specifics and raw log

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
func (it *Mrc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mrc20Transfer)
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
		it.Event = new(Mrc20Transfer)
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
func (it *Mrc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mrc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mrc20Transfer represents a Transfer event raised by the Mrc20 contract.
type Mrc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mrc20 *Mrc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Mrc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mrc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Mrc20TransferIterator{contract: _Mrc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mrc20 *Mrc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Mrc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mrc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mrc20Transfer)
				if err := _Mrc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mrc20 *Mrc20Filterer) ParseTransfer(log types.Log) (*Mrc20Transfer, error) {
	event := new(Mrc20Transfer)
	if err := _Mrc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Mrc20WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Mrc20 contract.
type Mrc20WithdrawIterator struct {
	Event *Mrc20Withdraw // Event containing the contract specifics and raw log

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
func (it *Mrc20WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mrc20Withdraw)
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
		it.Event = new(Mrc20Withdraw)
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
func (it *Mrc20WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mrc20WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mrc20Withdraw represents a Withdraw event raised by the Mrc20 contract.
type Mrc20Withdraw struct {
	Token   common.Address
	From    common.Address
	Amount  *big.Int
	Input1  *big.Int
	Output1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Mrc20 *Mrc20Filterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*Mrc20WithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Mrc20.contract.FilterLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Mrc20WithdrawIterator{contract: _Mrc20.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Mrc20 *Mrc20Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Mrc20Withdraw, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Mrc20.contract.WatchLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mrc20Withdraw)
				if err := _Mrc20.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Mrc20 *Mrc20Filterer) ParseWithdraw(log types.Log) (*Mrc20Withdraw, error) {
	event := new(Mrc20Withdraw)
	if err := _Mrc20.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
