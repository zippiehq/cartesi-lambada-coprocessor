// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractCoprocessorTaskManager

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

// ICoprocessorTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type ICoprocessorTaskManagerTask struct {
	BatchIndex uint32
	ProgramId  []byte
	InputHash  []byte
}

// ICoprocessorTaskManagerTaskBatch is an auto generated low-level Go binding around an user-defined struct.
type ICoprocessorTaskManagerTaskBatch struct {
	Index                     uint32
	BlockNumber               uint32
	MerkeRoot                 [32]byte
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// ICoprocessorTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type ICoprocessorTaskManagerTaskResponse struct {
	ResultCID  []byte
	OutputHash [32]byte
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

// ContractCoprocessorTaskManagerMetaData contains all meta data concerning the ContractCoprocessorTaskManager contract.
var ContractCoprocessorTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allBatchHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskOutputs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkValidTaskResponse\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"internalType\":\"structICoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structICoprocessorTaskManager.Task\",\"components\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"inputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structICoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"resultCID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextBatchIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseHash\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structICoprocessorTaskManager.Task\",\"components\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"inputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextBatchIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerNewTaskBatch\",\"inputs\":[{\"name\":\"batchRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondTask\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"internalType\":\"structICoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structICoprocessorTaskManager.Task\",\"components\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"inputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"taskProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structICoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"resultCID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskBatchRegistered\",\"inputs\":[{\"name\":\"batch\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICoprocessorTaskManager.TaskBatch\",\"components\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"merkeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICoprocessorTaskManager.Task\",\"components\":[{\"name\":\"batchIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"programId\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"inputHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"response\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structICoprocessorTaskManager.TaskResponse\",\"components\":[{\"name\":\"resultCID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outputHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005c4d38038062005c4d8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e0516101005161596b620002e2600039600061027a01526000818161051e01526120f70152600081816103eb01526122d9015260008181610412015281816124af015261267101526000818161043901528181610ed001528181611dc201528181611f5a0152612194015261596b6000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c8063683048351161011a5780638da5cb5b116100ad578063f0f19ee11161007c578063f0f19ee114610540578063f2fde38b1461054e578063f8c8765e14610561578063fa36f60c14610574578063fabc1cbc1461058757600080fd5b80638da5cb5b146104da578063b98d0908146104eb578063cefdc1d4146104f8578063df5cf7231461051957600080fd5b8063715018a6116100e9578063715018a61461049c5780637afa1eed146104a45780637e4fa700146104b7578063886f1195146104c757600080fd5b8063683048351461040d5780636d14a987146104345780636e1699ae1461045b5780636efb46361461047b57600080fd5b80633f685bf61161019d5780635ac86ab71161016c5780635ac86ab7146103785780635c155662146103ab5780635c317d82146103cb5780635c975abb146103de5780635df45946146103e657600080fd5b80633f685bf61461032a578063416c7e5e1461033d5780634f739f7414610350578063595c6a671461037057600080fd5b80631ad43189116101d95780631ad43189146102755780631d2d5d24146102b1578063245a7bfc146102df5780633563b0d11461030a57600080fd5b806310d67a2f1461020b5780631349087614610220578063136439dd14610233578063171f1d5b14610246575b600080fd5b61021e61021936600461433e565b61059a565b005b61021e61022e366004614808565b610656565b61021e6102413660046148dd565b61074a565b6102596102543660046148f6565b610889565b6040805192151583529015156020830152015b60405180910390f35b61029c7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026c565b6102d16102bf366004614947565b60ca6020526000908152604090205481565b60405190815260200161026c565b60cc546102f2906001600160a01b031681565b6040516001600160a01b03909116815260200161026c565b61031d610318366004614964565b610a13565b60405161026c9190614abf565b61021e610338366004614808565b610ea9565b61021e61034b366004614ae0565b610ece565b61036361035e366004614b3e565b611043565b60405161026c9190614c13565b61021e611769565b61039b610386366004614cdd565b606654600160ff9092169190911b9081161490565b604051901515815260200161026c565b6103be6103b9366004614cfa565b611830565b60405161026c9190614da6565b6102d16103d9366004614dea565b6119f8565b6066546102d1565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b6102d16104693660046148dd565b60cb6020526000908152604090205481565b61048e610489366004614e26565b611a0f565b60405161026c929190614ee6565b61021e612926565b60cd546102f2906001600160a01b031681565b60c95461029c9063ffffffff1681565b6065546102f2906001600160a01b031681565b6033546001600160a01b03166102f2565b60975461039b9060ff1681565b61050b610506366004614f2f565b61293a565b60405161026c929190614f71565b6102f27f000000000000000000000000000000000000000000000000000000000000000081565b60c95463ffffffff1661029c565b61021e61055c36600461433e565b612acc565b61021e61056f366004614f8a565b612b42565b61021e610582366004614fe6565b612c93565b61021e6105953660046148dd565b612e1b565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106119190615041565b6001600160a01b0316336001600160a01b03161461064a5760405162461bcd60e51b81526004016106419061505e565b60405180910390fd5b61065381612f77565b50565b60cc546001600160a01b031633146106b05760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610641565b60006106be8787878761306e565b90506106cc878785856132db565b6000836040516020016106df919061514b565b60408051601f198184030181528282528051602091820120600086815260cb90925291902081905591507fde0a98210b193f0e75b97832743430e00f558966c6917131cee6e590de548c4d9061073890899087906151c2565b60405180910390a15050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610792573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b691906151e7565b6107d25760405162461bcd60e51b815260040161064190615204565b6066548181161461084b5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610641565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001878760000151886020015188600001516000600281106108d1576108d161524c565b60200201518951600160200201518a602001516000600281106108f6576108f661524c565b60200201518b602001516001600281106109125761091261524c565b602090810291909101518c518d83015160405161096f9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6109929190615262565b9050610a056109ab6109a48884613581565b8690613618565b6109b36136ac565b6109fb6109ec856109e6604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613581565b6109f58c61376c565b90613618565b886201d4c06137fc565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a799190615041565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610abb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610adf9190615041565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b459190615041565b9050600086516001600160401b03811115610b6257610b626143e2565b604051908082528060200260200182016040528015610b9557816020015b6060815260200190600190039081610b805790505b50905060005b8751811015610e9d576000888281518110610bb857610bb861524c565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c19573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c419190810190615284565b905080516001600160401b03811115610c5c57610c5c6143e2565b604051908082528060200260200182016040528015610ca757816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610c7a5790505b50848481518110610cba57610cba61524c565b602002602001018190525060005b8151811015610e87576040518060600160405280876001600160a01b03166347b314e8858581518110610cfd57610cfd61524c565b60200260200101516040518263ffffffff1660e01b8152600401610d2391815260200190565b602060405180830381865afa158015610d40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d649190615041565b6001600160a01b03168152602001838381518110610d8457610d8461524c565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610db257610db261524c565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610e0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e329190615314565b6001600160601b0316815250858581518110610e5057610e5061524c565b60200260200101518281518110610e6957610e6961524c565b60200260200101819052508080610e7f90615353565b915050610cc8565b5050508080610e9590615353565b915050610b9b565b50979650505050505050565b6000610eb78787878761306e565b9050610ec5878785856132db565b50505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f509190615041565b6001600160a01b0316336001600160a01b031614610ffc5760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610641565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b61106e6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d29190615041565b90506110ff6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061112f908b908990899060040161536e565b600060405180830381865afa15801561114c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261117491908101906153b8565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906111a6908b908b908b90600401615446565b600060405180830381865afa1580156111c3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111eb91908101906153b8565b6040820152856001600160401b03811115611208576112086143e2565b60405190808252806020026020018201604052801561123b57816020015b60608152602001906001900390816112265790505b50606082015260005b60ff811687111561167a576000856001600160401b03811115611269576112696143e2565b604051908082528060200260200182016040528015611292578160200160208202803683370190505b5083606001518360ff16815181106112ac576112ac61524c565b602002602001018190525060005b8681101561157a5760008c6001600160a01b03166304ec63518a8a858181106112e5576112e561524c565b905060200201358e886000015186815181106113035761130361524c565b60200260200101516040518463ffffffff1660e01b81526004016113409392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561135d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113819190615466565b90506001600160c01b0381166114255760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610641565b8a8a8560ff1681811061143a5761143a61524c565b6001600160c01b03841692013560f81c9190911c60019081161415905061156757856001600160a01b031663dd9846b98a8a8581811061147c5761147c61524c565b905060200201358d8d8860ff168181106114985761149861524c565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa1580156114ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611512919061548f565b85606001518560ff168151811061152b5761152b61524c565b602002602001015184815181106115445761154461524c565b63ffffffff909216602092830291909101909101528261156381615353565b9350505b508061157281615353565b9150506112ba565b506000816001600160401b03811115611595576115956143e2565b6040519080825280602002602001820160405280156115be578160200160208202803683370190505b50905060005b8281101561163f5784606001518460ff16815181106115e5576115e561524c565b602002602001015181815181106115fe576115fe61524c565b60200260200101518282815181106116185761161861524c565b63ffffffff909216602092830291909101909101528061163781615353565b9150506115c4565b508084606001518460ff168151811061165a5761165a61524c565b602002602001018190525050508080611672906154ac565b915050611244565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116df9190615041565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611712908b908b908e906004016154cc565b600060405180830381865afa15801561172f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261175791908101906153b8565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156117b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d591906151e7565b6117f15760405162461bcd60e51b815260040161064190615204565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b81526004016118629291906154f6565b600060405180830381865afa15801561187f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118a791908101906153b8565b9050600084516001600160401b038111156118c4576118c46143e2565b6040519080825280602002602001820160405280156118ed578160200160208202803683370190505b50905060005b85518110156119ee57866001600160a01b03166304ec635187838151811061191d5761191d61524c565b6020026020010151878685815181106119385761193861524c565b60200260200101516040518463ffffffff1660e01b81526004016119759392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611992573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119b69190615466565b6001600160c01b03168282815181106119d1576119d161524c565b6020908102919091010152806119e681615353565b9150506118f3565b5095945050505050565b6000806000611a0684613a20565b95945050505050565b6040805180820190915260608082526020820152600084611a865760405162461bcd60e51b8152602060048201526037602482015260008051602061591683398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610641565b60408301515185148015611a9e575060a08301515185145b8015611aae575060c08301515185145b8015611abe575060e08301515185145b611b285760405162461bcd60e51b8152602060048201526041602482015260008051602061591683398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610641565b82515160208401515114611ba05760405162461bcd60e51b815260206004820152604460248201819052600080516020615916833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610641565b4363ffffffff168463ffffffff1610611c0f5760405162461bcd60e51b815260206004820152603c602482015260008051602061591683398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610641565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611c5057611c506143e2565b604051908082528060200260200182016040528015611c79578160200160208202803683370190505b506020820152866001600160401b03811115611c9757611c976143e2565b604051908082528060200260200182016040528015611cc0578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611cf457611cf46143e2565b604051908082528060200260200182016040528015611d1d578160200160208202803683370190505b5081526020860151516001600160401b03811115611d3d57611d3d6143e2565b604051908082528060200260200182016040528015611d66578160200160208202803683370190505b5081602001819052506000611e388a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611e0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e33919061553d565b613a66565b905060005b8760200151518110156120d357611e8288602001518281518110611e6357611e6361524c565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611e9857611e9861524c565b60209081029190910101528015611f58576020830151611eb960018361555a565b81518110611ec957611ec961524c565b602002602001015160001c83602001518281518110611eea57611eea61524c565b602002602001015160001c11611f58576040805162461bcd60e51b815260206004820152602481019190915260008051602061591683398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610641565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611f9d57611f9d61524c565b60200260200101518b8b600001518581518110611fbc57611fbc61524c565b60200260200101516040518463ffffffff1660e01b8152600401611ff99392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612016573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203a9190615466565b6001600160c01b0316836000015182815181106120595761205961524c565b6020026020010181815250506120bf6109a461209384866000015185815181106120855761208561524c565b602002602001015116613af9565b8a6020015184815181106120a9576120a961524c565b6020026020010151613b2490919063ffffffff16565b9450806120cb81615353565b915050611e3d565b50506120de83613c08565b60975490935060ff166000816120f5576000612177565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612153573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121779190615571565b905060005b8a8110156127f55782156122d7578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106121d3576121d361524c565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612213573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122379190615571565b612241919061558a565b116122d75760405162461bcd60e51b8152602060048201526066602482015260008051602061591683398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610641565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d848181106123185761231861524c565b9050013560f81c60f81b60f81c8c8c60a00151858151811061233c5761233c61524c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612398573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123bc91906155a2565b6001600160401b0319166123df8a604001518381518110611e6357611e6361524c565b67ffffffffffffffff19161461247b5760405162461bcd60e51b8152602060048201526061602482015260008051602061591683398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610641565b6124ab896040015182815181106124945761249461524c565b60200260200101518761361890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106124ee576124ee61524c565b9050013560f81c60f81b60f81c8c8c60c0015185815181106125125761251261524c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561256e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125929190615314565b856020015182815181106125a8576125a861524c565b6001600160601b039092166020928302919091018201528501518051829081106125d4576125d461524c565b6020026020010151856000015182815181106125f2576125f261524c565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156127e05761266a8660000151828151811061263c5761263c61524c565b60200260200101518f8f868181106126565761265661524c565b600192013560f81c9290921c811614919050565b156127ce577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106126b0576126b061524c565b9050013560f81c60f81b60f81c8e896020015185815181106126d4576126d461524c565b60200260200101518f60e0015188815181106126f2576126f261524c565b6020026020010151878151811061270b5761270b61524c565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561276f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127939190615314565b87518051859081106127a7576127a761524c565b602002602001018181516127bb91906155cd565b6001600160601b03169052506001909101905b806127d881615353565b915050612616565b505080806127ed90615353565b91505061217c565b50505060008061280f8c868a606001518b60800151610889565b91509150816128805760405162461bcd60e51b8152602060048201526043602482015260008051602061591683398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610641565b806128e15760405162461bcd60e51b8152602060048201526039602482015260008051602061591683398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610641565b505060008782602001516040516020016128fc9291906155f5565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b61292e613ca3565b6129386000613cfd565b565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106129755761297561524c565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906129b190889086906004016154f6565b600060405180830381865afa1580156129ce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526129f691908101906153b8565b600081518110612a0857612a0861524c565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612a74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a989190615466565b6001600160c01b031690506000612aae82613d4f565b905081612abc8a838a610a13565b9550955050505050935093915050565b612ad4613ca3565b6001600160a01b038116612b395760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610641565b61065381613cfd565b600054610100900460ff1615808015612b625750600054600160ff909116105b80612b7c5750303b158015612b7c575060005460ff166001145b612bdf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610641565b6000805460ff191660011790558015612c02576000805461ff0019166101001790555b612c0d856000613e1b565b612c1684613cfd565b60cc80546001600160a01b038086166001600160a01b03199283161790925560cd8054928516929091169190911790558015612c8c576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60cd546001600160a01b03163314612cf75760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610641565b6040805160a08101825260608082015260c95463ffffffff908116825243811660208084019190915282840188905290861660808301528251601f85018290048202810182019093528383529091908490849081908401838280828437600092018290525060608601949094525050604051612d789150839060200161568a565b60408051601f19818403018152918152815160209283012060c9805463ffffffff908116600090815260ca9095529284208290558054919450911691612dbd836156e6565b91906101000a81548163ffffffff021916908363ffffffff160217905550507f9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd9582604051612e0b919061568a565b60405180910390a1505050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e929190615041565b6001600160a01b0316336001600160a01b031614612ec25760405162461bcd60e51b81526004016106419061505e565b606654198119606654191614612f405760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610641565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161087e565b6001600160a01b0381166130055760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610641565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b600061307d6020850185614947565b63ffffffff166130906020870187614947565b63ffffffff16146130f85760405162461bcd60e51b815260206004820152602c60248201527f62617463682e696e64657820616e64207461736b2e6261746368496e6465782060448201526b0c8de40dcdee840dac2e8c6d60a31b6064820152608401610641565b60008560405160200161310b919061570a565b6040516020818303038152906040528051906020012090508060ca600088600001602081019061313b9190614947565b63ffffffff1663ffffffff16815260200190815260200160002054146131ae5760405162461bcd60e51b815260206004820152602260248201527f5461736b20626174636820686173206e6f74206265656e207265676973746572604482015261195960f21b6064820152608401610641565b6000806131ba87613a20565b915091508060001461320e5760405162461bcd60e51b815260206004820152601f60248201527f5461736b20726573706f6e736520616c726561647920726573706f6e646564006044820152606401610641565b600061321d6020890189615790565b61322a60408b018b615790565b60405160200161323d94939291906157d6565b60408051601f198184030181528282528051602091820120908301520160405160208183030381529060405280519060200120905061328287878b6040013584613f05565b6132ce5760405162461bcd60e51b815260206004820152601d60248201527f5461736b20646f6573206e6f742062656c6f6e6720746f2062617463680000006044820152606401610641565b5090979650505050505050565b6133166040518060a00160405280600063ffffffff168152602001606081526020016060815260200160608152602001600080191681525090565b6133236020850185614947565b63ffffffff1681526133386020850185615790565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050602082015261337e6040850185615790565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408201526133c18380615790565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060608601949094525050506020808501356080840152604051613416918491016157fd565b60408051601f198184030181529190528051602090910120905060008061345a8361344460608b018b615790565b61345460408d0160208e01614947565b89611a0f565b9150915060005b61346e60608a018a615790565b90508110156135765761348760a08a0160808b01614947565b60ff16836020015182815181106134a0576134a061524c565b60200260200101516134b2919061587c565b6001600160601b03166064846000015183815181106134d3576134d361524c565b60200260200101516001600160601b03166134ee91906158ab565b1015613564576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610641565b8061356e81615353565b915050613461565b505050505050505050565b604080518082019091526000808252602082015261359d61424f565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156135d0576135d2565bfe5b50806136105760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610641565b505092915050565b604080518082019091526000808252602082015261363461426d565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156135d05750806136105760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610641565b6136b461428b565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061379c6000805160206158f683398151915286615262565b90505b6137a881613f1d565b90935091506000805160206158f68339815191528283098314156137e2576040805180820190915290815260208101919091529392505050565b6000805160206158f683398151915260018208905061379f565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061382e6142b0565b60005b60028110156139f35760006138478260066158ab565b905084826002811061385b5761385b61524c565b6020020151518361386d83600061558a565b600c811061387d5761387d61524c565b60200201528482600281106138945761389461524c565b602002015160200151838260016138ab919061558a565b600c81106138bb576138bb61524c565b60200201528382600281106138d2576138d261524c565b60200201515151836138e583600261558a565b600c81106138f5576138f561524c565b602002015283826002811061390c5761390c61524c565b602002015151600160200201518361392583600361558a565b600c81106139355761393561524c565b602002015283826002811061394c5761394c61524c565b6020020151602001516000600281106139675761396761524c565b60200201518361397883600461558a565b600c81106139885761398861524c565b602002015283826002811061399f5761399f61524c565b6020020151602001516001600281106139ba576139ba61524c565b6020020151836139cb83600561558a565b600c81106139db576139db61524c565b602002015250806139eb81615353565b915050613831565b506139fc6142cf565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080600083604051602001613a3691906158ca565b60408051601f198184030181529181528151602092830120600081815260cb909352912054909590945092505050565b600080613a7284613f9f565b9050808360ff166001901b11613af05760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610641565b90505b92915050565b6000805b8215613af357613b0e60018461555a565b9092169180613b1c816158dd565b915050613afd565b60408051808201909152600080825260208201526102008261ffff1610613b805760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610641565b8161ffff1660011415613b94575081613af3565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613bfd57600161ffff871660ff83161c81161415613be057613bdd8484613618565b93505b613bea8384613618565b92506201fffe600192831b169101613bb0565b509195945050505050565b60408051808201909152600080825260208201528151158015613c2d57506020820151155b15613c4b575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206158f68339815191528460200151613c7e9190615262565b613c96906000805160206158f683398151915261555a565b905292915050565b919050565b6033546001600160a01b031633146129385760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610641565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613d5d84613af9565b61ffff166001600160401b03811115613d7857613d786143e2565b6040519080825280601f01601f191660200182016040528015613da2576020820181803683370190505b5090506000805b825182108015613dba575061010081105b15613e11576001811b935085841615613e01578060f81b838381518110613de357613de361524c565b60200101906001600160f81b031916908160001a9053508160010191505b613e0a81615353565b9050613da9565b5090949350505050565b6065546001600160a01b0316158015613e3c57506001600160a01b03821615155b613ebe5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610641565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613f0182612f77565b5050565b600082613f1386868561412c565b1495945050505050565b600080806000805160206158f683398151915260036000805160206158f6833981519152866000805160206158f6833981519152888909090890506000613f93827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206158f6833981519152614178565b91959194509092505050565b6000610100825111156140285760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610641565b815161403657506000919050565b6000808360008151811061404c5761404c61524c565b0160200151600160f89190911c81901b92505b84518110156141235784818151811061407a5761407a61524c565b0160200151600160f89190911c1b915082821161410f5760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610641565b9181179161411c81615353565b905061405f565b50909392505050565b600081815b8481101561416f5761415b8287878481811061414f5761414f61524c565b90506020020135614220565b91508061416781615353565b915050614131565b50949350505050565b6000806141836142cf565b61418b6142ed565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156135d05750826142155760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610641565b505195945050505050565b600081831061423c576000828152602084905260409020613af0565b6000838152602083905260409020613af0565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061429e61430b565b81526020016142ab61430b565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461065357600080fd5b60006020828403121561435057600080fd5b8135613af081614329565b600060a0828403121561436d57600080fd5b50919050565b60006060828403121561436d57600080fd5b60008083601f84011261439757600080fd5b5081356001600160401b038111156143ae57600080fd5b6020830191508360208260051b85010111156143c957600080fd5b9250929050565b60006040828403121561436d57600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561441a5761441a6143e2565b60405290565b60405161010081016001600160401b038111828210171561441a5761441a6143e2565b604051601f8201601f191681016001600160401b038111828210171561446b5761446b6143e2565b604052919050565b60006001600160401b0382111561448c5761448c6143e2565b5060051b60200190565b63ffffffff8116811461065357600080fd5b8035613c9e81614496565b600082601f8301126144c457600080fd5b813560206144d96144d483614473565b614443565b82815260059290921b840181019181810190868411156144f857600080fd5b8286015b8481101561451c57803561450f81614496565b83529183019183016144fc565b509695505050505050565b60006040828403121561453957600080fd5b6145416143f8565b9050813581526020820135602082015292915050565b600082601f83011261456857600080fd5b813560206145786144d483614473565b82815260069290921b8401810191818101908684111561459757600080fd5b8286015b8481101561451c576145ad8882614527565b83529183019160400161459b565b600082601f8301126145cc57600080fd5b604051604081018181106001600160401b03821117156145ee576145ee6143e2565b806040525080604084018581111561460557600080fd5b845b81811015613bfd578035835260209283019201614607565b60006080828403121561463157600080fd5b6146396143f8565b905061464583836145bb565b815261465483604084016145bb565b602082015292915050565b600082601f83011261467057600080fd5b813560206146806144d483614473565b82815260059290921b8401810191818101908684111561469f57600080fd5b8286015b8481101561451c5780356001600160401b038111156146c25760008081fd5b6146d08986838b01016144b3565b8452509183019183016146a3565b600061018082840312156146f157600080fd5b6146f9614420565b905081356001600160401b038082111561471257600080fd5b61471e858386016144b3565b8352602084013591508082111561473457600080fd5b61474085838601614557565b6020840152604084013591508082111561475957600080fd5b61476585838601614557565b6040840152614777856060860161461f565b60608401526147898560e08601614527565b60808401526101208401359150808211156147a357600080fd5b6147af858386016144b3565b60a08401526101408401359150808211156147c957600080fd5b6147d5858386016144b3565b60c08401526101608401359150808211156147ef57600080fd5b506147fc8482850161465f565b60e08301525092915050565b60008060008060008060a0878903121561482157600080fd5b86356001600160401b038082111561483857600080fd5b6148448a838b0161435b565b9750602089013591508082111561485a57600080fd5b6148668a838b01614373565b9650604089013591508082111561487c57600080fd5b6148888a838b01614385565b909650945060608901359150808211156148a157600080fd5b6148ad8a838b016143d0565b935060808901359150808211156148c357600080fd5b506148d089828a016146de565b9150509295509295509295565b6000602082840312156148ef57600080fd5b5035919050565b600080600080610120858703121561490d57600080fd5b8435935061491e8660208701614527565b925061492d866060870161461f565b915061493c8660e08701614527565b905092959194509250565b60006020828403121561495957600080fd5b8135613af081614496565b60008060006060848603121561497957600080fd5b833561498481614329565b92506020848101356001600160401b03808211156149a157600080fd5b818701915087601f8301126149b557600080fd5b8135818111156149c7576149c76143e2565b6149d9601f8201601f19168501614443565b915080825288848285010111156149ef57600080fd5b8084840185840137600084828401015250809450505050614a12604085016144a8565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614ab1578385038a52825180518087529087019087870190845b81811015614a9c57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614a58565b50509a87019a95505091850191600101614a3a565b509298975050505050505050565b602081526000613af06020830184614a1b565b801515811461065357600080fd5b600060208284031215614af257600080fd5b8135613af081614ad2565b60008083601f840112614b0f57600080fd5b5081356001600160401b03811115614b2657600080fd5b6020830191508360208285010111156143c957600080fd5b60008060008060008060808789031215614b5757600080fd5b8635614b6281614329565b95506020870135614b7281614496565b945060408701356001600160401b0380821115614b8e57600080fd5b614b9a8a838b01614afd565b90965094506060890135915080821115614bb357600080fd5b50614bc089828a01614385565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b83811015614c0857815163ffffffff1687529582019590820190600101614be6565b509495945050505050565b600060208083528351608082850152614c2f60a0850182614bd2565b905081850151601f1980868403016040870152614c4c8383614bd2565b92506040870151915080868403016060870152614c698383614bd2565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614cc05784878303018452614cae828751614bd2565b95880195938801939150600101614c94565b509998505050505050505050565b60ff8116811461065357600080fd5b600060208284031215614cef57600080fd5b8135613af081614cce565b600080600060608486031215614d0f57600080fd5b8335614d1a81614329565b92506020848101356001600160401b03811115614d3657600080fd5b8501601f81018713614d4757600080fd5b8035614d556144d482614473565b81815260059190911b82018301908381019089831115614d7457600080fd5b928401925b82841015614d9257833582529284019290840190614d79565b8096505050505050614a12604085016144a8565b6020808252825182820181905260009190848201906040850190845b81811015614dde57835183529284019291840191600101614dc2565b50909695505050505050565b600060208284031215614dfc57600080fd5b81356001600160401b03811115614e1257600080fd5b614e1e84828501614373565b949350505050565b600080600080600060808688031215614e3e57600080fd5b8535945060208601356001600160401b0380821115614e5c57600080fd5b614e6889838a01614afd565b909650945060408801359150614e7d82614496565b90925060608701359080821115614e9357600080fd5b50614ea0888289016146de565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614c085781516001600160601b031687529582019590820190600101614ec1565b6040815260008351604080840152614f016080840182614ead565b90506020850151603f19848303016060850152614f1e8282614ead565b925050508260208301529392505050565b600080600060608486031215614f4457600080fd5b8335614f4f81614329565b9250602084013591506040840135614f6681614496565b809150509250925092565b828152604060208201526000614e1e6040830184614a1b565b60008060008060808587031215614fa057600080fd5b8435614fab81614329565b93506020850135614fbb81614329565b92506040850135614fcb81614329565b91506060850135614fdb81614329565b939692955090935050565b60008060008060608587031215614ffc57600080fd5b84359350602085013561500e81614496565b925060408501356001600160401b0381111561502957600080fd5b61503587828801614afd565b95989497509550505050565b60006020828403121561505357600080fd5b8151613af081614329565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000808335601e198436030181126150bf57600080fd5b83016020810192503590506001600160401b038111156150de57600080fd5b8036038313156143c957600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061512282836150a8565b604085526151346040860182846150ed565b915050602083013560208501528091505092915050565b602081526000613af06020830184615116565b6000813561516b81614496565b63ffffffff16835261518060208301836150a8565b606060208601526151956060860182846150ed565b9150506151a560408401846150a8565b85830360408701526151b88382846150ed565b9695505050505050565b6040815260006151d5604083018561515e565b8281036020840152611a068185615116565b6000602082840312156151f957600080fd5b8151613af081614ad2565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261527f57634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561529757600080fd5b82516001600160401b038111156152ad57600080fd5b8301601f810185136152be57600080fd5b80516152cc6144d482614473565b81815260059190911b820183019083810190878311156152eb57600080fd5b928401925b82841015615309578351825292840192908401906152f0565b979650505050505050565b60006020828403121561532657600080fd5b81516001600160601b0381168114613af057600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156153675761536761533d565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561539b57600080fd5b8260051b8085606085013760009201606001918252509392505050565b600060208083850312156153cb57600080fd5b82516001600160401b038111156153e157600080fd5b8301601f810185136153f257600080fd5b80516154006144d482614473565b81815260059190911b8201830190838101908783111561541f57600080fd5b928401925b8284101561530957835161543781614496565b82529284019290840190615424565b63ffffffff84168152604060208201526000611a066040830184866150ed565b60006020828403121561547857600080fd5b81516001600160c01b0381168114613af057600080fd5b6000602082840312156154a157600080fd5b8151613af081614496565b600060ff821660ff8114156154c3576154c361533d565b60010192915050565b6040815260006154e06040830185876150ed565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156132ce57845183529383019391830191600101615521565b60006020828403121561554f57600080fd5b8151613af081614cce565b60008282101561556c5761556c61533d565b500390565b60006020828403121561558357600080fd5b5051919050565b6000821982111561559d5761559d61533d565b500190565b6000602082840312156155b457600080fd5b815167ffffffffffffffff1981168114613af057600080fd5b60006001600160601b03838116908316818110156155ed576155ed61533d565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561563057815185529382019390820190600101615614565b5092979650505050505050565b6000815180845260005b8181101561566357602081850181015186830182015201615647565b81811115615675576000602083870101525b50601f01601f19169290920160200192915050565b60208152600063ffffffff80845116602084015280602085015116604084015260408401516060840152606084015160a060808501526156cd60c085018261563d565b90508160808601511660a0850152809250505092915050565b600063ffffffff808316818114156157005761570061533d565b6001019392505050565b602081526000823561571b81614496565b63ffffffff80821660208501526020850135915061573882614496565b80821660408501526040850135606085015261575760608601866150a8565b925060a0608086015261576e60c0860184836150ed565b925050608085013561577f81614496565b1660a0939093019290925250919050565b6000808335601e198436030181126157a757600080fd5b8301803591506001600160401b038211156157c157600080fd5b6020019150368190038213156143c957600080fd5b6040815260006157ea6040830186886150ed565b82810360208401526153098185876150ed565b6020815263ffffffff82511660208201526000602083015160a0604084015261582960c084018261563d565b90506040840151601f1980858403016060860152615847838361563d565b9250606086015191508085840301608086015250615865828261563d565b915050608084015160a08401528091505092915050565b60006001600160601b03808316818516818304811182151516156158a2576158a261533d565b02949350505050565b60008160001904831182151516156158c5576158c561533d565b500290565b602081526000613af0602083018461515e565b600061ffff808316818114156157005761570061533d56fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212208708499b173525c17a18717c70541e80a4dafb4841292ad33065e866fdff381e64736f6c634300080c0033",
}

// ContractCoprocessorTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractCoprocessorTaskManagerMetaData.ABI instead.
var ContractCoprocessorTaskManagerABI = ContractCoprocessorTaskManagerMetaData.ABI

// ContractCoprocessorTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractCoprocessorTaskManagerMetaData.Bin instead.
var ContractCoprocessorTaskManagerBin = ContractCoprocessorTaskManagerMetaData.Bin

// DeployContractCoprocessorTaskManager deploys a new Ethereum contract, binding an instance of ContractCoprocessorTaskManager to it.
func DeployContractCoprocessorTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractCoprocessorTaskManager, error) {
	parsed, err := ContractCoprocessorTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractCoprocessorTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractCoprocessorTaskManager{ContractCoprocessorTaskManagerCaller: ContractCoprocessorTaskManagerCaller{contract: contract}, ContractCoprocessorTaskManagerTransactor: ContractCoprocessorTaskManagerTransactor{contract: contract}, ContractCoprocessorTaskManagerFilterer: ContractCoprocessorTaskManagerFilterer{contract: contract}}, nil
}

// ContractCoprocessorTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractCoprocessorTaskManager struct {
	ContractCoprocessorTaskManagerCaller     // Read-only binding to the contract
	ContractCoprocessorTaskManagerTransactor // Write-only binding to the contract
	ContractCoprocessorTaskManagerFilterer   // Log filterer for contract events
}

// ContractCoprocessorTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCoprocessorTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCoprocessorTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractCoprocessorTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCoprocessorTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractCoprocessorTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractCoprocessorTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractCoprocessorTaskManagerSession struct {
	Contract     *ContractCoprocessorTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractCoprocessorTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCoprocessorTaskManagerCallerSession struct {
	Contract *ContractCoprocessorTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ContractCoprocessorTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractCoprocessorTaskManagerTransactorSession struct {
	Contract     *ContractCoprocessorTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractCoprocessorTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractCoprocessorTaskManagerRaw struct {
	Contract *ContractCoprocessorTaskManager // Generic contract binding to access the raw methods on
}

// ContractCoprocessorTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCoprocessorTaskManagerCallerRaw struct {
	Contract *ContractCoprocessorTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractCoprocessorTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractCoprocessorTaskManagerTransactorRaw struct {
	Contract *ContractCoprocessorTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractCoprocessorTaskManager creates a new instance of ContractCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractCoprocessorTaskManager(address common.Address, backend bind.ContractBackend) (*ContractCoprocessorTaskManager, error) {
	contract, err := bindContractCoprocessorTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManager{ContractCoprocessorTaskManagerCaller: ContractCoprocessorTaskManagerCaller{contract: contract}, ContractCoprocessorTaskManagerTransactor: ContractCoprocessorTaskManagerTransactor{contract: contract}, ContractCoprocessorTaskManagerFilterer: ContractCoprocessorTaskManagerFilterer{contract: contract}}, nil
}

// NewContractCoprocessorTaskManagerCaller creates a new read-only instance of ContractCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractCoprocessorTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractCoprocessorTaskManagerCaller, error) {
	contract, err := bindContractCoprocessorTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerCaller{contract: contract}, nil
}

// NewContractCoprocessorTaskManagerTransactor creates a new write-only instance of ContractCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractCoprocessorTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractCoprocessorTaskManagerTransactor, error) {
	contract, err := bindContractCoprocessorTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerTransactor{contract: contract}, nil
}

// NewContractCoprocessorTaskManagerFilterer creates a new log filterer instance of ContractCoprocessorTaskManager, bound to a specific deployed contract.
func NewContractCoprocessorTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractCoprocessorTaskManagerFilterer, error) {
	contract, err := bindContractCoprocessorTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerFilterer{contract: contract}, nil
}

