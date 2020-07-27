// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawmanager

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

// WithdrawmanagerABI is the input ABI used to generate the binding from.
const WithdrawmanagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"exitWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptAmountOrNFTId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isRegularExit\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"predicate\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ownerExits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exitsQueues\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"HALF_EXIT_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitNft\",\"outputs\":[{\"internalType\":\"contractExitNFT\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRegularExit\",\"type\":\"bool\"}],\"name\":\"ExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ExitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldExitPeriod\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newExitPeriod\",\"type\":\"uint256\"}],\"name\":\"ExitPeriodUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"}],\"name\":\"ExitCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"createExitQueue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"offset\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"verifyTxInclusion\",\"type\":\"bool\"}],\"name\":\"verifyInclusion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrToken\",\"type\":\"uint256\"}],\"name\":\"startExitWithDepositedTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"exitor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exitAmountOrTokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isRegularExit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"priority\",\"type\":\"uint256\"}],\"name\":\"addExitToQueue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inputId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"challengeData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"adjudicatorPredicate\",\"type\":\"address\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"processExits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exitId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"utxoOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addInput\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"halfExitPeriod\",\"type\":\"uint256\"}],\"name\":\"updateExitPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Withdrawmanager is an auto generated Go binding around an Ethereum contract.
type Withdrawmanager struct {
	WithdrawmanagerCaller     // Read-only binding to the contract
	WithdrawmanagerTransactor // Write-only binding to the contract
	WithdrawmanagerFilterer   // Log filterer for contract events
}

// WithdrawmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawmanagerSession struct {
	Contract     *Withdrawmanager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawmanagerCallerSession struct {
	Contract *WithdrawmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WithdrawmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawmanagerTransactorSession struct {
	Contract     *WithdrawmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WithdrawmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawmanagerRaw struct {
	Contract *Withdrawmanager // Generic contract binding to access the raw methods on
}

// WithdrawmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawmanagerCallerRaw struct {
	Contract *WithdrawmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawmanagerTransactorRaw struct {
	Contract *WithdrawmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawmanager creates a new instance of Withdrawmanager, bound to a specific deployed contract.
func NewWithdrawmanager(address common.Address, backend bind.ContractBackend) (*Withdrawmanager, error) {
	contract, err := bindWithdrawmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdrawmanager{WithdrawmanagerCaller: WithdrawmanagerCaller{contract: contract}, WithdrawmanagerTransactor: WithdrawmanagerTransactor{contract: contract}, WithdrawmanagerFilterer: WithdrawmanagerFilterer{contract: contract}}, nil
}

// NewWithdrawmanagerCaller creates a new read-only instance of Withdrawmanager, bound to a specific deployed contract.
func NewWithdrawmanagerCaller(address common.Address, caller bind.ContractCaller) (*WithdrawmanagerCaller, error) {
	contract, err := bindWithdrawmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerCaller{contract: contract}, nil
}

// NewWithdrawmanagerTransactor creates a new write-only instance of Withdrawmanager, bound to a specific deployed contract.
func NewWithdrawmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawmanagerTransactor, error) {
	contract, err := bindWithdrawmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerTransactor{contract: contract}, nil
}

// NewWithdrawmanagerFilterer creates a new log filterer instance of Withdrawmanager, bound to a specific deployed contract.
func NewWithdrawmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawmanagerFilterer, error) {
	contract, err := bindWithdrawmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerFilterer{contract: contract}, nil
}

// bindWithdrawmanager binds a generic wrapper to an already deployed contract.
func bindWithdrawmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawmanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawmanager *WithdrawmanagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Withdrawmanager.Contract.WithdrawmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawmanager *WithdrawmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.WithdrawmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawmanager *WithdrawmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.WithdrawmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawmanager *WithdrawmanagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Withdrawmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawmanager *WithdrawmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawmanager *WithdrawmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.contract.Transact(opts, method, params...)
}

// HALFEXITPERIOD is a free data retrieval call binding the contract method 0xed4a0be8.
//
// Solidity: function HALF_EXIT_PERIOD() constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCaller) HALFEXITPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "HALF_EXIT_PERIOD")
	return *ret0, err
}

// HALFEXITPERIOD is a free data retrieval call binding the contract method 0xed4a0be8.
//
// Solidity: function HALF_EXIT_PERIOD() constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerSession) HALFEXITPERIOD() (*big.Int, error) {
	return _Withdrawmanager.Contract.HALFEXITPERIOD(&_Withdrawmanager.CallOpts)
}

// HALFEXITPERIOD is a free data retrieval call binding the contract method 0xed4a0be8.
//
// Solidity: function HALF_EXIT_PERIOD() constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCallerSession) HALFEXITPERIOD() (*big.Int, error) {
	return _Withdrawmanager.Contract.HALFEXITPERIOD(&_Withdrawmanager.CallOpts)
}

