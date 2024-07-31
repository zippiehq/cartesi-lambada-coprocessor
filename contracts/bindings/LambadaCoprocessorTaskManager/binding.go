// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractLambadaCoprocessorTaskManager

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// ILambadaCoprocessorTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type ILambadaCoprocessorTaskManagerTask struct {
	ProgramId []byte
	InputHash []byte
}

// ILambadaCoprocessorTaskManagerTaskBatch is an auto generated low-level Go binding around an user-defined struct.
type ILambadaCoprocessorTaskManagerTaskBatch struct {
	Index                     uint32
	BlockNumber               uint32
	MerkeRoot                 [32]byte
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// ILambadaCoprocessorTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type ILambadaCoprocessorTaskManagerTaskResponse struct {
	OutputHash [32]byte
}

// ILambadaCoprocessorTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type ILambadaCoprocessorTaskManagerTaskResponseMetadata struct {
	BatchIndex    uint32
	ProgramId     []byte
	TaskInputHash []byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractLambadaCoprocessorTaskManagerMetaData contains all meta data concerning the ContractLambadaCoprocessorTaskManager contract.
var ContractLambadaCoprocessorTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allBatchHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskOutputs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextBatchIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskOutputHash\",\"inputs\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskInputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextBatchIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerNewTaskBatch\",\"inputs\":[{\"name\":\"batchRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondTask\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.Task\",\"components\":[{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"inputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskBatchRegistered\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structILambadaCoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"responseMeta\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structILambadaCoprocessorTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskInputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"response\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structILambadaCoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005a9738038062005a978339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e051610100516157b5620002e2600039600061025c015260008181610500015261220b0152600081816103cd01526123ed0152600081816103f4015281816125c3015261278501526000818161041b01528181610ffb01528181611ed60152818161206e01526122a801526157b56000f3fe608060405234801561001057600080fd5b50600436106101fb5760003560e01c8063683048351161011a5780638da5cb5b116100ad578063f0f19ee11161007c578063f0f19ee114610522578063f2fde38b14610530578063f8c8765e14610543578063fa36f60c14610556578063fabc1cbc1461056957600080fd5b80638da5cb5b146104bc578063b98d0908146104cd578063cefdc1d4146104da578063df5cf723146104fb57600080fd5b8063715018a6116100e9578063715018a61461047e5780637afa1eed146104865780637e4fa70014610499578063886f1195146104a957600080fd5b806368304835146103ef5780636d14a987146104165780636e1699ae1461043d5780636efb46361461045d57600080fd5b80633e6a9264116101925780635ac86ab7116101615780635ac86ab71461036d5780635c155662146103a05780635c975abb146103c05780635df45946146103c857600080fd5b80633e6a92641461031f578063416c7e5e146103325780634f739f7414610345578063595c6a671461036557600080fd5b80631d2d5d24116101ce5780631d2d5d2414610293578063245a7bfc146102c15780632c3b965f146102ec5780633563b0d1146102ff57600080fd5b806310d67a2f14610200578063136439dd14610215578063171f1d5b146102285780631ad4318914610257575b600080fd5b61021361020e36600461423d565b61057c565b005b61021361022336600461425a565b610638565b61023b6102363660046143d8565b610777565b6040805192151583529015156020830152015b60405180910390f35b61027e7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161024e565b6102b36102a1366004614446565b60ca6020526000908152604090205481565b60405190815260200161024e565b60cc546102d4906001600160a01b031681565b6040516001600160a01b03909116815260200161024e565b6102b36102fa3660046144ab565b610901565b61031261030d36600461452d565b610923565b60405161024e9190614688565b61021361032d3660046149ad565b610db9565b610213610340366004614a88565b610ff9565b610358610353366004614aa5565b61116e565b60405161024e9190614b7a565b610213611894565b61039061037b366004614c44565b606654600160ff9092169190911b9081161490565b604051901515815260200161024e565b6103b36103ae366004614c61565b61195b565b60405161024e9190614d0d565b6066546102b3565b6102d47f000000000000000000000000000000000000000000000000000000000000000081565b6102d47f000000000000000000000000000000000000000000000000000000000000000081565b6102d47f000000000000000000000000000000000000000000000000000000000000000081565b6102b361044b36600461425a565b60cb6020526000908152604090205481565b61047061046b366004614d51565b611b23565b60405161024e929190614e11565b610213612a3a565b60cd546102d4906001600160a01b031681565b60c95461027e9063ffffffff1681565b6065546102d4906001600160a01b031681565b6033546001600160a01b03166102d4565b6097546103909060ff1681565b6104ed6104e8366004614e5a565b612a4e565b60405161024e929190614e9c565b6102d47f000000000000000000000000000000000000000000000000000000000000000081565b60c95463ffffffff1661027e565b61021361053e36600461423d565b612be0565b610213610551366004614ebd565b612c56565b610213610564366004614f19565b612da7565b61021361057736600461425a565b612f2f565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f39190614f74565b6001600160a01b0316336001600160a01b03161461062c5760405162461bcd60e51b815260040161062390614f91565b60405180910390fd5b6106358161308b565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610680573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a49190614fdb565b6106c05760405162461bcd60e51b815260040161062390614ff8565b606654818116146107395760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610623565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106107bf576107bf615040565b60200201518951600160200201518a602001516000600281106107e4576107e4615040565b60200201518b6020015160016002811061080057610800615040565b602090810291909101518c518d83015160405161085d9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108809190615056565b90506108f36108996108928884613182565b8690613219565b6108a16132ad565b6108e96108da856108d4604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613182565b6108e38c61336d565b90613219565b886201d4c06133fd565b909890975095505050505050565b6000806000806109148989898989613621565b9b9a5050505050505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610965573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109899190614f74565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ef9190614f74565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a559190614f74565b9050600086516001600160401b03811115610a7257610a72614273565b604051908082528060200260200182016040528015610aa557816020015b6060815260200190600190039081610a905790505b50905060005b8751811015610dad576000888281518110610ac857610ac8615040565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b29573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b519190810190615078565b905080516001600160401b03811115610b6c57610b6c614273565b604051908082528060200260200182016040528015610bb757816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610b8a5790505b50848481518110610bca57610bca615040565b602002602001018190525060005b8151811015610d97576040518060600160405280876001600160a01b03166347b314e8858581518110610c0d57610c0d615040565b60200260200101516040518263ffffffff1660e01b8152600401610c3391815260200190565b602060405180830381865afa158015610c50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c749190614f74565b6001600160a01b03168152602001838381518110610c9457610c94615040565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610cc257610cc2615040565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d429190615108565b6001600160601b0316815250858581518110610d6057610d60615040565b60200260200101518281518110610d7957610d79615040565b60200260200101819052508080610d8f90615147565b915050610bd8565b5050508080610da590615147565b915050610aab565b50979650505050505050565b60cc546001600160a01b03163314610e135760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610623565b600080610e2288888888613731565b91509150600084604051602001610e3c9135815260200190565b60408051601f1981840301815291905280516020909101209050600080610e8383610e6a60608e018e615162565b8e6020016020810190610e7d9190614446565b8a611b23565b9150915060005b610e9760608d018d615162565b9050811015610f9f57610eb060a08d0160808e01614446565b60ff1683602001518281518110610ec957610ec9615040565b6020026020010151610edb91906151a8565b6001600160601b0316606484600001518381518110610efc57610efc615040565b60200260200101516001600160601b0316610f1791906151d7565b1015610f8d576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610623565b80610f9781615147565b915050610e8a565b50600084815260cb60205260409081902088359055517f1a2d57dfac825bbb1241991344035c9a8a3f0d22b240201d472577d7554834b190610fe49087908a9061528a565b60405180910390a15050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611057573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107b9190614f74565b6001600160a01b0316336001600160a01b0316146111275760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610623565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b6111996040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111fd9190614f74565b905061122a6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061125a908b90899089906004016152ad565b600060405180830381865afa158015611277573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261129f91908101906152f7565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906112d1908b908b908b906004016153ae565b600060405180830381865afa1580156112ee573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261131691908101906152f7565b6040820152856001600160401b0381111561133357611333614273565b60405190808252806020026020018201604052801561136657816020015b60608152602001906001900390816113515790505b50606082015260005b60ff81168711156117a5576000856001600160401b0381111561139457611394614273565b6040519080825280602002602001820160405280156113bd578160200160208202803683370190505b5083606001518360ff16815181106113d7576113d7615040565b602002602001018190525060005b868110156116a55760008c6001600160a01b03166304ec63518a8a8581811061141057611410615040565b905060200201358e8860000151868151811061142e5761142e615040565b60200260200101516040518463ffffffff1660e01b815260040161146b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611488573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ac91906153ce565b90506001600160c01b0381166115505760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610623565b8a8a8560ff1681811061156557611565615040565b6001600160c01b03841692013560f81c9190911c60019081161415905061169257856001600160a01b031663dd9846b98a8a858181106115a7576115a7615040565b905060200201358d8d8860ff168181106115c3576115c3615040565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611619573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163d91906153f7565b85606001518560ff168151811061165657611656615040565b6020026020010151848151811061166f5761166f615040565b63ffffffff909216602092830291909101909101528261168e81615147565b9350505b508061169d81615147565b9150506113e5565b506000816001600160401b038111156116c0576116c0614273565b6040519080825280602002602001820160405280156116e9578160200160208202803683370190505b50905060005b8281101561176a5784606001518460ff168151811061171057611710615040565b6020026020010151818151811061172957611729615040565b602002602001015182828151811061174357611743615040565b63ffffffff909216602092830291909101909101528061176281615147565b9150506116ef565b508084606001518460ff168151811061178557611785615040565b60200260200101819052505050808061179d90615414565b91505061136f565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061180a9190614f74565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061183d908b908b908e90600401615434565b600060405180830381865afa15801561185a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261188291908101906152f7565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156118dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119009190614fdb565b61191c5760405162461bcd60e51b815260040161062390614ff8565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161198d92919061545e565b600060405180830381865afa1580156119aa573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119d291908101906152f7565b9050600084516001600160401b038111156119ef576119ef614273565b604051908082528060200260200182016040528015611a18578160200160208202803683370190505b50905060005b8551811015611b1957866001600160a01b03166304ec6351878381518110611a4857611a48615040565b602002602001015187868581518110611a6357611a63615040565b60200260200101516040518463ffffffff1660e01b8152600401611aa09392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611abd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ae191906153ce565b6001600160c01b0316828281518110611afc57611afc615040565b602090810291909101015280611b1181615147565b915050611a1e565b5095945050505050565b6040805180820190915260608082526020820152600084611b9a5760405162461bcd60e51b8152602060048201526037602482015260008051602061576083398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610623565b60408301515185148015611bb2575060a08301515185145b8015611bc2575060c08301515185145b8015611bd2575060e08301515185145b611c3c5760405162461bcd60e51b8152602060048201526041602482015260008051602061576083398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610623565b82515160208401515114611cb45760405162461bcd60e51b815260206004820152604460248201819052600080516020615760833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610623565b4363ffffffff168463ffffffff1610611d235760405162461bcd60e51b815260206004820152603c602482015260008051602061576083398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610623565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611d6457611d64614273565b604051908082528060200260200182016040528015611d8d578160200160208202803683370190505b506020820152866001600160401b03811115611dab57611dab614273565b604051908082528060200260200182016040528015611dd4578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611e0857611e08614273565b604051908082528060200260200182016040528015611e31578160200160208202803683370190505b5081526020860151516001600160401b03811115611e5157611e51614273565b604051908082528060200260200182016040528015611e7a578160200160208202803683370190505b5081602001819052506000611f4c8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611f23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f4791906154b2565b613965565b905060005b8760200151518110156121e757611f9688602001518281518110611f7757611f77615040565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611fac57611fac615040565b6020908102919091010152801561206c576020830151611fcd6001836154cf565b81518110611fdd57611fdd615040565b602002602001015160001c83602001518281518110611ffe57611ffe615040565b602002602001015160001c1161206c576040805162461bcd60e51b815260206004820152602481019190915260008051602061576083398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610623565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106120b1576120b1615040565b60200260200101518b8b6000015185815181106120d0576120d0615040565b60200260200101516040518463ffffffff1660e01b815260040161210d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561212a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061214e91906153ce565b6001600160c01b03168360000151828151811061216d5761216d615040565b6020026020010181815250506121d36108926121a7848660000151858151811061219957612199615040565b6020026020010151166139f8565b8a6020015184815181106121bd576121bd615040565b6020026020010151613a2390919063ffffffff16565b9450806121df81615147565b915050611f51565b50506121f283613b07565b60975490935060ff1660008161220957600061228b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612267573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061228b91906154e6565b905060005b8a8110156129095782156123eb578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106122e7576122e7615040565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612327573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061234b91906154e6565b61235591906154ff565b116123eb5760405162461bcd60e51b8152602060048201526066602482015260008051602061576083398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610623565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061242c5761242c615040565b9050013560f81c60f81b60f81c8c8c60a00151858151811061245057612450615040565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156124ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124d09190615517565b6001600160401b0319166124f38a604001518381518110611f7757611f77615040565b67ffffffffffffffff19161461258f5760405162461bcd60e51b8152602060048201526061602482015260008051602061576083398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610623565b6125bf896040015182815181106125a8576125a8615040565b60200260200101518761321990919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061260257612602615040565b9050013560f81c60f81b60f81c8c8c60c00151858151811061262657612626615040565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612682573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126a69190615108565b856020015182815181106126bc576126bc615040565b6001600160601b039092166020928302919091018201528501518051829081106126e8576126e8615040565b60200260200101518560000151828151811061270657612706615040565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156128f45761277e8660000151828151811061275057612750615040565b60200260200101518f8f8681811061276a5761276a615040565b600192013560f81c9290921c811614919050565b156128e2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106127c4576127c4615040565b9050013560f81c60f81b60f81c8e896020015185815181106127e8576127e8615040565b60200260200101518f60e00151888151811061280657612806615040565b6020026020010151878151811061281f5761281f615040565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612883573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128a79190615108565b87518051859081106128bb576128bb615040565b602002602001018181516128cf9190615542565b6001600160601b03169052506001909101905b806128ec81615147565b91505061272a565b5050808061290190615147565b915050612290565b5050506000806129238c868a606001518b60800151610777565b91509150816129945760405162461bcd60e51b8152602060048201526043602482015260008051602061576083398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610623565b806129f55760405162461bcd60e51b8152602060048201526039602482015260008051602061576083398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610623565b50506000878260200151604051602001612a1092919061556a565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612a42613ba2565b612a4c6000613bfc565b565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612a8957612a89615040565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612ac5908890869060040161545e565b600060405180830381865afa158015612ae2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612b0a91908101906152f7565b600081518110612b1c57612b1c615040565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612b88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bac91906153ce565b6001600160c01b031690506000612bc282613c4e565b905081612bd08a838a610923565b9550955050505050935093915050565b612be8613ba2565b6001600160a01b038116612c4d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610623565b61063581613bfc565b600054610100900460ff1615808015612c765750600054600160ff909116105b80612c905750303b158015612c90575060005460ff166001145b612cf35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610623565b6000805460ff191660011790558015612d16576000805461ff0019166101001790555b612d21856000613d1a565b612d2a84613bfc565b60cc80546001600160a01b038086166001600160a01b03199283161790925560cd8054928516929091169190911790558015612da0576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60cd546001600160a01b03163314612e0b5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610623565b6040805160a08101825260608082015260c95463ffffffff908116825243811660208084019190915282840188905290861660808301528251601f85018290048202810182019093528383529091908490849081908401838280828437600092018290525060608601949094525050604051612e8c915083906020016155b2565b60408051601f19818403018152918152815160209283012060c9805463ffffffff908116600090815260ca9095529284208290558054919450911691612ed18361560e565b91906101000a81548163ffffffff021916908363ffffffff160217905550507f9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd9582604051612f1f91906155b2565b60405180910390a1505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fa69190614f74565b6001600160a01b0316336001600160a01b031614612fd65760405162461bcd60e51b815260040161062390614f91565b6066541981196066541916146130545760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610623565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161076c565b6001600160a01b0381166131195760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610623565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261319e61414e565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156131d1576131d3565bfe5b50806132115760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610623565b505092915050565b604080518082019091526000808252602082015261323561416c565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156131d15750806132115760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610623565b6132b561418a565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061339d60008051602061574083398151915286615056565b90505b6133a981613e04565b90935091506000805160206157408339815191528283098314156133e3576040805180820190915290815260208101919091529392505050565b6000805160206157408339815191526001820890506133a0565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061342f6141af565b60005b60028110156135f45760006134488260066151d7565b905084826002811061345c5761345c615040565b6020020151518361346e8360006154ff565b600c811061347e5761347e615040565b602002015284826002811061349557613495615040565b602002015160200151838260016134ac91906154ff565b600c81106134bc576134bc615040565b60200201528382600281106134d3576134d3615040565b60200201515151836134e68360026154ff565b600c81106134f6576134f6615040565b602002015283826002811061350d5761350d615040565b60200201515160016020020151836135268360036154ff565b600c811061353657613536615040565b602002015283826002811061354d5761354d615040565b60200201516020015160006002811061356857613568615040565b6020020151836135798360046154ff565b600c811061358957613589615040565b60200201528382600281106135a0576135a0615040565b6020020151602001516001600281106135bb576135bb615040565b6020020151836135cc8360056154ff565b600c81106135dc576135dc615040565b602002015250806135ec81615147565b915050613432565b506135fd6141ce565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b61364b6040518060600160405280600063ffffffff16815260200160608152602001606081525090565b604080516060808201835260008083526020830182905292820152819063ffffffff89168152604080516020601f8a0181900481028201810190925288815290899089908190840183828082843760009201919091525050505060208083019190915260408051601f8801839004830281018301909152868152908790879081908401838280828437600092018290525060408087019590955293516136f993508592506020019050615632565b60408051808303601f190181529181528151602092830120600081815260cb909352912054929b909a50919850909650505050505050565b61375b6040518060600160405280600063ffffffff16815260200160608152602001606081525090565b6000808660405160200161376f9190615645565b6040516020818303038152906040528051906020012090508060ca600089600001602081019061379f9190614446565b63ffffffff1663ffffffff16815260200190815260200160002054146138125760405162461bcd60e51b815260206004820152602260248201527f5461736b20626174636820686173206e6f74206265656e207265676973746572604482015261195960f21b6064820152608401610623565b6000808061384261382660208c018c614446565b6138308b80615162565b61383d60208e018e615162565b613621565b925092509250806000146138985760405162461bcd60e51b815260206004820152601f60248201527f5461736b20726573706f6e736520616c726561647920726573706f6e646564006044820152606401610623565b60006138a48a80615162565b6138b160208d018d615162565b6040516020016138c49493929190615700565b60408051601f198184030181528282528051602091820120908301520160405160208183030381529060405280519060200120905061390989898d6040013584613e86565b6139555760405162461bcd60e51b815260206004820152601d60248201527f5461736b20646f6573206e6f742062656c6f6e6720746f2062617463680000006044820152606401610623565b5091999098509650505050505050565b60008061397184613e9e565b9050808360ff166001901b116139ef5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610623565b90505b92915050565b6000805b82156139f257613a0d6001846154cf565b9092169180613a1b81615727565b9150506139fc565b60408051808201909152600080825260208201526102008261ffff1610613a7f5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610623565b8161ffff1660011415613a935750816139f2565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613afc57600161ffff871660ff83161c81161415613adf57613adc8484613219565b93505b613ae98384613219565b92506201fffe600192831b169101613aaf565b509195945050505050565b60408051808201909152600080825260208201528151158015613b2c57506020820151155b15613b4a575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206157408339815191528460200151613b7d9190615056565b613b95906000805160206157408339815191526154cf565b905292915050565b919050565b6033546001600160a01b03163314612a4c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610623565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613c5c846139f8565b61ffff166001600160401b03811115613c7757613c77614273565b6040519080825280601f01601f191660200182016040528015613ca1576020820181803683370190505b5090506000805b825182108015613cb9575061010081105b15613d10576001811b935085841615613d00578060f81b838381518110613ce257613ce2615040565b60200101906001600160f81b031916908160001a9053508160010191505b613d0981615147565b9050613ca8565b5090949350505050565b6065546001600160a01b0316158015613d3b57506001600160a01b03821615155b613dbd5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610623565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613e008261308b565b5050565b60008080600080516020615740833981519152600360008051602061574083398151915286600080516020615740833981519152888909090890506000613e7a827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061574083398151915261402b565b91959194509092505050565b600082613e948686856140d3565b1495945050505050565b600061010082511115613f275760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610623565b8151613f3557506000919050565b60008083600081518110613f4b57613f4b615040565b0160200151600160f89190911c81901b92505b845181101561402257848181518110613f7957613f79615040565b0160200151600160f89190911c1b915082821161400e5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610623565b9181179161401b81615147565b9050613f5e565b50909392505050565b6000806140366141ce565b61403e6141ec565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156131d15750826140c85760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610623565b505195945050505050565b600081815b8481101561411657614102828787848181106140f6576140f6615040565b9050602002013561411f565b91508061410e81615147565b9150506140d8565b50949350505050565b600081831061413b5760008281526020849052604090206139ef565b60008381526020839052604090206139ef565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061419d61420a565b81526020016141aa61420a565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461063557600080fd5b60006020828403121561424f57600080fd5b81356139ef81614228565b60006020828403121561426c57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156142ab576142ab614273565b60405290565b60405161010081016001600160401b03811182821017156142ab576142ab614273565b604051601f8201601f191681016001600160401b03811182821017156142fc576142fc614273565b604052919050565b60006040828403121561431657600080fd5b61431e614289565b9050813581526020820135602082015292915050565b600082601f83011261434557600080fd5b604051604081018181106001600160401b038211171561436757614367614273565b806040525080604084018581111561437e57600080fd5b845b81811015613afc578035835260209283019201614380565b6000608082840312156143aa57600080fd5b6143b2614289565b90506143be8383614334565b81526143cd8360408401614334565b602082015292915050565b60008060008061012085870312156143ef57600080fd5b843593506144008660208701614304565b925061440f8660608701614398565b915061441e8660e08701614304565b905092959194509250565b63ffffffff8116811461063557600080fd5b8035613b9d81614429565b60006020828403121561445857600080fd5b81356139ef81614429565b60008083601f84011261447557600080fd5b5081356001600160401b0381111561448c57600080fd5b6020830191508360208285010111156144a457600080fd5b9250929050565b6000806000806000606086880312156144c357600080fd5b85356144ce81614429565b945060208601356001600160401b03808211156144ea57600080fd5b6144f689838a01614463565b9096509450604088013591508082111561450f57600080fd5b5061451c88828901614463565b969995985093965092949392505050565b60008060006060848603121561454257600080fd5b833561454d81614228565b92506020848101356001600160401b038082111561456a57600080fd5b818701915087601f83011261457e57600080fd5b81358181111561459057614590614273565b6145a2601f8201601f191685016142d4565b915080825288848285010111156145b857600080fd5b80848401858401376000848284010152508094505050506145db6040850161443b565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b8681101561467a578385038a52825180518087529087019087870190845b8181101561466557835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614621565b50509a87019a95505091850191600101614603565b509298975050505050505050565b6020815260006139ef60208301846145e4565b6000604082840312156146ad57600080fd5b50919050565b60008083601f8401126146c557600080fd5b5081356001600160401b038111156146dc57600080fd5b6020830191508360208260051b85010111156144a457600080fd5b6000602082840312156146ad57600080fd5b60006001600160401b0382111561472257614722614273565b5060051b60200190565b600082601f83011261473d57600080fd5b8135602061475261474d83614709565b6142d4565b82815260059290921b8401810191818101908684111561477157600080fd5b8286015b8481101561479557803561478881614429565b8352918301918301614775565b509695505050505050565b600082601f8301126147b157600080fd5b813560206147c161474d83614709565b82815260069290921b840181019181810190868411156147e057600080fd5b8286015b84811015614795576147f68882614304565b8352918301916040016147e4565b600082601f83011261481557600080fd5b8135602061482561474d83614709565b82815260059290921b8401810191818101908684111561484457600080fd5b8286015b848110156147955780356001600160401b038111156148675760008081fd5b6148758986838b010161472c565b845250918301918301614848565b6000610180828403121561489657600080fd5b61489e6142b1565b905081356001600160401b03808211156148b757600080fd5b6148c38583860161472c565b835260208401359150808211156148d957600080fd5b6148e5858386016147a0565b602084015260408401359150808211156148fe57600080fd5b61490a858386016147a0565b604084015261491c8560608601614398565b606084015261492e8560e08601614304565b608084015261012084013591508082111561494857600080fd5b6149548583860161472c565b60a084015261014084013591508082111561496e57600080fd5b61497a8583860161472c565b60c084015261016084013591508082111561499457600080fd5b506149a184828501614804565b60e08301525092915050565b60008060008060008060a087890312156149c657600080fd5b86356001600160401b03808211156149dd57600080fd5b9088019060a0828b0312156149f157600080fd5b90965060208801359080821115614a0757600080fd5b614a138a838b0161469b565b96506040890135915080821115614a2957600080fd5b614a358a838b016146b3565b9096509450849150614a4a8a60608b016146f7565b93506080890135915080821115614a6057600080fd5b50614a6d89828a01614883565b9150509295509295509295565b801515811461063557600080fd5b600060208284031215614a9a57600080fd5b81356139ef81614a7a565b60008060008060008060808789031215614abe57600080fd5b8635614ac981614228565b95506020870135614ad981614429565b945060408701356001600160401b0380821115614af557600080fd5b614b018a838b01614463565b90965094506060890135915080821115614b1a57600080fd5b50614b2789828a016146b3565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614b6f57815163ffffffff1687529582019590820190600101614b4d565b509495945050505050565b600060208083528351608082850152614b9660a0850182614b39565b905081850151601f1980868403016040870152614bb38383614b39565b92506040870151915080868403016060870152614bd08383614b39565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614c275784878303018452614c15828751614b39565b95880195938801939150600101614bfb565b509998505050505050505050565b60ff8116811461063557600080fd5b600060208284031215614c5657600080fd5b81356139ef81614c35565b600080600060608486031215614c7657600080fd5b8335614c8181614228565b92506020848101356001600160401b03811115614c9d57600080fd5b8501601f81018713614cae57600080fd5b8035614cbc61474d82614709565b81815260059190911b82018301908381019089831115614cdb57600080fd5b928401925b82841015614cf957833582529284019290840190614ce0565b80965050505050506145db6040850161443b565b6020808252825182820181905260009190848201906040850190845b81811015614d4557835183529284019291840191600101614d29565b50909695505050505050565b600080600080600060808688031215614d6957600080fd5b8535945060208601356001600160401b0380821115614d8757600080fd5b614d9389838a01614463565b909650945060408801359150614da882614429565b90925060608701359080821115614dbe57600080fd5b50614dcb88828901614883565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614b6f5781516001600160601b031687529582019590820190600101614dec565b6040815260008351604080840152614e2c6080840182614dd8565b90506020850151603f19848303016060850152614e498282614dd8565b925050508260208301529392505050565b600080600060608486031215614e6f57600080fd5b8335614e7a81614228565b9250602084013591506040840135614e9181614429565b809150509250925092565b828152604060208201526000614eb560408301846145e4565b949350505050565b60008060008060808587031215614ed357600080fd5b8435614ede81614228565b93506020850135614eee81614228565b92506040850135614efe81614228565b91506060850135614f0e81614228565b939692955090935050565b60008060008060608587031215614f2f57600080fd5b843593506020850135614f4181614429565b925060408501356001600160401b03811115614f5c57600080fd5b614f6887828801614463565b95989497509550505050565b600060208284031215614f8657600080fd5b81516139ef81614228565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614fed57600080fd5b81516139ef81614a7a565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261507357634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561508b57600080fd5b82516001600160401b038111156150a157600080fd5b8301601f810185136150b257600080fd5b80516150c061474d82614709565b81815260059190911b820183019083810190878311156150df57600080fd5b928401925b828410156150fd578351825292840192908401906150e4565b979650505050505050565b60006020828403121561511a57600080fd5b81516001600160601b03811681146139ef57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561515b5761515b615131565b5060010190565b6000808335601e1984360301811261517957600080fd5b8301803591506001600160401b0382111561519357600080fd5b6020019150368190038213156144a457600080fd5b60006001600160601b03808316818516818304811182151516156151ce576151ce615131565b02949350505050565b60008160001904831182151516156151f1576151f1615131565b500290565b6000815180845260005b8181101561521c57602081850181015186830182015201615200565b8181111561522e576000602083870101525b50601f01601f19169290920160200192915050565b63ffffffff8151168252600060208201516060602085015261526860608501826151f6565b90506040830151848203604086015261528182826151f6565b95945050505050565b60408152600061529d6040830185615243565b9050823560208301529392505050565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156152da57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561530a57600080fd5b82516001600160401b0381111561532057600080fd5b8301601f8101851361533157600080fd5b805161533f61474d82614709565b81815260059190911b8201830190838101908783111561535e57600080fd5b928401925b828410156150fd57835161537681614429565b82529284019290840190615363565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201526000615281604083018486615385565b6000602082840312156153e057600080fd5b81516001600160c01b03811681146139ef57600080fd5b60006020828403121561540957600080fd5b81516139ef81614429565b600060ff821660ff81141561542b5761542b615131565b60010192915050565b604081526000615448604083018587615385565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156154a557845183529383019391830191600101615489565b5090979650505050505050565b6000602082840312156154c457600080fd5b81516139ef81614c35565b6000828210156154e1576154e1615131565b500390565b6000602082840312156154f857600080fd5b5051919050565b6000821982111561551257615512615131565b500190565b60006020828403121561552957600080fd5b815167ffffffffffffffff19811681146139ef57600080fd5b60006001600160601b038381169083168181101561556257615562615131565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156155a557815185529382019390820190600101615589565b5092979650505050505050565b60208152600063ffffffff80845116602084015280602085015116604084015260408401516060840152606084015160a060808501526155f560c08501826151f6565b90508160808601511660a0850152809250505092915050565b600063ffffffff8083168181141561562857615628615131565b6001019392505050565b6020815260006139ef6020830184615243565b602081526000823561565681614429565b63ffffffff80821660208501526020850135915061567382614429565b80821660408501525050604083013560608301526060830135601e1984360301811261569e57600080fd5b830180356001600160401b038111156156b657600080fd5b8036038513156156c557600080fd5b60a060808501526156dd60c085018260208501615385565b9150506156ec6080850161443b565b63ffffffff811660a0850152509392505050565b604081526000615714604083018688615385565b82810360208401526150fd818587615385565b600061ffff808316818114156156285761562861513156fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220c8960f71c7df0bfa6c91a75b4877b2e579b2cf57e9d04e4c4860fdee7d22fd7164736f6c634300080c0033",
}

// ContractLambadaCoprocessorTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractLambadaCoprocessorTaskManagerMetaData.ABI instead.
var ContractLambadaCoprocessorTaskManagerABI = ContractLambadaCoprocessorTaskManagerMetaData.ABI

// ContractLambadaCoprocessorTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractLambadaCoprocessorTaskManagerMetaData.Bin instead.
var ContractLambadaCoprocessorTaskManagerBin = ContractLambadaCoprocessorTaskManagerMetaData.Bin

// DeployContractLambadaCoprocessorTaskManager deploys a new Ethereum contract, binding an instance of ContractLambadaCoprocessorTaskManager to it.
func DeployContractLambadaCoprocessorTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractLambadaCoprocessorTaskManager, error) {
	parsed, err := ContractLambadaCoprocessorTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractLambadaCoprocessorTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractLambadaCoprocessorTaskManager{ContractLambadaCoprocessorTaskManagerCaller: ContractLambadaCoprocessorTaskManagerCaller{contract: contract}, ContractLambadaCoprocessorTaskManagerTransactor: ContractLambadaCoprocessorTaskManagerTransactor{contract: contract}, ContractLambadaCoprocessorTaskManagerFilterer: ContractLambadaCoprocessorTaskManagerFilterer{contract: contract}}, nil
}

// ContractLambadaCoprocessorTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractLambadaCoprocessorTaskManager struct {
	ContractLambadaCoprocessorTaskManagerCaller     // Read-only binding to the contract
	ContractLambadaCoprocessorTaskManagerTransactor // Write-only binding to the contract
	ContractLambadaCoprocessorTaskManagerFilterer   // Log filterer for contract events
}

// ContractLambadaCoprocessorTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractLambadaCoprocessorTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractLambadaCoprocessorTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractLambadaCoprocessorTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractLambadaCoprocessorTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractLambadaCoprocessorTaskManagerSession struct {
	Contract     *ContractLambadaCoprocessorTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractLambadaCoprocessorTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractLambadaCoprocessorTaskManagerCallerSession struct {
	Contract *ContractLambadaCoprocessorTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// ContractLambadaCoprocessorTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractLambadaCoprocessorTaskManagerTransactorSession struct {
	Contract     *ContractLambadaCoprocessorTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// ContractLambadaCoprocessorTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractLambadaCoprocessorTaskManagerRaw struct {
	Contract *ContractLambadaCoprocessorTaskManager // Generic contract binding to access the raw methods on
}

// ContractLambadaCoprocessorTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorTaskManagerCallerRaw struct {
	Contract *ContractLambadaCoprocessorTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractLambadaCoprocessorTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractLambadaCoprocessorTaskManagerTransactorRaw struct {
	Contract *ContractLambadaCoprocessorTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractLambadaCoprocessorTaskManager creates a new instance of ContractLambadaCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorTaskManager(address common.Address, backend bind.ContractBackend) (*ContractLambadaCoprocessorTaskManager, error) {
	contract, err := bindContractLambadaCoprocessorTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManager{ContractLambadaCoprocessorTaskManagerCaller: ContractLambadaCoprocessorTaskManagerCaller{contract: contract}, ContractLambadaCoprocessorTaskManagerTransactor: ContractLambadaCoprocessorTaskManagerTransactor{contract: contract}, ContractLambadaCoprocessorTaskManagerFilterer: ContractLambadaCoprocessorTaskManagerFilterer{contract: contract}}, nil
}

// NewContractLambadaCoprocessorTaskManagerCaller creates a new read-only instance of ContractLambadaCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractLambadaCoprocessorTaskManagerCaller, error) {
	contract, err := bindContractLambadaCoprocessorTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerCaller{contract: contract}, nil
}

