// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractCoprocessorServiceManager

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

// ContractCoprocessorServiceManagerMetaData contains all meta data concerning the ContractCoprocessorServiceManager contract.
var ContractCoprocessorServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperatorsToWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"coprocessorTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICoprocessorTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableOperatorWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableOperatorWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_lambadaCoprocessorTaskManager\",\"type\":\"address\",\"internalType\":\"contractICoprocessorTaskManager\"},{\"name\":\"_operatorWhitelistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_operatorWhitelist\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperatorsFromWhitelist\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorWhitelistDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorWhitelistEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidOpeatroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperatorWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyInWhitelist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotInWhitelist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorWhitelistAlreadyDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorWhitelistAlreadyEnabled\",\"inputs\":[]}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001ccf38038062001ccf83398101604081905262000034916200014c565b6001600160a01b0380841660c052808316608052811660a0528282826200005a62000071565b506200006891505062000071565b505050620001a0565b600054610100900460ff1615620000de5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000131576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014957600080fd5b50565b6000806000606084860312156200016257600080fd5b83516200016f8162000133565b6020850151909350620001828162000133565b6040850151909250620001958162000133565b809150509250925092565b60805160a05160c051611a9d620002326000396000818161018801528181610c3b01528181610d0f0152610d8701526000818161076f015281816108cb0152818161096201528181610f8d0152818161111101526111b001526000818161059a01528181610629015281816106a901528181610b9701528181610cb301528181610ecb015261106c0152611a9d6000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063715018a611610097578063a98fb35511610066578063a98fb355146101eb578063d822af3c146101fe578063e481af9d14610211578063f2fde38b1461021957600080fd5b8063715018a6146101ac5780638da5cb5b146101b45780639926ee7d146101c5578063a364f4da146101d857600080fd5b806347b1b364116100d357806347b1b3641461014b5780634f962eca146101535780636052779e1461015b5780636b3aa72e1461018657600080fd5b80631be4b9f7146100fa57806325698b301461010f57806333cfb7b714610122575b600080fd5b61010d6101083660046114f4565b61022c565b005b61010d61011d36600461154b565b610364565b6101356101303660046115b5565b610575565b60405161014291906115d9565b60405180910390f35b61010d610a45565b61010d610ae2565b60975461016e906001600160a01b031681565b6040516001600160a01b039091168152602001610142565b7f000000000000000000000000000000000000000000000000000000000000000061016e565b61010d610b78565b6033546001600160a01b031661016e565b61010d6101d33660046116db565b610b8c565b61010d6101e63660046115b5565b610ca8565b61010d6101f9366004611786565b610d68565b61010d61020c3660046114f4565b610dbc565b610135610ec5565b61010d6102273660046115b5565b61128f565b6098546001600160a01b0316336001600160a01b03161461026057604051633d34d72960e01b815260040160405180910390fd5b60005b8181101561035f57600083838381811061027f5761027f6117d7565b905060200201602081019061029491906115b5565b90506001600160a01b0381166102bd57604051631c16af4560e31b815260040160405180910390fd5b6001600160a01b03811660009081526099602052604090205460ff16156102f757604051630ae43b6160e21b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916600117905590519182527f697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f910160405180910390a15061035881611803565b9050610263565b505050565b600054610100900460ff16158080156103845750600054600160ff909116105b8061039e5750303b15801561039e575060005460ff166001145b6104065760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610429576000805461ff0019166101001790555b609780546001600160a01b0319166001600160a01b03871617905560988054336001600160a81b031990911617600160a01b8615150217905560005b82811015610527576000848483818110610481576104816117d7565b905060200201602081019061049691906115b5565b90506001600160a01b0381166104bf57604051631c16af4560e31b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916600117905590519182527f697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f910160405180910390a15061052081611803565b9050610465565b50801561056e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156105e1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610605919061181e565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa158015610670573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106949190611837565b90506001600160c01b038116158061072e57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610705573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107299190611860565b60ff16155b1561074a57505060408051600081526020810190915292915050565b600061075e826001600160c01b0316611308565b90506000805b8251811015610834577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106107ae576107ae6117d7565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156107f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610816919061181e565b6108209083611883565b91508061082c81611803565b915050610764565b5060008167ffffffffffffffff81111561085057610850611626565b604051908082528060200260200182016040528015610879578160200160208202803683370190505b5090506000805b8451811015610a3857600085828151811061089d5761089d6117d7565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610912573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610936919061181e565b905060005b81811015610a22576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156109b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d4919061189b565b600001518686815181106109ea576109ea6117d7565b6001600160a01b039092166020928302919091019091015284610a0c81611803565b9550508080610a1a90611803565b91505061093b565b5050508080610a3090611803565b915050610880565b5090979650505050505050565b6098546001600160a01b0316336001600160a01b031614610a7957604051633d34d72960e01b815260040160405180910390fd5b609854600160a01b900460ff1615610aa45760405163b1a183d760e01b815260040160405180910390fd5b6098805460ff60a01b1916600160a01b1790556040517fa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb90600090a1565b6098546001600160a01b0316336001600160a01b031614610b1657604051633d34d72960e01b815260040160405180910390fd5b609854600160a01b900460ff16610b405760405163a0dd88cf60e01b815260040160405180910390fd5b6098805460ff60a01b191690556040517f87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb290600090a1565b610b806113cb565b610b8a6000611425565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610bd45760405162461bcd60e51b81526004016103fd9061190b565b609854600160a01b900460ff168015610c0657506001600160a01b03821660009081526099602052604090205460ff16155b15610c2457604051638065bea960e01b815260040160405180910390fd5b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d90610c7290859085906004016119d0565b600060405180830381600087803b158015610c8c57600080fd5b505af1158015610ca0573d6000803e3d6000fd5b505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cf05760405162461bcd60e51b81526004016103fd9061190b565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610d5457600080fd5b505af115801561056e573d6000803e3d6000fd5b610d706113cb565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610d3a908490600401611a1b565b6098546001600160a01b0316336001600160a01b031614610df057604051633d34d72960e01b815260040160405180910390fd5b60005b8181101561035f576000838383818110610e0f57610e0f6117d7565b9050602002016020810190610e2491906115b5565b6001600160a01b03811660009081526099602052604090205490915060ff16610e6057604051638065bea960e01b815260040160405180910390fd5b6001600160a01b038116600081815260996020908152604091829020805460ff1916905590519182527f90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a910160405180910390a150610ebe81611803565b9050610df3565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4b9190611860565b60ff16905080610f6957505060408051600081526020810190915290565b6000805b8281101561101e57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610fdc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611000919061181e565b61100a9083611883565b91508061101681611803565b915050610f6d565b5060008167ffffffffffffffff81111561103a5761103a611626565b604051908082528060200260200182016040528015611063578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ec9190611860565b60ff1681101561128557604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015611160573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611184919061181e565b905060005b81811015611270576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156111fe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611222919061189b565b60000151858581518110611238576112386117d7565b6001600160a01b03909216602092830291909101909101528361125a81611803565b945050808061126890611803565b915050611189565b5050808061127d90611803565b91505061106a565b5090949350505050565b6112976113cb565b6001600160a01b0381166112fc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103fd565b61130581611425565b50565b606060008061131684611477565b61ffff1667ffffffffffffffff81111561133257611332611626565b6040519080825280601f01601f19166020018201604052801561135c576020820181803683370190505b5090506000805b825182108015611374575061010081105b15611285576001811b9350858416156113bb578060f81b83838151811061139d5761139d6117d7565b60200101906001600160f81b031916908160001a9053508160010191505b6113c481611803565b9050611363565b6033546001600160a01b03163314610b8a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103fd565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b82156114a25761148c600184611a2e565b909216918061149a81611a45565b91505061147b565b92915050565b60008083601f8401126114ba57600080fd5b50813567ffffffffffffffff8111156114d257600080fd5b6020830191508360208260051b85010111156114ed57600080fd5b9250929050565b6000806020838503121561150757600080fd5b823567ffffffffffffffff81111561151e57600080fd5b61152a858286016114a8565b90969095509350505050565b6001600160a01b038116811461130557600080fd5b6000806000806060858703121561156157600080fd5b843561156c81611536565b93506020850135801515811461158157600080fd5b9250604085013567ffffffffffffffff81111561159d57600080fd5b6115a9878288016114a8565b95989497509550505050565b6000602082840312156115c757600080fd5b81356115d281611536565b9392505050565b6020808252825182820181905260009190848201906040850190845b8181101561161a5783516001600160a01b0316835292840192918401916001016115f5565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561165f5761165f611626565b60405290565b600067ffffffffffffffff8084111561168057611680611626565b604051601f8501601f19908116603f011681019082821181831017156116a8576116a8611626565b816040528093508581528686860111156116c157600080fd5b858560208301376000602087830101525050509392505050565b600080604083850312156116ee57600080fd5b82356116f981611536565b9150602083013567ffffffffffffffff8082111561171657600080fd5b908401906060828703121561172a57600080fd5b61173261163c565b82358281111561174157600080fd5b83019150601f8201871361175457600080fd5b61176387833560208501611665565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561179857600080fd5b813567ffffffffffffffff8111156117af57600080fd5b8201601f810184136117c057600080fd5b6117cf84823560208401611665565b949350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611817576118176117ed565b5060010190565b60006020828403121561183057600080fd5b5051919050565b60006020828403121561184957600080fd5b81516001600160c01b03811681146115d257600080fd5b60006020828403121561187257600080fd5b815160ff811681146115d257600080fd5b60008219821115611896576118966117ed565b500190565b6000604082840312156118ad57600080fd5b6040516040810181811067ffffffffffffffff821117156118d0576118d0611626565b60405282516118de81611536565b815260208301516bffffffffffffffffffffffff811681146118ff57600080fd5b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156119a95760208185018101518683018201520161198d565b818111156119bb576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b03831681526040602082015260008251606060408401526119fa60a0840182611983565b90506020840151606084015260408401516080840152809150509392505050565b6020815260006115d26020830184611983565b600082821015611a4057611a406117ed565b500390565b600061ffff80831681811415611a5d57611a5d6117ed565b600101939250505056fea264697066735822122097cd1941a7b4ba242fa1ab14f81adda9f31fb8c68d9ef9ce4543023c159a0b6064736f6c634300080c0033",
}

// ContractCoprocessorServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractCoprocessorServiceManagerMetaData.ABI instead.
var ContractCoprocessorServiceManagerABI = ContractCoprocessorServiceManagerMetaData.ABI

// ContractCoprocessorServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractCoprocessorServiceManagerMetaData.Bin instead.
var ContractCoprocessorServiceManagerBin = ContractCoprocessorServiceManagerMetaData.Bin

// DeployContractCoprocessorServiceManager deploys a new Ethereum contract, binding an instance of ContractCoprocessorServiceManager to it.
func DeployContractCoprocessorServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address) (common.Address, *types.Transaction, *ContractCoprocessorServiceManager, error) {
	parsed, err := ContractCoprocessorServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractCoprocessorServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractCoprocessorServiceManager{ContractCoprocessorServiceManagerCaller: ContractCoprocessorServiceManagerCaller{contract: contract}, ContractCoprocessorServiceManagerTransactor: ContractCoprocessorServiceManagerTransactor{contract: contract}, ContractCoprocessorServiceManagerFilterer: ContractCoprocessorServiceManagerFilterer{contract: contract}}, nil
}

// ContractCoprocessorServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractCoprocessorServiceManager struct {
	ContractCoprocessorServiceManagerCaller     // Read-only binding to the contract
	ContractCoprocessorServiceManagerTransactor // Write-only binding to the contract
	ContractCoprocessorServiceManagerFilterer   // Log filterer for contract events
}

// ContractCoprocessorServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCoprocessorServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCoprocessorServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractCoprocessorServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCoprocessorServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractCoprocessorServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCoprocessorServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractCoprocessorServiceManagerSession struct {
	Contract     *ContractCoprocessorServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractCoprocessorServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCoprocessorServiceManagerCallerSession struct {
	Contract *ContractCoprocessorServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// ContractCoprocessorServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractCoprocessorServiceManagerTransactorSession struct {
	Contract     *ContractCoprocessorServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// ContractCoprocessorServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractCoprocessorServiceManagerRaw struct {
	Contract *ContractCoprocessorServiceManager // Generic contract binding to access the raw methods on
}

// ContractCoprocessorServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCoprocessorServiceManagerCallerRaw struct {
	Contract *ContractCoprocessorServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractCoprocessorServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractCoprocessorServiceManagerTransactorRaw struct {
	Contract *ContractCoprocessorServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractCoprocessorServiceManager creates a new instance of ContractCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractCoprocessorServiceManager(address common.Address, backend bind.ContractBackend) (*ContractCoprocessorServiceManager, error) {
	contract, err := bindContractCoprocessorServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManager{ContractCoprocessorServiceManagerCaller: ContractCoprocessorServiceManagerCaller{contract: contract}, ContractCoprocessorServiceManagerTransactor: ContractCoprocessorServiceManagerTransactor{contract: contract}, ContractCoprocessorServiceManagerFilterer: ContractCoprocessorServiceManagerFilterer{contract: contract}}, nil
}

// NewContractCoprocessorServiceManagerCaller creates a new read-only instance of ContractCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractCoprocessorServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractCoprocessorServiceManagerCaller, error) {
	contract, err := bindContractCoprocessorServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerCaller{contract: contract}, nil
}