// ExitNft is a free data retrieval call binding the contract method 0xedeca09b.
//
// Solidity: function exitNft() constant returns(address)
func (_Withdrawmanager *WithdrawmanagerCaller) ExitNft(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "exitNft")
	return *ret0, err
}

// ExitNft is a free data retrieval call binding the contract method 0xedeca09b.
//
// Solidity: function exitNft() constant returns(address)
func (_Withdrawmanager *WithdrawmanagerSession) ExitNft() (common.Address, error) {
	return _Withdrawmanager.Contract.ExitNft(&_Withdrawmanager.CallOpts)
}

// ExitNft is a free data retrieval call binding the contract method 0xedeca09b.
//
// Solidity: function exitNft() constant returns(address)
func (_Withdrawmanager *WithdrawmanagerCallerSession) ExitNft() (common.Address, error) {
	return _Withdrawmanager.Contract.ExitNft(&_Withdrawmanager.CallOpts)
}

// ExitWindow is a free data retrieval call binding the contract method 0x1e29848b.
//
// Solidity: function exitWindow() constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCaller) ExitWindow(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "exitWindow")
	return *ret0, err
}

// ExitWindow is a free data retrieval call binding the contract method 0x1e29848b.
//
// Solidity: function exitWindow() constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerSession) ExitWindow() (*big.Int, error) {
	return _Withdrawmanager.Contract.ExitWindow(&_Withdrawmanager.CallOpts)
}

// ExitWindow is a free data retrieval call binding the contract method 0x1e29848b.
//
// Solidity: function exitWindow() constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCallerSession) ExitWindow() (*big.Int, error) {
	return _Withdrawmanager.Contract.ExitWindow(&_Withdrawmanager.CallOpts)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits(uint256 ) constant returns(uint256 receiptAmountOrNFTId, bytes32 txHash, address owner, address token, bool isRegularExit, address predicate)
