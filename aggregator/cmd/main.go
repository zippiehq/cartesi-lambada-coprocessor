package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string

	configFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "path to aggregator configuration file",
	}
	privKeyFlag = cli.StringFlag{
		Name:     "private-key",
		Required: true,
		Usage:    "Ethereum private key",
	}
	dbPwdFlag = cli.StringFlag{
		Name:     "database-password",
		Required: true,
		Usage:    "MySQL database password",
	}
)

func main() {
	app := cli.NewApp()
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Flags = []cli.Flag{configFlag, privKeyFlag, dbPwdFlag}
	app.Name = "lambada-coprocessor-aggregator"
	app.Usage = "Lambada Coprocessor Aggregator"

	app.Action = aggregatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func aggregatorMain(ctx *cli.Context) error {
	log.Println("Initializing Aggregator")

	var config aggregator.Config
	configPath := ctx.String(configFlag.Name)
	if err := sdkutils.ReadYamlConfig(configPath, &config); err != nil {
		return fmt.Errorf("failed to read configuration file - %s", err)
	}

	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize config to json - %s", err)
	}
	fmt.Println("Config:", string(configJson))

	privKey := ctx.String(privKeyFlag.Name)
	log, err := sdklogging.NewZapLogger(config.Environment)
	if err != nil {
		return fmt.Errorf("failed to create Zap logger - %s", err)
	}

	dbPwd := ctx.String(dbPwdFlag.Name)

	agg, err := aggregator.NewAggregator(privKey, dbPwd, config, log)
	if err != nil {
		return fmt.Errorf("failed to create Aggregator instance - %s", err)
	}

	err = agg.Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}
