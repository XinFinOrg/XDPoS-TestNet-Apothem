// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a3f63b465e1cf25f306b1eb1efefc8dac3c38993a7340f69d8b470c3bf599ff30029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// XDCRandomizeABI is the input ABI used to generate the binding from.
const XDCRandomizeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getSecret\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_secret\",\"type\":\"bytes32[]\"}],\"name\":\"setSecret\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"randomNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getOpening\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_opening\",\"type\":\"bytes32\"}],\"name\":\"setOpening\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_randomNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// XDCRandomizeBin is the compiled bytecode used for deploying new contracts.
const XDCRandomizeBin = `0x608060405234801561001057600080fd5b5060405160208061035b8339810160405251600055610327806100346000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663284180fc811461007157806334d38600146100ef578063ccbac9f514610146578063d442d6cc1461016d578063e11f5ba21461019b575b600080fd5b34801561007d57600080fd5b5061009f73ffffffffffffffffffffffffffffffffffffffff600435166101b3565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100db5781810151838201526020016100c3565b505050509050019250505060405180910390f35b3480156100fb57600080fd5b50604080516020600480358082013583810280860185019096528085526101449536959394602494938501929182918501908490808284375094975061022d9650505050505050565b005b34801561015257600080fd5b5061015b610251565b60408051918252519081900360200190f35b34801561017957600080fd5b5061015b73ffffffffffffffffffffffffffffffffffffffff60043516610257565b3480156101a757600080fd5b5061014460043561027f565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001602090815260409182902080548351818402810184019094528084526060939283018282801561022157602002820191906000526020600020905b8154815260019091019060200180831161020c575b50505050509050919050565b336000908152600160209081526040909120825161024d92840190610291565b5050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205490565b33600090815260026020526040902055565b8280548282559060005260206000209081019282156102ce579160200282015b828111156102ce57825182556020909201916001909101906102b1565b506102da9291506102de565b5090565b6102f891905b808211156102da57600081556001016102e4565b905600a165627a7a7230582027c1c755e5d546967a2057a4338eb36cf7ec5972758d7f7a8427abdfc890bd590029`

// DeployXDCRandomize deploys a new Ethereum contract, binding an instance of XDCRandomize to it.
func DeployXDCRandomize(auth *bind.TransactOpts, backend bind.ContractBackend, _randomNumber *big.Int) (common.Address, *types.Transaction, *XDCRandomize, error) {
	parsed, err := abi.JSON(strings.NewReader(XDCRandomizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XDCRandomizeBin), backend, _randomNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XDCRandomize{XDCRandomizeCaller: XDCRandomizeCaller{contract: contract}, XDCRandomizeTransactor: XDCRandomizeTransactor{contract: contract}, XDCRandomizeFilterer: XDCRandomizeFilterer{contract: contract}}, nil
}

// XDCRandomize is an auto generated Go binding around an Ethereum contract.
type XDCRandomize struct {
	XDCRandomizeCaller     // Read-only binding to the contract
	XDCRandomizeTransactor // Write-only binding to the contract
	XDCRandomizeFilterer   // Log filterer for contract events
}

// XDCRandomizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type XDCRandomizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDCRandomizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XDCRandomizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDCRandomizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XDCRandomizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XDCRandomizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XDCRandomizeSession struct {
	Contract     *XDCRandomize    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XDCRandomizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XDCRandomizeCallerSession struct {
	Contract *XDCRandomizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// XDCRandomizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XDCRandomizeTransactorSession struct {
	Contract     *XDCRandomizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// XDCRandomizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type XDCRandomizeRaw struct {
	Contract *XDCRandomize // Generic contract binding to access the raw methods on
}

// XDCRandomizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XDCRandomizeCallerRaw struct {
	Contract *XDCRandomizeCaller // Generic read-only contract binding to access the raw methods on
}

// XDCRandomizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XDCRandomizeTransactorRaw struct {
	Contract *XDCRandomizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXDCRandomize creates a new instance of XDCRandomize, bound to a specific deployed contract.
func NewXDCRandomize(address common.Address, backend bind.ContractBackend) (*XDCRandomize, error) {
	contract, err := bindXDCRandomize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XDCRandomize{XDCRandomizeCaller: XDCRandomizeCaller{contract: contract}, XDCRandomizeTransactor: XDCRandomizeTransactor{contract: contract}, XDCRandomizeFilterer: XDCRandomizeFilterer{contract: contract}}, nil
}

// NewXDCRandomizeCaller creates a new read-only instance of XDCRandomize, bound to a specific deployed contract.
func NewXDCRandomizeCaller(address common.Address, caller bind.ContractCaller) (*XDCRandomizeCaller, error) {
	contract, err := bindXDCRandomize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XDCRandomizeCaller{contract: contract}, nil
}

// NewXDCRandomizeTransactor creates a new write-only instance of XDCRandomize, bound to a specific deployed contract.
func NewXDCRandomizeTransactor(address common.Address, transactor bind.ContractTransactor) (*XDCRandomizeTransactor, error) {
	contract, err := bindXDCRandomize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XDCRandomizeTransactor{contract: contract}, nil
}

