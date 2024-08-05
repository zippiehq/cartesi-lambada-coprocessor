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
	ResultCID  []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allBatchHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskOutputs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkValidTaskResponse\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.Task\",\"components\":[{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"inputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"resultCID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextBatchIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseHash\",\"inputs\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskInputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextBatchIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerNewTaskBatch\",\"inputs\":[{\"name\":\"batchRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondTask\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.Task\",\"components\":[{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"inputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structILambadaCoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"resultCID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskBatchRegistered\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structILambadaCoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"responseMeta\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structILambadaCoprocessorTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskInputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"response\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structILambadaCoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"resultCID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005b6b38038062005b6b8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615889620002e2600039600061027a01526000818161050b015261201c0152600081816103c501526121fe0152600081816103ec015281816123d4015261259601526000818161041301528181610e0c01528181611ce701528181611e7f01526120b901526158896000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c80636e1699ae1161011a578063bda4422a116100ad578063f2fde38b1161007c578063f2fde38b1461053b578063f8ad09531461054e578063f8c8765e14610561578063fa36f60c14610574578063fabc1cbc1461058757600080fd5b8063bda4422a146104d2578063cefdc1d4146104e5578063df5cf72314610506578063f0f19ee11461052d57600080fd5b80637e4fa700116100e95780637e4fa70014610491578063886f1195146104a15780638da5cb5b146104b4578063b98d0908146104c557600080fd5b80636e1699ae146104355780636efb463614610455578063715018a6146104765780637afa1eed1461047e57600080fd5b8063416c7e5e1161019d5780635c1556621161016c5780635c155662146103985780635c975abb146103b85780635df45946146103c057806368304835146103e75780636d14a9871461040e57600080fd5b8063416c7e5e1461032a5780634f739f741461033d578063595c6a671461035d5780635ac86ab71461036557600080fd5b80631ad43189116101d95780631ad43189146102755780631d2d5d24146102b1578063245a7bfc146102df5780633563b0d11461030a57600080fd5b806302f0a1dc1461020b57806310d67a2f14610220578063136439dd14610233578063171f1d5b14610246575b600080fd5b61021e610219366004614735565b61059a565b005b61021e61022e36600461481f565b6105ef565b61021e61024136600461483c565b6106ab565b610259610254366004614855565b6107ea565b6040805192151583529015156020830152015b60405180910390f35b61029c7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026c565b6102d16102bf3660046148a6565b60ca6020526000908152604090205481565b60405190815260200161026c565b60cc546102f2906001600160a01b031681565b6040516001600160a01b03909116815260200161026c565b61031d6103183660046148c3565b610974565b60405161026c9190614a1e565b61021e610338366004614a3f565b610e0a565b61035061034b366004614a9d565b610f7f565b60405161026c9190614b72565b61021e6116a5565b610388610373366004614c3c565b606654600160ff9092169190911b9081161490565b604051901515815260200161026c565b6103ab6103a6366004614c59565b61176c565b60405161026c9190614d05565b6066546102d1565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b6102d161044336600461483c565b60cb6020526000908152604090205481565b610468610463366004614d49565b611934565b60405161026c929190614e09565b61021e61284b565b60cd546102f2906001600160a01b031681565b60c95461029c9063ffffffff1681565b6065546102f2906001600160a01b031681565b6033546001600160a01b03166102f2565b6097546103889060ff1681565b61021e6104e0366004614735565b61285f565b6104f86104f3366004614e52565b612958565b60405161026c929190614e94565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b60c95463ffffffff1661029c565b61021e61054936600461481f565b612aea565b6102d161055c366004614eb5565b612b60565b61021e61056f366004614f37565b612b82565b61021e610582366004614f93565b612cd3565b61021e61059536600461483c565b612e5b565b6000806105a988888888612fb7565b915091506000846040516020016105c09190615091565b6040516020818303038152906040528051906020012090506105e38982866131eb565b5050505b505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610642573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066691906150a4565b6001600160a01b0316336001600160a01b03161461069f5760405162461bcd60e51b8152600401610696906150c1565b60405180910390fd5b6106a881613331565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610717919061510b565b6107335760405162461bcd60e51b815260040161069690615128565b606654818116146107ac5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610696565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061083257610832615170565b60200201518951600160200201518a6020015160006002811061085757610857615170565b60200201518b6020015160016002811061087357610873615170565b602090810291909101518c518d8301516040516108d09a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108f39190615186565b905061096661090c6109058884613428565b86906134bf565b610914613553565b61095c61094d85610947604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613428565b6109568c613613565b906134bf565b886201d4c06136a3565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109da91906150a4565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4091906150a4565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa691906150a4565b9050600086516001600160401b03811115610ac357610ac361430f565b604051908082528060200260200182016040528015610af657816020015b6060815260200190600190039081610ae15790505b50905060005b8751811015610dfe576000888281518110610b1957610b19615170565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b7a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ba291908101906151a8565b905080516001600160401b03811115610bbd57610bbd61430f565b604051908082528060200260200182016040528015610c0857816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bdb5790505b50848481518110610c1b57610c1b615170565b602002602001018190525060005b8151811015610de8576040518060600160405280876001600160a01b03166347b314e8858581518110610c5e57610c5e615170565b60200260200101516040518263ffffffff1660e01b8152600401610c8491815260200190565b602060405180830381865afa158015610ca1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc591906150a4565b6001600160a01b03168152602001838381518110610ce557610ce5615170565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610d1357610d13615170565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d6f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d939190615238565b6001600160601b0316815250858581518110610db157610db1615170565b60200260200101518281518110610dca57610dca615170565b60200260200101819052508080610de090615277565b915050610c29565b5050508080610df690615277565b915050610afc565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8c91906150a4565b6001600160a01b0316336001600160a01b031614610f385760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610696565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610faa6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100e91906150a4565b905061103b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061106b908b9089908990600401615292565b600060405180830381865afa158015611088573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110b091908101906152dc565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906110e2908b908b908b9060040161536a565b600060405180830381865afa1580156110ff573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261112791908101906152dc565b6040820152856001600160401b038111156111445761114461430f565b60405190808252806020026020018201604052801561117757816020015b60608152602001906001900390816111625790505b50606082015260005b60ff81168711156115b6576000856001600160401b038111156111a5576111a561430f565b6040519080825280602002602001820160405280156111ce578160200160208202803683370190505b5083606001518360ff16815181106111e8576111e8615170565b602002602001018190525060005b868110156114b65760008c6001600160a01b03166304ec63518a8a8581811061122157611221615170565b905060200201358e8860000151868151811061123f5761123f615170565b60200260200101516040518463ffffffff1660e01b815260040161127c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611299573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bd9190615393565b90506001600160c01b0381166113615760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610696565b8a8a8560ff1681811061137657611376615170565b6001600160c01b03841692013560f81c9190911c6001908116141590506114a357856001600160a01b031663dd9846b98a8a858181106113b8576113b8615170565b905060200201358d8d8860ff168181106113d4576113d4615170565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561142a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144e91906153bc565b85606001518560ff168151811061146757611467615170565b6020026020010151848151811061148057611480615170565b63ffffffff909216602092830291909101909101528261149f81615277565b9350505b50806114ae81615277565b9150506111f6565b506000816001600160401b038111156114d1576114d161430f565b6040519080825280602002602001820160405280156114fa578160200160208202803683370190505b50905060005b8281101561157b5784606001518460ff168151811061152157611521615170565b6020026020010151818151811061153a5761153a615170565b602002602001015182828151811061155457611554615170565b63ffffffff909216602092830291909101909101528061157381615277565b915050611500565b508084606001518460ff168151811061159657611596615170565b6020026020010181905250505080806115ae906153d9565b915050611180565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161b91906150a4565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061164e908b908b908e906004016153f9565b600060405180830381865afa15801561166b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261169391908101906152dc565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156116ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611711919061510b565b61172d5760405162461bcd60e51b815260040161069690615128565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161179e929190615423565b600060405180830381865afa1580156117bb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117e391908101906152dc565b9050600084516001600160401b038111156118005761180061430f565b604051908082528060200260200182016040528015611829578160200160208202803683370190505b50905060005b855181101561192a57866001600160a01b03166304ec635187838151811061185957611859615170565b60200260200101518786858151811061187457611874615170565b60200260200101516040518463ffffffff1660e01b81526004016118b19392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156118ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f29190615393565b6001600160c01b031682828151811061190d5761190d615170565b60209081029190910101528061192281615277565b91505061182f565b5095945050505050565b60408051808201909152606080825260208201526000846119ab5760405162461bcd60e51b8152602060048201526037602482015260008051602061583483398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610696565b604083015151851480156119c3575060a08301515185145b80156119d3575060c08301515185145b80156119e3575060e08301515185145b611a4d5760405162461bcd60e51b8152602060048201526041602482015260008051602061583483398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610696565b82515160208401515114611ac55760405162461bcd60e51b815260206004820152604460248201819052600080516020615834833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610696565b4363ffffffff168463ffffffff1610611b345760405162461bcd60e51b815260206004820152603c602482015260008051602061583483398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610696565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611b7557611b7561430f565b604051908082528060200260200182016040528015611b9e578160200160208202803683370190505b506020820152866001600160401b03811115611bbc57611bbc61430f565b604051908082528060200260200182016040528015611be5578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611c1957611c1961430f565b604051908082528060200260200182016040528015611c42578160200160208202803683370190505b5081526020860151516001600160401b03811115611c6257611c6261430f565b604051908082528060200260200182016040528015611c8b578160200160208202803683370190505b5081602001819052506000611d5d8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611d34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d589190615477565b6138c7565b905060005b876020015151811015611ff857611da788602001518281518110611d8857611d88615170565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611dbd57611dbd615170565b60209081029190910101528015611e7d576020830151611dde600183615494565b81518110611dee57611dee615170565b602002602001015160001c83602001518281518110611e0f57611e0f615170565b602002602001015160001c11611e7d576040805162461bcd60e51b815260206004820152602481019190915260008051602061583483398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610696565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611ec257611ec2615170565b60200260200101518b8b600001518581518110611ee157611ee1615170565b60200260200101516040518463ffffffff1660e01b8152600401611f1e9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611f3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f5f9190615393565b6001600160c01b031683600001518281518110611f7e57611f7e615170565b602002602001018181525050611fe4610905611fb88486600001518581518110611faa57611faa615170565b60200260200101511661395a565b8a602001518481518110611fce57611fce615170565b602002602001015161398590919063ffffffff16565b945080611ff081615277565b915050611d62565b505061200383613a69565b60975490935060ff1660008161201a57600061209c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612078573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209c91906154ab565b905060005b8a81101561271a5782156121fc578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106120f8576120f8615170565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612138573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061215c91906154ab565b61216691906154c4565b116121fc5760405162461bcd60e51b8152602060048201526066602482015260008051602061583483398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610696565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061223d5761223d615170565b9050013560f81c60f81b60f81c8c8c60a00151858151811061226157612261615170565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156122bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122e191906154dc565b6001600160401b0319166123048a604001518381518110611d8857611d88615170565b67ffffffffffffffff1916146123a05760405162461bcd60e51b8152602060048201526061602482015260008051602061583483398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610696565b6123d0896040015182815181106123b9576123b9615170565b6020026020010151876134bf90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d8481811061241357612413615170565b9050013560f81c60f81b60f81c8c8c60c00151858151811061243757612437615170565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612493573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124b79190615238565b856020015182815181106124cd576124cd615170565b6001600160601b039092166020928302919091018201528501518051829081106124f9576124f9615170565b60200260200101518560000151828151811061251757612517615170565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156127055761258f8660000151828151811061256157612561615170565b60200260200101518f8f8681811061257b5761257b615170565b600192013560f81c9290921c811614919050565b156126f3577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106125d5576125d5615170565b9050013560f81c60f81b60f81c8e896020015185815181106125f9576125f9615170565b60200260200101518f60e00151888151811061261757612617615170565b6020026020010151878151811061263057612630615170565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612694573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126b89190615238565b87518051859081106126cc576126cc615170565b602002602001018181516126e09190615507565b6001600160601b03169052506001909101905b806126fd81615277565b91505061253b565b5050808061271290615277565b9150506120a1565b5050506000806127348c868a606001518b608001516107ea565b91509150816127a55760405162461bcd60e51b8152602060048201526043602482015260008051602061583483398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610696565b806128065760405162461bcd60e51b8152602060048201526039602482015260008051602061583483398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610696565b5050600087826020015160405160200161282192919061552f565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612853613b04565b61285d6000613b5e565b565b60cc546001600160a01b031633146128b95760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610696565b6000806128c888888888612fb7565b915091506000846040516020016128df9190615091565b6040516020818303038152906040528051906020012090506129028982866131eb565b600082815260cb602052604090819020829055517fde0a98210b193f0e75b97832743430e00f558966c6917131cee6e590de548c4d906129459085908890615602565b60405180910390a1505050505050505050565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061299357612993615170565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906129cf9088908690600401615423565b600060405180830381865afa1580156129ec573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612a1491908101906152dc565b600081518110612a2657612a26615170565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612a92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ab69190615393565b6001600160c01b031690506000612acc82613bb0565b905081612ada8a838a610974565b9550955050505050935093915050565b612af2613b04565b6001600160a01b038116612b575760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610696565b6106a881613b5e565b600080600080612b738989898989613c7c565b9b9a5050505050505050505050565b600054610100900460ff1615808015612ba25750600054600160ff909116105b80612bbc5750303b158015612bbc575060005460ff166001145b612c1f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610696565b6000805460ff191660011790558015612c42576000805461ff0019166101001790555b612c4d856000613d8c565b612c5684613b5e565b60cc80546001600160a01b038086166001600160a01b03199283161790925560cd8054928516929091169190911790558015612ccc576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60cd546001600160a01b03163314612d375760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610696565b6040805160a08101825260608082015260c95463ffffffff908116825243811660208084019190915282840188905290861660808301528251601f85018290048202810182019093528383529091908490849081908401838280828437600092018290525060608601949094525050604051612db891508390602001615627565b60408051601f19818403018152918152815160209283012060c9805463ffffffff908116600090815260ca9095529284208290558054919450911691612dfd83615683565b91906101000a81548163ffffffff021916908363ffffffff160217905550507f9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd9582604051612e4b9190615627565b60405180910390a1505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612eae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ed291906150a4565b6001600160a01b0316336001600160a01b031614612f025760405162461bcd60e51b8152600401610696906150c1565b606654198119606654191614612f805760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610696565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107df565b612fe16040518060600160405280600063ffffffff16815260200160608152602001606081525090565b60008086604051602001612ff591906156a7565b6040516020818303038152906040528051906020012090508060ca600089600001602081019061302591906148a6565b63ffffffff1663ffffffff16815260200190815260200160002054146130985760405162461bcd60e51b815260206004820152602260248201527f5461736b20626174636820686173206e6f74206265656e207265676973746572604482015261195960f21b6064820152608401610696565b600080806130c86130ac60208c018c6148a6565b6130b68b8061572d565b6130c360208e018e61572d565b613c7c565b9250925092508060001461311e5760405162461bcd60e51b815260206004820152601f60248201527f5461736b20726573706f6e736520616c726561647920726573706f6e646564006044820152606401610696565b600061312a8a8061572d565b61313760208d018d61572d565b60405160200161314a9493929190615773565b60408051601f198184030181528282528051602091820120908301520160405160208183030381529060405280519060200120905061318f89898d6040013584613e76565b6131db5760405162461bcd60e51b815260206004820152601d60248201527f5461736b20646f6573206e6f742062656c6f6e6720746f2062617463680000006044820152606401610696565b5091999098509650505050505050565b600080613215846131ff606088018861572d565b61320f60408a0160208b016148a6565b87611934565b9150915060005b613229606087018761572d565b90508110156105e75761324260a08701608088016148a6565b60ff168360200151828151811061325b5761325b615170565b602002602001015161326d919061579a565b6001600160601b031660648460000151838151811061328e5761328e615170565b60200260200101516001600160601b03166132a991906157c9565b101561331f576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610696565b8061332981615277565b91505061321c565b6001600160a01b0381166133bf5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610696565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526134446141c0565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561347757613479565bfe5b50806134b75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610696565b505092915050565b60408051808201909152600080825260208201526134db6141de565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156134775750806134b75760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610696565b61355b6141fc565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061364360008051602061581483398151915286615186565b90505b61364f81613e8e565b9093509150600080516020615814833981519152828309831415613689576040805180820190915290815260208101919091529392505050565b600080516020615814833981519152600182089050613646565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906136d5614221565b60005b600281101561389a5760006136ee8260066157c9565b905084826002811061370257613702615170565b602002015151836137148360006154c4565b600c811061372457613724615170565b602002015284826002811061373b5761373b615170565b6020020151602001518382600161375291906154c4565b600c811061376257613762615170565b602002015283826002811061377957613779615170565b602002015151518361378c8360026154c4565b600c811061379c5761379c615170565b60200201528382600281106137b3576137b3615170565b60200201515160016020020151836137cc8360036154c4565b600c81106137dc576137dc615170565b60200201528382600281106137f3576137f3615170565b60200201516020015160006002811061380e5761380e615170565b60200201518361381f8360046154c4565b600c811061382f5761382f615170565b602002015283826002811061384657613846615170565b60200201516020015160016002811061386157613861615170565b6020020151836138728360056154c4565b600c811061388257613882615170565b6020020152508061389281615277565b9150506136d8565b506138a3614240565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6000806138d384613f10565b9050808360ff166001901b116139515760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610696565b90505b92915050565b6000805b82156139545761396f600184615494565b909216918061397d816157e8565b91505061395e565b60408051808201909152600080825260208201526102008261ffff16106139e15760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610696565b8161ffff16600114156139f5575081613954565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613a5e57600161ffff871660ff83161c81161415613a4157613a3e84846134bf565b93505b613a4b83846134bf565b92506201fffe600192831b169101613a11565b509195945050505050565b60408051808201909152600080825260208201528151158015613a8e57506020820151155b15613aac575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206158148339815191528460200151613adf9190615186565b613af790600080516020615814833981519152615494565b905292915050565b919050565b6033546001600160a01b0316331461285d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610696565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613bbe8461395a565b61ffff166001600160401b03811115613bd957613bd961430f565b6040519080825280601f01601f191660200182016040528015613c03576020820181803683370190505b5090506000805b825182108015613c1b575061010081105b15613c72576001811b935085841615613c62578060f81b838381518110613c4457613c44615170565b60200101906001600160f81b031916908160001a9053508160010191505b613c6b81615277565b9050613c0a565b5090949350505050565b613ca66040518060600160405280600063ffffffff16815260200160608152602001606081525090565b604080516060808201835260008083526020830182905292820152819063ffffffff89168152604080516020601f8a0181900481028201810190925288815290899089908190840183828082843760009201919091525050505060208083019190915260408051601f880183900483028101830190915286815290879087908190840183828082843760009201829052506040808701959095529351613d5493508592506020019050615800565b60408051808303601f190181529181528151602092830120600081815260cb909352912054929b909a50919850909650505050505050565b6065546001600160a01b0316158015613dad57506001600160a01b03821615155b613e2f5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610696565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613e7282613331565b5050565b600082613e8486868561409d565b1495945050505050565b60008080600080516020615814833981519152600360008051602061581483398151915286600080516020615814833981519152888909090890506000613f04827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206158148339815191526140e9565b91959194509092505050565b600061010082511115613f995760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610696565b8151613fa757506000919050565b60008083600081518110613fbd57613fbd615170565b0160200151600160f89190911c81901b92505b845181101561409457848181518110613feb57613feb615170565b0160200151600160f89190911c1b91508282116140805760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610696565b9181179161408d81615277565b9050613fd0565b50909392505050565b600081815b848110156140e0576140cc828787848181106140c0576140c0615170565b90506020020135614191565b9150806140d881615277565b9150506140a2565b50949350505050565b6000806140f4614240565b6140fc61425e565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156134775750826141865760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610696565b505195945050505050565b60008183106141ad576000828152602084905260409020613951565b6000838152602083905260409020613951565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061420f61427c565b815260200161421c61427c565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b600060a082840312156142ac57600080fd5b50919050565b6000604082840312156142ac57600080fd5b60008083601f8401126142d657600080fd5b5081356001600160401b038111156142ed57600080fd5b6020830191508360208260051b850101111561430857600080fd5b9250929050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156143475761434761430f565b60405290565b60405161010081016001600160401b03811182821017156143475761434761430f565b604051601f8201601f191681016001600160401b03811182821017156143985761439861430f565b604052919050565b60006001600160401b038211156143b9576143b961430f565b5060051b60200190565b63ffffffff811681146106a857600080fd5b8035613aff816143c3565b600082601f8301126143f157600080fd5b81356020614406614401836143a0565b614370565b82815260059290921b8401810191818101908684111561442557600080fd5b8286015b8481101561444957803561443c816143c3565b8352918301918301614429565b509695505050505050565b60006040828403121561446657600080fd5b61446e614325565b9050813581526020820135602082015292915050565b600082601f83011261449557600080fd5b813560206144a5614401836143a0565b82815260069290921b840181019181810190868411156144c457600080fd5b8286015b84811015614449576144da8882614454565b8352918301916040016144c8565b600082601f8301126144f957600080fd5b604051604081018181106001600160401b038211171561451b5761451b61430f565b806040525080604084018581111561453257600080fd5b845b81811015613a5e578035835260209283019201614534565b60006080828403121561455e57600080fd5b614566614325565b905061457283836144e8565b815261458183604084016144e8565b602082015292915050565b600082601f83011261459d57600080fd5b813560206145ad614401836143a0565b82815260059290921b840181019181810190868411156145cc57600080fd5b8286015b848110156144495780356001600160401b038111156145ef5760008081fd5b6145fd8986838b01016143e0565b8452509183019183016145d0565b6000610180828403121561461e57600080fd5b61462661434d565b905081356001600160401b038082111561463f57600080fd5b61464b858386016143e0565b8352602084013591508082111561466157600080fd5b61466d85838601614484565b6020840152604084013591508082111561468657600080fd5b61469285838601614484565b60408401526146a4856060860161454c565b60608401526146b68560e08601614454565b60808401526101208401359150808211156146d057600080fd5b6146dc858386016143e0565b60a08401526101408401359150808211156146f657600080fd5b614702858386016143e0565b60c084015261016084013591508082111561471c57600080fd5b506147298482850161458c565b60e08301525092915050565b60008060008060008060a0878903121561474e57600080fd5b86356001600160401b038082111561476557600080fd5b6147718a838b0161429a565b9750602089013591508082111561478757600080fd5b6147938a838b016142b2565b965060408901359150808211156147a957600080fd5b6147b58a838b016142c4565b909650945060608901359150808211156147ce57600080fd5b6147da8a838b016142b2565b935060808901359150808211156147f057600080fd5b506147fd89828a0161460b565b9150509295509295509295565b6001600160a01b03811681146106a857600080fd5b60006020828403121561483157600080fd5b81356139518161480a565b60006020828403121561484e57600080fd5b5035919050565b600080600080610120858703121561486c57600080fd5b8435935061487d8660208701614454565b925061488c866060870161454c565b915061489b8660e08701614454565b905092959194509250565b6000602082840312156148b857600080fd5b8135613951816143c3565b6000806000606084860312156148d857600080fd5b83356148e38161480a565b92506020848101356001600160401b038082111561490057600080fd5b818701915087601f83011261491457600080fd5b8135818111156149265761492661430f565b614938601f8201601f19168501614370565b9150808252888482850101111561494e57600080fd5b8084840185840137600084828401015250809450505050614971604085016143d5565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614a10578385038a52825180518087529087019087870190845b818110156149fb57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b031690840152928901926060909201916001016149b7565b50509a87019a95505091850191600101614999565b509298975050505050505050565b602081526000613951602083018461497a565b80151581146106a857600080fd5b600060208284031215614a5157600080fd5b813561395181614a31565b60008083601f840112614a6e57600080fd5b5081356001600160401b03811115614a8557600080fd5b60208301915083602082850101111561430857600080fd5b60008060008060008060808789031215614ab657600080fd5b8635614ac18161480a565b95506020870135614ad1816143c3565b945060408701356001600160401b0380821115614aed57600080fd5b614af98a838b01614a5c565b90965094506060890135915080821115614b1257600080fd5b50614b1f89828a016142c4565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614b6757815163ffffffff1687529582019590820190600101614b45565b509495945050505050565b600060208083528351608082850152614b8e60a0850182614b31565b905081850151601f1980868403016040870152614bab8383614b31565b92506040870151915080868403016060870152614bc88383614b31565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614c1f5784878303018452614c0d828751614b31565b95880195938801939150600101614bf3565b509998505050505050505050565b60ff811681146106a857600080fd5b600060208284031215614c4e57600080fd5b813561395181614c2d565b600080600060608486031215614c6e57600080fd5b8335614c798161480a565b92506020848101356001600160401b03811115614c9557600080fd5b8501601f81018713614ca657600080fd5b8035614cb4614401826143a0565b81815260059190911b82018301908381019089831115614cd357600080fd5b928401925b82841015614cf157833582529284019290840190614cd8565b8096505050505050614971604085016143d5565b6020808252825182820181905260009190848201906040850190845b81811015614d3d57835183529284019291840191600101614d21565b50909695505050505050565b600080600080600060808688031215614d6157600080fd5b8535945060208601356001600160401b0380821115614d7f57600080fd5b614d8b89838a01614a5c565b909650945060408801359150614da0826143c3565b90925060608701359080821115614db657600080fd5b50614dc38882890161460b565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614b675781516001600160601b031687529582019590820190600101614de4565b6040815260008351604080840152614e246080840182614dd0565b90506020850151603f19848303016060850152614e418282614dd0565b925050508260208301529392505050565b600080600060608486031215614e6757600080fd5b8335614e728161480a565b9250602084013591506040840135614e89816143c3565b809150509250925092565b828152604060208201526000614ead604083018461497a565b949350505050565b600080600080600060608688031215614ecd57600080fd5b8535614ed8816143c3565b945060208601356001600160401b0380821115614ef457600080fd5b614f0089838a01614a5c565b90965094506040880135915080821115614f1957600080fd5b50614f2688828901614a5c565b969995985093965092949392505050565b60008060008060808587031215614f4d57600080fd5b8435614f588161480a565b93506020850135614f688161480a565b92506040850135614f788161480a565b91506060850135614f888161480a565b939692955090935050565b60008060008060608587031215614fa957600080fd5b843593506020850135614fbb816143c3565b925060408501356001600160401b03811115614fd657600080fd5b614fe287828801614a5c565b95989497509550505050565b6000808335601e1984360301811261500557600080fd5b83016020810192503590506001600160401b0381111561502457600080fd5b80360383131561430857600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006150688283614fee565b6040855261507a604086018284615033565b915050602083013560208501528091505092915050565b602081526000613951602083018461505c565b6000602082840312156150b657600080fd5b81516139518161480a565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561511d57600080fd5b815161395181614a31565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000826151a357634e487b7160e01b600052601260045260246000fd5b500690565b600060208083850312156151bb57600080fd5b82516001600160401b038111156151d157600080fd5b8301601f810185136151e257600080fd5b80516151f0614401826143a0565b81815260059190911b8201830190838101908783111561520f57600080fd5b928401925b8284101561522d57835182529284019290840190615214565b979650505050505050565b60006020828403121561524a57600080fd5b81516001600160601b038116811461395157600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561528b5761528b615261565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156152bf57600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156152ef57600080fd5b82516001600160401b0381111561530557600080fd5b8301601f8101851361531657600080fd5b8051615324614401826143a0565b81815260059190911b8201830190838101908783111561534357600080fd5b928401925b8284101561522d57835161535b816143c3565b82529284019290840190615348565b63ffffffff8416815260406020820152600061538a604083018486615033565b95945050505050565b6000602082840312156153a557600080fd5b81516001600160c01b038116811461395157600080fd5b6000602082840312156153ce57600080fd5b8151613951816143c3565b600060ff821660ff8114156153f0576153f0615261565b60010192915050565b60408152600061540d604083018587615033565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561546a5784518352938301939183019160010161544e565b5090979650505050505050565b60006020828403121561548957600080fd5b815161395181614c2d565b6000828210156154a6576154a6615261565b500390565b6000602082840312156154bd57600080fd5b5051919050565b600082198211156154d7576154d7615261565b500190565b6000602082840312156154ee57600080fd5b815167ffffffffffffffff198116811461395157600080fd5b60006001600160601b038381169083168181101561552757615527615261565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561556a5781518552938201939082019060010161554e565b5092979650505050505050565b6000815180845260005b8181101561559d57602081850181015186830182015201615581565b818111156155af576000602083870101525b50601f01601f19169290920160200192915050565b63ffffffff815116825260006020820151606060208501526155e96060850182615577565b90506040830151848203604086015261538a8282615577565b60408152600061561560408301856155c4565b828103602084015261538a818561505c565b60208152600063ffffffff80845116602084015280602085015116604084015260408401516060840152606084015160a0608085015261566a60c0850182615577565b90508160808601511660a0850152809250505092915050565b600063ffffffff8083168181141561569d5761569d615261565b6001019392505050565b60208152600082356156b8816143c3565b63ffffffff8082166020850152602085013591506156d5826143c3565b8082166040850152604085013560608501526156f46060860186614fee565b925060a0608086015261570b60c086018483615033565b925050608085013561571c816143c3565b1660a0939093019290925250919050565b6000808335601e1984360301811261574457600080fd5b8301803591506001600160401b0382111561575e57600080fd5b60200191503681900382131561430857600080fd5b604081526000615787604083018688615033565b828103602084015261522d818587615033565b60006001600160601b03808316818516818304811182151516156157c0576157c0615261565b02949350505050565b60008160001904831182151516156157e3576157e3615261565b500290565b600061ffff8083168181141561569d5761569d615261565b60208152600061395160208301846155c456fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220666a5869803efec382c33cae4986092440171253b59090426a2acfafaac9d2f064736f6c634300080c0033",
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

