// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import "@eigenlayer/contracts/core/StrategyManager.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";

import "../src/ERC20Mock.sol";
import {Utils} from "./utils/Utils.sol";
import {CoprocessorDeployer} from "./CoprocessorDeployer.s.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract CoprocessorDeployerDevnet is CoprocessorDeployer {
    function run() external {
        string memory paramFile = "./script/input/deployment_parameters_devnet.json";
        (EigenLayerContracts memory eigenLayer, DeploymentConfig memory config)
             = readDeploymentParameters(paramFile);

        vm.startBroadcast();

        (address erc20MockStrategy, address erc20Mock) = deployERC20MockStrategy(
            eigenLayer.proxyAdmin, 
            eigenLayer.pauserRegistry,
            eigenLayer.baseStrategy,
            eigenLayer.strategyManager
        );

        StrategyConfig[] memory strategyConfig = new StrategyConfig[](1);
        {
            strategyConfig[0].strategy = erc20MockStrategy;
            strategyConfig[0].weight = eigenLayer.wETH_Multiplier;
        }
        CoprocessorContracts memory contracts = deployAVS(eigenLayer, config, strategyConfig);

        vm.stopBroadcast();

        AuxContract[] memory auxContracts = new AuxContract[](2);
        {
            auxContracts[0].name = "erc20Mock";
            auxContracts[0].addr = erc20Mock;
            auxContracts[1].name = "erc20MockStrategy";
            auxContracts[1].addr = erc20MockStrategy;
        }
        string memory outputPath = "./script/output/coprocessor_deployment_output_devnet.json";
        writeDeploymentOutput(contracts, auxContracts, outputPath);
    }

    function deployERC20MockStrategy(
        ProxyAdmin eigenLayerProxyAdmin,
        PauserRegistry eigenLayerPauserReg,
        StrategyBaseTVLLimits baseStrategyImplementation,
        IStrategyManager strategyManager
    ) internal returns (address, address) {
        ERC20Mock erc20Mock = new ERC20Mock();
        // TODO(samlaf): any reason why we are using the strategybase with tvl limits instead of just using strategybase?
        // the maxPerDeposit and maxDeposits below are just arbitrary values.
        StrategyBaseTVLLimits erc20MockStrategy = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        StrategyBaseTVLLimits.initialize.selector,
                        1 ether, // maxPerDeposit
                        100 ether, // maxDeposits
                        IERC20(erc20Mock),
                        eigenLayerPauserReg
                    )
                )
            )
        );
        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = erc20MockStrategy;
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        thirdPartyTransfersForbiddenValues[0] = false;
        strategyManager.addStrategiesToDepositWhitelist(
            strats,
            thirdPartyTransfersForbiddenValues
        );

        return (address(erc20MockStrategy), address(erc20Mock));
    }
}