// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorshare

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorshareABI is the input ABI used to generate the binding from.
const ValidatorshareABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastCommissionUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingLogger\",\"outputs\":[{\"internalType\":\"contractStakingInfo\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakingLogger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCommissionRate\",\"type\":\"uint256\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawRewardsValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSharesToMint\",\"type\":\"uint256\"}],\"name\":\"buyVoucher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minClaimAmount\",\"type\":\"uint256\"}],\"name\":\"sellVoucher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getLiquidRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unstakeClaimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valPow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountToSlash\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"drain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlockContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lockContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Validatorshare is an auto generated Go binding around an Ethereum contract.
type Validatorshare struct {
	ValidatorshareCaller     // Read-only binding to the contract
	ValidatorshareTransactor // Write-only binding to the contract
	ValidatorshareFilterer   // Log filterer for contract events
}

// ValidatorshareCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorshareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorshareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorshareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorshareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorshareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorshareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorshareSession struct {
	Contract     *Validatorshare   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorshareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorshareCallerSession struct {
	Contract *ValidatorshareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ValidatorshareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorshareTransactorSession struct {
	Contract     *ValidatorshareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ValidatorshareRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorshareRaw struct {
	Contract *Validatorshare // Generic contract binding to access the raw methods on
}

// ValidatorshareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorshareCallerRaw struct {
	Contract *ValidatorshareCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorshareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorshareTransactorRaw struct {
	Contract *ValidatorshareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorshare creates a new instance of Validatorshare, bound to a specific deployed contract.
func NewValidatorshare(address common.Address, backend bind.ContractBackend) (*Validatorshare, error) {
	contract, err := bindValidatorshare(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorshare{ValidatorshareCaller: ValidatorshareCaller{contract: contract}, ValidatorshareTransactor: ValidatorshareTransactor{contract: contract}, ValidatorshareFilterer: ValidatorshareFilterer{contract: contract}}, nil
}

// NewValidatorshareCaller creates a new read-only instance of Validatorshare, bound to a specific deployed contract.
func NewValidatorshareCaller(address common.Address, caller bind.ContractCaller) (*ValidatorshareCaller, error) {
	contract, err := bindValidatorshare(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorshareCaller{contract: contract}, nil
}

// NewValidatorshareTransactor creates a new write-only instance of Validatorshare, bound to a specific deployed contract.
func NewValidatorshareTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorshareTransactor, error) {
	contract, err := bindValidatorshare(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorshareTransactor{contract: contract}, nil
}

// NewValidatorshareFilterer creates a new log filterer instance of Validatorshare, bound to a specific deployed contract.
func NewValidatorshareFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorshareFilterer, error) {
	contract, err := bindValidatorshare(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorshareFilterer{contract: contract}, nil
}

// bindValidatorshare binds a generic wrapper to an already deployed contract.
func bindValidatorshare(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorshareABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorshare *ValidatorshareRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validatorshare.Contract.ValidatorshareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorshare *ValidatorshareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.Contract.ValidatorshareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorshare *ValidatorshareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorshare.Contract.ValidatorshareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorshare *ValidatorshareCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validatorshare.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorshare *ValidatorshareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorshare *ValidatorshareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorshare.Contract.contract.Transact(opts, method, params...)
}

// ActiveAmount is a free data retrieval call binding the contract method 0x3a09bf44.
//
// Solidity: function activeAmount() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) ActiveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "activeAmount")
	return *ret0, err
}

// ActiveAmount is a free data retrieval call binding the contract method 0x3a09bf44.
//
// Solidity: function activeAmount() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) ActiveAmount() (*big.Int, error) {
	return _Validatorshare.Contract.ActiveAmount(&_Validatorshare.CallOpts)
}

// ActiveAmount is a free data retrieval call binding the contract method 0x3a09bf44.
//
// Solidity: function activeAmount() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) ActiveAmount() (*big.Int, error) {
	return _Validatorshare.Contract.ActiveAmount(&_Validatorshare.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Validatorshare *ValidatorshareSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.Allowance(&_Validatorshare.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.Allowance(&_Validatorshare.CallOpts, owner, spender)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address ) view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) AmountStaked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "amountStaked", arg0)
	return *ret0, err
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address ) view returns(uint256)
func (_Validatorshare *ValidatorshareSession) AmountStaked(arg0 common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.AmountStaked(&_Validatorshare.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address ) view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) AmountStaked(arg0 common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.AmountStaked(&_Validatorshare.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Validatorshare *ValidatorshareSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.BalanceOf(&_Validatorshare.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.BalanceOf(&_Validatorshare.CallOpts, owner)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) CommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "commissionRate")
	return *ret0, err
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) CommissionRate() (*big.Int, error) {
	return _Validatorshare.Contract.CommissionRate(&_Validatorshare.CallOpts)
}

// CommissionRate is a free data retrieval call binding the contract method 0x5ea1d6f8.
//
// Solidity: function commissionRate() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) CommissionRate() (*big.Int, error) {
	return _Validatorshare.Contract.CommissionRate(&_Validatorshare.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_Validatorshare *ValidatorshareCaller) Delegation(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "delegation")
	return *ret0, err
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_Validatorshare *ValidatorshareSession) Delegation() (bool, error) {
	return _Validatorshare.Contract.Delegation(&_Validatorshare.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_Validatorshare *ValidatorshareCallerSession) Delegation() (bool, error) {
	return _Validatorshare.Contract.Delegation(&_Validatorshare.CallOpts)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 share, uint256 withdrawEpoch)
func (_Validatorshare *ValidatorshareCaller) Delegators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Share         *big.Int
	WithdrawEpoch *big.Int
}, error) {
	ret := new(struct {
		Share         *big.Int
		WithdrawEpoch *big.Int
	})
	out := ret
	err := _Validatorshare.contract.Call(opts, out, "delegators", arg0)
	return *ret, err
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 share, uint256 withdrawEpoch)
func (_Validatorshare *ValidatorshareSession) Delegators(arg0 common.Address) (struct {
	Share         *big.Int
	WithdrawEpoch *big.Int
}, error) {
	return _Validatorshare.Contract.Delegators(&_Validatorshare.CallOpts, arg0)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 share, uint256 withdrawEpoch)