func (_Withdrawmanager *WithdrawmanagerCaller) Exits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ReceiptAmountOrNFTId *big.Int
	TxHash               [32]byte
	Owner                common.Address
	Token                common.Address
	IsRegularExit        bool
	Predicate            common.Address
}, error) {
	ret := new(struct {
		ReceiptAmountOrNFTId *big.Int
		TxHash               [32]byte
		Owner                common.Address
		Token                common.Address
		IsRegularExit        bool
		Predicate            common.Address
	})
	out := ret
	err := _Withdrawmanager.contract.Call(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits(uint256 ) constant returns(uint256 receiptAmountOrNFTId, bytes32 txHash, address owner, address token, bool isRegularExit, address predicate)
func (_Withdrawmanager *WithdrawmanagerSession) Exits(arg0 *big.Int) (struct {
	ReceiptAmountOrNFTId *big.Int
	TxHash               [32]byte
	Owner                common.Address
	Token                common.Address
	IsRegularExit        bool
	Predicate            common.Address
}, error) {
	return _Withdrawmanager.Contract.Exits(&_Withdrawmanager.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits(uint256 ) constant returns(uint256 receiptAmountOrNFTId, bytes32 txHash, address owner, address token, bool isRegularExit, address predicate)
func (_Withdrawmanager *WithdrawmanagerCallerSession) Exits(arg0 *big.Int) (struct {
	ReceiptAmountOrNFTId *big.Int
	TxHash               [32]byte
	Owner                common.Address
	Token                common.Address
	IsRegularExit        bool
	Predicate            common.Address
}, error) {
	return _Withdrawmanager.Contract.Exits(&_Withdrawmanager.CallOpts, arg0)
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues(address ) constant returns(address)
func (_Withdrawmanager *WithdrawmanagerCaller) ExitsQueues(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "exitsQueues", arg0)
	return *ret0, err
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues(address ) constant returns(address)
func (_Withdrawmanager *WithdrawmanagerSession) ExitsQueues(arg0 common.Address) (common.Address, error) {
	return _Withdrawmanager.Contract.ExitsQueues(&_Withdrawmanager.CallOpts, arg0)
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues(address ) constant returns(address)
func (_Withdrawmanager *WithdrawmanagerCallerSession) ExitsQueues(arg0 common.Address) (common.Address, error) {
	return _Withdrawmanager.Contract.ExitsQueues(&_Withdrawmanager.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Withdrawmanager *WithdrawmanagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Withdrawmanager *WithdrawmanagerSession) IsOwner() (bool, error) {
	return _Withdrawmanager.Contract.IsOwner(&_Withdrawmanager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Withdrawmanager *WithdrawmanagerCallerSession) IsOwner() (bool, error) {
	return _Withdrawmanager.Contract.IsOwner(&_Withdrawmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Withdrawmanager *WithdrawmanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Withdrawmanager *WithdrawmanagerSession) Owner() (common.Address, error) {
	return _Withdrawmanager.Contract.Owner(&_Withdrawmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Withdrawmanager *WithdrawmanagerCallerSession) Owner() (common.Address, error) {
	return _Withdrawmanager.Contract.Owner(&_Withdrawmanager.CallOpts)
}

// OwnerExits is a free data retrieval call binding the contract method 0x661429c8.
//
// Solidity: function ownerExits(bytes32 ) constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCaller) OwnerExits(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "ownerExits", arg0)
	return *ret0, err
}

// OwnerExits is a free data retrieval call binding the contract method 0x661429c8.
//
// Solidity: function ownerExits(bytes32 ) constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerSession) OwnerExits(arg0 [32]byte) (*big.Int, error) {
	return _Withdrawmanager.Contract.OwnerExits(&_Withdrawmanager.CallOpts, arg0)
}

// OwnerExits is a free data retrieval call binding the contract method 0x661429c8.
//
// Solidity: function ownerExits(bytes32 ) constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCallerSession) OwnerExits(arg0 [32]byte) (*big.Int, error) {
	return _Withdrawmanager.Contract.OwnerExits(&_Withdrawmanager.CallOpts, arg0)
}

// VerifyInclusion is a free data retrieval call binding the contract method 0xad1d8069.
//
// Solidity: function verifyInclusion(bytes data, uint8 offset, bool verifyTxInclusion) constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCaller) VerifyInclusion(opts *bind.CallOpts, data []byte, offset uint8, verifyTxInclusion bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Withdrawmanager.contract.Call(opts, out, "verifyInclusion", data, offset, verifyTxInclusion)
	return *ret0, err
}

// VerifyInclusion is a free data retrieval call binding the contract method 0xad1d8069.
//
// Solidity: function verifyInclusion(bytes data, uint8 offset, bool verifyTxInclusion) constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerSession) VerifyInclusion(data []byte, offset uint8, verifyTxInclusion bool) (*big.Int, error) {
	return _Withdrawmanager.Contract.VerifyInclusion(&_Withdrawmanager.CallOpts, data, offset, verifyTxInclusion)
}

// VerifyInclusion is a free data retrieval call binding the contract method 0xad1d8069.
//
// Solidity: function verifyInclusion(bytes data, uint8 offset, bool verifyTxInclusion) constant returns(uint256)
func (_Withdrawmanager *WithdrawmanagerCallerSession) VerifyInclusion(data []byte, offset uint8, verifyTxInclusion bool) (*big.Int, error) {
	return _Withdrawmanager.Contract.VerifyInclusion(&_Withdrawmanager.CallOpts, data, offset, verifyTxInclusion)
}

// AddExitToQueue is a paid mutator transaction binding the contract method 0xd931a869.
//
// Solidity: function addExitToQueue(address exitor, address childToken, address rootToken, uint256 exitAmountOrTokenId, bytes32 txHash, bool isRegularExit, uint256 priority) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) AddExitToQueue(opts *bind.TransactOpts, exitor common.Address, childToken common.Address, rootToken common.Address, exitAmountOrTokenId *big.Int, txHash [32]byte, isRegularExit bool, priority *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "addExitToQueue", exitor, childToken, rootToken, exitAmountOrTokenId, txHash, isRegularExit, priority)
}

// AddExitToQueue is a paid mutator transaction binding the contract method 0xd931a869.
//
// Solidity: function addExitToQueue(address exitor, address childToken, address rootToken, uint256 exitAmountOrTokenId, bytes32 txHash, bool isRegularExit, uint256 priority) returns()
func (_Withdrawmanager *WithdrawmanagerSession) AddExitToQueue(exitor common.Address, childToken common.Address, rootToken common.Address, exitAmountOrTokenId *big.Int, txHash [32]byte, isRegularExit bool, priority *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.AddExitToQueue(&_Withdrawmanager.TransactOpts, exitor, childToken, rootToken, exitAmountOrTokenId, txHash, isRegularExit, priority)
}

// AddExitToQueue is a paid mutator transaction binding the contract method 0xd931a869.
//
// Solidity: function addExitToQueue(address exitor, address childToken, address rootToken, uint256 exitAmountOrTokenId, bytes32 txHash, bool isRegularExit, uint256 priority) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) AddExitToQueue(exitor common.Address, childToken common.Address, rootToken common.Address, exitAmountOrTokenId *big.Int, txHash [32]byte, isRegularExit bool, priority *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.AddExitToQueue(&_Withdrawmanager.TransactOpts, exitor, childToken, rootToken, exitAmountOrTokenId, txHash, isRegularExit, priority)
}

// AddInput is a paid mutator transaction binding the contract method 0x22f192af.
//
// Solidity: function addInput(uint256 exitId, uint256 age, address utxoOwner, address token) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) AddInput(opts *bind.TransactOpts, exitId *big.Int, age *big.Int, utxoOwner common.Address, token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "addInput", exitId, age, utxoOwner, token)
}

// AddInput is a paid mutator transaction binding the contract method 0x22f192af.
//
// Solidity: function addInput(uint256 exitId, uint256 age, address utxoOwner, address token) returns()
func (_Withdrawmanager *WithdrawmanagerSession) AddInput(exitId *big.Int, age *big.Int, utxoOwner common.Address, token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.AddInput(&_Withdrawmanager.TransactOpts, exitId, age, utxoOwner, token)
}

// AddInput is a paid mutator transaction binding the contract method 0x22f192af.
//
// Solidity: function addInput(uint256 exitId, uint256 age, address utxoOwner, address token) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) AddInput(exitId *big.Int, age *big.Int, utxoOwner common.Address, token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.AddInput(&_Withdrawmanager.TransactOpts, exitId, age, utxoOwner, token)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x9492b0b8.
//
// Solidity: function challengeExit(uint256 exitId, uint256 inputId, bytes challengeData, address adjudicatorPredicate) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) ChallengeExit(opts *bind.TransactOpts, exitId *big.Int, inputId *big.Int, challengeData []byte, adjudicatorPredicate common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "challengeExit", exitId, inputId, challengeData, adjudicatorPredicate)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x9492b0b8.
//
// Solidity: function challengeExit(uint256 exitId, uint256 inputId, bytes challengeData, address adjudicatorPredicate) returns()
func (_Withdrawmanager *WithdrawmanagerSession) ChallengeExit(exitId *big.Int, inputId *big.Int, challengeData []byte, adjudicatorPredicate common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.ChallengeExit(&_Withdrawmanager.TransactOpts, exitId, inputId, challengeData, adjudicatorPredicate)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x9492b0b8.
//
// Solidity: function challengeExit(uint256 exitId, uint256 inputId, bytes challengeData, address adjudicatorPredicate) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) ChallengeExit(exitId *big.Int, inputId *big.Int, challengeData []byte, adjudicatorPredicate common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.ChallengeExit(&_Withdrawmanager.TransactOpts, exitId, inputId, challengeData, adjudicatorPredicate)
}

// CreateExitQueue is a paid mutator transaction binding the contract method 0x9145e6df.
//
// Solidity: function createExitQueue(address token) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) CreateExitQueue(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "createExitQueue", token)
}

// CreateExitQueue is a paid mutator transaction binding the contract method 0x9145e6df.
//
// Solidity: function createExitQueue(address token) returns()
func (_Withdrawmanager *WithdrawmanagerSession) CreateExitQueue(token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.CreateExitQueue(&_Withdrawmanager.TransactOpts, token)
}

// CreateExitQueue is a paid mutator transaction binding the contract method 0x9145e6df.
//
// Solidity: function createExitQueue(address token) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) CreateExitQueue(token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.CreateExitQueue(&_Withdrawmanager.TransactOpts, token)
}

