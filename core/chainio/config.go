package chainio

type AVSDeployment struct {
	Addresses AVSDeploymentAddresses `json:"addresses"`
}

type AVSDeploymentAddresses struct {
	RegistryCoordinator    string `json:"registryCoordinator"`
	OperatorStateRetriever string `json:"operatorStateRetriever"`
	ERC20Mock              string `json:"erc20Mock"`
	ERC20MockStrategy      string `json:"erc20MockStrategy"`
}

type AVSDeploymentParameters struct {
	StrategyManager         string `json:"strategyManager"`
	AVSDirectory            string `json:"avsDirectory"`
	DelegationManager       string `json:"delegationManager"`
	ProxyAdmin              string `json:"proxyAdmin"`
	PauserRegistry          string `json:"pauserRegistry"`
	BaseStrategy            string `json:"baseStrategyImplementation"`
	WETH                    string `json:"wETH"`
	WETHMultiplier          uint64 `json:"wETH_Multiplier"`
	RETH                    string `json:"rETH"`
	RETHMultiplier          uint64 `json:"rETH_Multiplier"`
	TaskResponseWindowBlock uint64 `json:"taskResponseWindowBlock"`
	TaskGenerator           string `json:"taskGenerator"`
	Aggregator              string `json:"aggregator"`
	Owner                   string `json:"owner"`
	Churner                 string `json:"churner"`
	Ejector                 string `json:"ejector"`
	Confirmer               string `json:"confirmer"`
	Whitelister             string `json:"whitelister"`
}