func (_Validatorshare *ValidatorshareCallerSession) Delegators(arg0 common.Address) (struct {
	Share         *big.Int
	WithdrawEpoch *big.Int
}, error) {
	return _Validatorshare.Contract.Delegators(&_Validatorshare.CallOpts, arg0)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "exchangeRate")
	return *ret0, err
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) ExchangeRate() (*big.Int, error) {
	return _Validatorshare.Contract.ExchangeRate(&_Validatorshare.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) ExchangeRate() (*big.Int, error) {
	return _Validatorshare.Contract.ExchangeRate(&_Validatorshare.CallOpts)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) GetLiquidRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "getLiquidRewards", user)
	return *ret0, err
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_Validatorshare *ValidatorshareSession) GetLiquidRewards(user common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.GetLiquidRewards(&_Validatorshare.CallOpts, user)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) GetLiquidRewards(user common.Address) (*big.Int, error) {
	return _Validatorshare.Contract.GetLiquidRewards(&_Validatorshare.CallOpts, user)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Validatorshare *ValidatorshareCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Validatorshare *ValidatorshareSession) Governance() (common.Address, error) {
	return _Validatorshare.Contract.Governance(&_Validatorshare.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Validatorshare *ValidatorshareCallerSession) Governance() (common.Address, error) {
	return _Validatorshare.Contract.Governance(&_Validatorshare.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Validatorshare *ValidatorshareCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Validatorshare *ValidatorshareSession) IsOwner() (bool, error) {
	return _Validatorshare.Contract.IsOwner(&_Validatorshare.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Validatorshare *ValidatorshareCallerSession) IsOwner() (bool, error) {
	return _Validatorshare.Contract.IsOwner(&_Validatorshare.CallOpts)
}

// LastCommissionUpdate is a free data retrieval call binding the contract method 0x0d58a5f3.
//
// Solidity: function lastCommissionUpdate() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) LastCommissionUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "lastCommissionUpdate")
	return *ret0, err
}

// LastCommissionUpdate is a free data retrieval call binding the contract method 0x0d58a5f3.
//
// Solidity: function lastCommissionUpdate() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) LastCommissionUpdate() (*big.Int, error) {
	return _Validatorshare.Contract.LastCommissionUpdate(&_Validatorshare.CallOpts)
}

// LastCommissionUpdate is a free data retrieval call binding the contract method 0x0d58a5f3.
//
// Solidity: function lastCommissionUpdate() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) LastCommissionUpdate() (*big.Int, error) {
	return _Validatorshare.Contract.LastCommissionUpdate(&_Validatorshare.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Validatorshare *ValidatorshareCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Validatorshare *ValidatorshareSession) Locked() (bool, error) {
	return _Validatorshare.Contract.Locked(&_Validatorshare.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Validatorshare *ValidatorshareCallerSession) Locked() (bool, error) {
	return _Validatorshare.Contract.Locked(&_Validatorshare.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "minAmount")
	return *ret0, err
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) MinAmount() (*big.Int, error) {
	return _Validatorshare.Contract.MinAmount(&_Validatorshare.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) MinAmount() (*big.Int, error) {
	return _Validatorshare.Contract.MinAmount(&_Validatorshare.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validatorshare *ValidatorshareCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validatorshare *ValidatorshareSession) Owner() (common.Address, error) {
	return _Validatorshare.Contract.Owner(&_Validatorshare.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validatorshare *ValidatorshareCallerSession) Owner() (common.Address, error) {
	return _Validatorshare.Contract.Owner(&_Validatorshare.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) Rewards(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "rewards")
	return *ret0, err
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) Rewards() (*big.Int, error) {
	return _Validatorshare.Contract.Rewards(&_Validatorshare.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x9ec5a894.
//
// Solidity: function rewards() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) Rewards() (*big.Int, error) {
	return _Validatorshare.Contract.Rewards(&_Validatorshare.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Validatorshare *ValidatorshareCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "stakeManager")
	return *ret0, err
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Validatorshare *ValidatorshareSession) StakeManager() (common.Address, error) {
	return _Validatorshare.Contract.StakeManager(&_Validatorshare.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_Validatorshare *ValidatorshareCallerSession) StakeManager() (common.Address, error) {
	return _Validatorshare.Contract.StakeManager(&_Validatorshare.CallOpts)
}

// StakingLogger is a free data retrieval call binding the contract method 0x3d94eb05.
//
// Solidity: function stakingLogger() view returns(address)
func (_Validatorshare *ValidatorshareCaller) StakingLogger(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "stakingLogger")
	return *ret0, err
}

// StakingLogger is a free data retrieval call binding the contract method 0x3d94eb05.
//
// Solidity: function stakingLogger() view returns(address)
func (_Validatorshare *ValidatorshareSession) StakingLogger() (common.Address, error) {
	return _Validatorshare.Contract.StakingLogger(&_Validatorshare.CallOpts)
}

// StakingLogger is a free data retrieval call binding the contract method 0x3d94eb05.
//
// Solidity: function stakingLogger() view returns(address)
func (_Validatorshare *ValidatorshareCallerSession) StakingLogger() (common.Address, error) {
	return _Validatorshare.Contract.StakingLogger(&_Validatorshare.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "totalStake")
	return *ret0, err
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) TotalStake() (*big.Int, error) {
	return _Validatorshare.Contract.TotalStake(&_Validatorshare.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) TotalStake() (*big.Int, error) {
	return _Validatorshare.Contract.TotalStake(&_Validatorshare.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) TotalSupply() (*big.Int, error) {
	return _Validatorshare.Contract.TotalSupply(&_Validatorshare.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) TotalSupply() (*big.Int, error) {
	return _Validatorshare.Contract.TotalSupply(&_Validatorshare.CallOpts)
}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) ValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "validatorId")
	return *ret0, err
}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) ValidatorId() (*big.Int, error) {
	return _Validatorshare.Contract.ValidatorId(&_Validatorshare.CallOpts)
}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) ValidatorId() (*big.Int, error) {
	return _Validatorshare.Contract.ValidatorId(&_Validatorshare.CallOpts)
}

// ValidatorRewards is a free data retrieval call binding the contract method 0x5d1e3616.
//
// Solidity: function validatorRewards() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) ValidatorRewards(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "validatorRewards")
	return *ret0, err
}

// ValidatorRewards is a free data retrieval call binding the contract method 0x5d1e3616.
//
// Solidity: function validatorRewards() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) ValidatorRewards() (*big.Int, error) {
	return _Validatorshare.Contract.ValidatorRewards(&_Validatorshare.CallOpts)
}

// ValidatorRewards is a free data retrieval call binding the contract method 0x5d1e3616.
//
// Solidity: function validatorRewards() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) ValidatorRewards() (*big.Int, error) {
	return _Validatorshare.Contract.ValidatorRewards(&_Validatorshare.CallOpts)
}

// WithdrawExchangeRate is a free data retrieval call binding the contract method 0xbfb18f29.
//
// Solidity: function withdrawExchangeRate() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) WithdrawExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "withdrawExchangeRate")
	return *ret0, err
}