// bindContractCoprocessorTaskManager binds a generic wrapper to an already deployed contract.
func bindContractCoprocessorTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractCoprocessorTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractCoprocessorTaskManager.Contract.ContractCoprocessorTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.ContractCoprocessorTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.ContractCoprocessorTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractCoprocessorTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractCoprocessorTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractCoprocessorTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractCoprocessorTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractCoprocessorTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Aggregator(&_ContractCoprocessorTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Aggregator(&_ContractCoprocessorTaskManager.CallOpts)
}

// AllBatchHashes is a free data retrieval call binding the contract method 0x1d2d5d24.
//
// Solidity: function allBatchHashes(uint32 ) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) AllBatchHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "allBatchHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllBatchHashes is a free data retrieval call binding the contract method 0x1d2d5d24.
//
// Solidity: function allBatchHashes(uint32 ) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) AllBatchHashes(arg0 uint32) ([32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.AllBatchHashes(&_ContractCoprocessorTaskManager.CallOpts, arg0)
}

// AllBatchHashes is a free data retrieval call binding the contract method 0x1d2d5d24.
//
// Solidity: function allBatchHashes(uint32 ) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) AllBatchHashes(arg0 uint32) ([32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.AllBatchHashes(&_ContractCoprocessorTaskManager.CallOpts, arg0)
}

