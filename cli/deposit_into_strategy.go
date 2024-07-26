package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
)

func DepositIntoStrategy(ctx *cli.Context) error {
	configPath := ctx.String(configFlag.Name)
	config := operator.Config{}
	if err := sdkutils.ReadYamlConfig(configPath, &config); err != nil {
		return fmt.Errorf("failed to read operator config - %s", err)
	}

	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize config to json - %s", err)
	}
	log.Println("Config:", string(configJson))

	blsPwd := ctx.String(blsPwdFlag.Name)
	ecdsaPwd := ctx.String(ecdsaPwdFlag.Name)
	operator, err := operator.NewOperator(blsPwd, ecdsaPwd, config)
	if err != nil {
		return fmt.Errorf("failed to create operator - %s", err)
	}

	strategyAddr := common.HexToAddress(strategyAddressFlag.Name)
	amount := ctx.Uint64(strategyDepositAmountFlag.Name)
	if err := operator.DepositIntoStrategy(strategyAddr, big.NewInt(int64(amount))); err != nil {
		return err
	}

	return nil
}
