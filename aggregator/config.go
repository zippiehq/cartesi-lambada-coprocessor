package aggregator

import sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

type Config struct {
	Environment sdklogging.LogLevel `yaml:"environment"`

	AVSDeploymentPath string `yaml:"avs_deployment_path"`

	EthHttpRpcUrl string `yaml:"eth_rpc_url"`
	EthWsRpcUrl   string `yaml:"eth_ws_url"`

	APIServerAddress string `yaml:"api_server_address"`
}
