// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSPubkeyRegistry} from "@eigenlayer-middleware/src/BLSPubkeyRegistry.sol";
import {BLSRegistryCoordinatorWithIndices} from "@eigenlayer-middleware/src/BLSRegistryCoordinatorWithIndices.sol";
import {BLSOperatorStateRetriever} from "@eigenlayer-middleware/src/BLSOperatorStateRetriever.sol";
import "@eigenlayer/contracts/libraries/BN254.sol";
import "./ILambadaCoprocessorTaskManager.sol";

contract LambadaCoprocessorTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    BLSOperatorStateRetriever,
    ILambadaCoprocessorTaskManager
{
    using BN254 for BN254.G1Point;

    /* CONSTANT */
    // The number of blocks from the task initialization within which the aggregator has to respond to
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;

    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    /* STORAGE */
    // Index of current task batch
    uint32 public nextBatchIndex;

    mapping(uint32 => bytes32) public allBatchHashes;
    mapping(bytes32 => bool) public allTaskResponses;

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
        IBLSRegistryCoordinatorWithIndices _registryCoordinator,
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
        newBatch.quorumThresholdPercentage = quorumThresholdPercentage;
        newBatch.quorumNumbers = quorumNumbers;

        bytes32 batchHash = keccak256(abi.encode(newBatch));
        allBatchHashes[nextBatchIndex] = batchHash;

        emit TaskBatchRegistered(newBatch);

        nextBatchIndex++;
    }

    function respondTask(
        TaskBatch calldata batch,
        Task calldata task,
        bytes32[] calldata taskProof,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        bytes32 responseMetaHash = verifyBatchTask(batch, task, taskProof);

        // Check the BLS signature.
        bytes32 responseHash = keccak256(abi.encode(taskResponse));
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
            responseHash,
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
        
        // Update task responses.
        allTaskResponses[responseMetaHash] = true;
    }

    function verifyBatchTask(
        TaskBatch calldata batch,
        Task calldata task,
        bytes32[] calldata taskProof
    ) internal returns (bytes32) {
         // Check that batch has been registered.
        bytes32 batchHash = keccak256(abi.encode(batch));
        require(
            allBatchHashes[batch.index] == batchHash,
            "Task batch has not been registered"
        );

        // Check that batch contains specified task.
        bytes32 taskHash = keccak256(abi.encode(task));
        require(
            MerkleProof.verify(taskProof, batch.merkeRoot, taskHash),
            "Task does not belong to batch"
        );

        // Check if task has been already responded.
        TaskResponseMetadata memory responseMeta;
        responseMeta.batchIndex = batch.index;
        responseMeta.taskInput = task.input;
        bytes32 responseMetaHash = keccak256(abi.encode(responseMeta));
        require(
            !allTaskResponses[responseMetaHash],
            "Task batch already responded"
        );

        return responseMetaHash;
    }

    /* FUNCTIONS */
    // NOTE: this function creates new task, assigns it a taskId
    /*
    function createNewTask(
        bytes calldata input,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyBatchGenerator {
        // create a new task struct
        Task memory newTask;
        newTask.input = input;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;

        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(latestTaskNum, newTask);
        latestTaskNum = latestTaskNum + 1;
    }
    */

    // NOTE: this function responds to existing tasks.
    /*
    function respondToTask(
        Task calldata task,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;

        // check that the task is valid, hasn't been responsed yet, and is being responsed in time
        require(
            keccak256(abi.encode(task)) ==
                allTaskHashes[taskResponse.referenceTaskIndex],
            "supplied task does not match the one recorded in the contract"
        );
        // some logical checks
        require(
            allTaskResponses[taskResponse.referenceTaskIndex] == bytes32(0),
            "Aggregator has already responded to the task"
        );
        require(
            uint32(block.number) <=
                taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "Aggregator has responded to the task too late"
        );

        // CHECKING SIGNATURES & WHETHER THRESHOLD IS MET OR NOT
        // calculate message which operators signed
        bytes32 message = keccak256(abi.encode(taskResponse));

        // check the BLS signature
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
                message,
                quorumNumbers,
                taskCreatedBlock,
                nonSignerStakesAndSignature
            );

        // check that signatories own at least a threshold percentage of each quourm
        for (uint i = 0; i < quorumNumbers.length; i++) {
            // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying
            // signed stake > total stake
            require(
                quorumStakeTotals.signedStakeForQuorum[i] *
                    _THRESHOLD_DENOMINATOR >=
                    quorumStakeTotals.totalStakeForQuorum[i] *
                        uint8(quorumThresholdPercentage),
                "Signatories do not own at least threshold percentage of a quorum"
            );
        }

        TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        // updating the storage with task responsea
        allTaskResponses[taskResponse.referenceTaskIndex] = keccak256(
            abi.encode(taskResponse, taskResponseMetadata)
        );

        // emitting event
        emit TaskResponded(taskResponse, taskResponseMetadata);
    }
    */
}
