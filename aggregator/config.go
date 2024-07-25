package aggregator

import sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

type Config struct {
	Environment sdklogging.LogLevel `yaml:"environment"`

	AVSDeploymentOutputPath string `yaml:"avs_deployment_output_path"`

	EthHttpRpcUrl string `yaml:"eth_rpc_url"`
	EthWsRpcUrl   string `yaml:"eth_ws_url"`

	APIServerAddress string `yaml:"api_server_address"`
}
