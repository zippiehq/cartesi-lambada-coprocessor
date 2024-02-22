package actions

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nikolalohinski/gonja"
	cp "github.com/otiai10/copy"
	"github.com/urfave/cli"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
	"github.com/zippiehq/cartesi-lambada-coprocessor/types"
)

// run beofer executing:
// anvil --load-state tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json --dump-state docker-compose/anvil-state.json

func GenerateDockerCompose(ctx *cli.Context) error {
	operatorCount := ctx.Uint("operators")
	if operatorCount == 0 {
		return errors.New("number of operators must be greater then 0")
	}

	// Clear "operators" directory
	if _, err := os.Stat("./docker-compose/operators"); err == nil {
		if err = os.RemoveAll("./docker-compose/operators"); err != nil {
			return err
		}
	}
	if err := os.Mkdir("./docker-compose/operators", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("./docker-compose/operators/keys", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("./docker-compose/operators/configs", os.ModePerm); err != nil {
		return err
	}
	if err := os.Mkdir("./docker-compose/operators/data", os.ModePerm); err != nil {
		return err
	}

	// Generate BLS and ecdsa keys.
	egnkeyCmd := fmt.Sprintf(
		"cd ./docker-compose/operators/keys && egnkey generate --key-type both --num-keys %d",
		operatorCount,
	)
	if _, _, err := runCommand(egnkeyCmd); err != nil {
		return err
	}

	// Read BLS and ECDSA keys.
	var blsKeyDir, ecdsaKeyDir string
	blsRegex, err := regexp.Compile("bls-")
	if err != nil {
		return err
	}
	ecdsaRegex, err := regexp.Compile("ecdsa-")
	if err != nil {
		return err
	}
	if err = filepath.Walk("./docker-compose/operators/keys", func(path string, info os.FileInfo, err error) error {
		if blsRegex.MatchString(info.Name()) {
			blsKeyDir = filepath.Join("./docker-compose/operators/keys", info.Name())
		}
		if ecdsaRegex.MatchString(info.Name()) {
			ecdsaKeyDir = filepath.Join("./docker-compose/operators/keys", info.Name())
		}
		return nil
	}); err != nil {
		return err
	}
	blsPwds, _, _, blsKeyPaths, err := readKeyDir(blsKeyDir)
	if err != nil {
		return err
	}
	ecdsaPwds, _, ecdsaKeys, ecdsaKeyPaths, err := readKeyDir(ecdsaKeyDir)
	if err != nil {
		return err
	}
	ecdsaAddrs := make([]string, len(ecdsaKeys))
	for i, keyData := range ecdsaKeys {
		key := struct {
			Address string `json:"address"`
		}{}

		if err := json.Unmarshal(keyData, &key); err != nil {
			return err
		}
		ecdsaAddrs[i] = fmt.Sprintf("0x%s", key.Address)
	}

	// Generate configuration files for each operator.
	configPaths := make([]string, operatorCount)
	for i := 0; i < int(operatorCount); i++ {
		configPaths[i] = filepath.Join(
			"./docker-compose/operators/configs",
			fmt.Sprintf("operator%d.yaml", i+1),
		)
		tmpl, err := gonja.FromFile("./docker-compose/operator-docker-compose.j2")
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		config, err := tmpl.Execute(gonja.Context{
			"address":        ecdsaAddrs[i],
			"ecdsa_key_path": ecdsaKeyPaths[i],
			"bls_key_path":   blsKeyPaths[i],
		})
		if err != nil {
			return err
		}
		configFile, err := os.Create(configPaths[i])
		if err != nil {
			return err
		}
		defer configFile.Close()
		if _, err = configFile.WriteString(config); err != nil {
			return err
		}
	}

	// Generate directories for docker volume mount.
	dataPaths := make([]string, operatorCount)
	for i := 0; i < int(operatorCount); i++ {
		dataPaths[i] = fmt.Sprintf("./docker-compose/operators/data/operator%d", i+1)
		if err := cp.Copy("./machine/data", dataPaths[i]); err != nil {
			return err
		}
	}

	// Update Anvil snapshot.

	// TODO: start anvil and terminate it gracefully?

	// Send eth to operator accounts.
	devAccount := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	for i := 0; i < int(operatorCount); i++ {
		sendFundsCmd := fmt.Sprintf(
			"cast send %s --value 10ether --private-key %s",
			ecdsaAddrs[i], devAccount,
		)
		if _, _, err := runCommand(sendFundsCmd); err != nil {
			return err
		}
	}

	// Operator must deposit into stragtegy, otherwise registration will fail.
	for i, path := range configPaths {
		if err := depositIntoStrategy(path, blsPwds[i], ecdsaPwds[i]); err != nil {
			return err
		}
	}

	// Docker compose.
	operators := make([]map[string]interface{}, operatorCount)
	for i := 1; i <= int(operatorCount); i++ {
		machine := fmt.Sprintf("machine%d", i)

		operators[i-1] = map[string]interface{}{
			"name":               fmt.Sprintf("operator%d", i),
			"machine":            machine,
			"ipfs_address":       fmt.Sprintf("%s:5001", machine),
			"lambada_address":    fmt.Sprintf("%s:3033", machine),
			"config_path":        configPaths[i-1],
			"data_path":          dataPaths[i-1],
			"bls_key_password":   blsPwds[i-1],
			"ecdsa_key_password": ecdsaPwds[i-1],
		}
	}
	cofingTmpl, err := gonja.FromFile("./docker-compose/docker-compose.j2")
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

func readKeyDir(dirPath string) ([]string, []string, [][]byte, []string, error) {
	pwds, err := readFileLines(filepath.Join(dirPath, "password.txt"))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	privKeys, err := readFileLines(filepath.Join(dirPath, "private_key_hex.txt"))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	keys := make([][]byte, len(pwds))
	keyPaths := make([]string, len(keys))
	if err = filepath.Walk(
		filepath.Join(dirPath, "keys"),
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			key, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			keyIdxStr := strings.Split(info.Name(), ".")[0]
			keyIdx, err := strconv.ParseInt(keyIdxStr, 0, 32)
			if err != nil {
				return err
			}
			keys[keyIdx-1] = key
			keyPaths[keyIdx-1] = path

			return nil
		}); err != nil {
		return nil, nil, nil, nil, err
	}

	return pwds, privKeys, keys, keyPaths, nil
}

func readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return lines, nil
			} else {
				return nil, err
			}
		}
		lines = append(lines, string(line[:len(line)-1]))
	}
}

func runCommand(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func depositIntoStrategy(configPath, blsPwd, ecdsaPwd string) error {
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	nodeConfig.EthRpcUrl = "http://127.0.0.1:8545"
	nodeConfig.EthWsUrl = "ws://127.0.0.1:8545"
	// need to make sure we don't register the operator on startup
	// when using the cli commands to register the operator.
	nodeConfig.RegisterOperatorOnStartup = false

	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", blsPwd)
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", ecdsaPwd)

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	strategyAddr := common.HexToAddress("0xa85233C63b9Ee964Add6F2cffe00Fd84eb32338f")
	err = operator.DepositIntoStrategy(strategyAddr, big.NewInt(10))
	if err != nil {
		return err
	}

	return nil
}
