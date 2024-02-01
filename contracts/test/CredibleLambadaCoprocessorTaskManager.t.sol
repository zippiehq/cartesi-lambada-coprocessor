// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/LambadaCoprocessorServiceManager.sol" as incsqsm;
import {LambadaCoprocessorTaskManager} from "../src/LambadaCoprocessorTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract LambadaCoprocessorTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.LambadaCoprocessorServiceManager sm;
    incsqsm.LambadaCoprocessorServiceManager smImplementation;
    LambadaCoprocessorTaskManager tm;
    LambadaCoprocessorTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new LambadaCoprocessorTaskManager(
            incsqsm.IBLSRegistryCoordinatorWithIndices(
                address(registryCoordinator)
            ),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = LambadaCoprocessorTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        serviceManagerOwner,
                        aggregator,
                        generator
                    )
                )
            )
        );
    }

    /*
    function testCreateRegisterBatch() public {
        TODO: implement
    }
    */
}