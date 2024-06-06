package actions

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
)

const ANVIL_DEV_ACCOUNT_PRIVATE_KEY = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const WETH_HOLESKY_ADDRESS = "0x94373a4919B3240D86eA41593D5eBa789FEF3848"

type strategyDepoist struct {
	address string
	amount  uint64
}

func SetupOperator(ctx *cli.Context) error {
	// Read configuration.
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	operatorConfig := config.OperatorConfig{}
	if err := sdkutils.ReadYamlConfig(configPath, &operatorConfig); err != nil {
		return fmt.Errorf("failed to read operator config - %s", err)
	}

	deploymentParamsPath := ctx.String("deployment-parameters")
	deploymentParams := config.AVSDeploymentParameters{}
	if err := sdkutils.ReadJsonConfig(deploymentParamsPath, &deploymentParams); err != nil {
		return fmt.Errorf("failed to read deployment parameters - %s", err)
	}

	blsPwd := ctx.String("bls-password")
	ecdsaPwd := ctx.String("ecdsa-password")
	operatorEcdsaPrivKey, err := sdkecdsa.ReadKey(operatorConfig.EcdsaPrivateKeyStorePath, ecdsaPwd)
	if err != nil {
		return fmt.Errorf("failed to read operator private key - %s", err)
	}

	// Init operator without starting it.
	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", blsPwd)
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", ecdsaPwd)
	operator, err := operator.NewOperatorFromConfig(operatorConfig)
	if err != nil {
		return fmt.Errorf("failed to init operator - %s", err)
	}

	if ctx.Bool("devnet") {
		// Fund operator with eth.
		sendFundsCmd := fmt.Sprintf(
			"cast send --private-key %s --value 20ether %s",
			ANVIL_DEV_ACCOUNT_PRIVATE_KEY, operatorConfig.OperatorAddress,
		)
		if _, _, err := runCommand(sendFundsCmd); err != nil {
			return fmt.Errorf("failed to send ether to operator on devnet - %s", err)
		}

		// Obtain wETH
		sendFundsCmd = fmt.Sprintf(
			"cast send --private-key %s --value 10ether %s 'deposit()'",
			ctx.String("private-key"), WETH_HOLESKY_ADDRESS,
		)
		if _, _, err := runCommand(sendFundsCmd); err != nil {
			return fmt.Errorf("failed to obtain WETH for operator on devnet - %s", err)
		}
	}

	// Register operator with Eigenlayer.
	if err := operator.RegisterOperatorWithEigenlayer(); err != nil {
		return fmt.Errorf("failed to regsiter operator with Eigenlayer")
	}

	// Deposit into strategies.
	deposits := []strategyDepoist{
		{
			address: deploymentParams.WETH,
			amount:  ctx.Uint64("deposit-amount-weth"),
		},
	}
	for _, deposit := range deposits {
		address := common.HexToAddress(deposit.address)
		amount := big.NewInt(int64(deposit.amount))
		if err = operator.DepositIntoStrategy(address, amount); err != nil {
			return fmt.Errorf("failed to desposit to strategy %s - %s", address, err)
		}
	}

	// Register operator with AVS.
	if err = operator.RegisterOperatorWithAvs(operatorEcdsaPrivKey); err != nil {
		return fmt.Errorf("failed to register operator with AVS - %s", err)
	}

	return nil
}
