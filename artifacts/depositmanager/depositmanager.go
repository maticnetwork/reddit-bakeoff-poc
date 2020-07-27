// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositmanager

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

// DepositmanagerABI is the input ABI used to generate the binding from.
const DepositmanagerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"MaxErc20DepositUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOrNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositBlockId\",\"type\":\"uint256\"}],\"name\":\"NewDepositBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"childChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"depositHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxErc20Deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootChain\",\"outputs\":[{\"internalType\":\"contractRootChain\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stateSender\",\"outputs\":[{\"internalType\":\"contractStateSender\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxDepositAmount\",\"type\":\"uint256\"}],\"name\":\"updateMaxErc20Deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOrNFTId\",\"type\":\"uint256\"}],\"name\":\"transferAssets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amountOrTokens\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositBulk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateChildChainAndStateSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20ForUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"depositERC721ForUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositEther\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rootChain\",\"type\":\"address\"}],\"name\":\"updateRootChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Depositmanager is an auto generated Go binding around an Ethereum contract.
type Depositmanager struct {
	DepositmanagerCaller     // Read-only binding to the contract
	DepositmanagerTransactor // Write-only binding to the contract
	DepositmanagerFilterer   // Log filterer for contract events
}

// DepositmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositmanagerSession struct {
	Contract     *Depositmanager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositmanagerCallerSession struct {
	Contract *DepositmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositmanagerTransactorSession struct {
	Contract     *DepositmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositmanagerRaw struct {
	Contract *Depositmanager // Generic contract binding to access the raw methods on
}

// DepositmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositmanagerCallerRaw struct {
	Contract *DepositmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// DepositmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositmanagerTransactorRaw struct {
	Contract *DepositmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositmanager creates a new instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanager(address common.Address, backend bind.ContractBackend) (*Depositmanager, error) {
	contract, err := bindDepositmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositmanager{DepositmanagerCaller: DepositmanagerCaller{contract: contract}, DepositmanagerTransactor: DepositmanagerTransactor{contract: contract}, DepositmanagerFilterer: DepositmanagerFilterer{contract: contract}}, nil
}

// NewDepositmanagerCaller creates a new read-only instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanagerCaller(address common.Address, caller bind.ContractCaller) (*DepositmanagerCaller, error) {
	contract, err := bindDepositmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerCaller{contract: contract}, nil
}

// NewDepositmanagerTransactor creates a new write-only instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositmanagerTransactor, error) {
	contract, err := bindDepositmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerTransactor{contract: contract}, nil
}

// NewDepositmanagerFilterer creates a new log filterer instance of Depositmanager, bound to a specific deployed contract.
func NewDepositmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositmanagerFilterer, error) {
	contract, err := bindDepositmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerFilterer{contract: contract}, nil
}

// bindDepositmanager binds a generic wrapper to an already deployed contract.
func bindDepositmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositmanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositmanager *DepositmanagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Depositmanager.Contract.DepositmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositmanager *DepositmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositmanager *DepositmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositmanager *DepositmanagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Depositmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositmanager *DepositmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositmanager *DepositmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositmanager.Contract.contract.Transact(opts, method, params...)
}

// ChildChain is a free data retrieval call binding the contract method 0x42fc47fb.
//
// Solidity: function childChain() constant returns(address)
func (_Depositmanager *DepositmanagerCaller) ChildChain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "childChain")
	return *ret0, err
}

// ChildChain is a free data retrieval call binding the contract method 0x42fc47fb.
//
// Solidity: function childChain() constant returns(address)
func (_Depositmanager *DepositmanagerSession) ChildChain() (common.Address, error) {
	return _Depositmanager.Contract.ChildChain(&_Depositmanager.CallOpts)
}