// CheckValidTaskResponse is a free data retrieval call binding the contract method 0x02f0a1dc.
//
// Solidity: function checkValidTaskResponse((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) CheckValidTaskResponse(opts *bind.CallOpts, batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) error {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "checkValidTaskResponse", batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)

	if err != nil {
		return err
	}

	return err

}

// CheckValidTaskResponse is a free data retrieval call binding the contract method 0x02f0a1dc.
//
// Solidity: function checkValidTaskResponse((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) CheckValidTaskResponse(batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) error {
	return _ContractLambadaCoprocessorTaskManager.Contract.CheckValidTaskResponse(&_ContractLambadaCoprocessorTaskManager.CallOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// CheckValidTaskResponse is a free data retrieval call binding the contract method 0x02f0a1dc.
//
// Solidity: function checkValidTaskResponse((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) CheckValidTaskResponse(batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) error {
	return _ContractLambadaCoprocessorTaskManager.Contract.CheckValidTaskResponse(&_ContractLambadaCoprocessorTaskManager.CallOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
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

// GetTaskResponseHash is a free data retrieval call binding the contract method 0xf8ad0953.
//
// Solidity: function getTaskResponseHash(uint32 batchIndex, bytes programId, bytes taskInputHash) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCaller) GetTaskResponseHash(opts *bind.CallOpts, batchIndex uint32, programId []byte, taskInputHash []byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractLambadaCoprocessorTaskManager.contract.Call(opts, &out, "getTaskResponseHash", batchIndex, programId, taskInputHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTaskResponseHash is a free data retrieval call binding the contract method 0xf8ad0953.
//
// Solidity: function getTaskResponseHash(uint32 batchIndex, bytes programId, bytes taskInputHash) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) GetTaskResponseHash(batchIndex uint32, programId []byte, taskInputHash []byte) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetTaskResponseHash(&_ContractLambadaCoprocessorTaskManager.CallOpts, batchIndex, programId, taskInputHash)
}

// GetTaskResponseHash is a free data retrieval call binding the contract method 0xf8ad0953.
//
// Solidity: function getTaskResponseHash(uint32 batchIndex, bytes programId, bytes taskInputHash) view returns(bytes32)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerCallerSession) GetTaskResponseHash(batchIndex uint32, programId []byte, taskInputHash []byte) ([32]byte, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.GetTaskResponseHash(&_ContractLambadaCoprocessorTaskManager.CallOpts, batchIndex, programId, taskInputHash)
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

// RespondTask is a paid mutator transaction binding the contract method 0xbda4422a.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerTransactor) RespondTask(opts *bind.TransactOpts, batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.contract.Transact(opts, "respondTask", batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// RespondTask is a paid mutator transaction binding the contract method 0xbda4422a.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerSession) RespondTask(batch ILambadaCoprocessorTaskManagerTaskBatch, task ILambadaCoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ILambadaCoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractLambadaCoprocessorTaskManager.Contract.RespondTask(&_ContractLambadaCoprocessorTaskManager.TransactOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// RespondTask is a paid mutator transaction binding the contract method 0xbda4422a.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
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

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xde0a98210b193f0e75b97832743430e00f558966c6917131cee6e590de548c4d.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) responseMeta, (bytes,bytes32) response)
func (_ContractLambadaCoprocessorTaskManager *ContractLambadaCoprocessorTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractLambadaCoprocessorTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractLambadaCoprocessorTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractLambadaCoprocessorTaskManagerTaskRespondedIterator{contract: _ContractLambadaCoprocessorTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xde0a98210b193f0e75b97832743430e00f558966c6917131cee6e590de548c4d.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) responseMeta, (bytes,bytes32) response)
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

// ParseTaskResponded is a log parse operation binding the contract event 0xde0a98210b193f0e75b97832743430e00f558966c6917131cee6e590de548c4d.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) responseMeta, (bytes,bytes32) response)
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
