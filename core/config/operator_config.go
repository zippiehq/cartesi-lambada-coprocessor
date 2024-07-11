package config

type OperatorConfig struct {
	// used to set the logger level (true = info, false = debug)
	Production bool `yaml:"production"`

	AVSDeploymentPath string `yaml:"avs_deployment_path"`

	BLSPrivateKeyStorePath   string `yaml:"bls_private_key_store_path"`
	ECDSAPrivateKeyStorePath string `yaml:"ecdsa_private_key_store_path"`

	EthRpcUrl string `yaml:"eth_rpc_url"`
	EthWsUrl  string `yaml:"eth_ws_url"`

	AggregatorServerIpPortAddress string `yaml:"aggregator_server_ip_port_address"`
	IPFSIpPortAddress             string `yaml:"ipfs_ip_port_address"`
	LambadaIpPortAddress          string `yaml:"lambada_ip_port_address"`

	EnableMetrics             bool   `yaml:"enable_metrics"`
	EigenMetricsIpPortAddress string `yaml:"eigen_metrics_ip_port_address"`

	EnableNodeApi        bool   `yaml:"enable_node_api"`
	NodeApiIpPortAddress string `yaml:"node_api_ip_port_address"`
}