// NewContractCoprocessorServiceManagerTransactor creates a new write-only instance of ContractCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractCoprocessorServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractCoprocessorServiceManagerTransactor, error) {
	contract, err := bindContractCoprocessorServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerTransactor{contract: contract}, nil
}

// NewContractCoprocessorServiceManagerFilterer creates a new log filterer instance of ContractCoprocessorServiceManager, bound to a specific deployed contract.
func NewContractCoprocessorServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractCoprocessorServiceManagerFilterer, error) {
	contract, err := bindContractCoprocessorServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerFilterer{contract: contract}, nil
}

// bindContractCoprocessorServiceManager binds a generic wrapper to an already deployed contract.
func bindContractCoprocessorServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractCoprocessorServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractCoprocessorServiceManager.Contract.ContractCoprocessorServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.ContractCoprocessorServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.ContractCoprocessorServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractCoprocessorServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.AvsDirectory(&_ContractCoprocessorServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.AvsDirectory(&_ContractCoprocessorServiceManager.CallOpts)
}

// CoprocessorTaskManager is a free data retrieval call binding the contract method 0x6052779e.
//
// Solidity: function coprocessorTaskManager() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCaller) CoprocessorTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorServiceManager.contract.Call(opts, &out, "coprocessorTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CoprocessorTaskManager is a free data retrieval call binding the contract method 0x6052779e.
//
// Solidity: function coprocessorTaskManager() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) CoprocessorTaskManager() (common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.CoprocessorTaskManager(&_ContractCoprocessorServiceManager.CallOpts)
}

// CoprocessorTaskManager is a free data retrieval call binding the contract method 0x6052779e.
//
// Solidity: function coprocessorTaskManager() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCallerSession) CoprocessorTaskManager() (common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.CoprocessorTaskManager(&_ContractCoprocessorServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractCoprocessorServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractCoprocessorServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.GetRestakeableStrategies(&_ContractCoprocessorServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.GetRestakeableStrategies(&_ContractCoprocessorServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) Owner() (common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.Owner(&_ContractCoprocessorServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractCoprocessorServiceManager.Contract.Owner(&_ContractCoprocessorServiceManager.CallOpts)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) AddOperatorsToWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "addOperatorsToWhitelist", operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.AddOperatorsToWhitelist(&_ContractCoprocessorServiceManager.TransactOpts, operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.AddOperatorsToWhitelist(&_ContractCoprocessorServiceManager.TransactOpts, operators)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractCoprocessorServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractCoprocessorServiceManager.TransactOpts, operator)
}

