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

const ANVIL_DEV_ACCOUNT_PRIVATE_KEY = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

func GenerateDockerCompose(ctx *cli.Context) error {
	network := ctx.String("network")
	if network != DEVNET && network != HOLESKY {
		return errors.New("invalid network type")
	}

	operatorCount := ctx.Uint("operators")
	if operatorCount == 0 {
		return errors.New("number of operators must be greater then 0")
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
			"funder_private_key":         ANVIL_DEV_ACCOUNT_PRIVATE_KEY,
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
	compose, err := cofingTmpl.Execute(gonja.Context{"operators": operators})
	if err != nil {
		return err
	}
	composeFile, err := os.Create("./docker-compose.yaml")
	if err != nil {
		return err
	}
	defer composeFile.Close()
	if _, err = composeFile.WriteString(compose); err != nil {
		return err
	}

	return nil
}
