package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string

	blsPwdFlag = cli.StringFlag{
		Name:     "bls-password",
		Required: true,
		Usage:    "BLS keystore password",
	}
	ecdsaPwdFlag = cli.StringFlag{
		Name:     "ecdsa-password",
		Required: true,
		Usage:    "ECDSA keystore password",
	}
	configFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "path to aggregator configuration file",
	}
)

func main() {
	app := cli.NewApp()
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Flags = []cli.Flag{configFlag, blsPwdFlag}
	app.Name = "lambada-coprocessor-operator"
	app.Usage = "Lambada Coprocessor Operator"

	app.Action = operatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func operatorMain(ctx *cli.Context) error {
	log.Println("Initializing Operator")

	var config operator.Config
	configPath := ctx.GlobalString(configFlag.Name)
	if err := sdkutils.ReadYamlConfig(configPath, &config); err != nil {
		return fmt.Errorf("failed to read configuration file - %s", err)
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
		return err
	}

	log.Println("initialized operator")

	log.Println("starting operator")

	err = operator.Start(context.Background())
	if err != nil {
		return err
	}

	log.Println("started operator")

	return nil
}
