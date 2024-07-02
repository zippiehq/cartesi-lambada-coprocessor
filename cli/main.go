package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	configFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "configuration file path",
	}
	strategyAddressFlag = cli.StringFlag{
		Name:     "strategy-address",
		Usage:    "address of strategy contract to deposit into",
		Required: true,
	}
	strategyDepositAmountFlag = cli.UintFlag{
		Name:     "strategy-deposit-amount",
		Usage:    "amount of tokens to deposit into strategy",
		Required: true,
	}
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "register-operator-with-eigenlayer",
			Aliases: []string{"rel"},
			Usage:   "registers operator with eigenlayer (this should be called via eigenlayer cli, not plugin, but keeping here for convenience for now)",
			Action:  RegisterOperatorWithEigenlayer,
			Flags:   []cli.Flag{configFlag},
		},
		{
			Name:    "deposit-into-strategy",
			Aliases: []string{"d"},
			Usage:   "deposit tokens into a strategy",
			Action:  DepositIntoStrategy,
			Flags: []cli.Flag{
				configFlag,
				strategyAddressFlag,
				strategyDepositAmountFlag,
			},
		},
		{
			Name:    "register-operator-with-avs",
			Aliases: []string{"r"},
			Usage:   "registers bls keys with pubkey-compendium, opts into slashing by avs service-manager, and registers operators with avs registry",
			Action:  RegisterOperatorWithAvs,
			Flags:   []cli.Flag{configFlag},
		},
		{
			Name:    "print-operator-status",
			Aliases: []string{"s"},
			Usage:   "prints operator status as viewed from incredible squaring contracts",
			Action:  PrintOperatorStatus,
			Flags:   []cli.Flag{configFlag},
		},
		{
			Name:    "generate-docker-compose",
			Aliases: []string{"g"},
			Usage:   "generate docker compose with specified number of operator nodes",
			Action:  GenerateDockerCompose,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "network",
					Usage:    "devent or holesky",
					Required: true,
				},
				cli.UintFlag{
					Name:     "operators",
					Usage:    "number of operators nodes",
					Required: true,
				},
			},
		},
		{
			Name:    "setup-operator",
			Aliases: []string{"s"},
			Usage:   "full registration and strategy deposit",
			Action:  SetupOperator,
			Flags: []cli.Flag{
				configFlag,
				cli.StringFlag{
					Name:     "deployment-parameters",
					Usage:    "deployment parameters file path",
					Required: true,
				},
				cli.StringFlag{
					Name:     "bls-password",
					Usage:    "operator bls password",
					Required: true,
				},
				cli.StringFlag{
					Name:     "ecdsa-password",
					Usage:    "operator ecsda password",
					Required: true,
				},
				strategyAddressFlag,
				strategyDepositAmountFlag,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