// ProcessExits is a paid mutator transaction binding the contract method 0x0f6795f2.
//
// Solidity: function processExits(address _token) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) ProcessExits(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "processExits", _token)
}

// ProcessExits is a paid mutator transaction binding the contract method 0x0f6795f2.
//
// Solidity: function processExits(address _token) returns()
func (_Withdrawmanager *WithdrawmanagerSession) ProcessExits(_token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.ProcessExits(&_Withdrawmanager.TransactOpts, _token)
}

// ProcessExits is a paid mutator transaction binding the contract method 0x0f6795f2.
//
// Solidity: function processExits(address _token) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) ProcessExits(_token common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.ProcessExits(&_Withdrawmanager.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Withdrawmanager *WithdrawmanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Withdrawmanager.Contract.RenounceOwnership(&_Withdrawmanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Withdrawmanager.Contract.RenounceOwnership(&_Withdrawmanager.TransactOpts)
}

// StartExitWithDepositedTokens is a paid mutator transaction binding the contract method 0x144a03b3.
//
// Solidity: function startExitWithDepositedTokens(uint256 depositId, address token, uint256 amountOrToken) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) StartExitWithDepositedTokens(opts *bind.TransactOpts, depositId *big.Int, token common.Address, amountOrToken *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "startExitWithDepositedTokens", depositId, token, amountOrToken)
}

