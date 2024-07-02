package main

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
)

func DepositIntoStrategy(ctx *cli.Context) error {

	configPath := ctx.String(configFlag.Name)
	nodeConfig := config.OperatorConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	// need to make sure we don't register the operator on startup
	// when using the cli commands to register the operator.
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	strategyAddr := common.HexToAddress(strategyAddressFlag.Name)
	amount := ctx.Uint64(strategyDepositAmountFlag.Name)

	err = operator.DepositIntoStrategy(strategyAddr, big.NewInt(int64(amount)))
	if err != nil {
		return err
	}

	return nil
}