// NewContractLambadaCoprocessorTaskManagerTransactor creates a new write-only instance of ContractLambadaCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractLambadaCoprocessorTaskManagerTransactor, error) {
	contract, err := bindContractLambadaCoprocessorTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerTransactor{contract: contract}, nil
}

// NewContractLambadaCoprocessorTaskManagerFilterer creates a new log filterer instance of ContractLambadaCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractLambadaCoprocessorTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractLambadaCoprocessorTaskManagerFilterer, error) {
	contract, err := bindContractLambadaCoprocessorTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerFilterer{contract: contract}, nil
}

// bindContractLambadaCoprocessorTaskManager binds a generic wrapper to an already deployed contract.
func bindContractLambadaCoprocessorTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractLambadaCoprocessorTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractLambadaCoprocessorTaskManager.Contract.ContractLambadaCoprocessorTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.ContractLambadaCoprocessorTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.ContractLambadaCoprocessorTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractLambadaCoprocessorTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Aggregator(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Aggregator(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// AllBatchHashes is a free data retrieval call binding the contract method 0x1d2d5d24.
//
// Solidity: function allBatchHashes(uint32 ) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) AllBatchHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "allBatchHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllBatchHashes is a free data retrieval call binding the contract method 0x1d2d5d24.
//
// Solidity: function allBatchHashes(uint32 ) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) AllBatchHashes(arg0 uint32) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.AllBatchHashes(&_ContractLambadaCoprocessorTaskManager.CallOpts, arg0)
}

// AllBatchHashes is a free data retrieval call binding the contract method 0x1d2d5d24.
//
// Solidity: function allBatchHashes(uint32 ) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) AllBatchHashes(arg0 uint32) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.AllBatchHashes(&_ContractLambadaCoprocessorTaskManager.CallOpts, arg0)
}

