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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_lambadaCoprocessorTaskManager\",\"type\":\"address\",\"internalType\":\"contractILambadaCoprocessorTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lambadaCoprocessorTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractILambadaCoprocessorTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200164d3803806200164d83398101604081905262000035916200014f565b6001600160a01b0380851660c052808416608052821660a0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e0516113f4620002596000396000818161017d015261069501526000818160ee0152818161078f0152818161086301526108e20152600081816103b401528181610510015281816105a7015281816109df01528181610b630152610c020152600081816101df0152818161026e015281816102ee0152818161073b015281816108070152818161091d0152610abe01526113f46000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639926ee7d116100715780639926ee7d1461013f578063a364f4da14610152578063a98fb35514610165578063cf7e745b14610178578063e481af9d1461019f578063f2fde38b146101a757600080fd5b806333cfb7b7146100ae57806338c8ee64146100d75780636b3aa72e146100ec578063715018a6146101265780638da5cb5b1461012e575b600080fd5b6100c16100bc366004610f0c565b6101ba565b6040516100ce9190610f30565b60405180910390f35b6100ea6100e5366004610f0c565b61068a565b005b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ce565b6100ea61071c565b6033546001600160a01b031661010e565b6100ea61014d366004611032565b610730565b6100ea610160366004610f0c565b6107fc565b6100ea6101733660046110dd565b6108c3565b61010e7f000000000000000000000000000000000000000000000000000000000000000081565b6100c1610917565b6100ea6101b5366004610f0c565b610ce1565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610226573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024a919061112e565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156102b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d99190611147565b90506001600160c01b038116158061037357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036e9190611170565b60ff16155b1561038f57505060408051600081526020810190915292915050565b60006103a3826001600160c01b0316610d57565b90506000805b8251811015610479577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106103f3576103f3611193565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610437573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045b919061112e565b61046590836111bf565b915080610471816111d7565b9150506103a9565b5060008167ffffffffffffffff81111561049557610495610f7d565b6040519080825280602002602001820160405280156104be578160200160208202803683370190505b5090506000805b845181101561067d5760008582815181106104e2576104e2611193565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057b919061112e565b905060005b81811015610667576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156105f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061991906111f2565b6000015186868151811061062f5761062f611193565b6001600160a01b039092166020928302919091019091015284610651816111d7565b955050808061065f906111d7565b915050610580565b5050508080610675906111d7565b9150506104c5565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107195760405162461bcd60e51b815260206004820152602960248201527f6e6f742066726f6d206c616d6261646120636f70726f636573736f7220746173604482015268359036b0b730b3b2b960b91b60648201526084015b60405180910390fd5b50565b610724610e1a565b61072e6000610e74565b565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107785760405162461bcd60e51b815260040161071090611262565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d906107c69085908590600401611327565b600060405180830381600087803b1580156107e057600080fd5b505af11580156107f4573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108445760405162461bcd60e51b815260040161071090611262565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b1580156108a857600080fd5b505af11580156108bc573d6000803e3d6000fd5b5050505050565b6108cb610e1a565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb3559061088e908490600401611372565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610979573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099d9190611170565b60ff169050806109bb57505060408051600081526020810190915290565b6000805b82811015610a7057604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610a2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a52919061112e565b610a5c90836111bf565b915080610a68816111d7565b9150506109bf565b5060008167ffffffffffffffff811115610a8c57610a8c610f7d565b604051908082528060200260200182016040528015610ab5578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3e9190611170565b60ff16811015610cd757604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610bb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bd6919061112e565b905060005b81811015610cc2576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610c50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7491906111f2565b60000151858581518110610c8a57610c8a611193565b6001600160a01b039092166020928302919091019091015283610cac816111d7565b9450508080610cba906111d7565b915050610bdb565b50508080610ccf906111d7565b915050610abc565b5090949350505050565b610ce9610e1a565b6001600160a01b038116610d4e5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610710565b61071981610e74565b6060600080610d6584610ec6565b61ffff1667ffffffffffffffff811115610d8157610d81610f7d565b6040519080825280601f01601f191660200182016040528015610dab576020820181803683370190505b5090506000805b825182108015610dc3575061010081105b15610cd7576001811b935085841615610e0a578060f81b838381518110610dec57610dec611193565b60200101906001600160f81b031916908160001a9053508160010191505b610e13816111d7565b9050610db2565b6033546001600160a01b0316331461072e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610710565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b8215610ef157610edb600184611385565b9092169180610ee98161139c565b915050610eca565b92915050565b6001600160a01b038116811461071957600080fd5b600060208284031215610f1e57600080fd5b8135610f2981610ef7565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610f715783516001600160a01b031683529284019291840191600101610f4c565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610fb657610fb6610f7d565b60405290565b600067ffffffffffffffff80841115610fd757610fd7610f7d565b604051601f8501601f19908116603f01168101908282118183101715610fff57610fff610f7d565b8160405280935085815286868601111561101857600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561104557600080fd5b823561105081610ef7565b9150602083013567ffffffffffffffff8082111561106d57600080fd5b908401906060828703121561108157600080fd5b611089610f93565b82358281111561109857600080fd5b83019150601f820187136110ab57600080fd5b6110ba87833560208501610fbc565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156110ef57600080fd5b813567ffffffffffffffff81111561110657600080fd5b8201601f8101841361111757600080fd5b61112684823560208401610fbc565b949350505050565b60006020828403121561114057600080fd5b5051919050565b60006020828403121561115957600080fd5b81516001600160c01b0381168114610f2957600080fd5b60006020828403121561118257600080fd5b815160ff81168114610f2957600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156111d2576111d26111a9565b500190565b60006000198214156111eb576111eb6111a9565b5060010190565b60006040828403121561120457600080fd5b6040516040810181811067ffffffffffffffff8211171561122757611227610f7d565b604052825161123581610ef7565b815260208301516bffffffffffffffffffffffff8116811461125657600080fd5b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b81811015611300576020818501810151868301820152016112e4565b81811115611312576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261135160a08401826112da565b90506020840151606084015260408401516080840152809150509392505050565b602081526000610f2960208301846112da565b600082821015611397576113976111a9565b500390565b600061ffff808316818114156113b4576113b46111a9565b600101939250505056fea26469706673582212206f66da1409dd2efa5a3decd8c9d32bdd2ab3584edfe53862df0c493731ef102664736f6c634300080c0033",
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
