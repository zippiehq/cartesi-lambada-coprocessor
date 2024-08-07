package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
)

type strategyDepoist struct {
	address string
	amount  uint64
}

func SetupOperator(ctx *cli.Context) error {
	configPath := ctx.String(configFlag.Name)
	operatorConfig := operator.Config{}
	if err := sdkutils.ReadYamlConfig(configPath, &operatorConfig); err != nil {
		return fmt.Errorf("failed to read operator config - %s", err)
	}

	blsPwd := ctx.String(blsPwdFlag.Name)
	ecdsaPwd := ctx.String(ecdsaPwdFlag.Name)
	operator, err := operator.NewOperator(blsPwd, ecdsaPwd, operatorConfig)
	if err != nil {
		return fmt.Errorf("failed to create operator - %s", err)
	}

	// Register operator with Eigenlayer.
	if err := operator.RegisterOperatorWithEigenlayer(); err != nil {
		return fmt.Errorf("failed to regsiter operator with Eigenlayer - %s", err)
	}

	// Deposit into strategies.
	deposits := []strategyDepoist{
		{
			address: ctx.String(strategyAddressFlag.Name),
			amount:  ctx.Uint64(strategyDepositAmountFlag.Name),
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
	if err = operator.RegisterOperatorWithAvs(); err != nil {
		return fmt.Errorf("failed to register operator with AVS - %s", err)
	}

	return nil
}
