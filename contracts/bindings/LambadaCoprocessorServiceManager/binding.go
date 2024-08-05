// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractLambadaCoprocessorServiceManager

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractLambadaCoprocessorServiceManagerMetaData contains all meta data concerning the ContractLambadaCoprocessorServiceManager contract.
var ContractLambadaCoprocessorServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorsToWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableOperatorWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableOperatorWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_lambadaCoprocessorTaskManager\",\"type\":\"address\",\"internalType\":\"contractILambadaCoprocessorTaskManager\"},{\"name\":\"_operatorWhitelistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_operatorWhitelist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lambadaCoprocessorTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILambadaCoprocessorTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorsFromWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorWhitelistDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorWhitelistEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidOpeatroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperatorWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyInWhitelist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotInWhitelist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorWhitelistAlreadyDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorWhitelistAlreadyEnabled\",\"inputs\":[]}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001f4938038062001f4983398101604081905262000034916200014c565b6001600160a01b0380841660c052808316608052811660a0528282826200005a62000071565b506200006891505062000071565b505050620001a0565b600054610100900460ff1615620000de5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000131576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014957600080fd5b50565b6000806000606084860312156200016257600080fd5b83516200016f8162000133565b6020850151909350620001828162000133565b6040850151909250620001958162000133565b809150509250925092565b60805160a05160c051611d17620002326000396000818161015d01528181610db701528181610e8b0152610f0301526000818161084b015281816109a701528181610a3e0152818161110e015281816112920152611331015260008181610676015281816107050152818161078501528181610c7301528181610e2f0152818161104c01526111ed0152611d176000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638da5cb5b11610097578063cf7e745b11610066578063cf7e745b146101e7578063d822af3c146101fa578063e481af9d1461020d578063f2fde38b1461021557600080fd5b80638da5cb5b1461019d5780639926ee7d146101ae578063a364f4da146101c1578063a98fb355146101d457600080fd5b806347b1b364116100d357806347b1b3641461014b5780634f962eca146101535780636b3aa72e1461015b578063715018a61461019557600080fd5b80631be4b9f7146100fa57806325698b301461010f57806333cfb7b714610122575b600080fd5b61010d610108366004611720565b610228565b005b61010d61011d366004611777565b6103ff565b6101356101303660046117e1565b610651565b6040516101429190611805565b60405180910390f35b61010d610b21565b61010d610bbe565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610142565b61010d610c54565b6033546001600160a01b031661017d565b61010d6101bc366004611907565b610c68565b61010d6101cf3660046117e1565b610e24565b61010d6101e23660046119b2565b610ee4565b60975461017d906001600160a01b031681565b61010d610208366004611720565b610f38565b610135611046565b61010d6102233660046117e1565b611410565b6098546001600160a01b0316336001600160a01b03161461025c57604051633d34d72960e01b815260040160405180910390fd5b60005b818110156103b957600083838381811061027b5761027b611a03565b905060200201602081019061029091906117e1565b90506001600160a01b0381166102b957604051631c16af4560e31b815260040160405180910390fd5b6001600160a01b03811660009081526099602052604090205460ff16156102f357604051630ae43b6160e21b815260040160405180910390fd5b6001600160a01b038116600090815260996020908152604091829020805460ff191660011790558151808301909252600e82526d1bdc195c985d1bdc88185919195960921b9082015261036c9085858581811061035257610352611a03565b905060200201602081019061036791906117e1565b611489565b6040516001600160a01b03821681527f697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f9060200160405180910390a1506103b281611a2f565b905061025f565b506040805180820190915260138152723bb434ba32b634b9ba1032b730b13632b2101960691b60208201526098546103fb9190600160a01b900460ff166114ce565b5050565b600054610100900460ff161580801561041f5750600054600160ff909116105b806104395750303b158015610439575060005460ff166001145b6104a15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156104c4576000805461ff0019166101001790555b609780546001600160a01b0319166001600160a01b03871617905560988054336001600160a81b031990911617600160a01b8615150217905560005b828110156105c257600084848381811061051c5761051c611a03565b905060200201602081019061053191906117e1565b90506001600160a01b03811661055a57604051631c16af4560e31b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916600117905590519182527f697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f910160405180910390a1506105bb81611a2f565b9050610500565b5060408051808201909152601381527277686974656c69737420656e61626c6564203160681b60208201526098546106049190600160a01b900460ff166114ce565b801561064a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156106bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e19190611a4a565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561074c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107709190611a63565b90506001600160c01b038116158061080a57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108059190611a8c565b60ff16155b1561082657505060408051600081526020810190915292915050565b600061083a826001600160c01b0316611513565b90506000805b8251811015610910577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061088a5761088a611a03565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156108ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f29190611a4a565b6108fc9083611aaf565b91508061090881611a2f565b915050610840565b5060008167ffffffffffffffff81111561092c5761092c611852565b604051908082528060200260200182016040528015610955578160200160208202803683370190505b5090506000805b8451811015610b1457600085828151811061097957610979611a03565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156109ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a129190611a4a565b905060005b81811015610afe576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610a8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab09190611ac7565b60000151868681518110610ac657610ac6611a03565b6001600160a01b039092166020928302919091019091015284610ae881611a2f565b9550508080610af690611a2f565b915050610a17565b5050508080610b0c90611a2f565b91505061095c565b5090979650505050505050565b6098546001600160a01b0316336001600160a01b031614610b5557604051633d34d72960e01b815260040160405180910390fd5b609854600160a01b900460ff1615610b805760405163b1a183d760e01b815260040160405180910390fd5b6098805460ff60a01b1916600160a01b1790556040517fa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb90600090a1565b6098546001600160a01b0316336001600160a01b031614610bf257604051633d34d72960e01b815260040160405180910390fd5b609854600160a01b900460ff16610c1c5760405163a0dd88cf60e01b815260040160405180910390fd5b6098805460ff60a01b191690556040517f87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb290600090a1565b610c5c6115d6565b610c666000611630565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cb05760405162461bcd60e51b815260040161049890611b37565b60408051808201909152601c81527f6f70657261746f722077686974656c69737420656e61626c65642033000000006020820152609854610cfb9190600160a01b900460ff166114ce565b60408051808201825260158152741bdc195c985d1bdc881a5b881dda1a5d195b1a5cdd605a1b6020808301919091526001600160a01b038516600090815260999091529190912054610d50919060ff166114ce565b609854600160a01b900460ff168015610d8257506001600160a01b03821660009081526099602052604090205460ff16155b15610da057604051638065bea960e01b815260040160405180910390fd5b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610dee9085908590600401611bfc565b600060405180830381600087803b158015610e0857600080fd5b505af1158015610e1c573d6000803e3d6000fd5b505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e6c5760405162461bcd60e51b815260040161049890611b37565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610ed057600080fd5b505af115801561064a573d6000803e3d6000fd5b610eec6115d6565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610eb6908490600401611c47565b6098546001600160a01b0316336001600160a01b031614610f6c57604051633d34d72960e01b815260040160405180910390fd5b60005b81811015611041576000838383818110610f8b57610f8b611a03565b9050602002016020810190610fa091906117e1565b6001600160a01b03811660009081526099602052604090205490915060ff16610fdc57604051638065bea960e01b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916905590519182527f90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a910160405180910390a15061103a81611a2f565b9050610f6f565b505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110cc9190611a8c565b60ff169050806110ea57505060408051600081526020810190915290565b6000805b8281101561119f57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa15801561115d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111819190611a4a565b61118b9083611aaf565b91508061119781611a2f565b9150506110ee565b5060008167ffffffffffffffff8111156111bb576111bb611852565b6040519080825280602002602001820160405280156111e4578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126d9190611a8c565b60ff1681101561140657604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa1580156112e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113059190611a4a565b905060005b818110156113f1576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561137f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a39190611ac7565b600001518585815181106113b9576113b9611a03565b6001600160a01b0390921660209283029190910190910152836113db81611a2f565b94505080806113e990611a2f565b91505061130a565b505080806113fe90611a2f565b9150506111eb565b5090949350505050565b6114186115d6565b6001600160a01b03811661147d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610498565b61148681611630565b50565b6103fb828260405160240161149f929190611c5a565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b179052611682565b6103fb82826040516024016114e4929190611c84565b60408051601f198184030181529190526020810180516001600160e01b031663c3b5563560e01b179052611682565b6060600080611521846116a3565b61ffff1667ffffffffffffffff81111561153d5761153d611852565b6040519080825280601f01601f191660200182016040528015611567576020820181803683370190505b5090506000805b82518210801561157f575061010081105b15611406576001811b9350858416156115c6578060f81b8383815181106115a8576115a8611a03565b60200101906001600160f81b031916908160001a9053508160010191505b6115cf81611a2f565b905061156e565b6033546001600160a01b03163314610c665760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610498565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b6000805b82156116ce576116b8600184611ca8565b90921691806116c681611cbf565b9150506116a7565b92915050565b60008083601f8401126116e657600080fd5b50813567ffffffffffffffff8111156116fe57600080fd5b6020830191508360208260051b850101111561171957600080fd5b9250929050565b6000806020838503121561173357600080fd5b823567ffffffffffffffff81111561174a57600080fd5b611756858286016116d4565b90969095509350505050565b6001600160a01b038116811461148657600080fd5b6000806000806060858703121561178d57600080fd5b843561179881611762565b9350602085013580151581146117ad57600080fd5b9250604085013567ffffffffffffffff8111156117c957600080fd5b6117d5878288016116d4565b95989497509550505050565b6000602082840312156117f357600080fd5b81356117fe81611762565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156118465783516001600160a01b031683529284019291840191600101611821565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561188b5761188b611852565b60405290565b600067ffffffffffffffff808411156118ac576118ac611852565b604051601f8501601f19908116603f011681019082821181831017156118d4576118d4611852565b816040528093508581528686860111156118ed57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561191a57600080fd5b823561192581611762565b9150602083013567ffffffffffffffff8082111561194257600080fd5b908401906060828703121561195657600080fd5b61195e611868565b82358281111561196d57600080fd5b83019150601f8201871361198057600080fd5b61198f87833560208501611891565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156119c457600080fd5b813567ffffffffffffffff8111156119db57600080fd5b8201601f810184136119ec57600080fd5b6119fb84823560208401611891565b949350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611a4357611a43611a19565b5060010190565b600060208284031215611a5c57600080fd5b5051919050565b600060208284031215611a7557600080fd5b81516001600160c01b03811681146117fe57600080fd5b600060208284031215611a9e57600080fd5b815160ff811681146117fe57600080fd5b60008219821115611ac257611ac2611a19565b500190565b600060408284031215611ad957600080fd5b6040516040810181811067ffffffffffffffff82111715611afc57611afc611852565b6040528251611b0a81611762565b815260208301516bffffffffffffffffffffffff81168114611b2b57600080fd5b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b81811015611bd557602081850181015186830182015201611bb9565b81811115611be7576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b0383168152604060208201526000825160606040840152611c2660a0840182611baf565b90506020840151606084015260408401516080840152809150509392505050565b6020815260006117fe6020830184611baf565b604081526000611c6d6040830185611baf565b905060018060a01b03831660208301529392505050565b604081526000611c976040830185611baf565b905082151560208301529392505050565b600082821015611cba57611cba611a19565b500390565b600061ffff80831681811415611cd757611cd7611a19565b600101939250505056fea2646970667358221220e3e9a0e08402e92c2f8b26ebf13b6af2b9e334d2dec8c6a7f9db19ae1b98a41564736f6c634300080c0033",
}

// ContractLambadaCoprocessorServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractLambadaCoprocessorServiceManagerMetaData.ABI instead.
var ContractLambadaCoprocessorServiceManagerABI = ContractLambadaCoprocessorServiceManagerMetaData.ABI

// ContractLambadaCoprocessorServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractLambadaCoprocessorServiceManagerMetaData.Bin instead.
var ContractLambadaCoprocessorServiceManagerBin = ContractLambadaCoprocessorServiceManagerMetaData.Bin

// DeployContractLambadaCoprocessorServiceManager deploys a new Ethereum contract, binding an instance of ContractLambadaCoprocessorServiceManager to it.
func DeployContractLambadaCoprocessorServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address) (common.Address, *types.Transaction, *ContractLambadaCoprocessorServiceManager, error) {
	parsed, err := ContractLambadaCoprocessorServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractLambadaCoprocessorServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractLambadaCoprocessorServiceManager{ContractLambadaCoprocessorServiceManagerCaller: ContractLambadaCoprocessorServiceManagerCaller{contract: contract}, ContractLambadaCoprocessorServiceManagerTransactor: ContractLambadaCoprocessorServiceManagerTransactor{contract: contract}, ContractLambadaCoprocessorServiceManagerFilterer: ContractLambadaCoprocessorServiceManagerFilterer{contract: contract}}, nil
}

// ContractLambadaCoprocessorServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractLambadaCoprocessorServiceManager struct {
	ContractLambadaCoprocessorServiceManagerCaller     // Read-only binding to the contract
	ContractLambadaCoprocessorServiceManagerTransactor // Write-only binding to the contract
	ContractLambadaCoprocessorServiceManagerFilterer   // Log filterer for contract events
}

// ContractLambadaCoprocessorServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractLambadaCoprocessorServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractLambadaCoprocessorServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractLambadaCoprocessorServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractLambadaCoprocessorServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractLambadaCoprocessorServiceManagerSession struct {
	Contract     *ContractLambadaCoprocessorServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                             // Call options to use throughout this session
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractLambadaCoprocessorServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractLambadaCoprocessorServiceManagerCallerSession struct {
	Contract *ContractLambadaCoprocessorServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                   // Call options to use throughout this session
}

// ContractLambadaCoprocessorServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractLambadaCoprocessorServiceManagerTransactorSession struct {
	Contract     *ContractLambadaCoprocessorServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                   // Transaction auth options to use throughout this session
}

// ContractLambadaCoprocessorServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractLambadaCoprocessorServiceManagerRaw struct {
	Contract *ContractLambadaCoprocessorServiceManager // Generic contract binding to access the raw methods on
}

// ContractLambadaCoprocessorServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorServiceManagerCallerRaw struct {
	Contract *ContractLambadaCoprocessorServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractLambadaCoprocessorServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorServiceManagerTransactorRaw struct {
	Contract *ContractLambadaCoprocessorServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractLambadaCoprocessorServiceManager creates a new instance of ContractLambadaCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorServiceManager(address common.Address, backend bind.ContractBackend) (*ContractLambadaCoprocessorServiceManager, error) {
	contract, err := bindContractLambadaCoprocessorServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManager{ContractLambadaCoprocessorServiceManagerCaller: ContractLambadaCoprocessorServiceManagerCaller{contract: contract}, ContractLambadaCoprocessorServiceManagerTransactor: ContractLambadaCoprocessorServiceManagerTransactor{contract: contract}, ContractLambadaCoprocessorServiceManagerFilterer: ContractLambadaCoprocessorServiceManagerFilterer{contract: contract}}, nil
}

// NewContractLambadaCoprocessorServiceManagerCaller creates a new read-only instance of ContractLambadaCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractLambadaCoprocessorServiceManagerCaller, error) {
	contract, err := bindContractLambadaCoprocessorServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerCaller{contract: contract}, nil
}

// NewContractLambadaCoprocessorServiceManagerTransactor creates a new write-only instance of ContractLambadaCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractLambadaCoprocessorServiceManagerTransactor, error) {
	contract, err := bindContractLambadaCoprocessorServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerTransactor{contract: contract}, nil
}

// NewContractLambadaCoprocessorServiceManagerFilterer creates a new log filterer instance of ContractLambadaCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractLambadaCoprocessorServiceManagerFilterer, error) {
	contract, err := bindContractLambadaCoprocessorServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerFilterer{contract: contract}, nil
}

// bindContractLambadaCoprocessorServiceManager binds a generic wrapper to an already deployed contract.
func bindContractLambadaCoprocessorServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractLambadaCoprocessorServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractLambadaCoprocessorServiceManager.Contract.ContractLambadaCoprocessorServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.ContractLambadaCoprocessorServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.ContractLambadaCoprocessorServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractLambadaCoprocessorServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.AvsDirectory(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.AvsDirectory(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractLambadaCoprocessorServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractLambadaCoprocessorServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.GetRestakeableStrategies(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.GetRestakeableStrategies(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// LambadaCoprocessorTaskManager is a free data retrieval call binding the contract method 0xcf7e745b.
//
// Solidity: function lambadaCoprocessorTaskManager() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCaller) LambadaCoprocessorTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorServiceManager.contract.Call(opts, &out, "lambadaCoprocessorTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LambadaCoprocessorTaskManager is a free data retrieval call binding the contract method 0xcf7e745b.
//
// Solidity: function lambadaCoprocessorTaskManager() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) LambadaCoprocessorTaskManager() (common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.LambadaCoprocessorTaskManager(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// LambadaCoprocessorTaskManager is a free data retrieval call binding the contract method 0xcf7e745b.
//
// Solidity: function lambadaCoprocessorTaskManager() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCallerSession) LambadaCoprocessorTaskManager() (common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.LambadaCoprocessorTaskManager(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) Owner() (common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.Owner(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.Owner(&_ContractLambadaCoprocessorServiceManager.CallOpts)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) AddOperatorsToWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "addOperatorsToWhitelist", operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.AddOperatorsToWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.AddOperatorsToWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operators)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operator)
}

// DisableOperatorWhitelist is a paid mutator transaction binding the contract method 0x4f962eca.
//
// Solidity: function disableOperatorWhitelist() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) DisableOperatorWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "disableOperatorWhitelist")
}

// DisableOperatorWhitelist is a paid mutator transaction binding the contract method 0x4f962eca.
//
// Solidity: function disableOperatorWhitelist() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) DisableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.DisableOperatorWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts)
}

// DisableOperatorWhitelist is a paid mutator transaction binding the contract method 0x4f962eca.
//
// Solidity: function disableOperatorWhitelist() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) DisableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.DisableOperatorWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts)
}