// AllTaskOutputs is a free data retrieval call binding the contract method 0x6e1699ae.
//
// Solidity: function allTaskOutputs(bytes32 ) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) AllTaskOutputs(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "allTaskOutputs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskOutputs is a free data retrieval call binding the contract method 0x6e1699ae.
//
// Solidity: function allTaskOutputs(bytes32 ) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) AllTaskOutputs(arg0 [32]byte) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.AllTaskOutputs(&_ContractLambadaCoprocessorTaskManager.CallOpts, arg0)
}

// AllTaskOutputs is a free data retrieval call binding the contract method 0x6e1699ae.
//
// Solidity: function allTaskOutputs(bytes32 ) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) AllTaskOutputs(arg0 [32]byte) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.AllTaskOutputs(&_ContractLambadaCoprocessorTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.BlsApkRegistry(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.BlsApkRegistry(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.CheckSignatures(&_ContractLambadaCoprocessorTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.CheckSignatures(&_ContractLambadaCoprocessorTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Delegation(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Delegation(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Generator() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Generator(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Generator(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetCheckSignaturesIndices(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetCheckSignaturesIndices(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetNextBatchIndex is a free data retrieval call binding the contract method 0xf0f19ee1.
//
// Solidity: function getNextBatchIndex() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) GetNextBatchIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "getNextBatchIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetNextBatchIndex is a free data retrieval call binding the contract method 0xf0f19ee1.
//
// Solidity: function getNextBatchIndex() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) GetNextBatchIndex() (uint32, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetNextBatchIndex(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// GetNextBatchIndex is a free data retrieval call binding the contract method 0xf0f19ee1.
//
// Solidity: function getNextBatchIndex() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) GetNextBatchIndex() (uint32, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetNextBatchIndex(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetOperatorState(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetOperatorState(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetOperatorState0(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetOperatorState0(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractLambadaCoprocessorTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskOutputHash is a free data retrieval call binding the contract method 0x2c3b965f.
//
// Solidity: function getTaskOutputHash(uint32 batchIndex, bytes programId, bytes taskInputHash) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) GetTaskOutputHash(opts *bind.CallOpts, batchIndex uint32, programId []byte, taskInputHash []byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "getTaskOutputHash", batchIndex, programId, taskInputHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTaskOutputHash is a free data retrieval call binding the contract method 0x2c3b965f.
//
// Solidity: function getTaskOutputHash(uint32 batchIndex, bytes programId, bytes taskInputHash) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) GetTaskOutputHash(batchIndex uint32, programId []byte, taskInputHash []byte) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetTaskOutputHash(&_ContractLambadaCoprocessorTaskManager.CallOpts, batchIndex, programId, taskInputHash)
}

