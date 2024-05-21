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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractLambadaCoprocessorServiceManagerMetaData contains all meta data concerning the ContractLambadaCoprocessorServiceManager contract.
var ContractLambadaCoprocessorServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_lambadaCoprocessorTaskManager\",\"type\":\"address\",\"internalType\":\"contractILambadaCoprocessorTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lambadaCoprocessorTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILambadaCoprocessorTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b3838038062001b3883398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118c1620002776000396000818161019b015261093b01526000818161065a015281816107b60152818161084d01528181610c4f01528181610dd30152610e720152600081816104850152818161051401528181610594015281816109e101528181610a7701528181610b8d0152610d2e01526000818161031501526103f301526000818161010c01528181610a3501528181610ad30152610b5201526118c16000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639926ee7d116100715780639926ee7d1461015d578063a364f4da14610170578063a98fb35514610183578063cf7e745b14610196578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a6146101445780638da5cb5b1461014c575b600080fd5b6100cc6100c7366004611167565b6101d8565b005b6100e16100dc3660046111f1565b610460565b6040516100ee9190611215565b60405180910390f35b6100cc6101053660046111f1565b610930565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ee565b6100cc6109c2565b6033546001600160a01b031661012c565b6100cc61016b366004611317565b6109d6565b6100cc61017e3660046111f1565b610a6c565b6100cc6101913660046113c2565b610b33565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6100e1610b87565b6100cc6101d33660046111f1565b610f51565b6101e0610fc7565b60005b818110156103db578282828181106101fd576101fd611413565b905060200281019061020f9190611429565b6102209060408101906020016111f1565b6001600160a01b03166323b872dd333086868681811061024257610242611413565b90506020028101906102549190611429565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf9190611459565b508282828181106102e2576102e2611413565b90506020028101906102f49190611429565b6103059060408101906020016111f1565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061034657610346611413565b90506020028101906103589190611429565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca9190611459565b506103d481611491565b90506101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042a9085908590600401611545565b600060405180830381600087803b15801561044457600080fd5b505af1158015610458573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f09190611653565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f919061166c565b90506001600160c01b038116158061061957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106149190611695565b60ff16155b1561063557505060408051600081526020810190915292915050565b6000610649826001600160c01b0316611021565b90506000805b825181101561071f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061069957610699611413565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107019190611653565b61070b90836116b8565b91508061071781611491565b91505061064f565b5060008167ffffffffffffffff81111561073b5761073b611262565b604051908082528060200260200182016040528015610764578160200160208202803683370190505b5090506000805b845181101561092357600085828151811061078857610788611413565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108219190611653565b905060005b8181101561090d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906116d0565b600001518686815181106108d5576108d5611413565b6001600160a01b0390921660209283029190910190910152846108f781611491565b955050808061090590611491565b915050610826565b505050808061091b90611491565b91505061076b565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109bf5760405162461bcd60e51b815260206004820152602960248201527f6e6f742066726f6d206c616d6261646120636f70726f636573736f7220746173604482015268359036b0b730b3b2b960b91b60648201526084015b60405180910390fd5b50565b6109ca610fc7565b6109d460006110e4565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a1e5760405162461bcd60e51b81526004016109b69061172f565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042a90859085906004016117f4565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ab45760405162461bcd60e51b81526004016109b69061172f565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b1857600080fd5b505af1158015610b2c573d6000803e3d6000fd5b5050505050565b610b3b610fc7565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610afe90849060040161183f565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610be9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0d9190611695565b60ff16905080610c2b57505060408051600081526020810190915290565b6000805b82811015610ce057604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610c9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc29190611653565b610ccc90836116b8565b915080610cd881611491565b915050610c2f565b5060008167ffffffffffffffff811115610cfc57610cfc611262565b604051908082528060200260200182016040528015610d25578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dae9190611695565b60ff16811015610f4757604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e469190611653565b905060005b81811015610f32576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610ec0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee491906116d0565b60000151858581518110610efa57610efa611413565b6001600160a01b039092166020928302919091019091015283610f1c81611491565b9450508080610f2a90611491565b915050610e4b565b50508080610f3f90611491565b915050610d2c565b5090949350505050565b610f59610fc7565b6001600160a01b038116610fbe5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109b6565b6109bf816110e4565b6033546001600160a01b031633146109d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109b6565b606060008061102f84611136565b61ffff1667ffffffffffffffff81111561104b5761104b611262565b6040519080825280601f01601f191660200182016040528015611075576020820181803683370190505b5090506000805b82518210801561108d575061010081105b15610f47576001811b9350858416156110d4578060f81b8383815181106110b6576110b6611413565b60200101906001600160f81b031916908160001a9053508160010191505b6110dd81611491565b905061107c565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b82156111615761114b600184611852565b909216918061115981611869565b91505061113a565b92915050565b6000806020838503121561117a57600080fd5b823567ffffffffffffffff8082111561119257600080fd5b818501915085601f8301126111a657600080fd5b8135818111156111b557600080fd5b8660208260051b85010111156111ca57600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109bf57600080fd5b60006020828403121561120357600080fd5b813561120e816111dc565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156112565783516001600160a01b031683529284019291840191600101611231565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561129b5761129b611262565b60405290565b600067ffffffffffffffff808411156112bc576112bc611262565b604051601f8501601f19908116603f011681019082821181831017156112e4576112e4611262565b816040528093508581528686860111156112fd57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561132a57600080fd5b8235611335816111dc565b9150602083013567ffffffffffffffff8082111561135257600080fd5b908401906060828703121561136657600080fd5b61136e611278565b82358281111561137d57600080fd5b83019150601f8201871361139057600080fd5b61139f878335602085016112a1565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113d457600080fd5b813567ffffffffffffffff8111156113eb57600080fd5b8201601f810184136113fc57600080fd5b61140b848235602084016112a1565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261143f57600080fd5b9190910192915050565b8035611454816111dc565b919050565b60006020828403121561146b57600080fd5b8151801515811461120e57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114a5576114a561147b565b5060010190565b6bffffffffffffffffffffffff811681146109bf57600080fd5b8183526000602080850194508260005b858110156115265781356114e9816111dc565b6001600160a01b0316875281830135611501816114ac565b6bffffffffffffffffffffffff168784015260409687019691909101906001016114d6565b509495945050505050565b803563ffffffff8116811461145457600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561164557878303603f190184528135368b9003609e1901811261158a57600080fd5b8a0160a0813536839003601e190181126115a357600080fd5b8201803567ffffffffffffffff8111156115bc57600080fd5b8060061b36038413156115ce57600080fd5b8287526115e0838801828c85016114c6565b925050506115ef888301611449565b6001600160a01b03168886015281870135878601526060611611818401611531565b63ffffffff16908601526080611628838201611531565b63ffffffff16950194909452509285019290850190600101611564565b509098975050505050505050565b60006020828403121561166557600080fd5b5051919050565b60006020828403121561167e57600080fd5b81516001600160c01b038116811461120e57600080fd5b6000602082840312156116a757600080fd5b815160ff8116811461120e57600080fd5b600082198211156116cb576116cb61147b565b500190565b6000604082840312156116e257600080fd5b6040516040810181811067ffffffffffffffff8211171561170557611705611262565b6040528251611713816111dc565b81526020830151611723816114ac565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117cd576020818501810151868301820152016117b1565b818111156117df576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261181e60a08401826117a7565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061120e60208301846117a7565b6000828210156118645761186461147b565b500390565b600061ffff808316818114156118815761188161147b565b600101939250505056fea264697066735822122047e5684d261b50c3b03c4b7b4accd3726e21aa77a58d3bfd1b1f88c982b717a364736f6c634300080c0033",
}

// ContractLambadaCoprocessorServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractLambadaCoprocessorServiceManagerMetaData.ABI instead.
var ContractLambadaCoprocessorServiceManagerABI = ContractLambadaCoprocessorServiceManagerMetaData.ABI

// ContractLambadaCoprocessorServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractLambadaCoprocessorServiceManagerMetaData.Bin instead.
var ContractLambadaCoprocessorServiceManagerBin = ContractLambadaCoprocessorServiceManagerMetaData.Bin

// DeployContractLambadaCoprocessorServiceManager deploys a new Ethereum contract, binding an instance of ContractLambadaCoprocessorServiceManager to it.
func DeployContractLambadaCoprocessorServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _lambadaCoprocessorTaskManager common.Address) (common.Address, *types.Transaction, *ContractLambadaCoprocessorServiceManager, error) {
	parsed, err := ContractLambadaCoprocessorServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractLambadaCoprocessorServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _lambadaCoprocessorTaskManager)
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

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.FreezeOperator(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.FreezeOperator(&_ContractLambadaCoprocessorServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.PayForRange(&_ContractLambadaCoprocessorServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractLambadaCoprocessorServiceManager *ContractLambadaCoprocessorServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorServiceManager.Contract.PayForRange(&_ContractLambadaCoprocessorServiceManager.TransactOpts, rangePayments)
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
