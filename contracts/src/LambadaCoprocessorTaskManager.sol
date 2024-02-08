// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/console.sol";
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
        newBatch.merkeRoot = batchRoot;
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

        // !!!
        console.log("smart contract - task proof:");
        for (uint i = 0; i < taskProof.length; i++) {
            console.logBytes32(taskProof[i]);
        }
        console.log("smart contract - batch root:");
        console.logBytes32(batch.merkeRoot);
        console.log("smart contract - task hash:");
        console.logBytes32(taskHash);
        
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
            "Task response already responded"
        );

        return responseMetaHash;
    }
}
