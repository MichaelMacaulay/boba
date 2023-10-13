// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"fmt"
	"math/big"
	"reflect"
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

// L2ERC721BridgeABI is the input ABI used to generate the binding from.
const L2ERC721BridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_otherBridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC721BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC721BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC721To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractCrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// L2ERC721BridgeBin is the compiled bytecode used for deploying new contracts.
var L2ERC721BridgeBin = "0x60a06040523480156200001157600080fd5b506040516200187f3803806200187f83398101604081905262000034916200024b565b806001600160a01b038116620000a95760405162461bcd60e51b815260206004820152602f60248201527f4552433732314272696467653a206f74686572206272696467652063616e6e6f60448201526e74206265206164647265737328302960881b60648201526084015b60405180910390fd5b6001600160a01b0316608052620000c16000620000c8565b506200027d565b600054600290610100900460ff16158015620000eb575060005460ff8083169116105b620001505760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401620000a0565b6000805461ffff191660ff8316176101001790556200016f82620001b4565b6000805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b600054610100900460ff16620002215760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401620000a0565b600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b6000602082840312156200025e57600080fd5b81516001600160a01b03811681146200027657600080fd5b9392505050565b6080516115d1620002ae6000396000818161016e015281816101dc015281816102d60152610ea101526115d16000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80637f46ddb211610076578063aa5574521161005b578063aa557452146101b4578063c4d66de8146101c7578063c89701a2146101da57600080fd5b80637f46ddb214610169578063927ede2d1461019057600080fd5b80633687011a146100a85780633cb747bf146100bd57806354fd4d501461010d578063761f449314610156575b600080fd5b6100bb6100b636600461124f565b610200565b005b6000546100e39062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101496040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b604051610104919061133d565b6100bb610164366004611350565b6102ac565b6100e37f000000000000000000000000000000000000000000000000000000000000000081565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff166100e3565b6100bb6101c23660046113e8565b6107fd565b6100bb6101d536600461145f565b6108b9565b7f00000000000000000000000000000000000000000000000000000000000000006100e3565b333b15610294576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102a48686333388888888610a03565b505050505050565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331480156103b457507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610378573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039c919061147c565b73ffffffffffffffffffffffffffffffffffffffff16145b610440576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f746865722062726964676500606482015260840161028b565b3073ffffffffffffffffffffffffffffffffffffffff8816036104e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c6600000000000000000000000000000000000000000000606482015260840161028b565b61050f877f74259ebf00000000000000000000000000000000000000000000000000000000610f8d565b61059b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e746560448201527f7266616365206973206e6f7420636f6d706c69616e7400000000000000000000606482015260840161028b565b8673ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060a919061147c565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16146106ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604b60248201527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433732312060648201527f6c6f63616c20746f6b656e000000000000000000000000000000000000000000608482015260a40161028b565b6040517fa144819400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905288169063a144819490604401600060405180830381600087803b15801561075a57600080fd5b505af115801561076e573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516107ec94939291906114e2565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff85166108a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f74206265206164647265737328302900000000000000000000000000000000606482015260840161028b565b6108b08787338888888888610a03565b50505050505050565b600054600290610100900460ff161580156108db575060005460ff8083169116105b610967576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161028b565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8316176101001790556109a182610fb0565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff8716610aa6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f742062652061646472657373283029000000000000000000000000000000606482015260840161028b565b6040517f6352211e0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff891690636352211e90602401602060405180830381865afa158015610b11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b35919061147c565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610bef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324552433732314272696467653a205769746864726177616c206973206e6f60448201527f74206265696e6720696e69746961746564206279204e4654206f776e65720000606482015260840161028b565b60008873ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c60919061147c565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610d1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6560448201527f73206e6f74206d6174636820676976656e2076616c7565000000000000000000606482015260840161028b565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018790528a1690639dc29fac90604401600060405180830381600087803b158015610d8d57600080fd5b505af1158015610da1573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a8989604051602401610dcc9796959493929190611522565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925260005491517f3dbb202b00000000000000000000000000000000000000000000000000000000815290925073ffffffffffffffffffffffffffffffffffffffff620100009092049190911690633dbb202b90610ecd907f00000000000000000000000000000000000000000000000000000000000000009085908a9060040161157f565b600060405180830381600087803b158015610ee757600080fd5b505af1158015610efb573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a8989604051610f7994939291906114e2565b60405180910390a450505050505050505050565b6000610f9883611094565b8015610fa95750610fa983836110f9565b9392505050565b600054610100900460ff16611047576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161028b565b6000805473ffffffffffffffffffffffffffffffffffffffff90921662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909216919091179055565b60006110c0827f01ffc9a7000000000000000000000000000000000000000000000000000000006110f9565b80156110f357506110f1827fffffffff000000000000000000000000000000000000000000000000000000006110f9565b155b92915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156111b1575060208210155b80156111bd5750600081115b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff811681146111ea57600080fd5b50565b803563ffffffff8116811461120157600080fd5b919050565b60008083601f84011261121857600080fd5b50813567ffffffffffffffff81111561123057600080fd5b60208301915083602082850101111561124857600080fd5b9250929050565b60008060008060008060a0878903121561126857600080fd5b8635611273816111c8565b95506020870135611283816111c8565b945060408701359350611298606088016111ed565b9250608087013567ffffffffffffffff8111156112b457600080fd5b6112c089828a01611206565b979a9699509497509295939492505050565b6000815180845260005b818110156112f8576020818501810151868301820152016112dc565b8181111561130a576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610fa960208301846112d2565b600080600080600080600060c0888a03121561136b57600080fd5b8735611376816111c8565b96506020880135611386816111c8565b95506040880135611396816111c8565b945060608801356113a6816111c8565b93506080880135925060a088013567ffffffffffffffff8111156113c957600080fd5b6113d58a828b01611206565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561140357600080fd5b873561140e816111c8565b9650602088013561141e816111c8565b9550604088013561142e816111c8565b945060608801359350611443608089016111ed565b925060a088013567ffffffffffffffff8111156113c957600080fd5b60006020828403121561147157600080fd5b8135610fa9816111c8565b60006020828403121561148e57600080fd5b8151610fa9816111c8565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152606060408201526000611518606083018486611499565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261157260c083018486611499565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006115ae60608301856112d2565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"

