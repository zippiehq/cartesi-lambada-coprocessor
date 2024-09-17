package aggregator

import sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

type Config struct {
	Environment sdklogging.LogLevel `yaml:"environment"`

	DatabaseAddress string `yaml:"database_address"`
	DatabaseName    string `yaml:"database_name"`
	DatabaseUser    string `yaml:"database_user"`

	AVSDeploymentOutputPath string `yaml:"avs_deployment_output_path"`

	EthHttpRpcUrl string `yaml:"eth_rpc_url"`
	EthWsRpcUrl   string `yaml:"eth_ws_url"`

	APIServerAddress string `yaml:"api_server_address"`

	QueryOperators          bool   `yaml:"query_operators"`
	QueryOperatorStartBlock uint64 `yaml:"query_operator_start_block"`
}