// AllTaskOutputs is a free data retrieval call binding the contract method 0x6e1699ae.
//
// Solidity: function allTaskOutputs(bytes32 ) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) AllTaskOutputs(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "allTaskOutputs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskOutputs is a free data retrieval call binding the contract method 0x6e1699ae.
//
// Solidity: function allTaskOutputs(bytes32 ) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) AllTaskOutputs(arg0 [32]byte) ([32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.AllTaskOutputs(&_ContractCoprocessorTaskManager.CallOpts, arg0)
}

// AllTaskOutputs is a free data retrieval call binding the contract method 0x6e1699ae.
//
// Solidity: function allTaskOutputs(bytes32 ) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) AllTaskOutputs(arg0 [32]byte) ([32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.AllTaskOutputs(&_ContractCoprocessorTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.BlsApkRegistry(&_ContractCoprocessorTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.BlsApkRegistry(&_ContractCoprocessorTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.CheckSignatures(&_ContractCoprocessorTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.CheckSignatures(&_ContractCoprocessorTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckValidTaskResponse is a free data retrieval call binding the contract method 0x3f685bf6.
//
// Solidity: function checkValidTaskResponse((uint32,uint32,bytes32,bytes,uint32) batch, (uint32,bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) CheckValidTaskResponse(opts *bind.CallOpts, batch ICoprocessorTaskManagerTaskBatch, task ICoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ICoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) error {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "checkValidTaskResponse", batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)

	if err != nil {
		return err
	}

	return err

}