// ChildChain is a free data retrieval call binding the contract method 0x42fc47fb.
//
// Solidity: function childChain() constant returns(address)
func (_Depositmanager *DepositmanagerCallerSession) ChildChain() (common.Address, error) {
	return _Depositmanager.Contract.ChildChain(&_Depositmanager.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) constant returns(bytes32 depositHash, uint256 createdAt)
func (_Depositmanager *DepositmanagerCaller) Deposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DepositHash [32]byte
	CreatedAt   *big.Int
}, error) {
	ret := new(struct {
		DepositHash [32]byte
		CreatedAt   *big.Int
	})
	out := ret
	err := _Depositmanager.contract.Call(opts, out, "deposits", arg0)
	return *ret, err
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) constant returns(bytes32 depositHash, uint256 createdAt)
func (_Depositmanager *DepositmanagerSession) Deposits(arg0 *big.Int) (struct {
	DepositHash [32]byte
	CreatedAt   *big.Int
}, error) {
	return _Depositmanager.Contract.Deposits(&_Depositmanager.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) constant returns(bytes32 depositHash, uint256 createdAt)
func (_Depositmanager *DepositmanagerCallerSession) Deposits(arg0 *big.Int) (struct {
	DepositHash [32]byte
	CreatedAt   *big.Int
}, error) {
	return _Depositmanager.Contract.Deposits(&_Depositmanager.CallOpts, arg0)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() constant returns(address)
func (_Depositmanager *DepositmanagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() constant returns(address)
func (_Depositmanager *DepositmanagerSession) Governance() (common.Address, error) {
	return _Depositmanager.Contract.Governance(&_Depositmanager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() constant returns(address)
func (_Depositmanager *DepositmanagerCallerSession) Governance() (common.Address, error) {
	return _Depositmanager.Contract.Governance(&_Depositmanager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Depositmanager *DepositmanagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Depositmanager *DepositmanagerSession) IsOwner() (bool, error) {
	return _Depositmanager.Contract.IsOwner(&_Depositmanager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Depositmanager *DepositmanagerCallerSession) IsOwner() (bool, error) {
	return _Depositmanager.Contract.IsOwner(&_Depositmanager.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Depositmanager *DepositmanagerCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Depositmanager *DepositmanagerSession) Locked() (bool, error) {
	return _Depositmanager.Contract.Locked(&_Depositmanager.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_Depositmanager *DepositmanagerCallerSession) Locked() (bool, error) {
	return _Depositmanager.Contract.Locked(&_Depositmanager.CallOpts)
}

// MaxErc20Deposit is a free data retrieval call binding the contract method 0xe7af7ba1.
//
// Solidity: function maxErc20Deposit() constant returns(uint256)
func (_Depositmanager *DepositmanagerCaller) MaxErc20Deposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "maxErc20Deposit")
	return *ret0, err
}

// MaxErc20Deposit is a free data retrieval call binding the contract method 0xe7af7ba1.
//
// Solidity: function maxErc20Deposit() constant returns(uint256)
func (_Depositmanager *DepositmanagerSession) MaxErc20Deposit() (*big.Int, error) {
	return _Depositmanager.Contract.MaxErc20Deposit(&_Depositmanager.CallOpts)
}

// MaxErc20Deposit is a free data retrieval call binding the contract method 0xe7af7ba1.
//
// Solidity: function maxErc20Deposit() constant returns(uint256)
func (_Depositmanager *DepositmanagerCallerSession) MaxErc20Deposit() (*big.Int, error) {
	return _Depositmanager.Contract.MaxErc20Deposit(&_Depositmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Depositmanager *DepositmanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Depositmanager *DepositmanagerSession) Owner() (common.Address, error) {
	return _Depositmanager.Contract.Owner(&_Depositmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Depositmanager *DepositmanagerCallerSession) Owner() (common.Address, error) {
	return _Depositmanager.Contract.Owner(&_Depositmanager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Depositmanager *DepositmanagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Depositmanager *DepositmanagerSession) Registry() (common.Address, error) {
	return _Depositmanager.Contract.Registry(&_Depositmanager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Depositmanager *DepositmanagerCallerSession) Registry() (common.Address, error) {
	return _Depositmanager.Contract.Registry(&_Depositmanager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Depositmanager *DepositmanagerCaller) RootChain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "rootChain")
	return *ret0, err
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Depositmanager *DepositmanagerSession) RootChain() (common.Address, error) {
	return _Depositmanager.Contract.RootChain(&_Depositmanager.CallOpts)
}

// RootChain is a free data retrieval call binding the contract method 0x987ab9db.
//
// Solidity: function rootChain() constant returns(address)
func (_Depositmanager *DepositmanagerCallerSession) RootChain() (common.Address, error) {
	return _Depositmanager.Contract.RootChain(&_Depositmanager.CallOpts)
}

// StateSender is a free data retrieval call binding the contract method 0xcb10f94c.
//
// Solidity: function stateSender() constant returns(address)
func (_Depositmanager *DepositmanagerCaller) StateSender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Depositmanager.contract.Call(opts, out, "stateSender")
	return *ret0, err
}

// StateSender is a free data retrieval call binding the contract method 0xcb10f94c.
//
// Solidity: function stateSender() constant returns(address)
func (_Depositmanager *DepositmanagerSession) StateSender() (common.Address, error) {
	return _Depositmanager.Contract.StateSender(&_Depositmanager.CallOpts)
}

// StateSender is a free data retrieval call binding the contract method 0xcb10f94c.
//
// Solidity: function stateSender() constant returns(address)
func (_Depositmanager *DepositmanagerCallerSession) StateSender() (common.Address, error) {
	return _Depositmanager.Contract.StateSender(&_Depositmanager.CallOpts)
}

// DepositBulk is a paid mutator transaction binding the contract method 0x7b1f7117.
//
// Solidity: function depositBulk(address[] _tokens, uint256[] _amountOrTokens, address _user) returns()
func (_Depositmanager *DepositmanagerTransactor) DepositBulk(opts *bind.TransactOpts, _tokens []common.Address, _amountOrTokens []*big.Int, _user common.Address) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "depositBulk", _tokens, _amountOrTokens, _user)
}

// DepositBulk is a paid mutator transaction binding the contract method 0x7b1f7117.
//
// Solidity: function depositBulk(address[] _tokens, uint256[] _amountOrTokens, address _user) returns()
func (_Depositmanager *DepositmanagerSession) DepositBulk(_tokens []common.Address, _amountOrTokens []*big.Int, _user common.Address) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositBulk(&_Depositmanager.TransactOpts, _tokens, _amountOrTokens, _user)
}

// DepositBulk is a paid mutator transaction binding the contract method 0x7b1f7117.
//
// Solidity: function depositBulk(address[] _tokens, uint256[] _amountOrTokens, address _user) returns()
func (_Depositmanager *DepositmanagerTransactorSession) DepositBulk(_tokens []common.Address, _amountOrTokens []*big.Int, _user common.Address) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositBulk(&_Depositmanager.TransactOpts, _tokens, _amountOrTokens, _user)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address _token, uint256 _amount) returns()
func (_Depositmanager *DepositmanagerTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "depositERC20", _token, _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address _token, uint256 _amount) returns()
func (_Depositmanager *DepositmanagerSession) DepositERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC20(&_Depositmanager.TransactOpts, _token, _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address _token, uint256 _amount) returns()
func (_Depositmanager *DepositmanagerTransactorSession) DepositERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC20(&_Depositmanager.TransactOpts, _token, _amount)
}

// DepositERC20ForUser is a paid mutator transaction binding the contract method 0x8b9e4f93.
//
// Solidity: function depositERC20ForUser(address _token, address _user, uint256 _amount) returns()
func (_Depositmanager *DepositmanagerTransactor) DepositERC20ForUser(opts *bind.TransactOpts, _token common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "depositERC20ForUser", _token, _user, _amount)
}

// DepositERC20ForUser is a paid mutator transaction binding the contract method 0x8b9e4f93.
//
// Solidity: function depositERC20ForUser(address _token, address _user, uint256 _amount) returns()
func (_Depositmanager *DepositmanagerSession) DepositERC20ForUser(_token common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC20ForUser(&_Depositmanager.TransactOpts, _token, _user, _amount)
}

// DepositERC20ForUser is a paid mutator transaction binding the contract method 0x8b9e4f93.
//
// Solidity: function depositERC20ForUser(address _token, address _user, uint256 _amount) returns()
func (_Depositmanager *DepositmanagerTransactorSession) DepositERC20ForUser(_token common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC20ForUser(&_Depositmanager.TransactOpts, _token, _user, _amount)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId) returns()
func (_Depositmanager *DepositmanagerTransactor) DepositERC721(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "depositERC721", _token, _tokenId)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId) returns()
func (_Depositmanager *DepositmanagerSession) DepositERC721(_token common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC721(&_Depositmanager.TransactOpts, _token, _tokenId)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xd29a4bf6.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId) returns()
func (_Depositmanager *DepositmanagerTransactorSession) DepositERC721(_token common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC721(&_Depositmanager.TransactOpts, _token, _tokenId)
}

// DepositERC721ForUser is a paid mutator transaction binding the contract method 0x072b1535.
//
// Solidity: function depositERC721ForUser(address _token, address _user, uint256 _tokenId) returns()
func (_Depositmanager *DepositmanagerTransactor) DepositERC721ForUser(opts *bind.TransactOpts, _token common.Address, _user common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "depositERC721ForUser", _token, _user, _tokenId)
}

// DepositERC721ForUser is a paid mutator transaction binding the contract method 0x072b1535.
//
// Solidity: function depositERC721ForUser(address _token, address _user, uint256 _tokenId) returns()
func (_Depositmanager *DepositmanagerSession) DepositERC721ForUser(_token common.Address, _user common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC721ForUser(&_Depositmanager.TransactOpts, _token, _user, _tokenId)
}

// DepositERC721ForUser is a paid mutator transaction binding the contract method 0x072b1535.
//
// Solidity: function depositERC721ForUser(address _token, address _user, uint256 _tokenId) returns()
func (_Depositmanager *DepositmanagerTransactorSession) DepositERC721ForUser(_token common.Address, _user common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositERC721ForUser(&_Depositmanager.TransactOpts, _token, _user, _tokenId)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns()
func (_Depositmanager *DepositmanagerTransactor) DepositEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "depositEther")
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns()
func (_Depositmanager *DepositmanagerSession) DepositEther() (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositEther(&_Depositmanager.TransactOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns()
func (_Depositmanager *DepositmanagerTransactorSession) DepositEther() (*types.Transaction, error) {
	return _Depositmanager.Contract.DepositEther(&_Depositmanager.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Depositmanager *DepositmanagerTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Depositmanager *DepositmanagerSession) Lock() (*types.Transaction, error) {
	return _Depositmanager.Contract.Lock(&_Depositmanager.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Depositmanager *DepositmanagerTransactorSession) Lock() (*types.Transaction, error) {
	return _Depositmanager.Contract.Lock(&_Depositmanager.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address _user, uint256 _tokenId, bytes ) returns(bytes4)
func (_Depositmanager *DepositmanagerTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, _user common.Address, _tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "onERC721Received", arg0, _user, _tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address _user, uint256 _tokenId, bytes ) returns(bytes4)
func (_Depositmanager *DepositmanagerSession) OnERC721Received(arg0 common.Address, _user common.Address, _tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Depositmanager.Contract.OnERC721Received(&_Depositmanager.TransactOpts, arg0, _user, _tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address _user, uint256 _tokenId, bytes ) returns(bytes4)
func (_Depositmanager *DepositmanagerTransactorSession) OnERC721Received(arg0 common.Address, _user common.Address, _tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Depositmanager.Contract.OnERC721Received(&_Depositmanager.TransactOpts, arg0, _user, _tokenId, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Depositmanager *DepositmanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Depositmanager *DepositmanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Depositmanager.Contract.RenounceOwnership(&_Depositmanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Depositmanager *DepositmanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Depositmanager.Contract.RenounceOwnership(&_Depositmanager.TransactOpts)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _user, uint256 _amount, bytes ) returns()
func (_Depositmanager *DepositmanagerTransactor) TokenFallback(opts *bind.TransactOpts, _user common.Address, _amount *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "tokenFallback", _user, _amount, arg2)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _user, uint256 _amount, bytes ) returns()
func (_Depositmanager *DepositmanagerSession) TokenFallback(_user common.Address, _amount *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Depositmanager.Contract.TokenFallback(&_Depositmanager.TransactOpts, _user, _amount, arg2)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _user, uint256 _amount, bytes ) returns()
func (_Depositmanager *DepositmanagerTransactorSession) TokenFallback(_user common.Address, _amount *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Depositmanager.Contract.TokenFallback(&_Depositmanager.TransactOpts, _user, _amount, arg2)
}

// TransferAssets is a paid mutator transaction binding the contract method 0x49f4cc17.
//
// Solidity: function transferAssets(address _token, address _user, uint256 _amountOrNFTId) returns()
func (_Depositmanager *DepositmanagerTransactor) TransferAssets(opts *bind.TransactOpts, _token common.Address, _user common.Address, _amountOrNFTId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "transferAssets", _token, _user, _amountOrNFTId)
}

// TransferAssets is a paid mutator transaction binding the contract method 0x49f4cc17.
//
// Solidity: function transferAssets(address _token, address _user, uint256 _amountOrNFTId) returns()
func (_Depositmanager *DepositmanagerSession) TransferAssets(_token common.Address, _user common.Address, _amountOrNFTId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.TransferAssets(&_Depositmanager.TransactOpts, _token, _user, _amountOrNFTId)
}

// TransferAssets is a paid mutator transaction binding the contract method 0x49f4cc17.
//
// Solidity: function transferAssets(address _token, address _user, uint256 _amountOrNFTId) returns()
func (_Depositmanager *DepositmanagerTransactorSession) TransferAssets(_token common.Address, _user common.Address, _amountOrNFTId *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.TransferAssets(&_Depositmanager.TransactOpts, _token, _user, _amountOrNFTId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Depositmanager *DepositmanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Depositmanager *DepositmanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Depositmanager.Contract.TransferOwnership(&_Depositmanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Depositmanager *DepositmanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Depositmanager.Contract.TransferOwnership(&_Depositmanager.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Depositmanager *DepositmanagerTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Depositmanager *DepositmanagerSession) Unlock() (*types.Transaction, error) {
	return _Depositmanager.Contract.Unlock(&_Depositmanager.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Depositmanager *DepositmanagerTransactorSession) Unlock() (*types.Transaction, error) {
	return _Depositmanager.Contract.Unlock(&_Depositmanager.TransactOpts)
}

// UpdateChildChainAndStateSender is a paid mutator transaction binding the contract method 0x42be8379.
//
// Solidity: function updateChildChainAndStateSender() returns()
func (_Depositmanager *DepositmanagerTransactor) UpdateChildChainAndStateSender(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "updateChildChainAndStateSender")
}

// UpdateChildChainAndStateSender is a paid mutator transaction binding the contract method 0x42be8379.
//
// Solidity: function updateChildChainAndStateSender() returns()
func (_Depositmanager *DepositmanagerSession) UpdateChildChainAndStateSender() (*types.Transaction, error) {
	return _Depositmanager.Contract.UpdateChildChainAndStateSender(&_Depositmanager.TransactOpts)
}

// UpdateChildChainAndStateSender is a paid mutator transaction binding the contract method 0x42be8379.
//
// Solidity: function updateChildChainAndStateSender() returns()
func (_Depositmanager *DepositmanagerTransactorSession) UpdateChildChainAndStateSender() (*types.Transaction, error) {
	return _Depositmanager.Contract.UpdateChildChainAndStateSender(&_Depositmanager.TransactOpts)
}

// UpdateMaxErc20Deposit is a paid mutator transaction binding the contract method 0x4b56c071.
//
// Solidity: function updateMaxErc20Deposit(uint256 maxDepositAmount) returns()
func (_Depositmanager *DepositmanagerTransactor) UpdateMaxErc20Deposit(opts *bind.TransactOpts, maxDepositAmount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "updateMaxErc20Deposit", maxDepositAmount)
}

// UpdateMaxErc20Deposit is a paid mutator transaction binding the contract method 0x4b56c071.
//
// Solidity: function updateMaxErc20Deposit(uint256 maxDepositAmount) returns()
func (_Depositmanager *DepositmanagerSession) UpdateMaxErc20Deposit(maxDepositAmount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.UpdateMaxErc20Deposit(&_Depositmanager.TransactOpts, maxDepositAmount)
}

// UpdateMaxErc20Deposit is a paid mutator transaction binding the contract method 0x4b56c071.
//
// Solidity: function updateMaxErc20Deposit(uint256 maxDepositAmount) returns()
func (_Depositmanager *DepositmanagerTransactorSession) UpdateMaxErc20Deposit(maxDepositAmount *big.Int) (*types.Transaction, error) {
	return _Depositmanager.Contract.UpdateMaxErc20Deposit(&_Depositmanager.TransactOpts, maxDepositAmount)
}

// UpdateRootChain is a paid mutator transaction binding the contract method 0xf2203711.
//
// Solidity: function updateRootChain(address _rootChain) returns()
func (_Depositmanager *DepositmanagerTransactor) UpdateRootChain(opts *bind.TransactOpts, _rootChain common.Address) (*types.Transaction, error) {
	return _Depositmanager.contract.Transact(opts, "updateRootChain", _rootChain)
}

// UpdateRootChain is a paid mutator transaction binding the contract method 0xf2203711.
//
// Solidity: function updateRootChain(address _rootChain) returns()
func (_Depositmanager *DepositmanagerSession) UpdateRootChain(_rootChain common.Address) (*types.Transaction, error) {
	return _Depositmanager.Contract.UpdateRootChain(&_Depositmanager.TransactOpts, _rootChain)
}

// UpdateRootChain is a paid mutator transaction binding the contract method 0xf2203711.
//
// Solidity: function updateRootChain(address _rootChain) returns()
func (_Depositmanager *DepositmanagerTransactorSession) UpdateRootChain(_rootChain common.Address) (*types.Transaction, error) {
	return _Depositmanager.Contract.UpdateRootChain(&_Depositmanager.TransactOpts, _rootChain)
}

// DepositmanagerMaxErc20DepositUpdateIterator is returned from FilterMaxErc20DepositUpdate and is used to iterate over the raw logs and unpacked data for MaxErc20DepositUpdate events raised by the Depositmanager contract.
type DepositmanagerMaxErc20DepositUpdateIterator struct {
	Event *DepositmanagerMaxErc20DepositUpdate // Event containing the contract specifics and raw log

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
func (it *DepositmanagerMaxErc20DepositUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositmanagerMaxErc20DepositUpdate)
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
		it.Event = new(DepositmanagerMaxErc20DepositUpdate)
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
func (it *DepositmanagerMaxErc20DepositUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositmanagerMaxErc20DepositUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositmanagerMaxErc20DepositUpdate represents a MaxErc20DepositUpdate event raised by the Depositmanager contract.
type DepositmanagerMaxErc20DepositUpdate struct {
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMaxErc20DepositUpdate is a free log retrieval operation binding the contract event 0x010c0265813c273963aa5e8683cf5c45a3b744ba6369c22af0958ec5fcf16b20.
//
// Solidity: event MaxErc20DepositUpdate(uint256 indexed oldLimit, uint256 indexed newLimit)
func (_Depositmanager *DepositmanagerFilterer) FilterMaxErc20DepositUpdate(opts *bind.FilterOpts, oldLimit []*big.Int, newLimit []*big.Int) (*DepositmanagerMaxErc20DepositUpdateIterator, error) {

	var oldLimitRule []interface{}
	for _, oldLimitItem := range oldLimit {
		oldLimitRule = append(oldLimitRule, oldLimitItem)
	}
	var newLimitRule []interface{}
	for _, newLimitItem := range newLimit {
		newLimitRule = append(newLimitRule, newLimitItem)
	}

	logs, sub, err := _Depositmanager.contract.FilterLogs(opts, "MaxErc20DepositUpdate", oldLimitRule, newLimitRule)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerMaxErc20DepositUpdateIterator{contract: _Depositmanager.contract, event: "MaxErc20DepositUpdate", logs: logs, sub: sub}, nil
}

// WatchMaxErc20DepositUpdate is a free log subscription operation binding the contract event 0x010c0265813c273963aa5e8683cf5c45a3b744ba6369c22af0958ec5fcf16b20.
//
// Solidity: event MaxErc20DepositUpdate(uint256 indexed oldLimit, uint256 indexed newLimit)
func (_Depositmanager *DepositmanagerFilterer) WatchMaxErc20DepositUpdate(opts *bind.WatchOpts, sink chan<- *DepositmanagerMaxErc20DepositUpdate, oldLimit []*big.Int, newLimit []*big.Int) (event.Subscription, error) {

	var oldLimitRule []interface{}
	for _, oldLimitItem := range oldLimit {
		oldLimitRule = append(oldLimitRule, oldLimitItem)
	}
	var newLimitRule []interface{}
	for _, newLimitItem := range newLimit {
		newLimitRule = append(newLimitRule, newLimitItem)
	}

	logs, sub, err := _Depositmanager.contract.WatchLogs(opts, "MaxErc20DepositUpdate", oldLimitRule, newLimitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositmanagerMaxErc20DepositUpdate)
				if err := _Depositmanager.contract.UnpackLog(event, "MaxErc20DepositUpdate", log); err != nil {
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

// ParseMaxErc20DepositUpdate is a log parse operation binding the contract event 0x010c0265813c273963aa5e8683cf5c45a3b744ba6369c22af0958ec5fcf16b20.
//
// Solidity: event MaxErc20DepositUpdate(uint256 indexed oldLimit, uint256 indexed newLimit)
func (_Depositmanager *DepositmanagerFilterer) ParseMaxErc20DepositUpdate(log types.Log) (*DepositmanagerMaxErc20DepositUpdate, error) {
	event := new(DepositmanagerMaxErc20DepositUpdate)
	if err := _Depositmanager.contract.UnpackLog(event, "MaxErc20DepositUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositmanagerNewDepositBlockIterator is returned from FilterNewDepositBlock and is used to iterate over the raw logs and unpacked data for NewDepositBlock events raised by the Depositmanager contract.
type DepositmanagerNewDepositBlockIterator struct {
	Event *DepositmanagerNewDepositBlock // Event containing the contract specifics and raw log

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
func (it *DepositmanagerNewDepositBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositmanagerNewDepositBlock)
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
		it.Event = new(DepositmanagerNewDepositBlock)
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
func (it *DepositmanagerNewDepositBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositmanagerNewDepositBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositmanagerNewDepositBlock represents a NewDepositBlock event raised by the Depositmanager contract.
type DepositmanagerNewDepositBlock struct {
	Owner          common.Address
	Token          common.Address
	AmountOrNFTId  *big.Int
	DepositBlockId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewDepositBlock is a free log retrieval operation binding the contract event 0x1dadc8d0683c6f9824e885935c1bec6f76816730dcec148dda8cf25a7b9f797b.
//
// Solidity: event NewDepositBlock(address indexed owner, address indexed token, uint256 amountOrNFTId, uint256 depositBlockId)
func (_Depositmanager *DepositmanagerFilterer) FilterNewDepositBlock(opts *bind.FilterOpts, owner []common.Address, token []common.Address) (*DepositmanagerNewDepositBlockIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Depositmanager.contract.FilterLogs(opts, "NewDepositBlock", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerNewDepositBlockIterator{contract: _Depositmanager.contract, event: "NewDepositBlock", logs: logs, sub: sub}, nil
}

// WatchNewDepositBlock is a free log subscription operation binding the contract event 0x1dadc8d0683c6f9824e885935c1bec6f76816730dcec148dda8cf25a7b9f797b.
//
// Solidity: event NewDepositBlock(address indexed owner, address indexed token, uint256 amountOrNFTId, uint256 depositBlockId)
func (_Depositmanager *DepositmanagerFilterer) WatchNewDepositBlock(opts *bind.WatchOpts, sink chan<- *DepositmanagerNewDepositBlock, owner []common.Address, token []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Depositmanager.contract.WatchLogs(opts, "NewDepositBlock", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositmanagerNewDepositBlock)
				if err := _Depositmanager.contract.UnpackLog(event, "NewDepositBlock", log); err != nil {
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

// ParseNewDepositBlock is a log parse operation binding the contract event 0x1dadc8d0683c6f9824e885935c1bec6f76816730dcec148dda8cf25a7b9f797b.
//
// Solidity: event NewDepositBlock(address indexed owner, address indexed token, uint256 amountOrNFTId, uint256 depositBlockId)
func (_Depositmanager *DepositmanagerFilterer) ParseNewDepositBlock(log types.Log) (*DepositmanagerNewDepositBlock, error) {
	event := new(DepositmanagerNewDepositBlock)
	if err := _Depositmanager.contract.UnpackLog(event, "NewDepositBlock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DepositmanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Depositmanager contract.
type DepositmanagerOwnershipTransferredIterator struct {
	Event *DepositmanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DepositmanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositmanagerOwnershipTransferred)
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
		it.Event = new(DepositmanagerOwnershipTransferred)
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
func (it *DepositmanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositmanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositmanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Depositmanager contract.
type DepositmanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Depositmanager *DepositmanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DepositmanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Depositmanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DepositmanagerOwnershipTransferredIterator{contract: _Depositmanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Depositmanager *DepositmanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DepositmanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Depositmanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositmanagerOwnershipTransferred)
				if err := _Depositmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Depositmanager *DepositmanagerFilterer) ParseOwnershipTransferred(log types.Log) (*DepositmanagerOwnershipTransferred, error) {
	event := new(DepositmanagerOwnershipTransferred)
	if err := _Depositmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
