// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BN254.sol";
import "@eigenlayer-middleware/src/BLSSignatureChecker.sol";

interface ILambadaCoprocessorTaskManager is IBLSSignatureChecker {
    event TaskBatchRegistered(TaskBatch batch);
    event TaskResponded(TaskResponse taskResponse);

    struct TaskBatch {
        uint32 index;
        uint32 blockNumber;
        bytes32 merkeRoot;
        // batch submitter decides on the criteria for a tasks in batch to be completed
        // note that this does not mean the task was "correctly" answered (i.e. the number was squared correctly)
        //      this is for the challenge logic to verify
        // task is completed (and contract will accept its TaskResponse) when each quorumNumbers specified here
        // are signed by at least quorumThresholdPercentage of the operators
        // note that we set the quorumThresholdPercentage to be the same for all quorumNumbers, but this could be changed
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
    }

    struct Task {
        bytes programId;
        bytes input;
    }

    // Task response is hashed and signed by operators.
    // these signatures are aggregated and sent to the contract as response.
    struct TaskResponse {
        // This is just the response that the operator has to compute by itself.
        bytes32 outputHash;
    }

    struct TaskResponseMetadata {
        uint32 batchIndex;
        bytes programId;
        bytes taskInput;
    }
  
    function registerNewTaskBatch(
        bytes32 batchRoot,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external;

    function respondTask(
        TaskBatch calldata batch,
        Task calldata task,
        bytes32[] calldata taskProof,
        TaskResponse calldata taskResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external;
}