// GetTaskOutputHash is a free data retrieval call binding the contract method 0x2c3b965f.
//
// Solidity: function getTaskOutputHash(uint32 batchIndex, bytes programId, bytes taskInputHash) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) GetTaskOutputHash(batchIndex uint32, programId []byte, taskInputHash []byte) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetTaskOutputHash(&_ContractLambadaCoprocessorTaskManager.CallOpts, batchIndex, programId, taskInputHash)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) NextBatchIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "nextBatchIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) NextBatchIndex() (uint32, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.NextBatchIndex(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) NextBatchIndex() (uint32, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.NextBatchIndex(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Owner() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Owner(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Owner(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Paused(&_ContractLambadaCoprocessorTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Paused(&_ContractLambadaCoprocessorTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Paused0(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Paused0(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.PauserRegistry(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.PauserRegistry(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RegistryCoordinator(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RegistryCoordinator(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.StakeRegistry(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.StakeRegistry(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.StaleStakesForbidden(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.StaleStakesForbidden(&_ContractLambadaCoprocessorTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.TrySignatureAndApkVerification(&_ContractLambadaCoprocessorTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.TrySignatureAndApkVerification(&_ContractLambadaCoprocessorTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Initialize(&_ContractLambadaCoprocessorTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Initialize(&_ContractLambadaCoprocessorTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Pause(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Pause(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.PauseAll(&_ContractLambadaCoprocessorTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.PauseAll(&_ContractLambadaCoprocessorTaskManager.TransactOpts)
}