// CheckValidTaskResponse is a free data retrieval call binding the contract method 0x3f685bf6.
//
// Solidity: function checkValidTaskResponse((uint32,uint32,bytes32,bytes,uint32) batch, (uint32,bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) CheckValidTaskResponse(batch ICoprocessorTaskManagerTaskBatch, task ICoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ICoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) error {
	return _ContractCoprocessorTaskManager.Contract.CheckValidTaskResponse(&_ContractCoprocessorTaskManager.CallOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// CheckValidTaskResponse is a free data retrieval call binding the contract method 0x3f685bf6.
//
// Solidity: function checkValidTaskResponse((uint32,uint32,bytes32,bytes,uint32) batch, (uint32,bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) view returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) CheckValidTaskResponse(batch ICoprocessorTaskManagerTaskBatch, task ICoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ICoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) error {
	return _ContractCoprocessorTaskManager.Contract.CheckValidTaskResponse(&_ContractCoprocessorTaskManager.CallOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Delegation(&_ContractCoprocessorTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Delegation(&_ContractCoprocessorTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Generator() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Generator(&_ContractCoprocessorTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Generator(&_ContractCoprocessorTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractCoprocessorTaskManager.Contract.GetCheckSignaturesIndices(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractCoprocessorTaskManager.Contract.GetCheckSignaturesIndices(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetNextBatchIndex is a free data retrieval call binding the contract method 0xf0f19ee1.
//
// Solidity: function getNextBatchIndex() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) GetNextBatchIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "getNextBatchIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetNextBatchIndex is a free data retrieval call binding the contract method 0xf0f19ee1.
//
// Solidity: function getNextBatchIndex() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) GetNextBatchIndex() (uint32, error) {
	return _ContractCoprocessorTaskManager.Contract.GetNextBatchIndex(&_ContractCoprocessorTaskManager.CallOpts)
}

// GetNextBatchIndex is a free data retrieval call binding the contract method 0xf0f19ee1.
//
// Solidity: function getNextBatchIndex() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) GetNextBatchIndex() (uint32, error) {
	return _ContractCoprocessorTaskManager.Contract.GetNextBatchIndex(&_ContractCoprocessorTaskManager.CallOpts)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractCoprocessorTaskManager.Contract.GetOperatorState(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractCoprocessorTaskManager.Contract.GetOperatorState(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractCoprocessorTaskManager.Contract.GetOperatorState0(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractCoprocessorTaskManager.Contract.GetOperatorState0(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractCoprocessorTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractCoprocessorTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractCoprocessorTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseHash is a free data retrieval call binding the contract method 0x5c317d82.
//
// Solidity: function getTaskResponseHash((uint32,bytes,bytes) task) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) GetTaskResponseHash(opts *bind.CallOpts, task ICoprocessorTaskManagerTask) ([32]byte, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "getTaskResponseHash", task)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTaskResponseHash is a free data retrieval call binding the contract method 0x5c317d82.
//
// Solidity: function getTaskResponseHash((uint32,bytes,bytes) task) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) GetTaskResponseHash(task ICoprocessorTaskManagerTask) ([32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.GetTaskResponseHash(&_ContractCoprocessorTaskManager.CallOpts, task)
}

// GetTaskResponseHash is a free data retrieval call binding the contract method 0x5c317d82.
//
// Solidity: function getTaskResponseHash((uint32,bytes,bytes) task) view returns(bytes32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) GetTaskResponseHash(task ICoprocessorTaskManagerTask) ([32]byte, error) {
	return _ContractCoprocessorTaskManager.Contract.GetTaskResponseHash(&_ContractCoprocessorTaskManager.CallOpts, task)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) NextBatchIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "nextBatchIndex")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) NextBatchIndex() (uint32, error) {
	return _ContractCoprocessorTaskManager.Contract.NextBatchIndex(&_ContractCoprocessorTaskManager.CallOpts)
}

// NextBatchIndex is a free data retrieval call binding the contract method 0x7e4fa700.
//
// Solidity: function nextBatchIndex() view returns(uint32)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) NextBatchIndex() (uint32, error) {
	return _ContractCoprocessorTaskManager.Contract.NextBatchIndex(&_ContractCoprocessorTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Owner() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Owner(&_ContractCoprocessorTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.Owner(&_ContractCoprocessorTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractCoprocessorTaskManager.Contract.Paused(&_ContractCoprocessorTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractCoprocessorTaskManager.Contract.Paused(&_ContractCoprocessorTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractCoprocessorTaskManager.Contract.Paused0(&_ContractCoprocessorTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractCoprocessorTaskManager.Contract.Paused0(&_ContractCoprocessorTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.PauserRegistry(&_ContractCoprocessorTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.PauserRegistry(&_ContractCoprocessorTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.RegistryCoordinator(&_ContractCoprocessorTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.RegistryCoordinator(&_ContractCoprocessorTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.StakeRegistry(&_ContractCoprocessorTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractCoprocessorTaskManager.Contract.StakeRegistry(&_ContractCoprocessorTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractCoprocessorTaskManager.Contract.StaleStakesForbidden(&_ContractCoprocessorTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractCoprocessorTaskManager.Contract.StaleStakesForbidden(&_ContractCoprocessorTaskManager.CallOpts)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractCoprocessorTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractCoprocessorTaskManager.Contract.TrySignatureAndApkVerification(&_ContractCoprocessorTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractCoprocessorTaskManager.Contract.TrySignatureAndApkVerification(&_ContractCoprocessorTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.Initialize(&_ContractCoprocessorTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.Initialize(&_ContractCoprocessorTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.Pause(&_ContractCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.Pause(&_ContractCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.PauseAll(&_ContractCoprocessorTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.PauseAll(&_ContractCoprocessorTaskManager.TransactOpts)
}

// RegisterNewTaskBatch is a paid mutator transaction binding the contract method 0xfa36f60c.
//
// Solidity: function registerNewTaskBatch(bytes32 batchRoot, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) RegisterNewTaskBatch(opts *bind.TransactOpts, batchRoot [32]byte, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "registerNewTaskBatch", batchRoot, quorumThresholdPercentage, quorumNumbers)
}

// RegisterNewTaskBatch is a paid mutator transaction binding the contract method 0xfa36f60c.
//
// Solidity: function registerNewTaskBatch(bytes32 batchRoot, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) RegisterNewTaskBatch(batchRoot [32]byte, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.RegisterNewTaskBatch(&_ContractCoprocessorTaskManager.TransactOpts, batchRoot, quorumThresholdPercentage, quorumNumbers)
}

// RegisterNewTaskBatch is a paid mutator transaction binding the contract method 0xfa36f60c.
//
// Solidity: function registerNewTaskBatch(bytes32 batchRoot, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) RegisterNewTaskBatch(batchRoot [32]byte, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.RegisterNewTaskBatch(&_ContractCoprocessorTaskManager.TransactOpts, batchRoot, quorumThresholdPercentage, quorumNumbers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.RenounceOwnership(&_ContractCoprocessorTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.RenounceOwnership(&_ContractCoprocessorTaskManager.TransactOpts)
}

// RespondTask is a paid mutator transaction binding the contract method 0x13490876.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (uint32,bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) RespondTask(opts *bind.TransactOpts, batch ICoprocessorTaskManagerTaskBatch, task ICoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ICoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "respondTask", batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// RespondTask is a paid mutator transaction binding the contract method 0x13490876.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (uint32,bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) RespondTask(batch ICoprocessorTaskManagerTaskBatch, task ICoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ICoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.RespondTask(&_ContractCoprocessorTaskManager.TransactOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// RespondTask is a paid mutator transaction binding the contract method 0x13490876.
//
// Solidity: function respondTask((uint32,uint32,bytes32,bytes,uint32) batch, (uint32,bytes,bytes) task, bytes32[] taskProof, (bytes,bytes32) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) RespondTask(batch ICoprocessorTaskManagerTaskBatch, task ICoprocessorTaskManagerTask, taskProof [][32]byte, taskResponse ICoprocessorTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.RespondTask(&_ContractCoprocessorTaskManager.TransactOpts, batch, task, taskProof, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.SetPauserRegistry(&_ContractCoprocessorTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.SetPauserRegistry(&_ContractCoprocessorTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.SetStaleStakesForbidden(&_ContractCoprocessorTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.SetStaleStakesForbidden(&_ContractCoprocessorTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.TransferOwnership(&_ContractCoprocessorTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.TransferOwnership(&_ContractCoprocessorTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.Unpause(&_ContractCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractCoprocessorTaskManager.Contract.Unpause(&_ContractCoprocessorTaskManager.TransactOpts, newPausedStatus)
}

// ContractCoprocessorTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerInitializedIterator struct {
	Event *ContractCoprocessorTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerInitialized)
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
		it.Event = new(ContractCoprocessorTaskManagerInitialized)
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
func (it *ContractCoprocessorTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerInitialized represents a Initialized event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractCoprocessorTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerInitializedIterator{contract: _ContractCoprocessorTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerInitialized)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractCoprocessorTaskManagerInitialized, error) {
	event := new(ContractCoprocessorTaskManagerInitialized)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerOwnershipTransferredIterator struct {
	Event *ContractCoprocessorTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractCoprocessorTaskManagerOwnershipTransferred)
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
func (it *ContractCoprocessorTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractCoprocessorTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerOwnershipTransferredIterator{contract: _ContractCoprocessorTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerOwnershipTransferred)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractCoprocessorTaskManagerOwnershipTransferred, error) {
	event := new(ContractCoprocessorTaskManagerOwnershipTransferred)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerPausedIterator struct {
	Event *ContractCoprocessorTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerPaused)
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
		it.Event = new(ContractCoprocessorTaskManagerPaused)
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
func (it *ContractCoprocessorTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerPaused represents a Paused event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractCoprocessorTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerPausedIterator{contract: _ContractCoprocessorTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerPaused)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParsePaused(log types.Log) (*ContractCoprocessorTaskManagerPaused, error) {
	event := new(ContractCoprocessorTaskManagerPaused)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerPauserRegistrySetIterator struct {
	Event *ContractCoprocessorTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractCoprocessorTaskManagerPauserRegistrySet)
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
func (it *ContractCoprocessorTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractCoprocessorTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerPauserRegistrySetIterator{contract: _ContractCoprocessorTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerPauserRegistrySet)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractCoprocessorTaskManagerPauserRegistrySet, error) {
	event := new(ContractCoprocessorTaskManagerPauserRegistrySet)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractCoprocessorTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractCoprocessorTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorTaskManagerTaskBatchRegisteredIterator is returned from FilterTaskBatchRegistered and is used to iterate over the raw logs and unpacked data for TaskBatchRegistered events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerTaskBatchRegisteredIterator struct {
	Event *ContractCoprocessorTaskManagerTaskBatchRegistered // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerTaskBatchRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerTaskBatchRegistered)
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
		it.Event = new(ContractCoprocessorTaskManagerTaskBatchRegistered)
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
func (it *ContractCoprocessorTaskManagerTaskBatchRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerTaskBatchRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerTaskBatchRegistered represents a TaskBatchRegistered event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerTaskBatchRegistered struct {
	Batch ICoprocessorTaskManagerTaskBatch
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTaskBatchRegistered is a free log retrieval operation binding the contract event 0x9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd95.
//
// Solidity: event TaskBatchRegistered((uint32,uint32,bytes32,bytes,uint32) batch)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterTaskBatchRegistered(opts *bind.FilterOpts) (*ContractCoprocessorTaskManagerTaskBatchRegisteredIterator, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "TaskBatchRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerTaskBatchRegisteredIterator{contract: _ContractCoprocessorTaskManager.contract, event: "TaskBatchRegistered", logs: logs, sub: sub}, nil
}

// WatchTaskBatchRegistered is a free log subscription operation binding the contract event 0x9e93280875f7d0d786d58b17dfd709c5bd47d7802bf3ecc67d4ce2b9a8a7bd95.
//
// Solidity: event TaskBatchRegistered((uint32,uint32,bytes32,bytes,uint32) batch)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchTaskBatchRegistered(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerTaskBatchRegistered) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "TaskBatchRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerTaskBatchRegistered)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "TaskBatchRegistered", log); err != nil {
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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParseTaskBatchRegistered(log types.Log) (*ContractCoprocessorTaskManagerTaskBatchRegistered, error) {
	event := new(ContractCoprocessorTaskManagerTaskBatchRegistered)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "TaskBatchRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerTaskRespondedIterator struct {
	Event *ContractCoprocessorTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerTaskResponded)
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
		it.Event = new(ContractCoprocessorTaskManagerTaskResponded)
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
func (it *ContractCoprocessorTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerTaskResponded represents a TaskResponded event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerTaskResponded struct {
	Task     ICoprocessorTaskManagerTask
	Response ICoprocessorTaskManagerTaskResponse
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xde0a98210b193f0e75b97832743430e00f558966c6917131cee6e590de548c4d.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) task, (bytes,bytes32) response)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractCoprocessorTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerTaskRespondedIterator{contract: _ContractCoprocessorTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xde0a98210b193f0e75b97832743430e00f558966c6917131cee6e590de548c4d.
//
// Solidity: event TaskResponded((uint32,bytes,bytes) task, (bytes,bytes32) response)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerTaskResponded)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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
// Solidity: event TaskResponded((uint32,bytes,bytes) task, (bytes,bytes32) response)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractCoprocessorTaskManagerTaskResponded, error) {
	event := new(ContractCoprocessorTaskManagerTaskResponded)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCoprocessorTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerUnpausedIterator struct {
	Event *ContractCoprocessorTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractCoprocessorTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCoprocessorTaskManagerUnpaused)
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
		it.Event = new(ContractCoprocessorTaskManagerUnpaused)
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
func (it *ContractCoprocessorTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCoprocessorTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCoprocessorTaskManagerUnpaused represents a Unpaused event raised by the ContractCoprocessorTaskManager contract.
type ContractCoprocessorTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractCoprocessorTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractCoprocessorTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractCoprocessorTaskManagerUnpausedIterator{contract: _ContractCoprocessorTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractCoprocessorTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractCoprocessorTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCoprocessorTaskManagerUnpaused)
				if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractCoprocessorTaskManager *ContractCoprocessorTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractCoprocessorTaskManagerUnpaused, error) {
	event := new(ContractCoprocessorTaskManagerUnpaused)
	if err := _ContractCoprocessorTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