// EnableOperatorWhitelist is a paid mutator transaction binding the contract method 0x47b1b364.
//
// Solidity: function enableOperatorWhitelist() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) EnableOperatorWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "enableOperatorWhitelist")
}

// EnableOperatorWhitelist is a paid mutator transaction binding the contract method 0x47b1b364.
//
// Solidity: function enableOperatorWhitelist() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) EnableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.EnableOperatorWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts)
}

// EnableOperatorWhitelist is a paid mutator transaction binding the contract method 0x47b1b364.
//
// Solidity: function enableOperatorWhitelist() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) EnableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.EnableOperatorWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x25698b30.
//
// Solidity: function initialize(address _lambadaCoprocessorTaskManager, bool _operatorWhitelistEnabled, address[] _operatorWhitelist) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _lambadaCoprocessorTaskManager common.Address, _operatorWhitelistEnabled bool, _operatorWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "initialize", _lambadaCoprocessorTaskManager, _operatorWhitelistEnabled, _operatorWhitelist)
}

// Initialize is a paid mutator transaction binding the contract method 0x25698b30.
//
// Solidity: function initialize(address _lambadaCoprocessorTaskManager, bool _operatorWhitelistEnabled, address[] _operatorWhitelist) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) Initialize(_lambadaCoprocessorTaskManager common.Address, _operatorWhitelistEnabled bool, _operatorWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.Initialize(&_ContractLambadaCoprocessorServiceManager.TransactOpts, _lambadaCoprocessorTaskManager, _operatorWhitelistEnabled, _operatorWhitelist)
}

// Initialize is a paid mutator transaction binding the contract method 0x25698b30.
//
// Solidity: function initialize(address _lambadaCoprocessorTaskManager, bool _operatorWhitelistEnabled, address[] _operatorWhitelist) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) Initialize(_lambadaCoprocessorTaskManager common.Address, _operatorWhitelistEnabled bool, _operatorWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.Initialize(&_ContractLambadaCoprocessorServiceManager.TransactOpts, _lambadaCoprocessorTaskManager, _operatorWhitelistEnabled, _operatorWhitelist)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.RegisterOperatorToAVS(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.RegisterOperatorToAVS(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operator, operatorSignature)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) RemoveOperatorsFromWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "removeOperatorsFromWhitelist", operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.RemoveOperatorsFromWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.RemoveOperatorsFromWhitelist(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.RenounceOwnership(&_ContractLambadaCoprocessorServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.RenounceOwnership(&_ContractLambadaCoprocessorServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.TransferOwnership(&_ContractLambadaCoprocessorServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.TransferOwnership(&_ContractLambadaCoprocessorServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.UpdateAVSMetadataURI(&_ContractLambadaCoprocessorServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.UpdateAVSMetadataURI(&_ContractLambadaCoprocessorServiceManager.TransactOpts, _metadataURI)
}

// ContractLambadaCoprocessorServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerInitializedIterator struct {
	Event *ContractLambadaCoprocessorServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorServiceManagerInitialized)
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
		it.Event = new(ContractLambadaCoprocessorServiceManagerInitialized)
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
func (it *ContractLambadaCoprocessorServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorServiceManagerInitialized represents a Initialized event raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractLambadaCoprocessorServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerInitializedIterator{contract: _ContractLambadaCoprocessorServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorServiceManagerInitialized)
				if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractLambadaCoprocessorServiceManagerInitialized, error) {
	event := new(ContractLambadaCoprocessorServiceManagerInitialized)
	if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelistIterator is returned from FilterOperatorAddedToWhitelist and is used to iterate over the raw logs and unpacked data for OperatorAddedToWhitelist events raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelistIterator struct {
	Event *ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist)
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
		it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist)
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
func (it *ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist represents a OperatorAddedToWhitelist event raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToWhitelist is a free log retrieval operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address operator)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) FilterOperatorAddedToWhitelist(opts *bind.FilterOpts) (*ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelistIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelistIterator{contract: _ContractLambadaCoprocessorServiceManager.contract, event: "OperatorAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToWhitelist is a free log subscription operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address operator)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) WatchOperatorAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist)
				if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorAddedToWhitelist", log); err != nil {
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

// ParseOperatorAddedToWhitelist is a log parse operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address operator)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) ParseOperatorAddedToWhitelist(log types.Log) (*ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist, error) {
	event := new(ContractLambadaCoprocessorServiceManagerOperatorAddedToWhitelist)
	if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorAddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator is returned from FilterOperatorRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromWhitelist events raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator struct {
	Event *ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist)
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
		it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist)
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
func (it *ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist represents a OperatorRemovedFromWhitelist event raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address operator)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) FilterOperatorRemovedFromWhitelist(opts *bind.FilterOpts) (*ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator{contract: _ContractLambadaCoprocessorServiceManager.contract, event: "OperatorRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromWhitelist is a free log subscription operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address operator)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) WatchOperatorRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist)
				if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorRemovedFromWhitelist", log); err != nil {
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

// ParseOperatorRemovedFromWhitelist is a log parse operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address operator)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) ParseOperatorRemovedFromWhitelist(log types.Log) (*ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist, error) {
	event := new(ContractLambadaCoprocessorServiceManagerOperatorRemovedFromWhitelist)
	if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabledIterator is returned from FilterOperatorWhitelistDisabled and is used to iterate over the raw logs and unpacked data for OperatorWhitelistDisabled events raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabledIterator struct {
	Event *ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled)
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
		it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled)
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
func (it *ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled represents a OperatorWhitelistDisabled event raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistDisabled is a free log retrieval operation binding the contract event 0x87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb2.
//
// Solidity: event OperatorWhitelistDisabled()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) FilterOperatorWhitelistDisabled(opts *bind.FilterOpts) (*ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabledIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorWhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabledIterator{contract: _ContractLambadaCoprocessorServiceManager.contract, event: "OperatorWhitelistDisabled", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistDisabled is a free log subscription operation binding the contract event 0x87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb2.
//
// Solidity: event OperatorWhitelistDisabled()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) WatchOperatorWhitelistDisabled(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorWhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled)
				if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistDisabled", log); err != nil {
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

// ParseOperatorWhitelistDisabled is a log parse operation binding the contract event 0x87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb2.
//
// Solidity: event OperatorWhitelistDisabled()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) ParseOperatorWhitelistDisabled(log types.Log) (*ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled, error) {
	event := new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistDisabled)
	if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabledIterator is returned from FilterOperatorWhitelistEnabled and is used to iterate over the raw logs and unpacked data for OperatorWhitelistEnabled events raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabledIterator struct {
	Event *ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled)
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
		it.Event = new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled)
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
func (it *ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled represents a OperatorWhitelistEnabled event raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistEnabled is a free log retrieval operation binding the contract event 0xa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb.
//
// Solidity: event OperatorWhitelistEnabled()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) FilterOperatorWhitelistEnabled(opts *bind.FilterOpts) (*ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabledIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorWhitelistEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabledIterator{contract: _ContractLambadaCoprocessorServiceManager.contract, event: "OperatorWhitelistEnabled", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistEnabled is a free log subscription operation binding the contract event 0xa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb.
//
// Solidity: event OperatorWhitelistEnabled()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) WatchOperatorWhitelistEnabled(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorWhitelistEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled)
				if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistEnabled", log); err != nil {
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

// ParseOperatorWhitelistEnabled is a log parse operation binding the contract event 0xa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb.
//
// Solidity: event OperatorWhitelistEnabled()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) ParseOperatorWhitelistEnabled(log types.Log) (*ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled, error) {
	event := new(ContractLambadaCoprocessorServiceManagerOperatorWhitelistEnabled)
	if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOwnershipTransferredIterator struct {
	Event *ContractLambadaCoprocessorServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractLambadaCoprocessorServiceManagerOwnershipTransferred)
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
func (it *ContractLambadaCoprocessorServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractLambadaCoprocessorServiceManager contract.
type ContractLambadaCoprocessorServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractLambadaCoprocessorServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorServiceManagerOwnershipTransferredIterator{contract: _ContractLambadaCoprocessorServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorServiceManagerOwnershipTransferred)
				if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractLambadaCoprocessorServiceManagerOwnershipTransferred, error) {
	event := new(ContractLambadaCoprocessorServiceManagerOwnershipTransferred)
	if err := _ContractLambadaCoprocessorServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
