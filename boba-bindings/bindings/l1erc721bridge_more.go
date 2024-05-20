// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/boba/boba-bindings/solc"
)

const L1ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"spacer_0_2_30\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_bytes30\"},{\"astId\":1003,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_contract(CrossDomainMessenger)1008\"},{\"astId\":1004,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_contract(StandardBridge)1009\"},{\"astId\":1005,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_uint256)46_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"49\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\"},{\"astId\":1007,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"superchainConfig\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_contract(SuperchainConfig)1010\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)46_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes30\":{\"encoding\":\"inplace\",\"label\":\"bytes30\",\"numberOfBytes\":\"30\"},\"t_contract(CrossDomainMessenger)1008\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(StandardBridge)1009\":{\"encoding\":\"inplace\",\"label\":\"contract StandardBridge\",\"numberOfBytes\":\"20\"},\"t_contract(SuperchainConfig)1010\":{\"encoding\":\"inplace\",\"label\":\"contract SuperchainConfig\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e mapping(uint256 =\u003e bool)))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L1ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100d45760003560e01c80635d93a3fc11610081578063927ede2d1161005b578063927ede2d14610231578063aa5574521461024f578063c89701a21461026257600080fd5b80635d93a3fc146101cc578063761f4493146102005780637f46ddb21461021357600080fd5b8063485cc955116100b2578063485cc9551461015857806354fd4d501461016b5780635c975abb146101b457600080fd5b806335e80ab3146100d95780633687011a146101235780633cb747bf14610138575b600080fd5b6032546100f99073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610136610131366004610fe1565b610282565b005b6001546100f99073ffffffffffffffffffffffffffffffffffffffff1681565b610136610166366004611064565b61032e565b6101a76040518060400160405280600581526020017f322e312e3000000000000000000000000000000000000000000000000000000081525081565b60405161011a9190611108565b6101bc610518565b604051901515815260200161011a565b6101bc6101da366004611122565b603160209081526000938452604080852082529284528284209052825290205460ff1681565b61013661020e366004611163565b6105b1565b60025473ffffffffffffffffffffffffffffffffffffffff166100f9565b60015473ffffffffffffffffffffffffffffffffffffffff166100f9565b61013661025d3660046111fb565b610a58565b6002546100f99073ffffffffffffffffffffffffffffffffffffffff1681565b333b15610316576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6103268686333388888888610b30565b505050505050565b600054610100900460ff161580801561034e5750600054600160ff909116105b806103685750303b158015610368575060005460ff166001145b6103f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161030d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561045257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790556104b083734200000000000000000000000000000000000014610e70565b801561051357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b603254604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa158015610588573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ac9190611272565b905090565b60015473ffffffffffffffffffffffffffffffffffffffff16331480156106865750600254600154604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9384169390921691636e296e45916004808201926020929091908290030181865afa15801561064a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066e9190611294565b73ffffffffffffffffffffffffffffffffffffffff16145b610712576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f746865722062726964676500606482015260840161030d565b61071a610518565b15610781576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4c314552433732314272696467653a2070617573656400000000000000000000604482015260640161030d565b3073ffffffffffffffffffffffffffffffffffffffff881603610826576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c314552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c6600000000000000000000000000000000000000000000606482015260840161030d565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152603160209081526040808320938a1683529281528282208683529052205460ff1615156001146108f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4c314552433732314272696467653a20546f6b656e204944206973206e6f742060448201527f657363726f77656420696e20746865204c312042726964676500000000000000606482015260840161030d565b73ffffffffffffffffffffffffffffffffffffffff87811660008181526031602090815260408083208b8616845282528083208884529091529081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f42842e0e000000000000000000000000000000000000000000000000000000008152306004820152918616602483015260448201859052906342842e0e90606401600060405180830381600087803b1580156109b557600080fd5b505af11580156109c9573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac87878787604051610a4794939291906112fa565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff8516610afb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f74206265206164647265737328302900000000000000000000000000000000606482015260840161030d565b610b0b8787338888888888610b30565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b73ffffffffffffffffffffffffffffffffffffffff8716610bd3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c314552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f742062652061646472657373283029000000000000000000000000000000606482015260840161030d565b600063761f449360e01b888a8989898888604051602401610bfa979695949392919061133a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000959095169490941790935273ffffffffffffffffffffffffffffffffffffffff8c81166000818152603186528381208e8416825286528381208b82529095529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590517f23b872dd000000000000000000000000000000000000000000000000000000008152908a166004820152306024820152604481018890529092506323b872dd90606401600060405180830381600087803b158015610d3a57600080fd5b505af1158015610d4e573d6000803e3d6000fd5b50506001546002546040517f3dbb202b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9283169450633dbb202b9350610db1929091169085908990600401611397565b600060405180830381600087803b158015610dcb57600080fd5b505af1158015610ddf573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a589898888604051610e5d94939291906112fa565b60405180910390a4505050505050505050565b600054610100900460ff16610f07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161030d565b6001805473ffffffffffffffffffffffffffffffffffffffff9384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560028054929093169116179055565b73ffffffffffffffffffffffffffffffffffffffff81168114610f7c57600080fd5b50565b803563ffffffff81168114610f9357600080fd5b919050565b60008083601f840112610faa57600080fd5b50813567ffffffffffffffff811115610fc257600080fd5b602083019150836020828501011115610fda57600080fd5b9250929050565b60008060008060008060a08789031215610ffa57600080fd5b863561100581610f5a565b9550602087013561101581610f5a565b94506040870135935061102a60608801610f7f565b9250608087013567ffffffffffffffff81111561104657600080fd5b61105289828a01610f98565b979a9699509497509295939492505050565b6000806040838503121561107757600080fd5b823561108281610f5a565b9150602083013561109281610f5a565b809150509250929050565b6000815180845260005b818110156110c3576020818501810151868301820152016110a7565b818111156110d5576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061111b602083018461109d565b9392505050565b60008060006060848603121561113757600080fd5b833561114281610f5a565b9250602084013561115281610f5a565b929592945050506040919091013590565b600080600080600080600060c0888a03121561117e57600080fd5b873561118981610f5a565b9650602088013561119981610f5a565b955060408801356111a981610f5a565b945060608801356111b981610f5a565b93506080880135925060a088013567ffffffffffffffff8111156111dc57600080fd5b6111e88a828b01610f98565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561121657600080fd5b873561122181610f5a565b9650602088013561123181610f5a565b9550604088013561124181610f5a565b94506060880135935061125660808901610f7f565b925060a088013567ffffffffffffffff8111156111dc57600080fd5b60006020828403121561128457600080fd5b8151801515811461111b57600080fd5b6000602082840312156112a657600080fd5b815161111b81610f5a565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006113306060830184866112b1565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261138a60c0830184866112b1565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006113c6606083018561109d565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1ERC721BridgeStorageLayoutJSON), L1ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1ERC721Bridge"] = L1ERC721BridgeStorageLayout
	deployedBytecodes["L1ERC721Bridge"] = L1ERC721BridgeDeployedBin
}