// RegisterNewTaskBatch is a paid mutator transaction binding the contract method 0xfa36f60c.
//
// Solidity: function registerNewTaskBatch(bytes32 batchRoot, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) RegisterNewTaskBatch(opts *bind.TransactOpts, batchRoot [32]byte, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "registerNewTaskBatch", batchRoot, quorumThresholdPercentage, quorumNumbers)
}

// RegisterNewTaskBatch is a paid mutator transaction binding the contract method 0xfa36f60c.
//
// Solidity: function registerNewTaskBatch(bytes32 batchRoot, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) RegisterNewTaskBatch(batchRoot [32]byte, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RegisterNewTaskBatch(&_ContractLambadaCoprocessorTaskManager.TransactOpts, batchRoot, quorumThresholdPercentage, quorumNumbers)
}

// RegisterNewTaskBatch is a paid mutator transaction binding the contract method 0xfa36f60c.
//
// Solidity: function registerNewTaskBatch(bytes32 batchRoot, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) RegisterNewTaskBatch(batchRoot [32]byte, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RegisterNewTaskBatch(&_ContractLambadaCoprocessorTaskManager.TransactOpts, batchRoot, quorumThresholdPercentage, quorumNumbers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RenounceOwnership(&_ContractLambadaCoprocessorTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RenounceOwnership(&_ContractLambadaCoprocessorTaskManager.TransactOpts)
}