// DeployL2ERC721Bridge deploys a new Ethereum contract, binding an instance of L2ERC721Bridge to it.
func DeployL2ERC721Bridge(auth *bind.TransactOpts, backend bind.ContractBackend, _otherBridge libcommon.Address) (libcommon.Address, types.Transaction, *L2ERC721Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ERC721BridgeABI))
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, libcommon.FromHex(L2ERC721BridgeBin), backend, _otherBridge)
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}
	return address, tx, &L2ERC721Bridge{L2ERC721BridgeCaller: L2ERC721BridgeCaller{contract: contract}, L2ERC721BridgeTransactor: L2ERC721BridgeTransactor{contract: contract}, L2ERC721BridgeFilterer: L2ERC721BridgeFilterer{contract: contract}}, nil
}

// L2ERC721Bridge is an auto generated Go binding around an Ethereum contract.
type L2ERC721Bridge struct {
	L2ERC721BridgeCaller     // Read-only binding to the contract
	L2ERC721BridgeTransactor // Write-only binding to the contract
	L2ERC721BridgeFilterer   // Log filterer for contract events
}

// L2ERC721BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ERC721BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ERC721BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ERC721BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC721BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ERC721BridgeSession struct {
	Contract     *L2ERC721Bridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2ERC721BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ERC721BridgeCallerSession struct {
	Contract *L2ERC721BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L2ERC721BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ERC721BridgeTransactorSession struct {
	Contract     *L2ERC721BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L2ERC721BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ERC721BridgeRaw struct {
	Contract *L2ERC721Bridge // Generic contract binding to access the raw methods on
}

// L2ERC721BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ERC721BridgeCallerRaw struct {
	Contract *L2ERC721BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L2ERC721BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ERC721BridgeTransactorRaw struct {
	Contract *L2ERC721BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ERC721Bridge creates a new instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721Bridge(address libcommon.Address, backend bind.ContractBackend) (*L2ERC721Bridge, error) {
	contract, err := bindL2ERC721Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ERC721Bridge{L2ERC721BridgeCaller: L2ERC721BridgeCaller{contract: contract}, L2ERC721BridgeTransactor: L2ERC721BridgeTransactor{contract: contract}, L2ERC721BridgeFilterer: L2ERC721BridgeFilterer{contract: contract}}, nil
}

// NewL2ERC721BridgeCaller creates a new read-only instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721BridgeCaller(address libcommon.Address, caller bind.ContractCaller) (*L2ERC721BridgeCaller, error) {
	contract, err := bindL2ERC721Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeCaller{contract: contract}, nil
}

// NewL2ERC721BridgeTransactor creates a new write-only instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721BridgeTransactor(address libcommon.Address, transactor bind.ContractTransactor) (*L2ERC721BridgeTransactor, error) {
	contract, err := bindL2ERC721Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeTransactor{contract: contract}, nil
}

// NewL2ERC721BridgeFilterer creates a new log filterer instance of L2ERC721Bridge, bound to a specific deployed contract.
func NewL2ERC721BridgeFilterer(address libcommon.Address, filterer bind.ContractFilterer) (*L2ERC721BridgeFilterer, error) {
	contract, err := bindL2ERC721Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeFilterer{contract: contract}, nil
}

// bindL2ERC721Bridge binds a generic wrapper to an already deployed contract.
func bindL2ERC721Bridge(address libcommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ERC721BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC721Bridge *L2ERC721BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC721Bridge.Contract.L2ERC721BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC721Bridge *L2ERC721BridgeRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.L2ERC721BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC721Bridge *L2ERC721BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.L2ERC721BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ERC721Bridge *L2ERC721BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ERC721Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ERC721Bridge *L2ERC721BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ERC721Bridge *L2ERC721BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) MESSENGER(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) MESSENGER() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.MESSENGER(&_L2ERC721Bridge.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) MESSENGER() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.MESSENGER(&_L2ERC721Bridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) OTHERBRIDGE(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) OTHERBRIDGE() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.OTHERBRIDGE(&_L2ERC721Bridge.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) OTHERBRIDGE() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.OTHERBRIDGE(&_L2ERC721Bridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) Messenger(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) Messenger() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.Messenger(&_L2ERC721Bridge.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) Messenger() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.Messenger(&_L2ERC721Bridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) OtherBridge(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "otherBridge")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeSession) OtherBridge() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.OtherBridge(&_L2ERC721Bridge.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) OtherBridge() (libcommon.Address, error) {
	return _L2ERC721Bridge.Contract.OtherBridge(&_L2ERC721Bridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ERC721Bridge *L2ERC721BridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ERC721Bridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ERC721Bridge *L2ERC721BridgeSession) Version() (string, error) {
	return _L2ERC721Bridge.Contract.Version(&_L2ERC721Bridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ERC721Bridge *L2ERC721BridgeCallerSession) Version() (string, error) {
	return _L2ERC721Bridge.Contract.Version(&_L2ERC721Bridge.CallOpts)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) BridgeERC721(opts *bind.TransactOpts, _localToken libcommon.Address, _remoteToken libcommon.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "bridgeERC721", _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) BridgeERC721(_localToken libcommon.Address, _remoteToken libcommon.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x3687011a.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) BridgeERC721(_localToken libcommon.Address, _remoteToken libcommon.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) BridgeERC721To(opts *bind.TransactOpts, _localToken libcommon.Address, _remoteToken libcommon.Address, _to libcommon.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "bridgeERC721To", _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) BridgeERC721To(_localToken libcommon.Address, _remoteToken libcommon.Address, _to libcommon.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721To(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// BridgeERC721To is a paid mutator transaction binding the contract method 0xaa557452.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) BridgeERC721To(_localToken libcommon.Address, _remoteToken libcommon.Address, _to libcommon.Address, _tokenId *big.Int, _minGasLimit uint32, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.BridgeERC721To(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _to, _tokenId, _minGasLimit, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) FinalizeBridgeERC721(opts *bind.TransactOpts, _localToken libcommon.Address, _remoteToken libcommon.Address, _from libcommon.Address, _to libcommon.Address, _tokenId *big.Int, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "finalizeBridgeERC721", _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) FinalizeBridgeERC721(_localToken libcommon.Address, _remoteToken libcommon.Address, _from libcommon.Address, _to libcommon.Address, _tokenId *big.Int, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.FinalizeBridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// FinalizeBridgeERC721 is a paid mutator transaction binding the contract method 0x761f4493.
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) FinalizeBridgeERC721(_localToken libcommon.Address, _remoteToken libcommon.Address, _from libcommon.Address, _to libcommon.Address, _tokenId *big.Int, _extraData []byte) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.FinalizeBridgeERC721(&_L2ERC721Bridge.TransactOpts, _localToken, _remoteToken, _from, _to, _tokenId, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _messenger) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactor) Initialize(opts *bind.TransactOpts, _messenger libcommon.Address) (types.Transaction, error) {
	return _L2ERC721Bridge.contract.Transact(opts, "initialize", _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _messenger) returns()
func (_L2ERC721Bridge *L2ERC721BridgeSession) Initialize(_messenger libcommon.Address) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.Initialize(&_L2ERC721Bridge.TransactOpts, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _messenger) returns()
func (_L2ERC721Bridge *L2ERC721BridgeTransactorSession) Initialize(_messenger libcommon.Address) (types.Transaction, error) {
	return _L2ERC721Bridge.Contract.Initialize(&_L2ERC721Bridge.TransactOpts, _messenger)
}

// BridgeERC721Params is an auto generated read-only Go binding of transcaction calldata params
type BridgeERC721Params struct {
	Param__localToken  libcommon.Address
	Param__remoteToken libcommon.Address
	Param__tokenId     *big.Int
	Param__minGasLimit uint32
	Param__extraData   []byte
}

// Parse BridgeERC721 method from calldata of a transaction
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func ParseBridgeERC721(calldata []byte) (*BridgeERC721Params, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(L2ERC721BridgeABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["bridgeERC721"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack bridgeERC721 params data: %w", err)
	}

	var paramsResult = new(BridgeERC721Params)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)
	out1 := *abi.ConvertType(out[1], new(libcommon.Address)).(*libcommon.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(uint32)).(*uint32)
	out4 := *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return &BridgeERC721Params{
		Param__localToken: out0, Param__remoteToken: out1, Param__tokenId: out2, Param__minGasLimit: out3, Param__extraData: out4,
	}, nil
}

// BridgeERC721ToParams is an auto generated read-only Go binding of transcaction calldata params
type BridgeERC721ToParams struct {
	Param__localToken  libcommon.Address
	Param__remoteToken libcommon.Address
	Param__to          libcommon.Address
	Param__tokenId     *big.Int
	Param__minGasLimit uint32
	Param__extraData   []byte
}

// Parse BridgeERC721To method from calldata of a transaction
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func ParseBridgeERC721To(calldata []byte) (*BridgeERC721ToParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(L2ERC721BridgeABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["bridgeERC721To"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack bridgeERC721To params data: %w", err)
	}

	var paramsResult = new(BridgeERC721ToParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)
	out1 := *abi.ConvertType(out[1], new(libcommon.Address)).(*libcommon.Address)
	out2 := *abi.ConvertType(out[2], new(libcommon.Address)).(*libcommon.Address)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(uint32)).(*uint32)
	out5 := *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return &BridgeERC721ToParams{
		Param__localToken: out0, Param__remoteToken: out1, Param__to: out2, Param__tokenId: out3, Param__minGasLimit: out4, Param__extraData: out5,
	}, nil
}

// FinalizeBridgeERC721Params is an auto generated read-only Go binding of transcaction calldata params
type FinalizeBridgeERC721Params struct {
	Param__localToken  libcommon.Address
	Param__remoteToken libcommon.Address
	Param__from        libcommon.Address
	Param__to          libcommon.Address
	Param__tokenId     *big.Int
	Param__extraData   []byte
}

// Parse FinalizeBridgeERC721 method from calldata of a transaction
//
// Solidity: function finalizeBridgeERC721(address _localToken, address _remoteToken, address _from, address _to, uint256 _tokenId, bytes _extraData) returns()
func ParseFinalizeBridgeERC721(calldata []byte) (*FinalizeBridgeERC721Params, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(L2ERC721BridgeABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["finalizeBridgeERC721"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack finalizeBridgeERC721 params data: %w", err)
	}

	var paramsResult = new(FinalizeBridgeERC721Params)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)
	out1 := *abi.ConvertType(out[1], new(libcommon.Address)).(*libcommon.Address)
	out2 := *abi.ConvertType(out[2], new(libcommon.Address)).(*libcommon.Address)
	out3 := *abi.ConvertType(out[3], new(libcommon.Address)).(*libcommon.Address)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return &FinalizeBridgeERC721Params{
		Param__localToken: out0, Param__remoteToken: out1, Param__from: out2, Param__to: out3, Param__tokenId: out4, Param__extraData: out5,
	}, nil
}

// InitializeParams is an auto generated read-only Go binding of transcaction calldata params

// Parse Initialize method from calldata of a transaction
//
// Solidity: function initialize(address _messenger) returns()

// L2ERC721BridgeERC721BridgeFinalizedIterator is returned from FilterERC721BridgeFinalized and is used to iterate over the raw logs and unpacked data for ERC721BridgeFinalized events raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeFinalizedIterator struct {
	Event *L2ERC721BridgeERC721BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *L2ERC721BridgeERC721BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721BridgeERC721BridgeFinalized)
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
		it.Event = new(L2ERC721BridgeERC721BridgeFinalized)
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
func (it *L2ERC721BridgeERC721BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721BridgeERC721BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721BridgeERC721BridgeFinalized represents a ERC721BridgeFinalized event raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeFinalized struct {
	LocalToken  libcommon.Address
	RemoteToken libcommon.Address
	From        libcommon.Address
	To          libcommon.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC721BridgeFinalized is a free log retrieval operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) FilterERC721BridgeFinalized(opts *bind.FilterOpts, localToken []libcommon.Address, remoteToken []libcommon.Address, from []libcommon.Address) (*L2ERC721BridgeERC721BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.FilterLogs(opts, "ERC721BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeERC721BridgeFinalizedIterator{contract: _L2ERC721Bridge.contract, event: "ERC721BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchERC721BridgeFinalized is a free log subscription operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) WatchERC721BridgeFinalized(opts *bind.WatchOpts, sink chan<- *L2ERC721BridgeERC721BridgeFinalized, localToken []libcommon.Address, remoteToken []libcommon.Address, from []libcommon.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.WatchLogs(opts, "ERC721BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721BridgeERC721BridgeFinalized)
				if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeFinalized", log); err != nil {
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

// ParseERC721BridgeFinalized is a log parse operation binding the contract event 0x1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) ParseERC721BridgeFinalized(log types.Log) (*L2ERC721BridgeERC721BridgeFinalized, error) {
	event := new(L2ERC721BridgeERC721BridgeFinalized)
	if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721BridgeERC721BridgeInitiatedIterator is returned from FilterERC721BridgeInitiated and is used to iterate over the raw logs and unpacked data for ERC721BridgeInitiated events raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeInitiatedIterator struct {
	Event *L2ERC721BridgeERC721BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *L2ERC721BridgeERC721BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721BridgeERC721BridgeInitiated)
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
		it.Event = new(L2ERC721BridgeERC721BridgeInitiated)
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
func (it *L2ERC721BridgeERC721BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721BridgeERC721BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721BridgeERC721BridgeInitiated represents a ERC721BridgeInitiated event raised by the L2ERC721Bridge contract.
type L2ERC721BridgeERC721BridgeInitiated struct {
	LocalToken  libcommon.Address
	RemoteToken libcommon.Address
	From        libcommon.Address
	To          libcommon.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC721BridgeInitiated is a free log retrieval operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) FilterERC721BridgeInitiated(opts *bind.FilterOpts, localToken []libcommon.Address, remoteToken []libcommon.Address, from []libcommon.Address) (*L2ERC721BridgeERC721BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.FilterLogs(opts, "ERC721BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeERC721BridgeInitiatedIterator{contract: _L2ERC721Bridge.contract, event: "ERC721BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchERC721BridgeInitiated is a free log subscription operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) WatchERC721BridgeInitiated(opts *bind.WatchOpts, sink chan<- *L2ERC721BridgeERC721BridgeInitiated, localToken []libcommon.Address, remoteToken []libcommon.Address, from []libcommon.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2ERC721Bridge.contract.WatchLogs(opts, "ERC721BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721BridgeERC721BridgeInitiated)
				if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeInitiated", log); err != nil {
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

// ParseERC721BridgeInitiated is a log parse operation binding the contract event 0xb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a5.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) ParseERC721BridgeInitiated(log types.Log) (*L2ERC721BridgeERC721BridgeInitiated, error) {
	event := new(L2ERC721BridgeERC721BridgeInitiated)
	if err := _L2ERC721Bridge.contract.UnpackLog(event, "ERC721BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ERC721BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2ERC721Bridge contract.
type L2ERC721BridgeInitializedIterator struct {
	Event *L2ERC721BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L2ERC721BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ERC721BridgeInitialized)
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
		it.Event = new(L2ERC721BridgeInitialized)
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
func (it *L2ERC721BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ERC721BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ERC721BridgeInitialized represents a Initialized event raised by the L2ERC721Bridge contract.
type L2ERC721BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2ERC721BridgeInitializedIterator, error) {

	logs, sub, err := _L2ERC721Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2ERC721BridgeInitializedIterator{contract: _L2ERC721Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2ERC721BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L2ERC721Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ERC721BridgeInitialized)
				if err := _L2ERC721Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2ERC721Bridge *L2ERC721BridgeFilterer) ParseInitialized(log types.Log) (*L2ERC721BridgeInitialized, error) {
	event := new(L2ERC721BridgeInitialized)
	if err := _L2ERC721Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}