// StartExitWithDepositedTokens is a paid mutator transaction binding the contract method 0x144a03b3.
//
// Solidity: function startExitWithDepositedTokens(uint256 depositId, address token, uint256 amountOrToken) returns()
func (_Withdrawmanager *WithdrawmanagerSession) StartExitWithDepositedTokens(depositId *big.Int, token common.Address, amountOrToken *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.StartExitWithDepositedTokens(&_Withdrawmanager.TransactOpts, depositId, token, amountOrToken)
}

// StartExitWithDepositedTokens is a paid mutator transaction binding the contract method 0x144a03b3.
//
// Solidity: function startExitWithDepositedTokens(uint256 depositId, address token, uint256 amountOrToken) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) StartExitWithDepositedTokens(depositId *big.Int, token common.Address, amountOrToken *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.StartExitWithDepositedTokens(&_Withdrawmanager.TransactOpts, depositId, token, amountOrToken)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Withdrawmanager *WithdrawmanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.TransferOwnership(&_Withdrawmanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.TransferOwnership(&_Withdrawmanager.TransactOpts, newOwner)
}

// UpdateExitPeriod is a paid mutator transaction binding the contract method 0x433c76bf.
//
// Solidity: function updateExitPeriod(uint256 halfExitPeriod) returns()
func (_Withdrawmanager *WithdrawmanagerTransactor) UpdateExitPeriod(opts *bind.TransactOpts, halfExitPeriod *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.contract.Transact(opts, "updateExitPeriod", halfExitPeriod)
}

// UpdateExitPeriod is a paid mutator transaction binding the contract method 0x433c76bf.
//
// Solidity: function updateExitPeriod(uint256 halfExitPeriod) returns()
func (_Withdrawmanager *WithdrawmanagerSession) UpdateExitPeriod(halfExitPeriod *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.UpdateExitPeriod(&_Withdrawmanager.TransactOpts, halfExitPeriod)
}

// UpdateExitPeriod is a paid mutator transaction binding the contract method 0x433c76bf.
//
// Solidity: function updateExitPeriod(uint256 halfExitPeriod) returns()
func (_Withdrawmanager *WithdrawmanagerTransactorSession) UpdateExitPeriod(halfExitPeriod *big.Int) (*types.Transaction, error) {
	return _Withdrawmanager.Contract.UpdateExitPeriod(&_Withdrawmanager.TransactOpts, halfExitPeriod)
}

