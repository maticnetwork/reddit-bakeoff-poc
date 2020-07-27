// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package childmaticcontract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChildmaticcontractABI is the input ABI used to generate the binding from.
const ChildmaticcontractABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferWithSig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"ecrecovery\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disabledHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenIdOrAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"getTokenTransferOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAINID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EIP712_DOMAIN_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output2\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output2\",\"type\":\"uint256\"}],\"name\":\"LogFeeTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_childChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setParent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Childmaticcontract is an auto generated Go binding around an Ethereum contract.
type Childmaticcontract struct {
	ChildmaticcontractCaller     // Read-only binding to the contract
	ChildmaticcontractTransactor // Write-only binding to the contract
	ChildmaticcontractFilterer   // Log filterer for contract events
}

// ChildmaticcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildmaticcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildmaticcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildmaticcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildmaticcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildmaticcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildmaticcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildmaticcontractSession struct {
	Contract     *Childmaticcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ChildmaticcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildmaticcontractCallerSession struct {
	Contract *ChildmaticcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ChildmaticcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildmaticcontractTransactorSession struct {
	Contract     *ChildmaticcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ChildmaticcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildmaticcontractRaw struct {
	Contract *Childmaticcontract // Generic contract binding to access the raw methods on
}

// ChildmaticcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildmaticcontractCallerRaw struct {
	Contract *ChildmaticcontractCaller // Generic read-only contract binding to access the raw methods on
}

// ChildmaticcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildmaticcontractTransactorRaw struct {
	Contract *ChildmaticcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildmaticcontract creates a new instance of Childmaticcontract, bound to a specific deployed contract.
func NewChildmaticcontract(address common.Address, backend bind.ContractBackend) (*Childmaticcontract, error) {
	contract, err := bindChildmaticcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Childmaticcontract{ChildmaticcontractCaller: ChildmaticcontractCaller{contract: contract}, ChildmaticcontractTransactor: ChildmaticcontractTransactor{contract: contract}, ChildmaticcontractFilterer: ChildmaticcontractFilterer{contract: contract}}, nil
}

// NewChildmaticcontractCaller creates a new read-only instance of Childmaticcontract, bound to a specific deployed contract.
func NewChildmaticcontractCaller(address common.Address, caller bind.ContractCaller) (*ChildmaticcontractCaller, error) {
	contract, err := bindChildmaticcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractCaller{contract: contract}, nil
}

// NewChildmaticcontractTransactor creates a new write-only instance of Childmaticcontract, bound to a specific deployed contract.
func NewChildmaticcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildmaticcontractTransactor, error) {
	contract, err := bindChildmaticcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractTransactor{contract: contract}, nil
}

// NewChildmaticcontractFilterer creates a new log filterer instance of Childmaticcontract, bound to a specific deployed contract.
func NewChildmaticcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildmaticcontractFilterer, error) {
	contract, err := bindChildmaticcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractFilterer{contract: contract}, nil
}

// bindChildmaticcontract binds a generic wrapper to an already deployed contract.
func bindChildmaticcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildmaticcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Childmaticcontract *ChildmaticcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Childmaticcontract.Contract.ChildmaticcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Childmaticcontract *ChildmaticcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.ChildmaticcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Childmaticcontract *ChildmaticcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.ChildmaticcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Childmaticcontract *ChildmaticcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Childmaticcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Childmaticcontract *ChildmaticcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Childmaticcontract *ChildmaticcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "CHAINID")
	return *ret0, err
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractSession) CHAINID() (*big.Int, error) {
	return _Childmaticcontract.Contract.CHAINID(&_Childmaticcontract.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0xcc79f97b.
//
// Solidity: function CHAINID() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCallerSession) CHAINID() (*big.Int, error) {
	return _Childmaticcontract.Contract.CHAINID(&_Childmaticcontract.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractCaller) EIP712DOMAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "EIP712_DOMAIN_HASH")
	return *ret0, err
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Childmaticcontract.Contract.EIP712DOMAINHASH(&_Childmaticcontract.CallOpts)
}

