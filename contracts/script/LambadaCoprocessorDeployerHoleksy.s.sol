// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {LambadaCoprocessorDeployer} from "./LambadaCoprocessorDeployer.s.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

contract LambadaCoprocessorDeployerHolesky is LambadaCoprocessorDeployer {
    function run() external {
        string memory paramFile = "./script/input/deployment_parameters_holesky.json";
        (EigenLayerContracts memory eigenLayer, DeploymentConfig memory config)
             = readDeploymentParameters(paramFile);

        StrategyConfig[] memory strategyConfig = new StrategyConfig[](1);
        {
            strategyConfig[0].strategy = eigenLayer.wETH;
            strategyConfig[0].weight = eigenLayer.wETH_Multiplier;
        }

        vm.startBroadcast();

        string calldata outputPath = "./script/output/lambada_coprocessor_deployment_output_holesky.json";
        deployAVS(
            eigenLayer,
            config,
            strategyConfig,
            outputPath
        );

        vm.stopBroadcast();
    }
}