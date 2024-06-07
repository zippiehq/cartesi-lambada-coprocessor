package config

type AVSDeploymentParameters struct {
	StrategyManager         string `json:"strategyManager"`
	AVSDirectory            string `json:"avsDirectory"`
	DelegationManager       string `json:"delegationManager"`
	TaskResponseWindowBlock uint64 `json:"taskResponseWindowBlock"`
	TaskGenerator           string `json:"taskGenerator"`
	Aggregator              string `json:"aggregator"`
	Owner                   string `json:"owner"`
	Churner                 string `json:"churner"`
	Ejector                 string `json:"ejector"`
	Confirmer               string `json:"confirmer"`
	Whitelister             string `json:"whitelister"`
	WETH                    string `json:"wETH"`
	WETHMultiplier          uint64 `json:"wETH_Multiplier"`
}
