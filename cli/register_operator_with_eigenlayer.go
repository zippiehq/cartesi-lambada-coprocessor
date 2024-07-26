package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/urfave/cli"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
)

func RegisterOperatorWithEigenlayer(ctx *cli.Context) error {
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

	if err = operator.RegisterOperatorWithEigenlayer(); err != nil {
		return err
	}

	return nil
}