// DisableOperatorWhitelist is a paid mutator transaction binding the contract method 0x4f962eca.
//
// Solidity: function disableOperatorWhitelist() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) DisableOperatorWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "disableOperatorWhitelist")
}

// DisableOperatorWhitelist is a paid mutator transaction binding the contract method 0x4f962eca.
//
// Solidity: function disableOperatorWhitelist() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) DisableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.DisableOperatorWhitelist(&_ContractCoprocessorServiceManager.TransactOpts)
}

// DisableOperatorWhitelist is a paid mutator transaction binding the contract method 0x4f962eca.
//
// Solidity: function disableOperatorWhitelist() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) DisableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.DisableOperatorWhitelist(&_ContractCoprocessorServiceManager.TransactOpts)
}

// EnableOperatorWhitelist is a paid mutator transaction binding the contract method 0x47b1b364.
//
// Solidity: function enableOperatorWhitelist() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) EnableOperatorWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "enableOperatorWhitelist")
}

// EnableOperatorWhitelist is a paid mutator transaction binding the contract method 0x47b1b364.
//
// Solidity: function enableOperatorWhitelist() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) EnableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.EnableOperatorWhitelist(&_ContractCoprocessorServiceManager.TransactOpts)
}

// EnableOperatorWhitelist is a paid mutator transaction binding the contract method 0x47b1b364.
//
// Solidity: function enableOperatorWhitelist() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) EnableOperatorWhitelist() (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.EnableOperatorWhitelist(&_ContractCoprocessorServiceManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x25698b30.
//
// Solidity: function initialize(address _lambadaCoprocessorTaskManager, bool _operatorWhitelistEnabled, address[] _operatorWhitelist) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _lambadaCoprocessorTaskManager common.Address, _operatorWhitelistEnabled bool, _operatorWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "initialize", _lambadaCoprocessorTaskManager, _operatorWhitelistEnabled, _operatorWhitelist)
}

// Initialize is a paid mutator transaction binding the contract method 0x25698b30.
//
// Solidity: function initialize(address _lambadaCoprocessorTaskManager, bool _operatorWhitelistEnabled, address[] _operatorWhitelist) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) Initialize(_lambadaCoprocessorTaskManager common.Address, _operatorWhitelistEnabled bool, _operatorWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.Initialize(&_ContractCoprocessorServiceManager.TransactOpts, _lambadaCoprocessorTaskManager, _operatorWhitelistEnabled, _operatorWhitelist)
}

// Initialize is a paid mutator transaction binding the contract method 0x25698b30.
//
// Solidity: function initialize(address _lambadaCoprocessorTaskManager, bool _operatorWhitelistEnabled, address[] _operatorWhitelist) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) Initialize(_lambadaCoprocessorTaskManager common.Address, _operatorWhitelistEnabled bool, _operatorWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.Initialize(&_ContractCoprocessorServiceManager.TransactOpts, _lambadaCoprocessorTaskManager, _operatorWhitelistEnabled, _operatorWhitelist)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.RegisterOperatorToAVS(&_ContractCoprocessorServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.RegisterOperatorToAVS(&_ContractCoprocessorServiceManager.TransactOpts, operator, operatorSignature)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) RemoveOperatorsFromWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "removeOperatorsFromWhitelist", operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.RemoveOperatorsFromWhitelist(&_ContractCoprocessorServiceManager.TransactOpts, operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.RemoveOperatorsFromWhitelist(&_ContractCoprocessorServiceManager.TransactOpts, operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.RenounceOwnership(&_ContractCoprocessorServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.RenounceOwnership(&_ContractCoprocessorServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.TransferOwnership(&_ContractCoprocessorServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.TransferOwnership(&_ContractCoprocessorServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.UpdateAVSMetadataURI(&_ContractCoprocessorServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractCoprocessorServiceManager.Contract.UpdateAVSMetadataURI(&_ContractCoprocessorServiceManager.TransactOpts, _metadataURI)
}

// ContractCoprocessorServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerInitializedIterator struct {
	Event *ContractCoprocessorServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorServiceManagerInitialized)
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
		it.Event = new(ContractCoprocessorServiceManagerInitialized)
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
func (it *ContractCoprocessorServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorServiceManagerInitialized represents a Initialized event raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractCoprocessorServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerInitializedIterator{contract: _ContractCoprocessorServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorServiceManagerInitialized)
				if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractCoprocessorServiceManagerInitialized, error) {
	event := new(ContractCoprocessorServiceManagerInitialized)
	if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorServiceManagerOperatorAddedToWhitelistIterator is returned from FilterOperatorAddedToWhitelist and is used to iterate over the raw logs and unpacked data for OperatorAddedToWhitelist events raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorAddedToWhitelistIterator struct {
	Event *ContractCoprocessorServiceManagerOperatorAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorServiceManagerOperatorAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorServiceManagerOperatorAddedToWhitelist)
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
		it.Event = new(ContractCoprocessorServiceManagerOperatorAddedToWhitelist)
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
func (it *ContractCoprocessorServiceManagerOperatorAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorServiceManagerOperatorAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorServiceManagerOperatorAddedToWhitelist represents a OperatorAddedToWhitelist event raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorAddedToWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToWhitelist is a free log retrieval operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address operator)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) FilterOperatorAddedToWhitelist(opts *bind.FilterOpts) (*ContractCoprocessorServiceManagerOperatorAddedToWhitelistIterator, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerOperatorAddedToWhitelistIterator{contract: _ContractCoprocessorServiceManager.contract, event: "OperatorAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToWhitelist is a free log subscription operation binding the contract event 0x697698203fae7e6d8a36e588dca13624ae9eba99dbec581047633c44c6e1142f.
//
// Solidity: event OperatorAddedToWhitelist(address operator)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) WatchOperatorAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorServiceManagerOperatorAddedToWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorAddedToWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorServiceManagerOperatorAddedToWhitelist)
				if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorAddedToWhitelist", log); err != nil {
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
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) ParseOperatorAddedToWhitelist(log types.Log) (*ContractCoprocessorServiceManagerOperatorAddedToWhitelist, error) {
	event := new(ContractCoprocessorServiceManagerOperatorAddedToWhitelist)
	if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorAddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator is returned from FilterOperatorRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromWhitelist events raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator struct {
	Event *ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist)
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
		it.Event = new(ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist)
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
func (it *ContractCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist represents a OperatorRemovedFromWhitelist event raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address operator)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) FilterOperatorRemovedFromWhitelist(opts *bind.FilterOpts) (*ContractCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerOperatorRemovedFromWhitelistIterator{contract: _ContractCoprocessorServiceManager.contract, event: "OperatorRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromWhitelist is a free log subscription operation binding the contract event 0x90ff506c206ebc7071c8746ca959ea01c75ddbd1464c1263398da71530aaf23a.
//
// Solidity: event OperatorRemovedFromWhitelist(address operator)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) WatchOperatorRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorRemovedFromWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist)
				if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorRemovedFromWhitelist", log); err != nil {
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
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) ParseOperatorRemovedFromWhitelist(log types.Log) (*ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist, error) {
	event := new(ContractCoprocessorServiceManagerOperatorRemovedFromWhitelist)
	if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorRemovedFromWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorServiceManagerOperatorWhitelistDisabledIterator is returned from FilterOperatorWhitelistDisabled and is used to iterate over the raw logs and unpacked data for OperatorWhitelistDisabled events raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorWhitelistDisabledIterator struct {
	Event *ContractCoprocessorServiceManagerOperatorWhitelistDisabled // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorServiceManagerOperatorWhitelistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorServiceManagerOperatorWhitelistDisabled)
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
		it.Event = new(ContractCoprocessorServiceManagerOperatorWhitelistDisabled)
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
func (it *ContractCoprocessorServiceManagerOperatorWhitelistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorServiceManagerOperatorWhitelistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorServiceManagerOperatorWhitelistDisabled represents a OperatorWhitelistDisabled event raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorWhitelistDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistDisabled is a free log retrieval operation binding the contract event 0x87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb2.
//
// Solidity: event OperatorWhitelistDisabled()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) FilterOperatorWhitelistDisabled(opts *bind.FilterOpts) (*ContractCoprocessorServiceManagerOperatorWhitelistDisabledIterator, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorWhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerOperatorWhitelistDisabledIterator{contract: _ContractCoprocessorServiceManager.contract, event: "OperatorWhitelistDisabled", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistDisabled is a free log subscription operation binding the contract event 0x87cee420ec6ba09c53f0070cf9d4a22595fe68850ce2634361c27981db313bb2.
//
// Solidity: event OperatorWhitelistDisabled()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) WatchOperatorWhitelistDisabled(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorServiceManagerOperatorWhitelistDisabled) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorWhitelistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorServiceManagerOperatorWhitelistDisabled)
				if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistDisabled", log); err != nil {
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
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) ParseOperatorWhitelistDisabled(log types.Log) (*ContractCoprocessorServiceManagerOperatorWhitelistDisabled, error) {
	event := new(ContractCoprocessorServiceManagerOperatorWhitelistDisabled)
	if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorServiceManagerOperatorWhitelistEnabledIterator is returned from FilterOperatorWhitelistEnabled and is used to iterate over the raw logs and unpacked data for OperatorWhitelistEnabled events raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorWhitelistEnabledIterator struct {
	Event *ContractCoprocessorServiceManagerOperatorWhitelistEnabled // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorServiceManagerOperatorWhitelistEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorServiceManagerOperatorWhitelistEnabled)
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
		it.Event = new(ContractCoprocessorServiceManagerOperatorWhitelistEnabled)
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
func (it *ContractCoprocessorServiceManagerOperatorWhitelistEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorServiceManagerOperatorWhitelistEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorServiceManagerOperatorWhitelistEnabled represents a OperatorWhitelistEnabled event raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOperatorWhitelistEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistEnabled is a free log retrieval operation binding the contract event 0xa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb.
//
// Solidity: event OperatorWhitelistEnabled()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) FilterOperatorWhitelistEnabled(opts *bind.FilterOpts) (*ContractCoprocessorServiceManagerOperatorWhitelistEnabledIterator, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.FilterLogs(opts, "OperatorWhitelistEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerOperatorWhitelistEnabledIterator{contract: _ContractCoprocessorServiceManager.contract, event: "OperatorWhitelistEnabled", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistEnabled is a free log subscription operation binding the contract event 0xa8044adf496c3243097ba8ce8fb75dd0787f49995031523b39d91cecec396cbb.
//
// Solidity: event OperatorWhitelistEnabled()
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) WatchOperatorWhitelistEnabled(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorServiceManagerOperatorWhitelistEnabled) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorServiceManager.contract.WatchLogs(opts, "OperatorWhitelistEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorServiceManagerOperatorWhitelistEnabled)
				if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistEnabled", log); err != nil {
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
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) ParseOperatorWhitelistEnabled(log types.Log) (*ContractCoprocessorServiceManagerOperatorWhitelistEnabled, error) {
	event := new(ContractCoprocessorServiceManagerOperatorWhitelistEnabled)
	if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OperatorWhitelistEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOwnershipTransferredIterator struct {
	Event *ContractCoprocessorServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractCoprocessorServiceManagerOwnershipTransferred)
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
func (it *ContractCoprocessorServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractCoprocessorServiceManager contract.
type ContractCoprocessorServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractCoprocessorServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractCoprocessorServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorServiceManagerOwnershipTransferredIterator{contract: _ContractCoprocessorServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractCoprocessorServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorServiceManagerOwnershipTransferred)
				if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractCoprocessorServiceManager *ContractCoprocessorServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractCoprocessorServiceManagerOwnershipTransferred, error) {
	event := new(ContractCoprocessorServiceManagerOwnershipTransferred)
	if err := _ContractCoprocessorServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
