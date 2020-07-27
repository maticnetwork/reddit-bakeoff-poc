// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subredditpoints

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

// SubredditpointsABI is the input ABI used to generate the binding from.
const SubredditpointsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"defaultOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addDefaultOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributionContract_\",\"type\":\"address\"}],\"name\":\"updateDistributionContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"childChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"distributionContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"distributionContract_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"subreddit_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"defaultOperators_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_rootToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_childChain\",\"type\":\"address\"}],\"name\":\"initializecontract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"authorizeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"subreddit\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"isOperatorFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeDefaultOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"revokeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"operatorBurn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"AuthorizedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenHolder\",\"type\":\"address\"}],\"name\":\"RevokedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DefaultOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"DefaultOperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output2\",\"type\":\"uint256\"}],\"name\":\"LogTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"input1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"output1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Subredditpoints is an auto generated Go binding around an Ethereum contract.
type Subredditpoints struct {
	SubredditpointsCaller     // Read-only binding to the contract
	SubredditpointsTransactor // Write-only binding to the contract
	SubredditpointsFilterer   // Log filterer for contract events
}

// SubredditpointsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubredditpointsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubredditpointsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubredditpointsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubredditpointsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubredditpointsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubredditpointsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubredditpointsSession struct {
	Contract     *Subredditpoints  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubredditpointsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubredditpointsCallerSession struct {
	Contract *SubredditpointsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SubredditpointsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubredditpointsTransactorSession struct {
	Contract     *SubredditpointsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SubredditpointsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubredditpointsRaw struct {
	Contract *Subredditpoints // Generic contract binding to access the raw methods on
}

// SubredditpointsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubredditpointsCallerRaw struct {
	Contract *SubredditpointsCaller // Generic read-only contract binding to access the raw methods on
}

// SubredditpointsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubredditpointsTransactorRaw struct {
	Contract *SubredditpointsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubredditpoints creates a new instance of Subredditpoints, bound to a specific deployed contract.
func NewSubredditpoints(address common.Address, backend bind.ContractBackend) (*Subredditpoints, error) {
	contract, err := bindSubredditpoints(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subredditpoints{SubredditpointsCaller: SubredditpointsCaller{contract: contract}, SubredditpointsTransactor: SubredditpointsTransactor{contract: contract}, SubredditpointsFilterer: SubredditpointsFilterer{contract: contract}}, nil
}

// NewSubredditpointsCaller creates a new read-only instance of Subredditpoints, bound to a specific deployed contract.
func NewSubredditpointsCaller(address common.Address, caller bind.ContractCaller) (*SubredditpointsCaller, error) {
	contract, err := bindSubredditpoints(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsCaller{contract: contract}, nil
}

// NewSubredditpointsTransactor creates a new write-only instance of Subredditpoints, bound to a specific deployed contract.
func NewSubredditpointsTransactor(address common.Address, transactor bind.ContractTransactor) (*SubredditpointsTransactor, error) {
	contract, err := bindSubredditpoints(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsTransactor{contract: contract}, nil
}

// NewSubredditpointsFilterer creates a new log filterer instance of Subredditpoints, bound to a specific deployed contract.
func NewSubredditpointsFilterer(address common.Address, filterer bind.ContractFilterer) (*SubredditpointsFilterer, error) {
	contract, err := bindSubredditpoints(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsFilterer{contract: contract}, nil
}

// bindSubredditpoints binds a generic wrapper to an already deployed contract.
func bindSubredditpoints(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubredditpointsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subredditpoints *SubredditpointsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subredditpoints.Contract.SubredditpointsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subredditpoints *SubredditpointsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subredditpoints.Contract.SubredditpointsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subredditpoints *SubredditpointsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subredditpoints.Contract.SubredditpointsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subredditpoints *SubredditpointsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Subredditpoints.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subredditpoints *SubredditpointsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subredditpoints.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subredditpoints *SubredditpointsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subredditpoints.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Subredditpoints *SubredditpointsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Subredditpoints *SubredditpointsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Subredditpoints.Contract.Allowance(&_Subredditpoints.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Subredditpoints *SubredditpointsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Subredditpoints.Contract.Allowance(&_Subredditpoints.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Subredditpoints *SubredditpointsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Subredditpoints *SubredditpointsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Subredditpoints.Contract.BalanceOf(&_Subredditpoints.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Subredditpoints *SubredditpointsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Subredditpoints.Contract.BalanceOf(&_Subredditpoints.CallOpts, account)
}

// ChildChain is a free data retrieval call binding the contract method 0x42fc47fb.
//
// Solidity: function childChain() constant returns(address)
func (_Subredditpoints *SubredditpointsCaller) ChildChain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "childChain")
	return *ret0, err
}

// ChildChain is a free data retrieval call binding the contract method 0x42fc47fb.
//
// Solidity: function childChain() constant returns(address)
func (_Subredditpoints *SubredditpointsSession) ChildChain() (common.Address, error) {
	return _Subredditpoints.Contract.ChildChain(&_Subredditpoints.CallOpts)
}

// ChildChain is a free data retrieval call binding the contract method 0x42fc47fb.
//
// Solidity: function childChain() constant returns(address)
func (_Subredditpoints *SubredditpointsCallerSession) ChildChain() (common.Address, error) {
	return _Subredditpoints.Contract.ChildChain(&_Subredditpoints.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() constant returns(address[])
func (_Subredditpoints *SubredditpointsCaller) DefaultOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "defaultOperators")
	return *ret0, err
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() constant returns(address[])
func (_Subredditpoints *SubredditpointsSession) DefaultOperators() ([]common.Address, error) {
	return _Subredditpoints.Contract.DefaultOperators(&_Subredditpoints.CallOpts)
}

// DefaultOperators is a free data retrieval call binding the contract method 0x06e48538.
//
// Solidity: function defaultOperators() constant returns(address[])
func (_Subredditpoints *SubredditpointsCallerSession) DefaultOperators() ([]common.Address, error) {
	return _Subredditpoints.Contract.DefaultOperators(&_Subredditpoints.CallOpts)
}

// DistributionContract is a free data retrieval call binding the contract method 0x5a4528c2.
//
// Solidity: function distributionContract() constant returns(address)
func (_Subredditpoints *SubredditpointsCaller) DistributionContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "distributionContract")
	return *ret0, err
}

// DistributionContract is a free data retrieval call binding the contract method 0x5a4528c2.
//
// Solidity: function distributionContract() constant returns(address)
func (_Subredditpoints *SubredditpointsSession) DistributionContract() (common.Address, error) {
	return _Subredditpoints.Contract.DistributionContract(&_Subredditpoints.CallOpts)
}

// DistributionContract is a free data retrieval call binding the contract method 0x5a4528c2.
//
// Solidity: function distributionContract() constant returns(address)
func (_Subredditpoints *SubredditpointsCallerSession) DistributionContract() (common.Address, error) {
	return _Subredditpoints.Contract.DistributionContract(&_Subredditpoints.CallOpts)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) constant returns(bool)
func (_Subredditpoints *SubredditpointsCaller) IsOperatorFor(opts *bind.CallOpts, operator common.Address, tokenHolder common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "isOperatorFor", operator, tokenHolder)
	return *ret0, err
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) constant returns(bool)
func (_Subredditpoints *SubredditpointsSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _Subredditpoints.Contract.IsOperatorFor(&_Subredditpoints.CallOpts, operator, tokenHolder)
}

// IsOperatorFor is a free data retrieval call binding the contract method 0xd95b6371.
//
// Solidity: function isOperatorFor(address operator, address tokenHolder) constant returns(bool)
func (_Subredditpoints *SubredditpointsCallerSession) IsOperatorFor(operator common.Address, tokenHolder common.Address) (bool, error) {
	return _Subredditpoints.Contract.IsOperatorFor(&_Subredditpoints.CallOpts, operator, tokenHolder)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Subredditpoints *SubredditpointsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Subredditpoints *SubredditpointsSession) IsOwner() (bool, error) {
	return _Subredditpoints.Contract.IsOwner(&_Subredditpoints.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Subredditpoints *SubredditpointsCallerSession) IsOwner() (bool, error) {
	return _Subredditpoints.Contract.IsOwner(&_Subredditpoints.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Subredditpoints *SubredditpointsCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Subredditpoints *SubredditpointsSession) Name() (string, error) {
	return _Subredditpoints.Contract.Name(&_Subredditpoints.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Subredditpoints *SubredditpointsCallerSession) Name() (string, error) {
	return _Subredditpoints.Contract.Name(&_Subredditpoints.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subredditpoints *SubredditpointsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subredditpoints *SubredditpointsSession) Owner() (common.Address, error) {
	return _Subredditpoints.Contract.Owner(&_Subredditpoints.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Subredditpoints *SubredditpointsCallerSession) Owner() (common.Address, error) {
	return _Subredditpoints.Contract.Owner(&_Subredditpoints.CallOpts)
}

// RootToken is a free data retrieval call binding the contract method 0x1f2d0065.
//
// Solidity: function rootToken() constant returns(address)
func (_Subredditpoints *SubredditpointsCaller) RootToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "rootToken")
	return *ret0, err
}

// RootToken is a free data retrieval call binding the contract method 0x1f2d0065.
//
// Solidity: function rootToken() constant returns(address)
func (_Subredditpoints *SubredditpointsSession) RootToken() (common.Address, error) {
	return _Subredditpoints.Contract.RootToken(&_Subredditpoints.CallOpts)
}

// RootToken is a free data retrieval call binding the contract method 0x1f2d0065.
//
// Solidity: function rootToken() constant returns(address)
func (_Subredditpoints *SubredditpointsCallerSession) RootToken() (common.Address, error) {
	return _Subredditpoints.Contract.RootToken(&_Subredditpoints.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() constant returns(string)
func (_Subredditpoints *SubredditpointsCaller) Subreddit(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "subreddit")
	return *ret0, err
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() constant returns(string)
func (_Subredditpoints *SubredditpointsSession) Subreddit() (string, error) {
	return _Subredditpoints.Contract.Subreddit(&_Subredditpoints.CallOpts)
}

// Subreddit is a free data retrieval call binding the contract method 0xbdc330cb.
//
// Solidity: function subreddit() constant returns(string)
func (_Subredditpoints *SubredditpointsCallerSession) Subreddit() (string, error) {
	return _Subredditpoints.Contract.Subreddit(&_Subredditpoints.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Subredditpoints *SubredditpointsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Subredditpoints *SubredditpointsSession) Symbol() (string, error) {
	return _Subredditpoints.Contract.Symbol(&_Subredditpoints.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Subredditpoints *SubredditpointsCallerSession) Symbol() (string, error) {
	return _Subredditpoints.Contract.Symbol(&_Subredditpoints.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Subredditpoints *SubredditpointsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Subredditpoints.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Subredditpoints *SubredditpointsSession) TotalSupply() (*big.Int, error) {
	return _Subredditpoints.Contract.TotalSupply(&_Subredditpoints.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Subredditpoints *SubredditpointsCallerSession) TotalSupply() (*big.Int, error) {
	return _Subredditpoints.Contract.TotalSupply(&_Subredditpoints.CallOpts)
}

// AddDefaultOperator is a paid mutator transaction binding the contract method 0x1b7d83e5.
//
// Solidity: function addDefaultOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactor) AddDefaultOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "addDefaultOperator", operator)
}

// AddDefaultOperator is a paid mutator transaction binding the contract method 0x1b7d83e5.
//
// Solidity: function addDefaultOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsSession) AddDefaultOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.AddDefaultOperator(&_Subredditpoints.TransactOpts, operator)
}

// AddDefaultOperator is a paid mutator transaction binding the contract method 0x1b7d83e5.
//
// Solidity: function addDefaultOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) AddDefaultOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.AddDefaultOperator(&_Subredditpoints.TransactOpts, operator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Approve(&_Subredditpoints.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Approve(&_Subredditpoints.TransactOpts, spender, amount)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactor) AuthorizeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "authorizeOperator", operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.AuthorizeOperator(&_Subredditpoints.TransactOpts, operator)
}

// AuthorizeOperator is a paid mutator transaction binding the contract method 0x959b8c3f.
//
// Solidity: function authorizeOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) AuthorizeOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.AuthorizeOperator(&_Subredditpoints.TransactOpts, operator)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes userData) returns()
func (_Subredditpoints *SubredditpointsTransactor) Burn(opts *bind.TransactOpts, amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "burn", amount, userData)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes userData) returns()
func (_Subredditpoints *SubredditpointsSession) Burn(amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Burn(&_Subredditpoints.TransactOpts, amount, userData)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(uint256 amount, bytes userData) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) Burn(amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Burn(&_Subredditpoints.TransactOpts, amount, userData)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Subredditpoints *SubredditpointsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Subredditpoints *SubredditpointsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.DecreaseAllowance(&_Subredditpoints.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Subredditpoints *SubredditpointsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.DecreaseAllowance(&_Subredditpoints.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Subredditpoints *SubredditpointsTransactor) Deposit(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "deposit", user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Subredditpoints *SubredditpointsSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Deposit(&_Subredditpoints.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address user, uint256 amount) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) Deposit(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Deposit(&_Subredditpoints.TransactOpts, user, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Subredditpoints *SubredditpointsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Subredditpoints *SubredditpointsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.IncreaseAllowance(&_Subredditpoints.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Subredditpoints *SubredditpointsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.IncreaseAllowance(&_Subredditpoints.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Subredditpoints *SubredditpointsTransactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Subredditpoints *SubredditpointsSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Initialize(&_Subredditpoints.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Initialize(&_Subredditpoints.TransactOpts, sender)
}

// Initializecontract is a paid mutator transaction binding the contract method 0x938f893a.
//
// Solidity: function initializecontract(address owner_, address distributionContract_, string subreddit_, string name_, string symbol_, address[] defaultOperators_, address _rootToken, address _childChain) returns()
func (_Subredditpoints *SubredditpointsTransactor) Initializecontract(opts *bind.TransactOpts, owner_ common.Address, distributionContract_ common.Address, subreddit_ string, name_ string, symbol_ string, defaultOperators_ []common.Address, _rootToken common.Address, _childChain common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "initializecontract", owner_, distributionContract_, subreddit_, name_, symbol_, defaultOperators_, _rootToken, _childChain)
}

// Initializecontract is a paid mutator transaction binding the contract method 0x938f893a.
//
// Solidity: function initializecontract(address owner_, address distributionContract_, string subreddit_, string name_, string symbol_, address[] defaultOperators_, address _rootToken, address _childChain) returns()
func (_Subredditpoints *SubredditpointsSession) Initializecontract(owner_ common.Address, distributionContract_ common.Address, subreddit_ string, name_ string, symbol_ string, defaultOperators_ []common.Address, _rootToken common.Address, _childChain common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Initializecontract(&_Subredditpoints.TransactOpts, owner_, distributionContract_, subreddit_, name_, symbol_, defaultOperators_, _rootToken, _childChain)
}

// Initializecontract is a paid mutator transaction binding the contract method 0x938f893a.
//
// Solidity: function initializecontract(address owner_, address distributionContract_, string subreddit_, string name_, string symbol_, address[] defaultOperators_, address _rootToken, address _childChain) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) Initializecontract(owner_ common.Address, distributionContract_ common.Address, subreddit_ string, name_ string, symbol_ string, defaultOperators_ []common.Address, _rootToken common.Address, _childChain common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Initializecontract(&_Subredditpoints.TransactOpts, owner_, distributionContract_, subreddit_, name_, symbol_, defaultOperators_, _rootToken, _childChain)
}

// Mint is a paid mutator transaction binding the contract method 0xab89013b.
//
// Solidity: function mint(address operator, address account, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsTransactor) Mint(opts *bind.TransactOpts, operator common.Address, account common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "mint", operator, account, amount, userData, operatorData)
}

// Mint is a paid mutator transaction binding the contract method 0xab89013b.
//
// Solidity: function mint(address operator, address account, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsSession) Mint(operator common.Address, account common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Mint(&_Subredditpoints.TransactOpts, operator, account, amount, userData, operatorData)
}

// Mint is a paid mutator transaction binding the contract method 0xab89013b.
//
// Solidity: function mint(address operator, address account, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) Mint(operator common.Address, account common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Mint(&_Subredditpoints.TransactOpts, operator, account, amount, userData, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsTransactor) OperatorBurn(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "operatorBurn", account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.OperatorBurn(&_Subredditpoints.TransactOpts, account, amount, data, operatorData)
}

// OperatorBurn is a paid mutator transaction binding the contract method 0xfc673c4f.
//
// Solidity: function operatorBurn(address account, uint256 amount, bytes data, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) OperatorBurn(account common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.OperatorBurn(&_Subredditpoints.TransactOpts, account, amount, data, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsTransactor) OperatorSend(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "operatorSend", sender, recipient, amount, userData, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.OperatorSend(&_Subredditpoints.TransactOpts, sender, recipient, amount, userData, operatorData)
}

// OperatorSend is a paid mutator transaction binding the contract method 0x62ad1b83.
//
// Solidity: function operatorSend(address sender, address recipient, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) OperatorSend(sender common.Address, recipient common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Subredditpoints.Contract.OperatorSend(&_Subredditpoints.TransactOpts, sender, recipient, amount, userData, operatorData)
}

// RemoveDefaultOperator is a paid mutator transaction binding the contract method 0xe699e8c3.
//
// Solidity: function removeDefaultOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactor) RemoveDefaultOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "removeDefaultOperator", operator)
}

// RemoveDefaultOperator is a paid mutator transaction binding the contract method 0xe699e8c3.
//
// Solidity: function removeDefaultOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsSession) RemoveDefaultOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.RemoveDefaultOperator(&_Subredditpoints.TransactOpts, operator)
}

// RemoveDefaultOperator is a paid mutator transaction binding the contract method 0xe699e8c3.
//
// Solidity: function removeDefaultOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) RemoveDefaultOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.RemoveDefaultOperator(&_Subredditpoints.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subredditpoints *SubredditpointsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subredditpoints *SubredditpointsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subredditpoints.Contract.RenounceOwnership(&_Subredditpoints.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subredditpoints *SubredditpointsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subredditpoints.Contract.RenounceOwnership(&_Subredditpoints.TransactOpts)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactor) RevokeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "revokeOperator", operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.RevokeOperator(&_Subredditpoints.TransactOpts, operator)
}

// RevokeOperator is a paid mutator transaction binding the contract method 0xfad8b32a.
//
// Solidity: function revokeOperator(address operator) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) RevokeOperator(operator common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.RevokeOperator(&_Subredditpoints.TransactOpts, operator)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Transfer(&_Subredditpoints.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.Transfer(&_Subredditpoints.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.TransferFrom(&_Subredditpoints.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Subredditpoints *SubredditpointsTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Subredditpoints.Contract.TransferFrom(&_Subredditpoints.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subredditpoints *SubredditpointsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subredditpoints *SubredditpointsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.TransferOwnership(&_Subredditpoints.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.TransferOwnership(&_Subredditpoints.TransactOpts, newOwner)
}

// UpdateDistributionContract is a paid mutator transaction binding the contract method 0x403f1ffa.
//
// Solidity: function updateDistributionContract(address distributionContract_) returns()
func (_Subredditpoints *SubredditpointsTransactor) UpdateDistributionContract(opts *bind.TransactOpts, distributionContract_ common.Address) (*types.Transaction, error) {
	return _Subredditpoints.contract.Transact(opts, "updateDistributionContract", distributionContract_)
}

// UpdateDistributionContract is a paid mutator transaction binding the contract method 0x403f1ffa.
//
// Solidity: function updateDistributionContract(address distributionContract_) returns()
func (_Subredditpoints *SubredditpointsSession) UpdateDistributionContract(distributionContract_ common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.UpdateDistributionContract(&_Subredditpoints.TransactOpts, distributionContract_)
}

// UpdateDistributionContract is a paid mutator transaction binding the contract method 0x403f1ffa.
//
// Solidity: function updateDistributionContract(address distributionContract_) returns()
func (_Subredditpoints *SubredditpointsTransactorSession) UpdateDistributionContract(distributionContract_ common.Address) (*types.Transaction, error) {
	return _Subredditpoints.Contract.UpdateDistributionContract(&_Subredditpoints.TransactOpts, distributionContract_)
}

// SubredditpointsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Subredditpoints contract.
type SubredditpointsApprovalIterator struct {
	Event *SubredditpointsApproval // Event containing the contract specifics and raw log

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
func (it *SubredditpointsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsApproval)
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
		it.Event = new(SubredditpointsApproval)
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
func (it *SubredditpointsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsApproval represents a Approval event raised by the Subredditpoints contract.
type SubredditpointsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Subredditpoints *SubredditpointsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SubredditpointsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsApprovalIterator{contract: _Subredditpoints.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Subredditpoints *SubredditpointsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SubredditpointsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsApproval)
				if err := _Subredditpoints.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Subredditpoints *SubredditpointsFilterer) ParseApproval(log types.Log) (*SubredditpointsApproval, error) {
	event := new(SubredditpointsApproval)
	if err := _Subredditpoints.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsAuthorizedOperatorIterator is returned from FilterAuthorizedOperator and is used to iterate over the raw logs and unpacked data for AuthorizedOperator events raised by the Subredditpoints contract.
type SubredditpointsAuthorizedOperatorIterator struct {
	Event *SubredditpointsAuthorizedOperator // Event containing the contract specifics and raw log

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
func (it *SubredditpointsAuthorizedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsAuthorizedOperator)
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
		it.Event = new(SubredditpointsAuthorizedOperator)
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
func (it *SubredditpointsAuthorizedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsAuthorizedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsAuthorizedOperator represents a AuthorizedOperator event raised by the Subredditpoints contract.
type SubredditpointsAuthorizedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedOperator is a free log retrieval operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Subredditpoints *SubredditpointsFilterer) FilterAuthorizedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*SubredditpointsAuthorizedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsAuthorizedOperatorIterator{contract: _Subredditpoints.contract, event: "AuthorizedOperator", logs: logs, sub: sub}, nil
}

// WatchAuthorizedOperator is a free log subscription operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Subredditpoints *SubredditpointsFilterer) WatchAuthorizedOperator(opts *bind.WatchOpts, sink chan<- *SubredditpointsAuthorizedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "AuthorizedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsAuthorizedOperator)
				if err := _Subredditpoints.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
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

// ParseAuthorizedOperator is a log parse operation binding the contract event 0xf4caeb2d6ca8932a215a353d0703c326ec2d81fc68170f320eb2ab49e9df61f9.
//
// Solidity: event AuthorizedOperator(address indexed operator, address indexed tokenHolder)
func (_Subredditpoints *SubredditpointsFilterer) ParseAuthorizedOperator(log types.Log) (*SubredditpointsAuthorizedOperator, error) {
	event := new(SubredditpointsAuthorizedOperator)
	if err := _Subredditpoints.contract.UnpackLog(event, "AuthorizedOperator", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the Subredditpoints contract.
type SubredditpointsBurnedIterator struct {
	Event *SubredditpointsBurned // Event containing the contract specifics and raw log

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
func (it *SubredditpointsBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsBurned)
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
		it.Event = new(SubredditpointsBurned)
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
func (it *SubredditpointsBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsBurned represents a Burned event raised by the Subredditpoints contract.
type SubredditpointsBurned struct {
	Operator     common.Address
	From         common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) FilterBurned(opts *bind.FilterOpts, operator []common.Address, from []common.Address) (*SubredditpointsBurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsBurnedIterator{contract: _Subredditpoints.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *SubredditpointsBurned, operator []common.Address, from []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "Burned", operatorRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsBurned)
				if err := _Subredditpoints.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0xa78a9be3a7b862d26933ad85fb11d80ef66b8f972d7cbba06621d583943a4098.
//
// Solidity: event Burned(address indexed operator, address indexed from, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) ParseBurned(log types.Log) (*SubredditpointsBurned, error) {
	event := new(SubredditpointsBurned)
	if err := _Subredditpoints.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsDefaultOperatorAddedIterator is returned from FilterDefaultOperatorAdded and is used to iterate over the raw logs and unpacked data for DefaultOperatorAdded events raised by the Subredditpoints contract.
type SubredditpointsDefaultOperatorAddedIterator struct {
	Event *SubredditpointsDefaultOperatorAdded // Event containing the contract specifics and raw log

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
func (it *SubredditpointsDefaultOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsDefaultOperatorAdded)
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
		it.Event = new(SubredditpointsDefaultOperatorAdded)
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
func (it *SubredditpointsDefaultOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsDefaultOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsDefaultOperatorAdded represents a DefaultOperatorAdded event raised by the Subredditpoints contract.
type SubredditpointsDefaultOperatorAdded struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDefaultOperatorAdded is a free log retrieval operation binding the contract event 0xb1c266c79163ba5da9a4deba2a5d1313547c94f50a5db49c1d04d9cc24a66c9d.
//
// Solidity: event DefaultOperatorAdded(address indexed operator)
func (_Subredditpoints *SubredditpointsFilterer) FilterDefaultOperatorAdded(opts *bind.FilterOpts, operator []common.Address) (*SubredditpointsDefaultOperatorAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "DefaultOperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsDefaultOperatorAddedIterator{contract: _Subredditpoints.contract, event: "DefaultOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchDefaultOperatorAdded is a free log subscription operation binding the contract event 0xb1c266c79163ba5da9a4deba2a5d1313547c94f50a5db49c1d04d9cc24a66c9d.
//
// Solidity: event DefaultOperatorAdded(address indexed operator)
func (_Subredditpoints *SubredditpointsFilterer) WatchDefaultOperatorAdded(opts *bind.WatchOpts, sink chan<- *SubredditpointsDefaultOperatorAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "DefaultOperatorAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsDefaultOperatorAdded)
				if err := _Subredditpoints.contract.UnpackLog(event, "DefaultOperatorAdded", log); err != nil {
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

// ParseDefaultOperatorAdded is a log parse operation binding the contract event 0xb1c266c79163ba5da9a4deba2a5d1313547c94f50a5db49c1d04d9cc24a66c9d.
//
// Solidity: event DefaultOperatorAdded(address indexed operator)
func (_Subredditpoints *SubredditpointsFilterer) ParseDefaultOperatorAdded(log types.Log) (*SubredditpointsDefaultOperatorAdded, error) {
	event := new(SubredditpointsDefaultOperatorAdded)
	if err := _Subredditpoints.contract.UnpackLog(event, "DefaultOperatorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsDefaultOperatorRemovedIterator is returned from FilterDefaultOperatorRemoved and is used to iterate over the raw logs and unpacked data for DefaultOperatorRemoved events raised by the Subredditpoints contract.
type SubredditpointsDefaultOperatorRemovedIterator struct {
	Event *SubredditpointsDefaultOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *SubredditpointsDefaultOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsDefaultOperatorRemoved)
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
		it.Event = new(SubredditpointsDefaultOperatorRemoved)
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
func (it *SubredditpointsDefaultOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsDefaultOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsDefaultOperatorRemoved represents a DefaultOperatorRemoved event raised by the Subredditpoints contract.
type SubredditpointsDefaultOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDefaultOperatorRemoved is a free log retrieval operation binding the contract event 0xec42077e4a379228eef9f58422079e786670be900ea3257e9d2e2227d0b56abd.
//
// Solidity: event DefaultOperatorRemoved(address indexed operator)
func (_Subredditpoints *SubredditpointsFilterer) FilterDefaultOperatorRemoved(opts *bind.FilterOpts, operator []common.Address) (*SubredditpointsDefaultOperatorRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "DefaultOperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsDefaultOperatorRemovedIterator{contract: _Subredditpoints.contract, event: "DefaultOperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchDefaultOperatorRemoved is a free log subscription operation binding the contract event 0xec42077e4a379228eef9f58422079e786670be900ea3257e9d2e2227d0b56abd.
//
// Solidity: event DefaultOperatorRemoved(address indexed operator)
func (_Subredditpoints *SubredditpointsFilterer) WatchDefaultOperatorRemoved(opts *bind.WatchOpts, sink chan<- *SubredditpointsDefaultOperatorRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "DefaultOperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsDefaultOperatorRemoved)
				if err := _Subredditpoints.contract.UnpackLog(event, "DefaultOperatorRemoved", log); err != nil {
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

// ParseDefaultOperatorRemoved is a log parse operation binding the contract event 0xec42077e4a379228eef9f58422079e786670be900ea3257e9d2e2227d0b56abd.
//
// Solidity: event DefaultOperatorRemoved(address indexed operator)
func (_Subredditpoints *SubredditpointsFilterer) ParseDefaultOperatorRemoved(log types.Log) (*SubredditpointsDefaultOperatorRemoved, error) {
	event := new(SubredditpointsDefaultOperatorRemoved)
	if err := _Subredditpoints.contract.UnpackLog(event, "DefaultOperatorRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Subredditpoints contract.
type SubredditpointsDepositIterator struct {
	Event *SubredditpointsDeposit // Event containing the contract specifics and raw log

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
func (it *SubredditpointsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsDeposit)
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
		it.Event = new(SubredditpointsDeposit)
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
func (it *SubredditpointsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsDeposit represents a Deposit event raised by the Subredditpoints contract.
type SubredditpointsDeposit struct {
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
func (_Subredditpoints *SubredditpointsFilterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*SubredditpointsDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsDepositIterator{contract: _Subredditpoints.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Subredditpoints *SubredditpointsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SubredditpointsDeposit, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "Deposit", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsDeposit)
				if err := _Subredditpoints.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Subredditpoints *SubredditpointsFilterer) ParseDeposit(log types.Log) (*SubredditpointsDeposit, error) {
	event := new(SubredditpointsDeposit)
	if err := _Subredditpoints.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsLogTransferIterator is returned from FilterLogTransfer and is used to iterate over the raw logs and unpacked data for LogTransfer events raised by the Subredditpoints contract.
type SubredditpointsLogTransferIterator struct {
	Event *SubredditpointsLogTransfer // Event containing the contract specifics and raw log

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
func (it *SubredditpointsLogTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsLogTransfer)
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
		it.Event = new(SubredditpointsLogTransfer)
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
func (it *SubredditpointsLogTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsLogTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsLogTransfer represents a LogTransfer event raised by the Subredditpoints contract.
type SubredditpointsLogTransfer struct {
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
func (_Subredditpoints *SubredditpointsFilterer) FilterLogTransfer(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*SubredditpointsLogTransferIterator, error) {

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

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsLogTransferIterator{contract: _Subredditpoints.contract, event: "LogTransfer", logs: logs, sub: sub}, nil
}

// WatchLogTransfer is a free log subscription operation binding the contract event 0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4.
//
// Solidity: event LogTransfer(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 input1, uint256 input2, uint256 output1, uint256 output2)
func (_Subredditpoints *SubredditpointsFilterer) WatchLogTransfer(opts *bind.WatchOpts, sink chan<- *SubredditpointsLogTransfer, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "LogTransfer", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsLogTransfer)
				if err := _Subredditpoints.contract.UnpackLog(event, "LogTransfer", log); err != nil {
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
func (_Subredditpoints *SubredditpointsFilterer) ParseLogTransfer(log types.Log) (*SubredditpointsLogTransfer, error) {
	event := new(SubredditpointsLogTransfer)
	if err := _Subredditpoints.contract.UnpackLog(event, "LogTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Subredditpoints contract.
type SubredditpointsMintedIterator struct {
	Event *SubredditpointsMinted // Event containing the contract specifics and raw log

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
func (it *SubredditpointsMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsMinted)
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
		it.Event = new(SubredditpointsMinted)
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
func (it *SubredditpointsMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsMinted represents a Minted event raised by the Subredditpoints contract.
type SubredditpointsMinted struct {
	Operator     common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) FilterMinted(opts *bind.FilterOpts, operator []common.Address, to []common.Address) (*SubredditpointsMintedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsMintedIterator{contract: _Subredditpoints.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *SubredditpointsMinted, operator []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "Minted", operatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsMinted)
				if err := _Subredditpoints.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x2fe5be0146f74c5bce36c0b80911af6c7d86ff27e89d5cfa61fc681327954e5d.
//
// Solidity: event Minted(address indexed operator, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) ParseMinted(log types.Log) (*SubredditpointsMinted, error) {
	event := new(SubredditpointsMinted)
	if err := _Subredditpoints.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Subredditpoints contract.
type SubredditpointsOwnershipTransferredIterator struct {
	Event *SubredditpointsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SubredditpointsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsOwnershipTransferred)
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
		it.Event = new(SubredditpointsOwnershipTransferred)
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
func (it *SubredditpointsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsOwnershipTransferred represents a OwnershipTransferred event raised by the Subredditpoints contract.
type SubredditpointsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subredditpoints *SubredditpointsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubredditpointsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsOwnershipTransferredIterator{contract: _Subredditpoints.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subredditpoints *SubredditpointsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubredditpointsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsOwnershipTransferred)
				if err := _Subredditpoints.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Subredditpoints *SubredditpointsFilterer) ParseOwnershipTransferred(log types.Log) (*SubredditpointsOwnershipTransferred, error) {
	event := new(SubredditpointsOwnershipTransferred)
	if err := _Subredditpoints.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsRevokedOperatorIterator is returned from FilterRevokedOperator and is used to iterate over the raw logs and unpacked data for RevokedOperator events raised by the Subredditpoints contract.
type SubredditpointsRevokedOperatorIterator struct {
	Event *SubredditpointsRevokedOperator // Event containing the contract specifics and raw log

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
func (it *SubredditpointsRevokedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsRevokedOperator)
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
		it.Event = new(SubredditpointsRevokedOperator)
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
func (it *SubredditpointsRevokedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsRevokedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsRevokedOperator represents a RevokedOperator event raised by the Subredditpoints contract.
type SubredditpointsRevokedOperator struct {
	Operator    common.Address
	TokenHolder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevokedOperator is a free log retrieval operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Subredditpoints *SubredditpointsFilterer) FilterRevokedOperator(opts *bind.FilterOpts, operator []common.Address, tokenHolder []common.Address) (*SubredditpointsRevokedOperatorIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsRevokedOperatorIterator{contract: _Subredditpoints.contract, event: "RevokedOperator", logs: logs, sub: sub}, nil
}

// WatchRevokedOperator is a free log subscription operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Subredditpoints *SubredditpointsFilterer) WatchRevokedOperator(opts *bind.WatchOpts, sink chan<- *SubredditpointsRevokedOperator, operator []common.Address, tokenHolder []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenHolderRule []interface{}
	for _, tokenHolderItem := range tokenHolder {
		tokenHolderRule = append(tokenHolderRule, tokenHolderItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "RevokedOperator", operatorRule, tokenHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsRevokedOperator)
				if err := _Subredditpoints.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
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

// ParseRevokedOperator is a log parse operation binding the contract event 0x50546e66e5f44d728365dc3908c63bc5cfeeab470722c1677e3073a6ac294aa1.
//
// Solidity: event RevokedOperator(address indexed operator, address indexed tokenHolder)
func (_Subredditpoints *SubredditpointsFilterer) ParseRevokedOperator(log types.Log) (*SubredditpointsRevokedOperator, error) {
	event := new(SubredditpointsRevokedOperator)
	if err := _Subredditpoints.contract.UnpackLog(event, "RevokedOperator", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Subredditpoints contract.
type SubredditpointsSentIterator struct {
	Event *SubredditpointsSent // Event containing the contract specifics and raw log

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
func (it *SubredditpointsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsSent)
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
		it.Event = new(SubredditpointsSent)
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
func (it *SubredditpointsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsSent represents a Sent event raised by the Subredditpoints contract.
type SubredditpointsSent struct {
	Operator     common.Address
	From         common.Address
	To           common.Address
	Amount       *big.Int
	Data         []byte
	OperatorData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) FilterSent(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*SubredditpointsSentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsSentIterator{contract: _Subredditpoints.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *SubredditpointsSent, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "Sent", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsSent)
				if err := _Subredditpoints.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x06b541ddaa720db2b10a4d0cdac39b8d360425fc073085fac19bc82614677987.
//
// Solidity: event Sent(address indexed operator, address indexed from, address indexed to, uint256 amount, bytes data, bytes operatorData)
func (_Subredditpoints *SubredditpointsFilterer) ParseSent(log types.Log) (*SubredditpointsSent, error) {
	event := new(SubredditpointsSent)
	if err := _Subredditpoints.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Subredditpoints contract.
type SubredditpointsTransferIterator struct {
	Event *SubredditpointsTransfer // Event containing the contract specifics and raw log

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
func (it *SubredditpointsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsTransfer)
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
		it.Event = new(SubredditpointsTransfer)
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
func (it *SubredditpointsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsTransfer represents a Transfer event raised by the Subredditpoints contract.
type SubredditpointsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Subredditpoints *SubredditpointsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SubredditpointsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsTransferIterator{contract: _Subredditpoints.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Subredditpoints *SubredditpointsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SubredditpointsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsTransfer)
				if err := _Subredditpoints.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Subredditpoints *SubredditpointsFilterer) ParseTransfer(log types.Log) (*SubredditpointsTransfer, error) {
	event := new(SubredditpointsTransfer)
	if err := _Subredditpoints.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SubredditpointsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Subredditpoints contract.
type SubredditpointsWithdrawIterator struct {
	Event *SubredditpointsWithdraw // Event containing the contract specifics and raw log

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
func (it *SubredditpointsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubredditpointsWithdraw)
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
		it.Event = new(SubredditpointsWithdraw)
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
func (it *SubredditpointsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubredditpointsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubredditpointsWithdraw represents a Withdraw event raised by the Subredditpoints contract.
type SubredditpointsWithdraw struct {
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
func (_Subredditpoints *SubredditpointsFilterer) FilterWithdraw(opts *bind.FilterOpts, token []common.Address, from []common.Address) (*SubredditpointsWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Subredditpoints.contract.FilterLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &SubredditpointsWithdrawIterator{contract: _Subredditpoints.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed token, address indexed from, uint256 amount, uint256 input1, uint256 output1)
func (_Subredditpoints *SubredditpointsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SubredditpointsWithdraw, token []common.Address, from []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Subredditpoints.contract.WatchLogs(opts, "Withdraw", tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubredditpointsWithdraw)
				if err := _Subredditpoints.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Subredditpoints *SubredditpointsFilterer) ParseWithdraw(log types.Log) (*SubredditpointsWithdraw, error) {
	event := new(SubredditpointsWithdraw)
	if err := _Subredditpoints.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
