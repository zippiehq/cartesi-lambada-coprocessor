package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"

	"github.com/nikolalohinski/gonja"
	cp "github.com/otiai10/copy"
	"github.com/urfave/cli"
)

const (
	DEVNET  = "devnet"
	HOLESKY = "holesky"
	MAINNET = "mainnet"
)

var (
	NETWORKS = []string{DEVNET, HOLESKY, MAINNET}

	DEPLOYMENT_PARAM_PATHS = map[string]string{
		DEVNET:  "./contracts/script/input/deployment_parameters_devnet.json",
		HOLESKY: "./contracts/script/input/deployment_parameters_holesky.json",
		MAINNET: "./contracts/script/input/deployment_parameters_mainnet.json",
	}

	DEPLOYMENT_OUTPUT_PATHS = map[string]string{
		DEVNET:  "./contracts/script/output/lambada_coprocessor_deployment_output_devnet.json",
		HOLESKY: "./contracts/script/output/lambada_coprocessor_deployment_output_holesky.json",
		MAINNET: "./contracts/script/output/lambada_coprocessor_deployment_output_mainnet.json",
	}

	COMPOSE_FILE_PATHS = map[string]string{
		DEVNET:  "./docker-compose.yaml",
		HOLESKY: "./docker-compose-holesky.yaml",
		MAINNET: "./docker-compose-mainnet.yaml",
	}

	ANVIL_DEV_ACCOUNT_PRIVATE_KEYS = []string{
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
)

func GenerateDockerCompose(ctx *cli.Context) error {
	operatorCount := ctx.Uint("operators")
	if operatorCount == 0 {
		return errors.New("number of operators must be greater then 0")
	}
	if operatorCount > 10 {
		return errors.New("number of operators must be less then 10")
	}

	deploymentParams := map[string]config.AVSDeploymentParameters{}
	for network, paramPath := range DEPLOYMENT_PARAM_PATHS {
		var params config.AVSDeploymentParameters
		if err := sdkutils.ReadJsonConfig(paramPath, &params); err != nil {
			return fmt.Errorf("failed to read deployment parameters - %s", err)
		}
		deploymentParams[network] = params
	}

	var devnetDeploymentOutput chainio.AVSDeployment
	if err := sdkutils.ReadJsonConfig(DEPLOYMENT_OUTPUT_PATHS[DEVNET], &devnetDeploymentOutput); err != nil {
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

	// Read keystore passwords
	blsPwds, err := readLines("./tests/nodes/operators-keystore/bls/password.txt")
	if err != nil {
		return err
	}
	ecdsaPwds, err := readLines("./tests/nodes/operators-keystore/ecdsa/password.txt")
	if err != nil {
		return err
	}
	// Read private keys
	blsPrivKeys, err := readLines("./tests/nodes/operators-keystore/bls/private_key_hex.txt")
	if err != nil {
		return err
	}
	ecdsaPrivKeys, err := readLines("./tests/nodes/operators-keystore/ecdsa/private_key_hex.txt")
	if err != nil {
		return err
	}

	// Fill operator keystore information
	blsKeys := make([]OperatorBLSKey, int(operatorCount))
	ecdsaKeys := make([]OperatorECDSAKey, int(operatorCount))
	for i := 0; i < int(operatorCount); i++ {
		blsKeys[i] = OperatorBLSKey{
			FilePath:   fmt.Sprintf("./tests/nodes/operators-keystore/bls/keys/%d.bls.key.json", i+1),
			Password:   blsPwds[i],
			PrivateKey: blsPrivKeys[i],
		}
		ecdsaKeys[i] = OperatorECDSAKey{
			FilePath:   fmt.Sprintf("./tests/nodes/operators-keystore/ecdsa/keys/%d.ecdsa.key.json", i+1),
			Password:   ecdsaPwds[i],
			PrivateKey: ecdsaPrivKeys[i],
		}
		ecdsaKey, err := ecdsa.ReadKey(ecdsaKeys[i].FilePath, ecdsaKeys[i].Password)
		if err != nil {
			return err
		}
		ecdsaKeys[i].Address = crypto.PubkeyToAddress(ecdsaKey.PublicKey).Hex()
	}

	// Generate configuration files for each operator
	operatorConfigs := map[string][]string{}
	for _, network := range NETWORKS {
		configs := make([]string, operatorCount)
		for i := range configs {
			configs[i] = filepath.Join(operatorDirs[i], fmt.Sprintf("config-%s.yaml", network))
			tmpl, err := gonja.FromFile("./tests/jinja/operator-docker-compose.j2")
			if err != nil {
				return err
			}
			machine := fmt.Sprintf("machine%d", i+1)
			config, err := tmpl.Execute(gonja.Context{
				"avs_deployment_path":               DEPLOYMENT_OUTPUT_PATHS[network],
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
		operatorConfigs[network] = configs
	}

	// Generate operator startup scripts
	strategy_address := map[string]string{
		DEVNET:  devnetDeploymentOutput.Addresses.ERC20MockStrategy,
		HOLESKY: deploymentParams[HOLESKY].WETH,
		MAINNET: deploymentParams[MAINNET].RETH,
	}
	operatorScripts := map[string][]string{}
	for _, network := range NETWORKS {
		scripts := make([]string, operatorCount)
		for i := range scripts {
			scripts[i] = filepath.Join(operatorDirs[i], fmt.Sprintf("run-%s.sh", network))
			tmpl, err := gonja.FromFile("./tests/jinja/run-operator.j2")
			if err != nil {
				return err
			}
			script, err := tmpl.Execute(gonja.Context{
				"operator_private_key":    ecdsaKeys[i].PrivateKey,
				"operator_address":        ecdsaKeys[i].Address,
				"operator_bls_password":   blsKeys[i].Password,
				"operator_ecdsa_password": ecdsaKeys[i].Password,
				"operator_config":         operatorConfigs[network][i],
				"strategy_address":        strategy_address[network],
				"strategy_deposit_amount": 10,
			})
			if err != nil {
				return err
			}
			scriptFile, err := os.Create(scripts[i])
			if err != nil {
				return err
			}
			defer scriptFile.Close()
			if _, err = scriptFile.WriteString(script); err != nil {
				return err
			}
			exec.Command("chmod", "+x", scripts[i])
		}
		operatorScripts[network] = scripts
	}

	// Generate anvil startup scripts
	operators := make([]map[string]interface{}, operatorCount)
	for i := 1; i <= int(operatorCount); i++ {
		operators[i-1] = map[string]interface{}{
			"private_key": ecdsaKeys[i-1].PrivateKey,
			"address":     ecdsaKeys[i-1].Address,
		}
	}
	for _, network := range NETWORKS {
		tmpl, err := gonja.FromFile(fmt.Sprintf("./tests/jinja/run-anvil-%s.j2", network))
		if err != nil {
			return err
		}
		script, err := tmpl.Execute(gonja.Context{
			"funder_private_key":      ANVIL_DEV_ACCOUNT_PRIVATE_KEYS[0],
			"strategy_deposit_amount": 10,
			"erc20_mock_address":      devnetDeploymentOutput.Addresses.ERC20Mock,
			"operators":               operators,
		})
		if err != nil {
			return err
		}
		scriptFile, err := os.Create(fmt.Sprintf("./tests/anvil/run-anvil-%s.sh", network))
		if err != nil {
			return err
		}
		defer scriptFile.Close()
		if _, err = scriptFile.WriteString(script); err != nil {
			return err
		}
		exec.Command("chmod", "+x", script)
	}

	// Generate directories for machine data
	machineData := make([]string, operatorCount)
	for i := range machineData {
		machineData[i] = filepath.Join(operatorDirs[i], "data")
		if err := cp.Copy("./machine/data", machineData[i]); err != nil {
			return err
		}
	}

	// Generate compose files
	composeOperators := map[string][]map[string]interface{}{}
	for _, network := range NETWORKS {
		operators := make([]map[string]interface{}, operatorCount)
		for i := 1; i <= int(operatorCount); i++ {
			operators[i-1] = map[string]interface{}{
				"name":           fmt.Sprintf("operator%d", i),
				"machine":        fmt.Sprintf("machine%d", i),
				"machine_data":   machineData[i-1],
				"run_script":     operatorScripts[network][i-1],
				"host_ipfs_port": 5000 + i,
			}
		}
		composeOperators[network] = operators
	}
	for network, filePath := range COMPOSE_FILE_PATHS {
		tmpl, err := gonja.FromFile("./tests/jinja/docker-compose.j2")
		if err != nil {
			return err
		}
		compose, err := tmpl.Execute(gonja.Context{
			"network":   network,
			"operators": composeOperators[network],
		})
		if err != nil {
			return err
		}
		composeFile, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer composeFile.Close()
		if _, err = composeFile.WriteString(compose); err != nil {
			return err
		}
	}

	return nil
}

func readLines(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
