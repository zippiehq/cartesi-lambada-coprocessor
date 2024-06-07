// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import "@eigenlayer/contracts/core/DelegationManager.sol";
import {IAVSDirectory, AVSDirectory} from "@eigenlayer/contracts/core/AVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import "@eigenlayer/contracts/core/StrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry, IRegistryCoordinator, RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {LambadaCoprocessorServiceManager, IServiceManager} from "../src/LambadaCoprocessorServiceManager.sol";
import {LambadaCoprocessorTaskManager, ILambadaCoprocessorTaskManager} from "../src/LambadaCoprocessorTaskManager.sol";
import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

/*
forge script script/LambadaCoprocessorDeployerHolesky.s.sol:LambadaCoprocessorDeployerHolesky --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv

For local fork use 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 (anvil dev account).
For real network - specifiy private key with of real account with non-zero balance.
*/

// 
contract LambadaCoprocessorDeployerHolesky is Script, Utils {
    struct EigenLayerContracts {
        IAVSDirectory avsDirectory;
        DelegationManager delegationManager;
        StrategyManager strategyManager;
        address wETH;
        uint96 wETH_Multiplier;
    }

    struct LambadaCoprocessorContracts {
        LambadaCoprocessorTaskManager taskManager;
        LambadaCoprocessorTaskManager taskManagerImplementation;
        LambadaCoprocessorServiceManager serviceManager;
        LambadaCoprocessorServiceManager serviceManagerImplementation;
        IRegistryCoordinator registryCoordinator;
        IRegistryCoordinator registryCoordinatorImplementation;
        IIndexRegistry indexRegistry;
        IIndexRegistry indexRegistryImplementation;
        IStakeRegistry stakeRegistry;
        IStakeRegistry stakeRegistryImplementation;
        BLSApkRegistry apkRegistry;
        BLSApkRegistry apkRegistryImplementation;
        OperatorStateRetriever operatorStateRetriever;
    }

    struct TokenAndWeight {
        address token;
        uint96 weight;
    }

    struct DeploymentConfig {
        uint32 taskResponseWindowBlock;
        address taskGenerator;
        address aggregator;
        address communityMultisig;
        address pauser;
        address churner;
        address ejector;
        address whitelister;
        address confirmer;
        uint256 numQuorum;
        uint32 maxOperatorCount;
        uint16 kickBIPsOfOperatorStake; // an operator needs to have kickBIPsOfOperatorStake / 10000 times the stake of the operator with the least stake to kick them out
        uint16 kickBIPsOfTotalStake; // an operator needs to have less than kickBIPsOfTotalStake / 10000 of the total stake to be kicked out
        uint96 minimumStake;
    }

    function run() external {
        EigenLayerContracts memory eigenLayer;
        DeploymentConfig memory config;

        {
            string memory configData = vm.readFile("./script/input/parameters.holesky.json");
            
            eigenLayer.strategyManager = StrategyManager(stdJson.readAddress(configData, ".strategyManager"));
            eigenLayer.avsDirectory = AVSDirectory(stdJson.readAddress(configData, ".avsDirectory"));
            eigenLayer.delegationManager = DelegationManager(stdJson.readAddress(configData, ".delegationManager"));
            eigenLayer.wETH = stdJson.readAddress(configData, ".wETH");
            eigenLayer.wETH_Multiplier = uint96(stdJson.readUint(configData, ".wETH_Multiplier"));

            {
                config.taskResponseWindowBlock = uint32(stdJson.readUint(configData, ".taskResponseWindowBlock"));
                config.taskGenerator = stdJson.readAddress(configData, ".taskGenerator");
                config.aggregator = stdJson.readAddress(configData, ".aggregator");
                config.communityMultisig = stdJson.readAddress(configData, ".owner");
                config.churner = stdJson.readAddress(configData, ".churner");
                config.ejector = stdJson.readAddress(configData, ".ejector");
                config.confirmer = stdJson.readAddress(configData, ".confirmer");
                config.whitelister = stdJson.readAddress(configData, ".whitelister");
            }
        }

        config.numQuorum = 1;
        config.maxOperatorCount = 50;
        config.kickBIPsOfOperatorStake = 11000;
        config.kickBIPsOfTotalStake = 1001;
        config.minimumStake = 0;

        TokenAndWeight[] memory strategyConfig = new TokenAndWeight[](1);
        {
            strategyConfig[0].token = eigenLayer.wETH;
            strategyConfig[0].weight = eigenLayer.wETH_Multiplier;
        }

       vm.startBroadcast();

        // deploy proxy admin for ability to upgrade proxy contracts
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        
        // deploy pauser registry
        PauserRegistry pauserRegistry;
        {
            address[] memory pausers = new address[](1);
            pausers[0] = config.communityMultisig;
            pauserRegistry = new PauserRegistry(pausers, config.communityMultisig);
        }

        EmptyContract emptyContract = new EmptyContract();

        LambadaCoprocessorContracts memory contracts;
        
        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        contracts.indexRegistry = IIndexRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );

        contracts.stakeRegistry = IStakeRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin),""))
        );
        contracts.apkRegistry = BLSApkRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        contracts.registryCoordinator = RegistryCoordinator(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        contracts.taskManager = LambadaCoprocessorTaskManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        contracts.serviceManager = LambadaCoprocessorServiceManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );       

        contracts.operatorStateRetriever = new OperatorStateRetriever();

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        contracts.indexRegistryImplementation = new IndexRegistry(contracts.registryCoordinator);
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(contracts.indexRegistry))),
            address(contracts.indexRegistryImplementation)
        );

        contracts.stakeRegistryImplementation = new StakeRegistry(
            contracts.registryCoordinator, eigenLayer.delegationManager
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(contracts.stakeRegistry))),
            address(contracts.stakeRegistryImplementation)
        );

        contracts.apkRegistryImplementation = new BLSApkRegistry(contracts.registryCoordinator);
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(contracts.apkRegistry))),
            address(contracts.apkRegistryImplementation)
        );

        contracts.registryCoordinatorImplementation = new RegistryCoordinator(
            contracts.serviceManager,
            IStakeRegistry(address(contracts.stakeRegistry)),
            IBLSApkRegistry(address(contracts.apkRegistry)),
            IndexRegistry(address(contracts.indexRegistry))
        );

        {
            // for each quorum to setup, we need to define
            // QuorumOperatorSetParam, minimumStakeForQuorum, and strategyParams
            IRegistryCoordinator.OperatorSetParam[] memory quorumsOperatorSetParams
                = new IRegistryCoordinator.OperatorSetParam[](config.numQuorum);
            for (uint i = 0; i < config.numQuorum; i++) {
                // hard code these for now
                quorumsOperatorSetParams[i] = IRegistryCoordinator.OperatorSetParam({
                    maxOperatorCount: config.maxOperatorCount,
                    kickBIPsOfOperatorStake: config.kickBIPsOfOperatorStake,
                    kickBIPsOfTotalStake: config.kickBIPsOfTotalStake
                });
            }

            // set to 0 for every quorum
            uint96[] memory minimumStakeForQuourm = new uint96[](config.numQuorum);
            for (uint256 i = 0; i < config.numQuorum; i++) {
                minimumStakeForQuourm[i] = config.minimumStake;
            }

            IStakeRegistry.StrategyParams[][] memory strategyParams =
                new IStakeRegistry.StrategyParams[][](config.numQuorum);
            for (uint i = 0; i < config.numQuorum; i++) {
                IStakeRegistry.StrategyParams[] memory params =
                    new IStakeRegistry.StrategyParams[](strategyConfig.length);
                for (uint j = 0; j < strategyConfig.length; j++) {
                    params[j] = IStakeRegistry
                        .StrategyParams({
                            strategy: IStrategy(strategyConfig[j].token),
                            multiplier: strategyConfig[j].weight
                        });
                }
                strategyParams[i] = params;
            }

            // initialize registry coordinator
            proxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(contracts.registryCoordinator))),
                address(contracts.registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    RegistryCoordinator.initialize.selector,
                    config.communityMultisig,
                    config.churner,
                    config.ejector,
                    pauserRegistry,
                    0, // 0 initialPausedStatus means everything unpaused
                    quorumsOperatorSetParams,
                    minimumStakeForQuourm,
                    strategyParams
                )
            );
        }

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        contracts.taskManagerImplementation = new LambadaCoprocessorTaskManager(
            contracts.registryCoordinator,
            config.taskResponseWindowBlock
        );
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(contracts.taskManager))),
            address(contracts.taskManagerImplementation),
            abi.encodeWithSelector(
                LambadaCoprocessorTaskManager.initialize.selector,
                pauserRegistry,
                config.communityMultisig,
                config.aggregator,
                config.taskGenerator
            )
        );
        
        contracts.serviceManagerImplementation = new LambadaCoprocessorServiceManager(
            eigenLayer.avsDirectory,
            contracts.registryCoordinator,
            contracts.stakeRegistry,
            contracts.taskManager
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(contracts.serviceManager))),
            address(contracts.serviceManagerImplementation)
        );

        vm.stopBroadcast();

        string memory parent_object = "parent object";
        string memory addresses = "addresses";
        vm.serializeAddress(addresses, "taskManager", address(contracts.taskManager));
        vm.serializeAddress(addresses, "serviceManagerImpl", address(contracts.taskManagerImplementation));
        vm.serializeAddress(addresses, "serviceManager", address(contracts.serviceManager));
        vm.serializeAddress(addresses, "serviceManagerImpl", address(contracts.serviceManagerImplementation));
        vm.serializeAddress(addresses, "registryCoordinator", address(contracts.registryCoordinator));
        vm.serializeAddress(addresses, "registryCoordinatorImpl", address(contracts.registryCoordinatorImplementation));
        vm.serializeAddress(addresses, "indexRegistry", address(contracts.indexRegistry));
        vm.serializeAddress(addresses, "indexRegistryImpl", address(contracts.indexRegistryImplementation));
        vm.serializeAddress(addresses, "stakeRegistry", address(contracts.stakeRegistry));
        vm.serializeAddress(addresses, "stakeRegistryImpl", address(contracts.stakeRegistryImplementation));
        vm.serializeAddress(addresses, "apkRegistry", address(contracts.apkRegistry));
        vm.serializeAddress(addresses, "apkRegistryImpl", address(contracts.apkRegistryImplementation));
        vm.serializeAddress(addresses, "operatorStateRetriever", address(contracts.operatorStateRetriever));
        vm.serializeAddress(addresses, "pauserRegistry", address(pauserRegistry));
        vm.serializeAddress(addresses, "proxyAdmin", address(proxyAdmin));
        string memory addresses_output = vm.serializeAddress(addresses, "emptyContract", address(emptyContract));
        string memory finalJson = vm.serializeString(parent_object, addresses, addresses_output);
        vm.writeJson(finalJson, "./script/output/lambada_coprocessor_deployment_output.holesky.json");
    }
}