// NewXDCRandomizeFilterer creates a new log filterer instance of XDCRandomize, bound to a specific deployed contract.
func NewXDCRandomizeFilterer(address common.Address, filterer bind.ContractFilterer) (*XDCRandomizeFilterer, error) {
	contract, err := bindXDCRandomize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XDCRandomizeFilterer{contract: contract}, nil
}

// bindXDCRandomize binds a generic wrapper to an already deployed contract.
func bindXDCRandomize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XDCRandomizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XDCRandomize *XDCRandomizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XDCRandomize.Contract.XDCRandomizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XDCRandomize *XDCRandomizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XDCRandomize.Contract.XDCRandomizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XDCRandomize *XDCRandomizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XDCRandomize.Contract.XDCRandomizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XDCRandomize *XDCRandomizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XDCRandomize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XDCRandomize *XDCRandomizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XDCRandomize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XDCRandomize *XDCRandomizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XDCRandomize.Contract.contract.Transact(opts, method, params...)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_XDCRandomize *XDCRandomizeCaller) GetOpening(opts *bind.CallOpts, _validator common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _XDCRandomize.contract.Call(opts, out, "getOpening", _validator)
	return *ret0, err
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_XDCRandomize *XDCRandomizeSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _XDCRandomize.Contract.GetOpening(&_XDCRandomize.CallOpts, _validator)
}

// GetOpening is a free data retrieval call binding the contract method 0xd442d6cc.
//
// Solidity: function getOpening(_validator address) constant returns(bytes32)
func (_XDCRandomize *XDCRandomizeCallerSession) GetOpening(_validator common.Address) ([32]byte, error) {
	return _XDCRandomize.Contract.GetOpening(&_XDCRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_XDCRandomize *XDCRandomizeCaller) GetSecret(opts *bind.CallOpts, _validator common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _XDCRandomize.contract.Call(opts, out, "getSecret", _validator)
	return *ret0, err
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_XDCRandomize *XDCRandomizeSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _XDCRandomize.Contract.GetSecret(&_XDCRandomize.CallOpts, _validator)
}

// GetSecret is a free data retrieval call binding the contract method 0x284180fc.
//
// Solidity: function getSecret(_validator address) constant returns(bytes32[])
func (_XDCRandomize *XDCRandomizeCallerSession) GetSecret(_validator common.Address) ([][32]byte, error) {
	return _XDCRandomize.Contract.GetSecret(&_XDCRandomize.CallOpts, _validator)
}

// RandomNumber is a free data retrieval call binding the contract method 0xccbac9f5.
//
// Solidity: function randomNumber() constant returns(uint256)
func (_XDCRandomize *XDCRandomizeCaller) RandomNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _XDCRandomize.contract.Call(opts, out, "randomNumber")
	return *ret0, err
}

// RandomNumber is a free data retrieval call binding the contract method 0xccbac9f5.
//
// Solidity: function randomNumber() constant returns(uint256)
func (_XDCRandomize *XDCRandomizeSession) RandomNumber() (*big.Int, error) {
	return _XDCRandomize.Contract.RandomNumber(&_XDCRandomize.CallOpts)
}

// RandomNumber is a free data retrieval call binding the contract method 0xccbac9f5.
//
// Solidity: function randomNumber() constant returns(uint256)
func (_XDCRandomize *XDCRandomizeCallerSession) RandomNumber() (*big.Int, error) {
	return _XDCRandomize.Contract.RandomNumber(&_XDCRandomize.CallOpts)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_XDCRandomize *XDCRandomizeTransactor) SetOpening(opts *bind.TransactOpts, _opening [32]byte) (*types.Transaction, error) {
	return _XDCRandomize.contract.Transact(opts, "setOpening", _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_XDCRandomize *XDCRandomizeSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _XDCRandomize.Contract.SetOpening(&_XDCRandomize.TransactOpts, _opening)
}

// SetOpening is a paid mutator transaction binding the contract method 0xe11f5ba2.
//
// Solidity: function setOpening(_opening bytes32) returns()
func (_XDCRandomize *XDCRandomizeTransactorSession) SetOpening(_opening [32]byte) (*types.Transaction, error) {
	return _XDCRandomize.Contract.SetOpening(&_XDCRandomize.TransactOpts, _opening)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_XDCRandomize *XDCRandomizeTransactor) SetSecret(opts *bind.TransactOpts, _secret [][32]byte) (*types.Transaction, error) {
	return _XDCRandomize.contract.Transact(opts, "setSecret", _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_XDCRandomize *XDCRandomizeSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _XDCRandomize.Contract.SetSecret(&_XDCRandomize.TransactOpts, _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x34d38600.
//
// Solidity: function setSecret(_secret bytes32[]) returns()
func (_XDCRandomize *XDCRandomizeTransactorSession) SetSecret(_secret [][32]byte) (*types.Transaction, error) {
	return _XDCRandomize.Contract.SetSecret(&_XDCRandomize.TransactOpts, _secret)
}