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
            incsqsm.IRegistryCoordinator(address(registryCoordinator)),
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
                        registryCoordinatorOwner,
                        aggregator,
                        generator
                    )
                )
            )
        );
    }

    function testRegisterNewTaskBatch() public {
        bytes32 batchRoot = 0xdf10939438762484978d10ad34e9186e9b3bd1d2f4b1fef91dd1b245b02818b8;
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator, generator);
        tm.registerNewTaskBatch(batchRoot, 100, quorumNumbers);
        assertEq(tm.nextBatchIndex(), 1);
    }
}