// EIP712DOMAINHASH is a free data retrieval call binding the contract method 0xe306f779.
//
// Solidity: function EIP712_DOMAIN_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractCallerSession) EIP712DOMAINHASH() ([32]byte, error) {
	return _Childmaticcontract.Contract.EIP712DOMAINHASH(&_Childmaticcontract.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractCaller) EIP712DOMAINSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "EIP712_DOMAIN_SCHEMA_HASH")
	return *ret0, err
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractSession) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _Childmaticcontract.Contract.EIP712DOMAINSCHEMAHASH(&_Childmaticcontract.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractCallerSession) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _Childmaticcontract.Contract.EIP712DOMAINSCHEMAHASH(&_Childmaticcontract.CallOpts)
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractCaller) EIP712TOKENTRANSFERORDERSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH")
	return *ret0, err
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractSession) EIP712TOKENTRANSFERORDERSCHEMAHASH() ([32]byte, error) {
	return _Childmaticcontract.Contract.EIP712TOKENTRANSFERORDERSCHEMAHASH(&_Childmaticcontract.CallOpts)
}

// EIP712TOKENTRANSFERORDERSCHEMAHASH is a free data retrieval call binding the contract method 0xabceeba2.
//
// Solidity: function EIP712_TOKEN_TRANSFER_ORDER_SCHEMA_HASH() view returns(bytes32)
func (_Childmaticcontract *ChildmaticcontractCallerSession) EIP712TOKENTRANSFERORDERSCHEMAHASH() ([32]byte, error) {
	return _Childmaticcontract.Contract.EIP712TOKENTRANSFERORDERSCHEMAHASH(&_Childmaticcontract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Childmaticcontract.Contract.BalanceOf(&_Childmaticcontract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Childmaticcontract.Contract.BalanceOf(&_Childmaticcontract.CallOpts, account)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCaller) CurrentSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "currentSupply")
	return *ret0, err
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractSession) CurrentSupply() (*big.Int, error) {
	return _Childmaticcontract.Contract.CurrentSupply(&_Childmaticcontract.CallOpts)
}

// CurrentSupply is a free data retrieval call binding the contract method 0x771282f6.
//
// Solidity: function currentSupply() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCallerSession) CurrentSupply() (*big.Int, error) {
	return _Childmaticcontract.Contract.CurrentSupply(&_Childmaticcontract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Childmaticcontract *ChildmaticcontractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Childmaticcontract *ChildmaticcontractSession) Decimals() (uint8, error) {
	return _Childmaticcontract.Contract.Decimals(&_Childmaticcontract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Childmaticcontract *ChildmaticcontractCallerSession) Decimals() (uint8, error) {
	return _Childmaticcontract.Contract.Decimals(&_Childmaticcontract.CallOpts)
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) view returns(bool)
func (_Childmaticcontract *ChildmaticcontractCaller) DisabledHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "disabledHashes", arg0)
	return *ret0, err
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) view returns(bool)
func (_Childmaticcontract *ChildmaticcontractSession) DisabledHashes(arg0 [32]byte) (bool, error) {
	return _Childmaticcontract.Contract.DisabledHashes(&_Childmaticcontract.CallOpts, arg0)
}

// DisabledHashes is a free data retrieval call binding the contract method 0xacd06cb3.
//
// Solidity: function disabledHashes(bytes32 ) view returns(bool)
func (_Childmaticcontract *ChildmaticcontractCallerSession) DisabledHashes(arg0 [32]byte) (bool, error) {
	return _Childmaticcontract.Contract.DisabledHashes(&_Childmaticcontract.CallOpts, arg0)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) pure returns(address result)
