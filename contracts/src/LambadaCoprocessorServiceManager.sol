// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import {ILambadaCoprocessorTaskManager} from "./ILambadaCoprocessorTaskManager.sol";

/**
 * @title Primary entrypoint for procuring services from IncredibleSquaring.
 * @author Layr Labs, Inc.
 */
contract LambadaCoprocessorServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    ILambadaCoprocessorTaskManager
        public immutable lambadaCoprocessorTaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyLambadaCoprocessorTaskManager() {
        require(
            msg.sender == address(lambadaCoprocessorTaskManager),
            "not from lambada coprocessor task manager"
        );
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        ILambadaCoprocessorTaskManager _lambadaCoprocessorTaskManager
    )
        ServiceManagerBase(
            _avsDirectory,
            IPaymentCoordinator(address(0)), // inc-sq doesn't need to deal with payments
            _registryCoordinator,
            _stakeRegistry
        )
    {
        lambadaCoprocessorTaskManager = _lambadaCoprocessorTaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlyLambadaCoprocessorTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