// WithdrawmanagerExitCancelledIterator is returned from FilterExitCancelled and is used to iterate over the raw logs and unpacked data for ExitCancelled events raised by the Withdrawmanager contract.
type WithdrawmanagerExitCancelledIterator struct {
	Event *WithdrawmanagerExitCancelled // Event containing the contract specifics and raw log

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
func (it *WithdrawmanagerExitCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawmanagerExitCancelled)
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
		it.Event = new(WithdrawmanagerExitCancelled)
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
func (it *WithdrawmanagerExitCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawmanagerExitCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawmanagerExitCancelled represents a ExitCancelled event raised by the Withdrawmanager contract.
type WithdrawmanagerExitCancelled struct {
	ExitId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitCancelled is a free log retrieval operation binding the contract event 0x93a8052a01c184f88312af177ab8fae2e56a9973b6aa4bdc62dfcf744e09d041.
//
// Solidity: event ExitCancelled(uint256 indexed exitId)
func (_Withdrawmanager *WithdrawmanagerFilterer) FilterExitCancelled(opts *bind.FilterOpts, exitId []*big.Int) (*WithdrawmanagerExitCancelledIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _Withdrawmanager.contract.FilterLogs(opts, "ExitCancelled", exitIdRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerExitCancelledIterator{contract: _Withdrawmanager.contract, event: "ExitCancelled", logs: logs, sub: sub}, nil
}

// WatchExitCancelled is a free log subscription operation binding the contract event 0x93a8052a01c184f88312af177ab8fae2e56a9973b6aa4bdc62dfcf744e09d041.
//
// Solidity: event ExitCancelled(uint256 indexed exitId)
func (_Withdrawmanager *WithdrawmanagerFilterer) WatchExitCancelled(opts *bind.WatchOpts, sink chan<- *WithdrawmanagerExitCancelled, exitId []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}

	logs, sub, err := _Withdrawmanager.contract.WatchLogs(opts, "ExitCancelled", exitIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawmanagerExitCancelled)
				if err := _Withdrawmanager.contract.UnpackLog(event, "ExitCancelled", log); err != nil {
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

// ParseExitCancelled is a log parse operation binding the contract event 0x93a8052a01c184f88312af177ab8fae2e56a9973b6aa4bdc62dfcf744e09d041.
//
// Solidity: event ExitCancelled(uint256 indexed exitId)
func (_Withdrawmanager *WithdrawmanagerFilterer) ParseExitCancelled(log types.Log) (*WithdrawmanagerExitCancelled, error) {
	event := new(WithdrawmanagerExitCancelled)
	if err := _Withdrawmanager.contract.UnpackLog(event, "ExitCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WithdrawmanagerExitPeriodUpdateIterator is returned from FilterExitPeriodUpdate and is used to iterate over the raw logs and unpacked data for ExitPeriodUpdate events raised by the Withdrawmanager contract.
type WithdrawmanagerExitPeriodUpdateIterator struct {
	Event *WithdrawmanagerExitPeriodUpdate // Event containing the contract specifics and raw log

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
func (it *WithdrawmanagerExitPeriodUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawmanagerExitPeriodUpdate)
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
		it.Event = new(WithdrawmanagerExitPeriodUpdate)
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
func (it *WithdrawmanagerExitPeriodUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawmanagerExitPeriodUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawmanagerExitPeriodUpdate represents a ExitPeriodUpdate event raised by the Withdrawmanager contract.
type WithdrawmanagerExitPeriodUpdate struct {
	OldExitPeriod *big.Int
	NewExitPeriod *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExitPeriodUpdate is a free log retrieval operation binding the contract event 0x06b98f3947a8966918fef150b41170e78ba1d91dd2b1d2fd48a59c91ffbd66a1.
//
// Solidity: event ExitPeriodUpdate(uint256 indexed oldExitPeriod, uint256 indexed newExitPeriod)
func (_Withdrawmanager *WithdrawmanagerFilterer) FilterExitPeriodUpdate(opts *bind.FilterOpts, oldExitPeriod []*big.Int, newExitPeriod []*big.Int) (*WithdrawmanagerExitPeriodUpdateIterator, error) {

	var oldExitPeriodRule []interface{}
	for _, oldExitPeriodItem := range oldExitPeriod {
		oldExitPeriodRule = append(oldExitPeriodRule, oldExitPeriodItem)
	}
	var newExitPeriodRule []interface{}
	for _, newExitPeriodItem := range newExitPeriod {
		newExitPeriodRule = append(newExitPeriodRule, newExitPeriodItem)
	}

	logs, sub, err := _Withdrawmanager.contract.FilterLogs(opts, "ExitPeriodUpdate", oldExitPeriodRule, newExitPeriodRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerExitPeriodUpdateIterator{contract: _Withdrawmanager.contract, event: "ExitPeriodUpdate", logs: logs, sub: sub}, nil
}

// WatchExitPeriodUpdate is a free log subscription operation binding the contract event 0x06b98f3947a8966918fef150b41170e78ba1d91dd2b1d2fd48a59c91ffbd66a1.
//
// Solidity: event ExitPeriodUpdate(uint256 indexed oldExitPeriod, uint256 indexed newExitPeriod)
func (_Withdrawmanager *WithdrawmanagerFilterer) WatchExitPeriodUpdate(opts *bind.WatchOpts, sink chan<- *WithdrawmanagerExitPeriodUpdate, oldExitPeriod []*big.Int, newExitPeriod []*big.Int) (event.Subscription, error) {

	var oldExitPeriodRule []interface{}
	for _, oldExitPeriodItem := range oldExitPeriod {
		oldExitPeriodRule = append(oldExitPeriodRule, oldExitPeriodItem)
	}
	var newExitPeriodRule []interface{}
	for _, newExitPeriodItem := range newExitPeriod {
		newExitPeriodRule = append(newExitPeriodRule, newExitPeriodItem)
	}

	logs, sub, err := _Withdrawmanager.contract.WatchLogs(opts, "ExitPeriodUpdate", oldExitPeriodRule, newExitPeriodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawmanagerExitPeriodUpdate)
				if err := _Withdrawmanager.contract.UnpackLog(event, "ExitPeriodUpdate", log); err != nil {
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

// ParseExitPeriodUpdate is a log parse operation binding the contract event 0x06b98f3947a8966918fef150b41170e78ba1d91dd2b1d2fd48a59c91ffbd66a1.
//
// Solidity: event ExitPeriodUpdate(uint256 indexed oldExitPeriod, uint256 indexed newExitPeriod)
func (_Withdrawmanager *WithdrawmanagerFilterer) ParseExitPeriodUpdate(log types.Log) (*WithdrawmanagerExitPeriodUpdate, error) {
	event := new(WithdrawmanagerExitPeriodUpdate)
	if err := _Withdrawmanager.contract.UnpackLog(event, "ExitPeriodUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WithdrawmanagerExitStartedIterator is returned from FilterExitStarted and is used to iterate over the raw logs and unpacked data for ExitStarted events raised by the Withdrawmanager contract.
type WithdrawmanagerExitStartedIterator struct {
	Event *WithdrawmanagerExitStarted // Event containing the contract specifics and raw log

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
func (it *WithdrawmanagerExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawmanagerExitStarted)
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
		it.Event = new(WithdrawmanagerExitStarted)
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
func (it *WithdrawmanagerExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawmanagerExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawmanagerExitStarted represents a ExitStarted event raised by the Withdrawmanager contract.
type WithdrawmanagerExitStarted struct {
	Exitor        common.Address
	ExitId        *big.Int
	Token         common.Address
	Amount        *big.Int
	IsRegularExit bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExitStarted is a free log retrieval operation binding the contract event 0xaa5303fdad123ab5ecaefaf69137bf8632257839546d43a3b3dd148cc2879d6f.
//
// Solidity: event ExitStarted(address indexed exitor, uint256 indexed exitId, address indexed token, uint256 amount, bool isRegularExit)
func (_Withdrawmanager *WithdrawmanagerFilterer) FilterExitStarted(opts *bind.FilterOpts, exitor []common.Address, exitId []*big.Int, token []common.Address) (*WithdrawmanagerExitStartedIterator, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}
	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawmanager.contract.FilterLogs(opts, "ExitStarted", exitorRule, exitIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerExitStartedIterator{contract: _Withdrawmanager.contract, event: "ExitStarted", logs: logs, sub: sub}, nil
}

// WatchExitStarted is a free log subscription operation binding the contract event 0xaa5303fdad123ab5ecaefaf69137bf8632257839546d43a3b3dd148cc2879d6f.
//
// Solidity: event ExitStarted(address indexed exitor, uint256 indexed exitId, address indexed token, uint256 amount, bool isRegularExit)
func (_Withdrawmanager *WithdrawmanagerFilterer) WatchExitStarted(opts *bind.WatchOpts, sink chan<- *WithdrawmanagerExitStarted, exitor []common.Address, exitId []*big.Int, token []common.Address) (event.Subscription, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}
	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawmanager.contract.WatchLogs(opts, "ExitStarted", exitorRule, exitIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawmanagerExitStarted)
				if err := _Withdrawmanager.contract.UnpackLog(event, "ExitStarted", log); err != nil {
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

// ParseExitStarted is a log parse operation binding the contract event 0xaa5303fdad123ab5ecaefaf69137bf8632257839546d43a3b3dd148cc2879d6f.
//
// Solidity: event ExitStarted(address indexed exitor, uint256 indexed exitId, address indexed token, uint256 amount, bool isRegularExit)
func (_Withdrawmanager *WithdrawmanagerFilterer) ParseExitStarted(log types.Log) (*WithdrawmanagerExitStarted, error) {
	event := new(WithdrawmanagerExitStarted)
	if err := _Withdrawmanager.contract.UnpackLog(event, "ExitStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WithdrawmanagerExitUpdatedIterator is returned from FilterExitUpdated and is used to iterate over the raw logs and unpacked data for ExitUpdated events raised by the Withdrawmanager contract.
type WithdrawmanagerExitUpdatedIterator struct {
	Event *WithdrawmanagerExitUpdated // Event containing the contract specifics and raw log

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
func (it *WithdrawmanagerExitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawmanagerExitUpdated)
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
		it.Event = new(WithdrawmanagerExitUpdated)
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
func (it *WithdrawmanagerExitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawmanagerExitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawmanagerExitUpdated represents a ExitUpdated event raised by the Withdrawmanager contract.
type WithdrawmanagerExitUpdated struct {
	ExitId *big.Int
	Age    *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitUpdated is a free log retrieval operation binding the contract event 0x87d2daa6e85f166015ebbcf09f5ee4bc50f93677579339fe128e3561a6807cb6.
//
// Solidity: event ExitUpdated(uint256 indexed exitId, uint256 indexed age, address signer)
func (_Withdrawmanager *WithdrawmanagerFilterer) FilterExitUpdated(opts *bind.FilterOpts, exitId []*big.Int, age []*big.Int) (*WithdrawmanagerExitUpdatedIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}
	var ageRule []interface{}
	for _, ageItem := range age {
		ageRule = append(ageRule, ageItem)
	}

	logs, sub, err := _Withdrawmanager.contract.FilterLogs(opts, "ExitUpdated", exitIdRule, ageRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerExitUpdatedIterator{contract: _Withdrawmanager.contract, event: "ExitUpdated", logs: logs, sub: sub}, nil
}

// WatchExitUpdated is a free log subscription operation binding the contract event 0x87d2daa6e85f166015ebbcf09f5ee4bc50f93677579339fe128e3561a6807cb6.
//
// Solidity: event ExitUpdated(uint256 indexed exitId, uint256 indexed age, address signer)
func (_Withdrawmanager *WithdrawmanagerFilterer) WatchExitUpdated(opts *bind.WatchOpts, sink chan<- *WithdrawmanagerExitUpdated, exitId []*big.Int, age []*big.Int) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}
	var ageRule []interface{}
	for _, ageItem := range age {
		ageRule = append(ageRule, ageItem)
	}

	logs, sub, err := _Withdrawmanager.contract.WatchLogs(opts, "ExitUpdated", exitIdRule, ageRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawmanagerExitUpdated)
				if err := _Withdrawmanager.contract.UnpackLog(event, "ExitUpdated", log); err != nil {
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

// ParseExitUpdated is a log parse operation binding the contract event 0x87d2daa6e85f166015ebbcf09f5ee4bc50f93677579339fe128e3561a6807cb6.
//
// Solidity: event ExitUpdated(uint256 indexed exitId, uint256 indexed age, address signer)
func (_Withdrawmanager *WithdrawmanagerFilterer) ParseExitUpdated(log types.Log) (*WithdrawmanagerExitUpdated, error) {
	event := new(WithdrawmanagerExitUpdated)
	if err := _Withdrawmanager.contract.UnpackLog(event, "ExitUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WithdrawmanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Withdrawmanager contract.
type WithdrawmanagerOwnershipTransferredIterator struct {
	Event *WithdrawmanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WithdrawmanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawmanagerOwnershipTransferred)
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
		it.Event = new(WithdrawmanagerOwnershipTransferred)
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
func (it *WithdrawmanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawmanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawmanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Withdrawmanager contract.
type WithdrawmanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Withdrawmanager *WithdrawmanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WithdrawmanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Withdrawmanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerOwnershipTransferredIterator{contract: _Withdrawmanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Withdrawmanager *WithdrawmanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WithdrawmanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Withdrawmanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawmanagerOwnershipTransferred)
				if err := _Withdrawmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Withdrawmanager *WithdrawmanagerFilterer) ParseOwnershipTransferred(log types.Log) (*WithdrawmanagerOwnershipTransferred, error) {
	event := new(WithdrawmanagerOwnershipTransferred)
	if err := _Withdrawmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WithdrawmanagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Withdrawmanager contract.
type WithdrawmanagerWithdrawIterator struct {
	Event *WithdrawmanagerWithdraw // Event containing the contract specifics and raw log

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
func (it *WithdrawmanagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawmanagerWithdraw)
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
		it.Event = new(WithdrawmanagerWithdraw)
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
func (it *WithdrawmanagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawmanagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawmanagerWithdraw represents a Withdraw event raised by the Withdrawmanager contract.
type WithdrawmanagerWithdraw struct {
	ExitId *big.Int
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfeb2000dca3e617cd6f3a8bbb63014bb54a124aac6ccbf73ee7229b4cd01f120.
//
// Solidity: event Withdraw(uint256 indexed exitId, address indexed user, address indexed token, uint256 amount)
func (_Withdrawmanager *WithdrawmanagerFilterer) FilterWithdraw(opts *bind.FilterOpts, exitId []*big.Int, user []common.Address, token []common.Address) (*WithdrawmanagerWithdrawIterator, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawmanager.contract.FilterLogs(opts, "Withdraw", exitIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawmanagerWithdrawIterator{contract: _Withdrawmanager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfeb2000dca3e617cd6f3a8bbb63014bb54a124aac6ccbf73ee7229b4cd01f120.
//
// Solidity: event Withdraw(uint256 indexed exitId, address indexed user, address indexed token, uint256 amount)
func (_Withdrawmanager *WithdrawmanagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *WithdrawmanagerWithdraw, exitId []*big.Int, user []common.Address, token []common.Address) (event.Subscription, error) {

	var exitIdRule []interface{}
	for _, exitIdItem := range exitId {
		exitIdRule = append(exitIdRule, exitIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawmanager.contract.WatchLogs(opts, "Withdraw", exitIdRule, userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawmanagerWithdraw)
				if err := _Withdrawmanager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfeb2000dca3e617cd6f3a8bbb63014bb54a124aac6ccbf73ee7229b4cd01f120.
//
// Solidity: event Withdraw(uint256 indexed exitId, address indexed user, address indexed token, uint256 amount)
func (_Withdrawmanager *WithdrawmanagerFilterer) ParseWithdraw(log types.Log) (*WithdrawmanagerWithdraw, error) {
	event := new(WithdrawmanagerWithdraw)
	if err := _Withdrawmanager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