// WithdrawExchangeRate is a free data retrieval call binding the contract method 0xbfb18f29.
//
// Solidity: function withdrawExchangeRate() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) WithdrawExchangeRate() (*big.Int, error) {
	return _Validatorshare.Contract.WithdrawExchangeRate(&_Validatorshare.CallOpts)
}

// WithdrawExchangeRate is a free data retrieval call binding the contract method 0xbfb18f29.
//
// Solidity: function withdrawExchangeRate() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) WithdrawExchangeRate() (*big.Int, error) {
	return _Validatorshare.Contract.WithdrawExchangeRate(&_Validatorshare.CallOpts)
}

// WithdrawPool is a free data retrieval call binding the contract method 0x5c42c733.
//
// Solidity: function withdrawPool() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) WithdrawPool(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "withdrawPool")
	return *ret0, err
}

// WithdrawPool is a free data retrieval call binding the contract method 0x5c42c733.
//
// Solidity: function withdrawPool() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) WithdrawPool() (*big.Int, error) {
	return _Validatorshare.Contract.WithdrawPool(&_Validatorshare.CallOpts)
}

// WithdrawPool is a free data retrieval call binding the contract method 0x5c42c733.
//
// Solidity: function withdrawPool() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) WithdrawPool() (*big.Int, error) {
	return _Validatorshare.Contract.WithdrawPool(&_Validatorshare.CallOpts)
}

