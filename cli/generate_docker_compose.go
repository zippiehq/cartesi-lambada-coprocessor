package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"

	"github.com/nikolalohinski/gonja"
	cp "github.com/otiai10/copy"
	"github.com/urfave/cli"
)

const DEVNET = "devnet"
const HOLESKY = "holesky"
const MAINNET = "mainnet"

var ANVIL_DEV_ACCOUNT_PRIVATE_KEYS = []string{
	"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
	"0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
	"0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
	"0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba",
	"0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e",
	"0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356",
	"0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97",
	"0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
}

func GenerateDockerCompose(ctx *cli.Context) error {
	network := ctx.String("network")
	if network != DEVNET && network != HOLESKY && network != MAINNET {
		return errors.New("invalid network type")
	}

	operatorCount := ctx.Uint("operators")
	if operatorCount == 0 {
		return errors.New("number of operators must be greater then 0")
	}
	if operatorCount > 10 {
		return errors.New("number of operators must be less then 10")
	}

	// Read deployment parameters and output
	var deploymentParamPath, deploymentOuputPath string
	if network == DEVNET {
		deploymentParamPath = "./contracts/script/input/deployment_parameters_devnet.json"
		deploymentOuputPath = "./contracts/script/output/lambada_coprocessor_deployment_output_devnet.json"
	} else {
		deploymentParamPath = "./contracts/script/input/deployment_parameters_holesky.json"
		deploymentOuputPath = "./contracts/script/output/lambada_coprocessor_deployment_output_holesky.json"
	}
	var deploymentParams config.AVSDeploymentParameters
	if err := sdkutils.ReadJsonConfig(deploymentParamPath, &deploymentParams); err != nil {
		return fmt.Errorf("failed to read deployment parameters - %s", err)
	}
	var deploymentOutput config.AVSDeployment
	if err := sdkutils.ReadJsonConfig(deploymentOuputPath, &deploymentOutput); err != nil {
		return fmt.Errorf("failed to read deployment output - %s", err)
	}

	// Clear "operators" directory
	if _, err := os.Stat("./tests/nodes/operators"); err != nil {
		return err
	} else {
		if err = os.RemoveAll("./tests/nodes/operators"); err != nil {
			return err
		}
	}

	// Create directory for each operator
	operatorDirs := make([]string, int(operatorCount))
	for i := range operatorDirs {
		operatorDirs[i] = fmt.Sprintf("./tests/nodes/operators/%d", i+1)
		if err := os.MkdirAll(operatorDirs[i], os.ModePerm); err != nil {
			return err
		}
	}

	// Generate BLS and ECDSA keys
	blsKeys := make([]OperatorBLSKey, int(operatorCount))
	ecdsaKeys := make([]OperatorECDSAKey, int(operatorCount))
	for i := 0; i < int(operatorCount); i++ {
		var err error
		if blsKeys[i], err = GenerateBLSKey(operatorDirs[i]); err != nil {
			return err
		}
		if ecdsaKeys[i], err = GenerateECDSAKey(operatorDirs[i]); err != nil {
			return err
		}
	}

	// Generate configuration files for each operator
	configs := make([]string, operatorCount)
	for i := range configs {
		configs[i] = filepath.Join(operatorDirs[i], "config.yaml")
		tmpl, err := gonja.FromFile("./tests/jinja/operator-docker-compose.j2")
		if err != nil {
			return err
		}
		machine := fmt.Sprintf("machine%d", i+1)
		config, err := tmpl.Execute(gonja.Context{
			"address":                           ecdsaKeys[i].Address,
			"registry_coordinator_address":      deploymentOutput.Addresses.RegistryCoordinator,
			"operator_state_retriever_address":  deploymentOutput.Addresses.OperatorStateRetriever,
			"ecdsa_private_key_store_path":      ecdsaKeys[i].FilePath,
			"bls_private_key_store_path":        blsKeys[i].FilePath,
			"eth_rpc_url":                       "http://anvil:8545",
			"eth_ws_url":                        "ws://anvil:8545",
			"aggregator_server_ip_port_address": "lambada-coprocessor-aggregator:8090",
			"ipfs_ip_port_address":              fmt.Sprintf("%s:5001", machine),
			"lambada_ip_port_address":           fmt.Sprintf("%s:3033", machine),
			"eigen_metrics_ip_port_address":     "0.0.0.0:9090",
			"node_api_ip_port_address":          "0.0.0.0:9010",
		})
		if err != nil {
			return err
		}
		configFile, err := os.Create(configs[i])
		if err != nil {
			return err
		}
		defer configFile.Close()
		if _, err = configFile.WriteString(config); err != nil {
			return err
		}
	}

	// Generate startup scripts
	var strategy_address string
	if network == DEVNET {
		strategy_address = deploymentOutput.Addresses.ERC20MockStrategy
	} else {
		strategy_address = deploymentParams.WETH
	}
	runScripts := make([]string, operatorCount)
	for i := range runScripts {
		runScripts[i] = filepath.Join(operatorDirs[i], "run.sh")
		tmpl, err := gonja.FromFile("./tests/jinja/run-operator.j2")
		if err != nil {
			return err
		}
		script, err := tmpl.Execute(gonja.Context{
			"operator_delay":             i,
			"deployment_parameters_path": deploymentParamPath,
			"network":                    network,
			"funder_private_key":         ANVIL_DEV_ACCOUNT_PRIVATE_KEYS[i],
			"operator_private_key":       ecdsaKeys[i].PrivateKey,
			"operator_address":           ecdsaKeys[i].Address,
			"operator_bls_password":      blsKeys[i].Password,
			"operator_ecdsa_password":    ecdsaKeys[i].Password,
			"operator_config":            configs[i],
			"strategy_address":           strategy_address,
			"strategy_deposit_amount":    10,
			"erc20_mock_address":         deploymentOutput.Addresses.ERC20Mock,
		})
		if err != nil {
			return err
		}
		scriptFile, err := os.Create(runScripts[i])
		if err != nil {
			return err
		}
		defer scriptFile.Close()
		if _, err = scriptFile.WriteString(script); err != nil {
			return err
		}
	}

	// Generate directories for machine data
	machineData := make([]string, operatorCount)
	for i := range machineData {
		machineData[i] = filepath.Join(operatorDirs[i], "data")
		if err := cp.Copy("./machine/data", machineData[i]); err != nil {
			return err
		}
	}

	// Docker compose file
	operators := make([]map[string]interface{}, operatorCount)
	for i := 1; i <= int(operatorCount); i++ {
		machine := fmt.Sprintf("machine%d", i)
		operators[i-1] = map[string]interface{}{
			"name":         fmt.Sprintf("operator%d", i),
			"machine":      machine,
			"machine_data": machineData[i-1],
			"run_script":   runScripts[i-1],
		}
	}
	cofingTmpl, err := gonja.FromFile("./tests/jinja/docker-compose.j2")
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	compose, err := cofingTmpl.Execute(gonja.Context{"network": network, "operators": operators})
	if err != nil {
		return err
	}

	var composeFilePath string
	switch network {
	case DEVNET:
		composeFilePath = "./docker-compose.yaml"
	case HOLESKY:
		composeFilePath = "./docker-compose-holesky.yaml"
	case MAINNET:
		composeFilePath = "./docker-compose-mainnet.yaml"
	}
	composeFile, err := os.Create(composeFilePath)
	if err != nil {
		return err
	}
	defer composeFile.Close()
	if _, err = composeFile.WriteString(compose); err != nil {
		return err
	}

	return nil
}
