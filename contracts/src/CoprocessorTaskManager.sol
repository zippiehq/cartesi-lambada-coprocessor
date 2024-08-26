// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";

import "./ICoprocessorTaskManager.sol";

contract CoprocessorTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ICoprocessorTaskManager
{
    struct TaskResponseSigHashData {
        uint32 batchIndex;
        bytes programId;
        bytes inputHash;
        bytes resultCID;
        bytes32 outputHash;
    }

    using BN254 for BN254.G1Point;

    /* CONSTANT */
    // The number of blocks from the task initialization within which the aggregator has to respond to
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;

    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    /* STORAGE */
    // Index of current task batch
    uint32 public nextBatchIndex;

    mapping(uint32 => bytes32) public allBatchHashes;
    mapping(bytes32 => bytes32) public allTaskOutputs;

    address public aggregator;
    address public generator;

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    // onlyBatchGenerator is used to restrict createNewTask from only being called by a permissioned entity
    // in a real world scenario, this would be removed by instead making createNewTask a payable function
    modifier onlyBatchGenerator() {
        require(msg.sender == generator, "Task generator must be the caller");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock
    ) BLSSignatureChecker(_registryCoordinator) {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator,
        address _generator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
        generator = _generator;
    }

    function registerNewTaskBatch(
        bytes32 batchRoot,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyBatchGenerator {
        TaskBatch memory newBatch;
        newBatch.index = nextBatchIndex;
        newBatch.blockNumber = uint32(block.number);
        newBatch.merkeRoot = batchRoot;
        newBatch.quorumThresholdPercentage = quorumThresholdPercentage;
        newBatch.quorumNumbers = quorumNumbers;

        bytes32 batchHash = keccak256(abi.encode(newBatch));
        allBatchHashes[nextBatchIndex] = batchHash;

        nextBatchIndex++;

        emit TaskBatchRegistered(newBatch);
    }

    function respondTask(
        TaskBatch calldata batch,
        Task calldata task,
        bytes32[] calldata taskProof,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        bytes32 taskHash = verifyBatchTask(batch, task, taskProof);
        checkBLSSig(batch, task, taskResponse, nonSignerStakesAndSignature);

        // Mark task as responded.
        bytes32 responseHash = keccak256(abi.encode(taskResponse));
        allTaskOutputs[taskHash] = responseHash;

        emit TaskResponded(task, taskResponse);
    }

    function checkValidTaskResponse(
        TaskBatch calldata batch,
        Task calldata task,
        bytes32[] calldata taskProof,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external view {
        bytes32 taskHash = verifyBatchTask(batch, task, taskProof);
        checkBLSSig(batch, task, taskResponse, nonSignerStakesAndSignature);

        return;
    }

    function verifyBatchTask(
        TaskBatch calldata batch,
        Task calldata task,
        bytes32[] calldata taskProof
    ) internal view returns (bytes32) {
        require(
            batch.index == task.batchIndex,
            "batch.index and task.batchIndex do not match"
        );
        
        // Check that batch has been registered.
        bytes32 batchHash = keccak256(abi.encode(batch));
        require(
            allBatchHashes[batch.index] == batchHash,
            "Task batch has not been registered"
        );

        // Check if task has been already responded.
        (bytes32 taskHash, bytes32 responseHash) = taskResponseHash(task);
        require(
            responseHash == "",
            "Task response already responded"
        );

        // Check that batch contains specified task.
        bytes32 batchTaskHash = keccak256(bytes.concat(keccak256(abi.encode(task.programId, task.inputHash))));
        require(
            MerkleProof.verifyCalldata(taskProof, batch.merkeRoot, batchTaskHash),
            "Task does not belong to batch"
        );

        return taskHash;
    }

    function checkBLSSig(
        TaskBatch calldata batch,
        Task calldata task,
        TaskResponse calldata response,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) internal view {
        TaskResponseSigHashData memory hashData;
        hashData.batchIndex = task.batchIndex;
        hashData.programId = task.programId;
        hashData.inputHash = task.inputHash;
        hashData.resultCID = response.resultCID;
        hashData.outputHash = response.outputHash;
        
        bytes32 responseSigHash = keccak256(abi.encode(hashData));
        
        // Check the BLS signature.
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
            responseSigHash,
            batch.quorumNumbers,
            batch.blockNumber,
            nonSignerStakesAndSignature
        );

        // Check that signatories own at least a threshold percentage of each quourm.
        for (uint i = 0; i < batch.quorumNumbers.length; i++) {
            // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying
            // signed stake > total stake
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * _THRESHOLD_DENOMINATOR >=
                quorumStakeTotals.totalStakeForQuorum[i] * uint8(batch.quorumThresholdPercentage),
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }
    }
    function getTaskResponseHash(Task calldata task) external view returns (bytes32) {
        (bytes32 taskHash, bytes32 responseHash) = taskResponseHash(task);
        return responseHash;
    }

    function taskResponseHash(Task calldata task) internal view returns (bytes32, bytes32) {
        bytes32 taskHash = keccak256(abi.encode(task));
        return (taskHash, allTaskOutputs[taskHash]);
    }

    function getNextBatchIndex() external view returns (uint32) {
        return nextBatchIndex;
    }
}