// WithdrawShares is a free data retrieval call binding the contract method 0x8d086da4.
//
// Solidity: function withdrawShares() view returns(uint256)
func (_Validatorshare *ValidatorshareCaller) WithdrawShares(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validatorshare.contract.Call(opts, out, "withdrawShares")
	return *ret0, err
}

// WithdrawShares is a free data retrieval call binding the contract method 0x8d086da4.
//
// Solidity: function withdrawShares() view returns(uint256)
func (_Validatorshare *ValidatorshareSession) WithdrawShares() (*big.Int, error) {
	return _Validatorshare.Contract.WithdrawShares(&_Validatorshare.CallOpts)
}

// WithdrawShares is a free data retrieval call binding the contract method 0x8d086da4.
//
// Solidity: function withdrawShares() view returns(uint256)
func (_Validatorshare *ValidatorshareCallerSession) WithdrawShares() (*big.Int, error) {
	return _Validatorshare.Contract.WithdrawShares(&_Validatorshare.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Approve(&_Validatorshare.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Approve(&_Validatorshare.TransactOpts, spender, value)
}

// BuyVoucher is a paid mutator transaction binding the contract method 0x6ab15071.
//
// Solidity: function buyVoucher(uint256 _amount, uint256 _minSharesToMint) returns()
func (_Validatorshare *ValidatorshareTransactor) BuyVoucher(opts *bind.TransactOpts, _amount *big.Int, _minSharesToMint *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "buyVoucher", _amount, _minSharesToMint)
}

// BuyVoucher is a paid mutator transaction binding the contract method 0x6ab15071.
//
// Solidity: function buyVoucher(uint256 _amount, uint256 _minSharesToMint) returns()
func (_Validatorshare *ValidatorshareSession) BuyVoucher(_amount *big.Int, _minSharesToMint *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.BuyVoucher(&_Validatorshare.TransactOpts, _amount, _minSharesToMint)
}

// BuyVoucher is a paid mutator transaction binding the contract method 0x6ab15071.
//
// Solidity: function buyVoucher(uint256 _amount, uint256 _minSharesToMint) returns()
func (_Validatorshare *ValidatorshareTransactorSession) BuyVoucher(_amount *big.Int, _minSharesToMint *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.BuyVoucher(&_Validatorshare.TransactOpts, _amount, _minSharesToMint)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Validatorshare *ValidatorshareTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Validatorshare *ValidatorshareSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.DecreaseAllowance(&_Validatorshare.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Validatorshare *ValidatorshareTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.DecreaseAllowance(&_Validatorshare.TransactOpts, spender, subtractedValue)
}

// Drain is a paid mutator transaction binding the contract method 0xabf59fc9.
//
// Solidity: function drain(address token, address destination, uint256 amount) returns()
func (_Validatorshare *ValidatorshareTransactor) Drain(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "drain", token, destination, amount)
}

// Drain is a paid mutator transaction binding the contract method 0xabf59fc9.
//
// Solidity: function drain(address token, address destination, uint256 amount) returns()
func (_Validatorshare *ValidatorshareSession) Drain(token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Drain(&_Validatorshare.TransactOpts, token, destination, amount)
}

// Drain is a paid mutator transaction binding the contract method 0xabf59fc9.
//
// Solidity: function drain(address token, address destination, uint256 amount) returns()
func (_Validatorshare *ValidatorshareTransactorSession) Drain(token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Drain(&_Validatorshare.TransactOpts, token, destination, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Validatorshare *ValidatorshareTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Validatorshare *ValidatorshareSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.IncreaseAllowance(&_Validatorshare.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Validatorshare *ValidatorshareTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.IncreaseAllowance(&_Validatorshare.TransactOpts, spender, addedValue)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Validatorshare *ValidatorshareTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Validatorshare *ValidatorshareSession) Lock() (*types.Transaction, error) {
	return _Validatorshare.Contract.Lock(&_Validatorshare.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Validatorshare *ValidatorshareTransactorSession) Lock() (*types.Transaction, error) {
	return _Validatorshare.Contract.Lock(&_Validatorshare.TransactOpts)
}

// LockContract is a paid mutator transaction binding the contract method 0x753868e3.
//
// Solidity: function lockContract() returns(uint256)
func (_Validatorshare *ValidatorshareTransactor) LockContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "lockContract")
}

// LockContract is a paid mutator transaction binding the contract method 0x753868e3.
//
// Solidity: function lockContract() returns(uint256)
func (_Validatorshare *ValidatorshareSession) LockContract() (*types.Transaction, error) {
	return _Validatorshare.Contract.LockContract(&_Validatorshare.TransactOpts)
}

// LockContract is a paid mutator transaction binding the contract method 0x753868e3.
//
// Solidity: function lockContract() returns(uint256)
func (_Validatorshare *ValidatorshareTransactorSession) LockContract() (*types.Transaction, error) {
	return _Validatorshare.Contract.LockContract(&_Validatorshare.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validatorshare *ValidatorshareTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validatorshare *ValidatorshareSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validatorshare.Contract.RenounceOwnership(&_Validatorshare.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validatorshare *ValidatorshareTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validatorshare.Contract.RenounceOwnership(&_Validatorshare.TransactOpts)
}

// Restake is a paid mutator transaction binding the contract method 0x4f91440d.
//
// Solidity: function restake() returns()
func (_Validatorshare *ValidatorshareTransactor) Restake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "restake")
}

// Restake is a paid mutator transaction binding the contract method 0x4f91440d.
//
// Solidity: function restake() returns()
func (_Validatorshare *ValidatorshareSession) Restake() (*types.Transaction, error) {
	return _Validatorshare.Contract.Restake(&_Validatorshare.TransactOpts)
}

// Restake is a paid mutator transaction binding the contract method 0x4f91440d.
//
// Solidity: function restake() returns()
func (_Validatorshare *ValidatorshareTransactorSession) Restake() (*types.Transaction, error) {
	return _Validatorshare.Contract.Restake(&_Validatorshare.TransactOpts)
}

// SellVoucher is a paid mutator transaction binding the contract method 0x6c751897.
//
// Solidity: function sellVoucher(uint256 _minClaimAmount) returns()
func (_Validatorshare *ValidatorshareTransactor) SellVoucher(opts *bind.TransactOpts, _minClaimAmount *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "sellVoucher", _minClaimAmount)
}

// SellVoucher is a paid mutator transaction binding the contract method 0x6c751897.
//
// Solidity: function sellVoucher(uint256 _minClaimAmount) returns()
func (_Validatorshare *ValidatorshareSession) SellVoucher(_minClaimAmount *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.SellVoucher(&_Validatorshare.TransactOpts, _minClaimAmount)
}

// SellVoucher is a paid mutator transaction binding the contract method 0x6c751897.
//
// Solidity: function sellVoucher(uint256 _minClaimAmount) returns()
func (_Validatorshare *ValidatorshareTransactorSession) SellVoucher(_minClaimAmount *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.SellVoucher(&_Validatorshare.TransactOpts, _minClaimAmount)
}

// Slash is a paid mutator transaction binding the contract method 0xa22a6428.
//
// Solidity: function slash(uint256 valPow, uint256 totalAmountToSlash) returns(uint256)
func (_Validatorshare *ValidatorshareTransactor) Slash(opts *bind.TransactOpts, valPow *big.Int, totalAmountToSlash *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "slash", valPow, totalAmountToSlash)
}