func (_Childmaticcontract *ChildmaticcontractCaller) Ecrecovery(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "ecrecovery", hash, sig)
	return *ret0, err
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) pure returns(address result)
func (_Childmaticcontract *ChildmaticcontractSession) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _Childmaticcontract.Contract.Ecrecovery(&_Childmaticcontract.CallOpts, hash, sig)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(bytes32 hash, bytes sig) pure returns(address result)
func (_Childmaticcontract *ChildmaticcontractCallerSession) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _Childmaticcontract.Contract.Ecrecovery(&_Childmaticcontract.CallOpts, hash, sig)
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) view returns(bytes32 orderHash)
func (_Childmaticcontract *ChildmaticcontractCaller) GetTokenTransferOrderHash(opts *bind.CallOpts, spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "getTokenTransferOrderHash", spender, tokenIdOrAmount, data, expiration)
	return *ret0, err
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) view returns(bytes32 orderHash)
func (_Childmaticcontract *ChildmaticcontractSession) GetTokenTransferOrderHash(spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	return _Childmaticcontract.Contract.GetTokenTransferOrderHash(&_Childmaticcontract.CallOpts, spender, tokenIdOrAmount, data, expiration)
}

// GetTokenTransferOrderHash is a free data retrieval call binding the contract method 0xb789543c.
//
// Solidity: function getTokenTransferOrderHash(address spender, uint256 tokenIdOrAmount, bytes32 data, uint256 expiration) view returns(bytes32 orderHash)
func (_Childmaticcontract *ChildmaticcontractCallerSession) GetTokenTransferOrderHash(spender common.Address, tokenIdOrAmount *big.Int, data [32]byte, expiration *big.Int) ([32]byte, error) {
	return _Childmaticcontract.Contract.GetTokenTransferOrderHash(&_Childmaticcontract.CallOpts, spender, tokenIdOrAmount, data, expiration)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Childmaticcontract *ChildmaticcontractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Childmaticcontract *ChildmaticcontractSession) IsOwner() (bool, error) {
	return _Childmaticcontract.Contract.IsOwner(&_Childmaticcontract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Childmaticcontract *ChildmaticcontractCallerSession) IsOwner() (bool, error) {
	return _Childmaticcontract.Contract.IsOwner(&_Childmaticcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Childmaticcontract *ChildmaticcontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Childmaticcontract *ChildmaticcontractSession) Name() (string, error) {
	return _Childmaticcontract.Contract.Name(&_Childmaticcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Childmaticcontract *ChildmaticcontractCallerSession) Name() (string, error) {
	return _Childmaticcontract.Contract.Name(&_Childmaticcontract.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_Childmaticcontract *ChildmaticcontractCaller) NetworkId(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "networkId")
	return *ret0, err
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_Childmaticcontract *ChildmaticcontractSession) NetworkId() ([]byte, error) {
	return _Childmaticcontract.Contract.NetworkId(&_Childmaticcontract.CallOpts)
}

// NetworkId is a free data retrieval call binding the contract method 0x9025e64c.
//
// Solidity: function networkId() view returns(bytes)
func (_Childmaticcontract *ChildmaticcontractCallerSession) NetworkId() ([]byte, error) {
	return _Childmaticcontract.Contract.NetworkId(&_Childmaticcontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Childmaticcontract *ChildmaticcontractSession) Owner() (common.Address, error) {
	return _Childmaticcontract.Contract.Owner(&_Childmaticcontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCallerSession) Owner() (common.Address, error) {
	return _Childmaticcontract.Contract.Owner(&_Childmaticcontract.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCaller) Parent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "parent")
	return *ret0, err
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_Childmaticcontract *ChildmaticcontractSession) Parent() (common.Address, error) {
	return _Childmaticcontract.Contract.Parent(&_Childmaticcontract.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0x60f96a8f.
//
// Solidity: function parent() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCallerSession) Parent() (common.Address, error) {
	return _Childmaticcontract.Contract.Parent(&_Childmaticcontract.CallOpts)
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCaller) ParentOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "parentOwner")
	return *ret0, err
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() view returns(address)
func (_Childmaticcontract *ChildmaticcontractSession) ParentOwner() (common.Address, error) {
	return _Childmaticcontract.Contract.ParentOwner(&_Childmaticcontract.CallOpts)
}

// ParentOwner is a free data retrieval call binding the contract method 0x7019d41a.
//
// Solidity: function parentOwner() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCallerSession) ParentOwner() (common.Address, error) {
	return _Childmaticcontract.Contract.ParentOwner(&_Childmaticcontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Childmaticcontract *ChildmaticcontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Childmaticcontract *ChildmaticcontractSession) Symbol() (string, error) {
	return _Childmaticcontract.Contract.Symbol(&_Childmaticcontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Childmaticcontract *ChildmaticcontractCallerSession) Symbol() (string, error) {
	return _Childmaticcontract.Contract.Symbol(&_Childmaticcontract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Childmaticcontract *ChildmaticcontractSession) Token() (common.Address, error) {
	return _Childmaticcontract.Contract.Token(&_Childmaticcontract.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Childmaticcontract *ChildmaticcontractCallerSession) Token() (common.Address, error) {
	return _Childmaticcontract.Contract.Token(&_Childmaticcontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Childmaticcontract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractSession) TotalSupply() (*big.Int, error) {
	return _Childmaticcontract.Contract.TotalSupply(&_Childmaticcontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Childmaticcontract *ChildmaticcontractCallerSession) TotalSupply() (*big.Int, error) {
	return _Childmaticcontract.Contract.TotalSupply(&_Childmaticcontract.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Childmaticcontract *ChildmaticcontractTransactor) Deposit(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "deposit", user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Childmaticcontract *ChildmaticcontractSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Deposit(&_Childmaticcontract.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Childmaticcontract *ChildmaticcontractTransactorSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Deposit(&_Childmaticcontract.TransactOpts, user, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_Childmaticcontract *ChildmaticcontractTransactor) Initialize(opts *bind.TransactOpts, _childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "initialize", _childChain, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_Childmaticcontract *ChildmaticcontractSession) Initialize(_childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Initialize(&_Childmaticcontract.TransactOpts, _childChain, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _childChain, address _token) returns()
func (_Childmaticcontract *ChildmaticcontractTransactorSession) Initialize(_childChain common.Address, _token common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Initialize(&_Childmaticcontract.TransactOpts, _childChain, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Childmaticcontract *ChildmaticcontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Childmaticcontract *ChildmaticcontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Childmaticcontract.Contract.RenounceOwnership(&_Childmaticcontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Childmaticcontract *ChildmaticcontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Childmaticcontract.Contract.RenounceOwnership(&_Childmaticcontract.TransactOpts)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_Childmaticcontract *ChildmaticcontractTransactor) SetParent(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "setParent", arg0)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_Childmaticcontract *ChildmaticcontractSession) SetParent(arg0 common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.SetParent(&_Childmaticcontract.TransactOpts, arg0)
}

// SetParent is a paid mutator transaction binding the contract method 0x1499c592.
//
// Solidity: function setParent(address ) returns()
func (_Childmaticcontract *ChildmaticcontractTransactorSession) SetParent(arg0 common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.SetParent(&_Childmaticcontract.TransactOpts, arg0)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) payable returns(bool)
func (_Childmaticcontract *ChildmaticcontractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) payable returns(bool)
func (_Childmaticcontract *ChildmaticcontractSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Transfer(&_Childmaticcontract.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) payable returns(bool)
func (_Childmaticcontract *ChildmaticcontractTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Transfer(&_Childmaticcontract.TransactOpts, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Childmaticcontract *ChildmaticcontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Childmaticcontract *ChildmaticcontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.TransferOwnership(&_Childmaticcontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Childmaticcontract *ChildmaticcontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.TransferOwnership(&_Childmaticcontract.TransactOpts, newOwner)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_Childmaticcontract *ChildmaticcontractTransactor) TransferWithSig(opts *bind.TransactOpts, sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "transferWithSig", sig, amount, data, expiration, to)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_Childmaticcontract *ChildmaticcontractSession) TransferWithSig(sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.TransferWithSig(&_Childmaticcontract.TransactOpts, sig, amount, data, expiration, to)
}

// TransferWithSig is a paid mutator transaction binding the contract method 0x19d27d9c.
//
// Solidity: function transferWithSig(bytes sig, uint256 amount, bytes32 data, uint256 expiration, address to) returns(address from)
func (_Childmaticcontract *ChildmaticcontractTransactorSession) TransferWithSig(sig []byte, amount *big.Int, data [32]byte, expiration *big.Int, to common.Address) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.TransferWithSig(&_Childmaticcontract.TransactOpts, sig, amount, data, expiration, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_Childmaticcontract *ChildmaticcontractTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_Childmaticcontract *ChildmaticcontractSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Withdraw(&_Childmaticcontract.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) payable returns()
func (_Childmaticcontract *ChildmaticcontractTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Childmaticcontract.Contract.Withdraw(&_Childmaticcontract.TransactOpts, amount)
}

// ChildmaticcontractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Childmaticcontract contract.
type ChildmaticcontractDepositIterator struct {
	Event *ChildmaticcontractDeposit // Event containing the contract specifics and raw log

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
func (it *ChildmaticcontractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildmaticcontractDeposit)
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
		it.Event = new(ChildmaticcontractDeposit)
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
func (it *ChildmaticcontractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildmaticcontractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildmaticcontractDeposit represents a Deposit event raised by the Childmaticcontract contract.
type ChildmaticcontractDeposit struct {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*ChildmaticcontractDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Childmaticcontract.contract.FilterLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractDepositIterator{contract: _Childmaticcontract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Childmaticcontract *ChildmaticcontractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ChildmaticcontractDeposit, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Childmaticcontract.contract.WatchLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildmaticcontractDeposit)
				if err := _Childmaticcontract.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) ParseDeposit(log types.Log) (*ChildmaticcontractDeposit, error) {
	event := new(ChildmaticcontractDeposit)
	if err := _Childmaticcontract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChildmaticcontractLogFeeTransferIterator is returned from FilterLogFeeTransfer and is used to iterate over the raw logs and unpacked data for LogFeeTransfer events raised by the Childmaticcontract contract.
type ChildmaticcontractLogFeeTransferIterator struct {
	Event *ChildmaticcontractLogFeeTransfer // Event containing the contract specifics and raw log

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
func (it *ChildmaticcontractLogFeeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildmaticcontractLogFeeTransfer)
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
		it.Event = new(ChildmaticcontractLogFeeTransfer)
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
func (it *ChildmaticcontractLogFeeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildmaticcontractLogFeeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildmaticcontractLogFeeTransfer represents a LogFeeTransfer event raised by the Childmaticcontract contract.
type ChildmaticcontractLogFeeTransfer struct {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) FilterLogFeeTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*ChildmaticcontractLogFeeTransferIterator, error) {

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

	logs, sub, err := _Childmaticcontract.contract.FilterLogs(opts, "LogFeeTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractLogFeeTransferIterator{contract: _Childmaticcontract.contract, event: "LogFeeTransfer", logs: logs, sub: sub}, nil
}

// WatchLogFeeTransfer is a free log subscription operation binding the contract event 0x4dfe1bbbcf077ddc3e01291eea2d5c70c2b422b415d95645b9adcfd678cb1d63.
//
// Solidity: event LogFeeTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Childmaticcontract *ChildmaticcontractFilterer) WatchLogFeeTransfer(opts *bind.WatchOpts, sink chan<- *ChildmaticcontractLogFeeTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Childmaticcontract.contract.WatchLogs(opts, "LogFeeTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildmaticcontractLogFeeTransfer)
				if err := _Childmaticcontract.contract.UnpackLog(event, "LogFeeTransfer", log); err != nil {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) ParseLogFeeTransfer(log types.Log) (*ChildmaticcontractLogFeeTransfer, error) {
	event := new(ChildmaticcontractLogFeeTransfer)
	if err := _Childmaticcontract.contract.UnpackLog(event, "LogFeeTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChildmaticcontractLogTransferIterator is returned from FilterLogTransfer and is used to iterate over the raw logs and unpacked data for LogTransfer events raised by the Childmaticcontract contract.
type ChildmaticcontractLogTransferIterator struct {
	Event *ChildmaticcontractLogTransfer // Event containing the contract specifics and raw log

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
func (it *ChildmaticcontractLogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildmaticcontractLogTransfer)
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
		it.Event = new(ChildmaticcontractLogTransfer)
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
func (it *ChildmaticcontractLogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildmaticcontractLogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildmaticcontractLogTransfer represents a LogTransfer event raised by the Childmaticcontract contract.
type ChildmaticcontractLogTransfer struct {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) FilterLogTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*ChildmaticcontractLogTransferIterator, error) {

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

	logs, sub, err := _Childmaticcontract.contract.FilterLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractLogTransferIterator{contract: _Childmaticcontract.contract, event: "LogTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTransfer is a free log subscription operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Childmaticcontract *ChildmaticcontractFilterer) WatchLogTransfer(opts *bind.WatchOpts, sink chan<- *ChildmaticcontractLogTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Childmaticcontract.contract.WatchLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildmaticcontractLogTransfer)
				if err := _Childmaticcontract.contract.UnpackLog(event, "LogTransfer", log); err != nil {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) ParseLogTransfer(log types.Log) (*ChildmaticcontractLogTransfer, error) {
	event := new(ChildmaticcontractLogTransfer)
	if err := _Childmaticcontract.contract.UnpackLog(event, "LogTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChildmaticcontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Childmaticcontract contract.
type ChildmaticcontractOwnershipTransferredIterator struct {
	Event *ChildmaticcontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChildmaticcontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildmaticcontractOwnershipTransferred)
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
		it.Event = new(ChildmaticcontractOwnershipTransferred)
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
func (it *ChildmaticcontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildmaticcontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildmaticcontractOwnershipTransferred represents a OwnershipTransferred event raised by the Childmaticcontract contract.
type ChildmaticcontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Childmaticcontract *ChildmaticcontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChildmaticcontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Childmaticcontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractOwnershipTransferredIterator{contract: _Childmaticcontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Childmaticcontract *ChildmaticcontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChildmaticcontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Childmaticcontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildmaticcontractOwnershipTransferred)
				if err := _Childmaticcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) ParseOwnershipTransferred(log types.Log) (*ChildmaticcontractOwnershipTransferred, error) {
	event := new(ChildmaticcontractOwnershipTransferred)
	if err := _Childmaticcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChildmaticcontractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Childmaticcontract contract.
type ChildmaticcontractTransferIterator struct {
	Event *ChildmaticcontractTransfer // Event containing the contract specifics and raw log

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
func (it *ChildmaticcontractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildmaticcontractTransfer)
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
		it.Event = new(ChildmaticcontractTransfer)
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
func (it *ChildmaticcontractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildmaticcontractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildmaticcontractTransfer represents a Transfer event raised by the Childmaticcontract contract.
type ChildmaticcontractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Childmaticcontract *ChildmaticcontractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChildmaticcontractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Childmaticcontract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractTransferIterator{contract: _Childmaticcontract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Childmaticcontract *ChildmaticcontractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChildmaticcontractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Childmaticcontract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildmaticcontractTransfer)
				if err := _Childmaticcontract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) ParseTransfer(log types.Log) (*ChildmaticcontractTransfer, error) {
	event := new(ChildmaticcontractTransfer)
	if err := _Childmaticcontract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChildmaticcontractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Childmaticcontract contract.
type ChildmaticcontractWithdrawIterator struct {
	Event *ChildmaticcontractWithdraw // Event containing the contract specifics and raw log

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
func (it *ChildmaticcontractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildmaticcontractWithdraw)
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
		it.Event = new(ChildmaticcontractWithdraw)
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
func (it *ChildmaticcontractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildmaticcontractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildmaticcontractWithdraw represents a Withdraw event raised by the Childmaticcontract contract.
type ChildmaticcontractWithdraw struct {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*ChildmaticcontractWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Childmaticcontract.contract.FilterLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ChildmaticcontractWithdrawIterator{contract: _Childmaticcontract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Childmaticcontract *ChildmaticcontractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ChildmaticcontractWithdraw, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Childmaticcontract.contract.WatchLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildmaticcontractWithdraw)
				if err := _Childmaticcontract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Childmaticcontract *ChildmaticcontractFilterer) ParseWithdraw(log types.Log) (*ChildmaticcontractWithdraw, error) {
	event := new(ChildmaticcontractWithdraw)
	if err := _Childmaticcontract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