// RespondTask is a paid mutator transaction binding the contract method 0x3e6a9264.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) RespondTask(opts *bind.TransactOpts, batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "respondTask", batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// RespondTask is a paid mutator transaction binding the contract method 0x3e6a9264.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) RespondTask(batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RespondTask(&_ContractLambadaCoprocessorTaskManager.TransactOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// RespondTask is a paid mutator transaction binding the contract method 0x3e6a9264.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) RespondTask(batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RespondTask(&_ContractLambadaCoprocessorTaskManager.TransactOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.SetPauserRegistry(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.SetPauserRegistry(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.SetStaleStakesForbidden(&_ContractLambadaCoprocessorTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.SetStaleStakesForbidden(&_ContractLambadaCoprocessorTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.TransferOwnership(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.TransferOwnership(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Unpause(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.Unpause(&_ContractLambadaCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// ContractLambadaCoprocessorTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerInitializedIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerInitialized)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerInitialized)
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
func (it *ContractLambadaCoprocessorTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerInitialized represents a Initialized event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractLambadaCoprocessorTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerInitializedIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerInitialized)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractLambadaCoprocessorTaskManagerInitialized, error) {
	event := new(ContractLambadaCoprocessorTaskManagerInitialized)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerOwnershipTransferredIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerOwnershipTransferred)
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
func (it *ContractLambadaCoprocessorTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractLambadaCoprocessorTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerOwnershipTransferredIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerOwnershipTransferred)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractLambadaCoprocessorTaskManagerOwnershipTransferred, error) {
	event := new(ContractLambadaCoprocessorTaskManagerOwnershipTransferred)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerPausedIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerPaused)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerPaused)
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
func (it *ContractLambadaCoprocessorTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerPaused represents a Paused event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractLambadaCoprocessorTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerPausedIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerPaused)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParsePaused(log types.Log) (*ContractLambadaCoprocessorTaskManagerPaused, error) {
	event := new(ContractLambadaCoprocessorTaskManagerPaused)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerPauserRegistrySetIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerPauserRegistrySet)
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
func (it *ContractLambadaCoprocessorTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractLambadaCoprocessorTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerPauserRegistrySetIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerPauserRegistrySet)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractLambadaCoprocessorTaskManagerPauserRegistrySet, error) {
	event := new(ContractLambadaCoprocessorTaskManagerPauserRegistrySet)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractLambadaCoprocessorTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorTaskManagerTaskBatchRegisteredIterator is returned from FilterTaskBatchRegistered and is used to iterate over the raw logs and unpacked data for TaskBatchRegistered events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerTaskBatchRegisteredIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerTaskBatchRegistered // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerTaskBatchRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerTaskBatchRegistered)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerTaskBatchRegistered)
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
func (it *ContractLambadaCoprocessorTaskManagerTaskBatchRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerTaskBatchRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerTaskBatchRegistered represents a TaskBatchRegistered event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerTaskBatchRegistered struct {
	Batch ILambadaCoprocessorTaskManagerTaskBatch
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTaskBatchRegistered is a free log retrieval operation binding the contract event 0x9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd95.
//
// Solidity: event TaskBatchRegistered((uint32,uint32,bytes32,bytes,uint32) batch)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterTaskBatchRegistered(opts *bind.FilterOpts) (*ContractLambadaCoprocessorTaskManagerTaskBatchRegisteredIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "TaskBatchRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerTaskBatchRegisteredIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "TaskBatchRegistered", logs: logs, sub: sub}, nil
}

// WatchTaskBatchRegistered is a free log subscription operation binding the contract event 0x9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd95.
//
// Solidity: event TaskBatchRegistered((uint32,uint32,bytes32,bytes,uint32) batch)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchTaskBatchRegistered(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerTaskBatchRegistered) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "TaskBatchRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerTaskBatchRegistered)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "TaskBatchRegistered", log); err != nil {
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

// ParseTaskBatchRegistered is a log parse operation binding the contract event 0x9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd95.
//
// Solidity: event TaskBatchRegistered((uint32,uint32,bytes32,bytes,uint32) batch)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParseTaskBatchRegistered(log types.Log) (*ContractLambadaCoprocessorTaskManagerTaskBatchRegistered, error) {
	event := new(ContractLambadaCoprocessorTaskManagerTaskBatchRegistered)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "TaskBatchRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerTaskRespondedIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerTaskResponded)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerTaskResponded)
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
func (it *ContractLambadaCoprocessorTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerTaskResponded represents a TaskResponded event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerTaskResponded struct {
	ResponseMeta ILambadaCoprocessorTaskManagerTaskResponseMetadata
	Response     ILambadaCoprocessorTaskManagerTaskResponse
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x1a2d57dfac825bbb1241991344035c9a8a3f0d22b240201d472577d7554834b1.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) responseMeta, (bytes32) response)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractLambadaCoprocessorTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerTaskRespondedIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x1a2d57dfac825bbb1241991344035c9a8a3f0d22b240201d472577d7554834b1.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) responseMeta, (bytes32) response)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerTaskResponded)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0x1a2d57dfac825bbb1241991344035c9a8a3f0d22b240201d472577d7554834b1.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) responseMeta, (bytes32) response)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractLambadaCoprocessorTaskManagerTaskResponded, error) {
	event := new(ContractLambadaCoprocessorTaskManagerTaskResponded)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLambadaCoprocessorTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerUnpausedIterator struct {
	Event *ContractLambadaCoprocessorTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractLambadaCoprocessorTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLambadaCoprocessorTaskManagerUnpaused)
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
		it.Event = new(ContractLambadaCoprocessorTaskManagerUnpaused)
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
func (it *ContractLambadaCoprocessorTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLambadaCoprocessorTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLambadaCoprocessorTaskManagerUnpaused represents a Unpaused event raised by the ContractLambadaCoprocessorTaskManager contract.
type ContractLambadaCoprocessorTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractLambadaCoprocessorTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerUnpausedIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractLambadaCoprocessorTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLambadaCoprocessorTaskManagerUnpaused)
				if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractLambadaCoprocessorTaskManagerUnpaused, error) {
	event := new(ContractLambadaCoprocessorTaskManagerUnpaused)
	if err := _ContractLambadaCoprocessorTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
