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
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001ccb38038062001ccb83398101604081905262000034916200014c565b6001600160a01b0380841660c052808316608052811660a0528282826200005a62000071565b506200006891505062000071565b505050620001a0565b600054610100900460ff1615620000de5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000131576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014957600080fd5b50565b6000806000606084860312156200016257600080fd5b83516200016f8162000133565b6020850151909350620001828162000133565b6040850151909250620001958162000133565b809150509250925092565b60805160a05160c051611a99620002326000396000818161015d01528181610c3701528181610d0b0152610d8301526000818161076b015281816108c70152818161095e01528181610f890152818161110d01526111ac01526000818161059601528181610625015281816106a501528181610b9301528181610caf01528181610ec701526110680152611a996000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638da5cb5b11610097578063cf7e745b11610066578063cf7e745b146101e7578063d822af3c146101fa578063e481af9d1461020d578063f2fde38b1461021557600080fd5b80638da5cb5b1461019d5780639926ee7d146101ae578063a364f4da146101c1578063a98fb355146101d457600080fd5b806347b1b364116100d357806347b1b3641461014b5780634f962eca146101535780636b3aa72e1461015b578063715018a61461019557600080fd5b80631be4b9f7146100fa57806325698b301461010f57806333cfb7b714610122575b600080fd5b61010d6101083660046114f0565b610228565b005b61010d61011d366004611547565b610360565b6101356101303660046115b1565b610571565b60405161014291906115d5565b60405180910390f35b61010d610a41565b61010d610ade565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610142565b61010d610b74565b6033546001600160a01b031661017d565b61010d6101bc3660046116d7565b610b88565b61010d6101cf3660046115b1565b610ca4565b61010d6101e2366004611782565b610d64565b60975461017d906001600160a01b031681565b61010d6102083660046114f0565b610db8565b610135610ec1565b61010d6102233660046115b1565b61128b565b6098546001600160a01b0316336001600160a01b03161461025c57604051633d34d72960e01b815260040160405180910390fd5b60005b8181101561035b57600083838381811061027b5761027b6117d3565b905060200201602081019061029091906115b1565b90506001600160a01b0381166102b957604051631c16af4560e31b815260040160405180910390fd5b6001600160a01b03811660009081526099602052604090205460ff16156102f357604051630ae43b6160e21b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916600117905590519182527f697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f910160405180910390a150610354816117ff565b905061025f565b505050565b600054610100900460ff16158080156103805750600054600160ff909116105b8061039a5750303b15801561039a575060005460ff166001145b6104025760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610425576000805461ff0019166101001790555b609780546001600160a01b0319166001600160a01b03871617905560988054336001600160a81b031990911617600160a01b8615150217905560005b8281101561052357600084848381811061047d5761047d6117d3565b905060200201602081019061049291906115b1565b90506001600160a01b0381166104bb57604051631c16af4560e31b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916600117905590519182527f697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f910160405180910390a15061051c816117ff565b9050610461565b50801561056a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156105dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610601919061181a565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561066c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106909190611833565b90506001600160c01b038116158061072a57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610701573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610725919061185c565b60ff16155b1561074657505060408051600081526020810190915292915050565b600061075a826001600160c01b0316611304565b90506000805b8251811015610830577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106107aa576107aa6117d3565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156107ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610812919061181a565b61081c908361187f565b915080610828816117ff565b915050610760565b5060008167ffffffffffffffff81111561084c5761084c611622565b604051908082528060200260200182016040528015610875578160200160208202803683370190505b5090506000805b8451811015610a34576000858281518110610899576108996117d3565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa15801561090e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610932919061181a565b905060005b81811015610a1e576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156109ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d09190611897565b600001518686815181106109e6576109e66117d3565b6001600160a01b039092166020928302919091019091015284610a08816117ff565b9550508080610a16906117ff565b915050610937565b5050508080610a2c906117ff565b91505061087c565b5090979650505050505050565b6098546001600160a01b0316336001600160a01b031614610a7557604051633d34d72960e01b815260040160405180910390fd5b609854600160a01b900460ff1615610aa05760405163b1a183d760e01b815260040160405180910390fd5b6098805460ff60a01b1916600160a01b1790556040517fa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb90600090a1565b6098546001600160a01b0316336001600160a01b031614610b1257604051633d34d72960e01b815260040160405180910390fd5b609854600160a01b900460ff16610b3c5760405163a0dd88cf60e01b815260040160405180910390fd5b6098805460ff60a01b191690556040517f87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb290600090a1565b610b7c6113c7565b610b866000611421565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610bd05760405162461bcd60e51b81526004016103f990611907565b609854600160a01b900460ff168015610c0257506001600160a01b03821660009081526099602052604090205460ff16155b15610c2057604051638065bea960e01b815260040160405180910390fd5b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610c6e90859085906004016119cc565b600060405180830381600087803b158015610c8857600080fd5b505af1158015610c9c573d6000803e3d6000fd5b505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cec5760405162461bcd60e51b81526004016103f990611907565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610d5057600080fd5b505af115801561056a573d6000803e3d6000fd5b610d6c6113c7565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610d36908490600401611a17565b6098546001600160a01b0316336001600160a01b031614610dec57604051633d34d72960e01b815260040160405180910390fd5b60005b8181101561035b576000838383818110610e0b57610e0b6117d3565b9050602002016020810190610e2091906115b1565b6001600160a01b03811660009081526099602052604090205490915060ff16610e5c57604051638065bea960e01b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916905590519182527f90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a910160405180910390a150610eba816117ff565b9050610def565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f47919061185c565b60ff16905080610f6557505060408051600081526020810190915290565b6000805b8281101561101a57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffc919061181a565b611006908361187f565b915080611012816117ff565b915050610f69565b5060008167ffffffffffffffff81111561103657611036611622565b60405190808252806020026020018201604052801561105f578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e8919061185c565b60ff1681101561128157604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa15801561115c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611180919061181a565b905060005b8181101561126c576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156111fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121e9190611897565b60000151858581518110611234576112346117d3565b6001600160a01b039092166020928302919091019091015283611256816117ff565b9450508080611264906117ff565b915050611185565b50508080611279906117ff565b915050611066565b5090949350505050565b6112936113c7565b6001600160a01b0381166112f85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103f9565b61130181611421565b50565b606060008061131284611473565b61ffff1667ffffffffffffffff81111561132e5761132e611622565b6040519080825280601f01601f191660200182016040528015611358576020820181803683370190505b5090506000805b825182108015611370575061010081105b15611281576001811b9350858416156113b7578060f81b838381518110611399576113996117d3565b60200101906001600160f81b031916908160001a9053508160010191505b6113c0816117ff565b905061135f565b6033546001600160a01b03163314610b865760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103f9565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b821561149e57611488600184611a2a565b909216918061149681611a41565b915050611477565b92915050565b60008083601f8401126114b657600080fd5b50813567ffffffffffffffff8111156114ce57600080fd5b6020830191508360208260051b85010111156114e957600080fd5b9250929050565b6000806020838503121561150357600080fd5b823567ffffffffffffffff81111561151a57600080fd5b611526858286016114a4565b90969095509350505050565b6001600160a01b038116811461130157600080fd5b6000806000806060858703121561155d57600080fd5b843561156881611532565b93506020850135801515811461157d57600080fd5b9250604085013567ffffffffffffffff81111561159957600080fd5b6115a5878288016114a4565b95989497509550505050565b6000602082840312156115c357600080fd5b81356115ce81611532565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156116165783516001600160a01b0316835292840192918401916001016115f1565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561165b5761165b611622565b60405290565b600067ffffffffffffffff8084111561167c5761167c611622565b604051601f8501601f19908116603f011681019082821181831017156116a4576116a4611622565b816040528093508581528686860111156116bd57600080fd5b858560208301376000602087830101525050509392505050565b600080604083850312156116ea57600080fd5b82356116f581611532565b9150602083013567ffffffffffffffff8082111561171257600080fd5b908401906060828703121561172657600080fd5b61172e611638565b82358281111561173d57600080fd5b83019150601f8201871361175057600080fd5b61175f87833560208501611661565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561179457600080fd5b813567ffffffffffffffff8111156117ab57600080fd5b8201601f810184136117bc57600080fd5b6117cb84823560208401611661565b949350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611813576118136117e9565b5060010190565b60006020828403121561182c57600080fd5b5051919050565b60006020828403121561184557600080fd5b81516001600160c01b03811681146115ce57600080fd5b60006020828403121561186e57600080fd5b815160ff811681146115ce57600080fd5b60008219821115611892576118926117e9565b500190565b6000604082840312156118a957600080fd5b6040516040810181811067ffffffffffffffff821117156118cc576118cc611622565b60405282516118da81611532565b815260208301516bffffffffffffffffffffffff811681146118fb57600080fd5b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156119a557602081850181015186830182015201611989565b818111156119b7576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b03831681526040602082015260008251606060408401526119f660a084018261197f565b90506020840151606084015260408401516080840152809150509392505050565b6020815260006115ce602083018461197f565b600082821015611a3c57611a3c6117e9565b500390565b600061ffff80831681811415611a5957611a596117e9565b600101939250505056fea2646970667358221220b54c5ee97ae8a24a66d92677bbf9f4cf66e7a573c8d2aa65136fe5d074678d8b64736f6c634300080c0033",
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
