package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	blsPwdFlag = cli.StringFlag{
		Name:     "bls-password",
		Usage:    "BLS keystore password",
		Required: true,
	}
	ecdsaPwdFlag = cli.StringFlag{
		Name:     "ecdsa-password",
		Usage:    "ECDSA keystore password",
		Required: true,
	}
	configFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "path to configuration file",
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
			Name:   "register-operator-with-eigenlayer",
			Usage:  "registers operator with eigenlayer",
			Action: RegisterOperatorWithEigenlayer,
			Flags:  []cli.Flag{configFlag},
		},
		{
			Name:   "deposit-into-strategy",
			Usage:  "deposit tokens into a strategy",
			Action: DepositIntoStrategy,
			Flags: []cli.Flag{
				configFlag,
				strategyAddressFlag,
				strategyDepositAmountFlag,
			},
		},
		{
			Name:   "register-operator-with-avs",
			Usage:  "registers bls keys with pubkey-compendium, opts into slashing by avs service-manager, and registers operators with avs registry",
			Action: RegisterOperatorWithAvs,
			Flags:  []cli.Flag{configFlag},
		},
		{
			Name:   "setup-operator",
			Usage:  "full registration and strategy deposit",
			Action: SetupOperator,
			Flags: []cli.Flag{
				configFlag,
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
		{
			Name:   "print-operator-status",
			Usage:  "prints operator status as viewed from incredible squaring contracts",
			Action: PrintOperatorStatus,
			Flags:  []cli.Flag{configFlag},
		},
		{
			Name:   "generate-docker-compose",
			Usage:  "generate docker compose with specified number of operator nodes",
			Action: GenerateDockerCompose,
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:     "operators",
					Usage:    "number of operators nodes",
					Required: true,
				},
			},
		},
		{
			Name:   "generate-operator",
			Usage:  "generate operator keys and configuration file",
			Action: GenerateOperator,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "operator-dir",
					Usage:    "output directory with operator keys and configuration file",
					Required: true,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