// Slash is a paid mutator transaction binding the contract method 0xa22a6428.
//
// Solidity: function slash(uint256 valPow, uint256 totalAmountToSlash) returns(uint256)
func (_Validatorshare *ValidatorshareSession) Slash(valPow *big.Int, totalAmountToSlash *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Slash(&_Validatorshare.TransactOpts, valPow, totalAmountToSlash)
}

// Slash is a paid mutator transaction binding the contract method 0xa22a6428.
//
// Solidity: function slash(uint256 valPow, uint256 totalAmountToSlash) returns(uint256)
func (_Validatorshare *ValidatorshareTransactorSession) Slash(valPow *big.Int, totalAmountToSlash *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Slash(&_Validatorshare.TransactOpts, valPow, totalAmountToSlash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Transfer(&_Validatorshare.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.Transfer(&_Validatorshare.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.TransferFrom(&_Validatorshare.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Validatorshare *ValidatorshareTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.TransferFrom(&_Validatorshare.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validatorshare *ValidatorshareTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validatorshare *ValidatorshareSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validatorshare.Contract.TransferOwnership(&_Validatorshare.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validatorshare *ValidatorshareTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validatorshare.Contract.TransferOwnership(&_Validatorshare.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Validatorshare *ValidatorshareTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Validatorshare *ValidatorshareSession) Unlock() (*types.Transaction, error) {
	return _Validatorshare.Contract.Unlock(&_Validatorshare.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Validatorshare *ValidatorshareTransactorSession) Unlock() (*types.Transaction, error) {
	return _Validatorshare.Contract.Unlock(&_Validatorshare.TransactOpts)
}

// UnlockContract is a paid mutator transaction binding the contract method 0x22f0f2f9.
//
// Solidity: function unlockContract() returns(uint256)
func (_Validatorshare *ValidatorshareTransactor) UnlockContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "unlockContract")
}

// UnlockContract is a paid mutator transaction binding the contract method 0x22f0f2f9.
//
// Solidity: function unlockContract() returns(uint256)
func (_Validatorshare *ValidatorshareSession) UnlockContract() (*types.Transaction, error) {
	return _Validatorshare.Contract.UnlockContract(&_Validatorshare.TransactOpts)
}

// UnlockContract is a paid mutator transaction binding the contract method 0x22f0f2f9.
//
// Solidity: function unlockContract() returns(uint256)
func (_Validatorshare *ValidatorshareTransactorSession) UnlockContract() (*types.Transaction, error) {
	return _Validatorshare.Contract.UnlockContract(&_Validatorshare.TransactOpts)
}

// UnstakeClaimTokens is a paid mutator transaction binding the contract method 0x8d16a14a.
//
// Solidity: function unstakeClaimTokens() returns()
func (_Validatorshare *ValidatorshareTransactor) UnstakeClaimTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "unstakeClaimTokens")
}

// UnstakeClaimTokens is a paid mutator transaction binding the contract method 0x8d16a14a.
//
// Solidity: function unstakeClaimTokens() returns()
func (_Validatorshare *ValidatorshareSession) UnstakeClaimTokens() (*types.Transaction, error) {
	return _Validatorshare.Contract.UnstakeClaimTokens(&_Validatorshare.TransactOpts)
}

// UnstakeClaimTokens is a paid mutator transaction binding the contract method 0x8d16a14a.
//
// Solidity: function unstakeClaimTokens() returns()
func (_Validatorshare *ValidatorshareTransactorSession) UnstakeClaimTokens() (*types.Transaction, error) {
	return _Validatorshare.Contract.UnstakeClaimTokens(&_Validatorshare.TransactOpts)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 newCommissionRate) returns()
func (_Validatorshare *ValidatorshareTransactor) UpdateCommissionRate(opts *bind.TransactOpts, newCommissionRate *big.Int) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "updateCommissionRate", newCommissionRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 newCommissionRate) returns()
func (_Validatorshare *ValidatorshareSession) UpdateCommissionRate(newCommissionRate *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.UpdateCommissionRate(&_Validatorshare.TransactOpts, newCommissionRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 newCommissionRate) returns()
func (_Validatorshare *ValidatorshareTransactorSession) UpdateCommissionRate(newCommissionRate *big.Int) (*types.Transaction, error) {
	return _Validatorshare.Contract.UpdateCommissionRate(&_Validatorshare.TransactOpts, newCommissionRate)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_Validatorshare *ValidatorshareTransactor) WithdrawRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "withdrawRewards")
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_Validatorshare *ValidatorshareSession) WithdrawRewards() (*types.Transaction, error) {
	return _Validatorshare.Contract.WithdrawRewards(&_Validatorshare.TransactOpts)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_Validatorshare *ValidatorshareTransactorSession) WithdrawRewards() (*types.Transaction, error) {
	return _Validatorshare.Contract.WithdrawRewards(&_Validatorshare.TransactOpts)
}

// WithdrawRewardsValidator is a paid mutator transaction binding the contract method 0x32ba2e53.
//
// Solidity: function withdrawRewardsValidator() returns(uint256)
func (_Validatorshare *ValidatorshareTransactor) WithdrawRewardsValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorshare.contract.Transact(opts, "withdrawRewardsValidator")
}

// WithdrawRewardsValidator is a paid mutator transaction binding the contract method 0x32ba2e53.
//
// Solidity: function withdrawRewardsValidator() returns(uint256)
func (_Validatorshare *ValidatorshareSession) WithdrawRewardsValidator() (*types.Transaction, error) {
	return _Validatorshare.Contract.WithdrawRewardsValidator(&_Validatorshare.TransactOpts)
}

// WithdrawRewardsValidator is a paid mutator transaction binding the contract method 0x32ba2e53.
//
// Solidity: function withdrawRewardsValidator() returns(uint256)
func (_Validatorshare *ValidatorshareTransactorSession) WithdrawRewardsValidator() (*types.Transaction, error) {
	return _Validatorshare.Contract.WithdrawRewardsValidator(&_Validatorshare.TransactOpts)
}

// ValidatorshareApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Validatorshare contract.
type ValidatorshareApprovalIterator struct {
	Event *ValidatorshareApproval // Event containing the contract specifics and raw log

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
func (it *ValidatorshareApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorshareApproval)
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
		it.Event = new(ValidatorshareApproval)
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
func (it *ValidatorshareApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorshareApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorshareApproval represents a Approval event raised by the Validatorshare contract.
type ValidatorshareApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Validatorshare *ValidatorshareFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ValidatorshareApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Validatorshare.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorshareApprovalIterator{contract: _Validatorshare.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Validatorshare *ValidatorshareFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ValidatorshareApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Validatorshare.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorshareApproval)
				if err := _Validatorshare.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Validatorshare *ValidatorshareFilterer) ParseApproval(log types.Log) (*ValidatorshareApproval, error) {
	event := new(ValidatorshareApproval)
	if err := _Validatorshare.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorshareOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Validatorshare contract.
type ValidatorshareOwnershipTransferredIterator struct {
	Event *ValidatorshareOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorshareOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorshareOwnershipTransferred)
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
		it.Event = new(ValidatorshareOwnershipTransferred)
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
func (it *ValidatorshareOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorshareOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorshareOwnershipTransferred represents a OwnershipTransferred event raised by the Validatorshare contract.
type ValidatorshareOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validatorshare *ValidatorshareFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorshareOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validatorshare.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorshareOwnershipTransferredIterator{contract: _Validatorshare.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validatorshare *ValidatorshareFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorshareOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validatorshare.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorshareOwnershipTransferred)
				if err := _Validatorshare.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Validatorshare *ValidatorshareFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorshareOwnershipTransferred, error) {
	event := new(ValidatorshareOwnershipTransferred)
	if err := _Validatorshare.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorshareTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Validatorshare contract.
type ValidatorshareTransferIterator struct {
	Event *ValidatorshareTransfer // Event containing the contract specifics and raw log

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
func (it *ValidatorshareTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorshareTransfer)
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
		it.Event = new(ValidatorshareTransfer)
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
func (it *ValidatorshareTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorshareTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorshareTransfer represents a Transfer event raised by the Validatorshare contract.
type ValidatorshareTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Validatorshare *ValidatorshareFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ValidatorshareTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Validatorshare.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorshareTransferIterator{contract: _Validatorshare.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Validatorshare *ValidatorshareFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ValidatorshareTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Validatorshare.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorshareTransfer)
				if err := _Validatorshare.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Validatorshare *ValidatorshareFilterer) ParseTransfer(log types.Log) (*ValidatorshareTransfer, error) {
	event := new(ValidatorshareTransfer)
	if err := _Validatorshare.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
