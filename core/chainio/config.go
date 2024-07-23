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
