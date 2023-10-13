// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ledgerwatch/erigon"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/accounts/abi"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = libcommon.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AlphabetVMABI is the input ABI used to generate the binding from.
const AlphabetVMABI = "[{\"inputs\":[{\"internalType\":\"Claim\",\"name\":\"_absolutePrestate\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIPreimageOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_stateData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"step\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"postState_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AlphabetVMBin is the compiled bytecode used for deploying new contracts.
var AlphabetVMBin = "0x60a060405234801561001057600080fd5b50604051610a5d380380610a5d83398101604081905261002f91610090565b608081905260405161004090610083565b604051809103906000f08015801561005c573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055506100a9565b6106458061041883390190565b6000602082840312156100a257600080fd5b5051919050565b6080516103556100c3600039600060af01526103556000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80637dc0d1d01461003b578063f8e0cb9614610085575b600080fd5b60005461005b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610098610093366004610212565b6100a6565b60405190815260200161007c565b600080600060087f0000000000000000000000000000000000000000000000000000000000000000901b600888886040516100e292919061027e565b6040518091039020901b0361010857600091506101018688018861028e565b9050610127565b610114868801886102a7565b909250905081610123816102f8565b9250505b81610133826001610330565b604080516020810193909352820152606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f010000000000000000000000000000000000000000000000000000000000000017979650505050505050565b60008083601f8401126101db57600080fd5b50813567ffffffffffffffff8111156101f357600080fd5b60208301915083602082850101111561020b57600080fd5b9250929050565b6000806000806040858703121561022857600080fd5b843567ffffffffffffffff8082111561024057600080fd5b61024c888389016101c9565b9096509450602087013591508082111561026557600080fd5b50610272878288016101c9565b95989497509550505050565b8183823760009101908152919050565b6000602082840312156102a057600080fd5b5035919050565b600080604083850312156102ba57600080fd5b50508035926020909101359150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610329576103296102c9565b5060010190565b60008219821115610343576103436102c9565b50019056fea164736f6c634300080f000a608060405234801561001057600080fd5b50610625806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063e03110e111610050578063e03110e114610106578063e15926111461012e578063fef2b4ed1461014357600080fd5b806361238bde146100775780638542cf50146100b55780639a1f5e7f146100f3575b600080fd5b6100a26100853660046104d1565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6100e36100c33660046104d1565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100ac565b6100a26101013660046104f3565b610163565b6101196101143660046104d1565b610236565b604080519283526020830191909152016100ac565b61014161013c366004610525565b610327565b005b6100a26101513660046105a1565b60006020819052908152604090205481565b600061016e85610430565b905061017b8360086105e9565b8211806101885750602083115b156101bf576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845287528083209483529386528382205581815293849052922055919050565b6000828152600260209081526040808320848452909152812054819060ff166102bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b50600083815260208181526040909120546102db8160086105e9565b6102e68560206105e9565b1061030457836102f78260086105e9565b6103019190610601565b91505b506000938452600160209081526040808620948652939052919092205492909150565b604435600080600883018611156103465763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8216176104cb81600090815233602052604090207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b92915050565b600080604083850312156104e457600080fd5b50508035926020909101359150565b6000806000806080858703121561050957600080fd5b5050823594602084013594506040840135936060013592509050565b60008060006040848603121561053a57600080fd5b83359250602084013567ffffffffffffffff8082111561055957600080fd5b818601915086601f83011261056d57600080fd5b81358181111561057c57600080fd5b87602082850101111561058e57600080fd5b6020830194508093505050509250925092565b6000602082840312156105b357600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156105fc576105fc6105ba565b500190565b600082821015610613576106136105ba565b50039056fea164736f6c634300080f000a"

// DeployAlphabetVM deploys a new Ethereum contract, binding an instance of AlphabetVM to it.
func DeployAlphabetVM(auth *bind.TransactOpts, backend bind.ContractBackend, _absolutePrestate [32]byte) (libcommon.Address, types.Transaction, *AlphabetVM, error) {
	parsed, err := abi.JSON(strings.NewReader(AlphabetVMABI))
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, libcommon.FromHex(AlphabetVMBin), backend, _absolutePrestate)
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}
	return address, tx, &AlphabetVM{AlphabetVMCaller: AlphabetVMCaller{contract: contract}, AlphabetVMTransactor: AlphabetVMTransactor{contract: contract}, AlphabetVMFilterer: AlphabetVMFilterer{contract: contract}}, nil
}

// AlphabetVM is an auto generated Go binding around an Ethereum contract.
type AlphabetVM struct {
	AlphabetVMCaller     // Read-only binding to the contract
	AlphabetVMTransactor // Write-only binding to the contract
	AlphabetVMFilterer   // Log filterer for contract events
}

// AlphabetVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlphabetVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphabetVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlphabetVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphabetVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlphabetVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphabetVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlphabetVMSession struct {
	Contract     *AlphabetVM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlphabetVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlphabetVMCallerSession struct {
	Contract *AlphabetVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AlphabetVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlphabetVMTransactorSession struct {
	Contract     *AlphabetVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AlphabetVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlphabetVMRaw struct {
	Contract *AlphabetVM // Generic contract binding to access the raw methods on
}

// AlphabetVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlphabetVMCallerRaw struct {
	Contract *AlphabetVMCaller // Generic read-only contract binding to access the raw methods on
}

// AlphabetVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlphabetVMTransactorRaw struct {
	Contract *AlphabetVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlphabetVM creates a new instance of AlphabetVM, bound to a specific deployed contract.
func NewAlphabetVM(address libcommon.Address, backend bind.ContractBackend) (*AlphabetVM, error) {
	contract, err := bindAlphabetVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AlphabetVM{AlphabetVMCaller: AlphabetVMCaller{contract: contract}, AlphabetVMTransactor: AlphabetVMTransactor{contract: contract}, AlphabetVMFilterer: AlphabetVMFilterer{contract: contract}}, nil
}

// NewAlphabetVMCaller creates a new read-only instance of AlphabetVM, bound to a specific deployed contract.
func NewAlphabetVMCaller(address libcommon.Address, caller bind.ContractCaller) (*AlphabetVMCaller, error) {
	contract, err := bindAlphabetVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlphabetVMCaller{contract: contract}, nil
}

// NewAlphabetVMTransactor creates a new write-only instance of AlphabetVM, bound to a specific deployed contract.
func NewAlphabetVMTransactor(address libcommon.Address, transactor bind.ContractTransactor) (*AlphabetVMTransactor, error) {
	contract, err := bindAlphabetVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlphabetVMTransactor{contract: contract}, nil
}

// NewAlphabetVMFilterer creates a new log filterer instance of AlphabetVM, bound to a specific deployed contract.
func NewAlphabetVMFilterer(address libcommon.Address, filterer bind.ContractFilterer) (*AlphabetVMFilterer, error) {
	contract, err := bindAlphabetVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlphabetVMFilterer{contract: contract}, nil
}

// bindAlphabetVM binds a generic wrapper to an already deployed contract.
func bindAlphabetVM(address libcommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AlphabetVMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlphabetVM *AlphabetVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlphabetVM.Contract.AlphabetVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlphabetVM *AlphabetVMRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _AlphabetVM.Contract.AlphabetVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlphabetVM *AlphabetVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _AlphabetVM.Contract.AlphabetVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlphabetVM *AlphabetVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlphabetVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlphabetVM *AlphabetVMTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _AlphabetVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlphabetVM *AlphabetVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _AlphabetVM.Contract.contract.Transact(opts, method, params...)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_AlphabetVM *AlphabetVMCaller) Oracle(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _AlphabetVM.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_AlphabetVM *AlphabetVMSession) Oracle() (libcommon.Address, error) {
	return _AlphabetVM.Contract.Oracle(&_AlphabetVM.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_AlphabetVM *AlphabetVMCallerSession) Oracle() (libcommon.Address, error) {
	return _AlphabetVM.Contract.Oracle(&_AlphabetVM.CallOpts)
}

// Step is a free data retrieval call binding the contract method 0xf8e0cb96.
//
// Solidity: function step(bytes _stateData, bytes ) view returns(bytes32 postState_)
func (_AlphabetVM *AlphabetVMCaller) Step(opts *bind.CallOpts, _stateData []byte, arg1 []byte) ([32]byte, error) {
	var out []interface{}
	err := _AlphabetVM.contract.Call(opts, &out, "step", _stateData, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Step is a free data retrieval call binding the contract method 0xf8e0cb96.
//
// Solidity: function step(bytes _stateData, bytes ) view returns(bytes32 postState_)
func (_AlphabetVM *AlphabetVMSession) Step(_stateData []byte, arg1 []byte) ([32]byte, error) {
	return _AlphabetVM.Contract.Step(&_AlphabetVM.CallOpts, _stateData, arg1)
}

// Step is a free data retrieval call binding the contract method 0xf8e0cb96.
//
// Solidity: function step(bytes _stateData, bytes ) view returns(bytes32 postState_)
func (_AlphabetVM *AlphabetVMCallerSession) Step(_stateData []byte, arg1 []byte) ([32]byte, error) {
	return _AlphabetVM.Contract.Step(&_AlphabetVM.CallOpts, _stateData, arg1)
}