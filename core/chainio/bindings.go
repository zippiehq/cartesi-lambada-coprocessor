package chainio

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"

	servicemanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorServiceManager"
	taskmanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
)

type AvsManagersBindings struct {
	TaskManager    *taskmanager.ContractLambadaCoprocessorTaskManager
	ServiceManager *servicemanager.ContractLambadaCoprocessorServiceManager
}

func NewAvsManagersBindings(deployment AVSDeployment, ethclient eth.Client) (*AvsManagersBindings, error) {
	contractRegistryCoordinator, err := regcoord.NewContractRegistryCoordinator(
		common.HexToAddress(deployment.Addresses.RegistryCoordinator),
		ethclient,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RegistryCoordinator binding - %s", err)
	}

	serviceManagerAddr, err := contractRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get ServiceManager address - %s", err)
	}
	contractServiceManager, err := servicemanager.NewContractLambadaCoprocessorServiceManager(serviceManagerAddr, ethclient)
	if err != nil {
		return nil, fmt.Errorf("failed to create LambadaCoprocessorServiceManager binding - %s", err)
	}

	taskManagerAddr, err := contractServiceManager.LambadaCoprocessorTaskManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get LambadaCorpocessorTaskManager address - %s", err)
	}
	contractTaskManager, err := taskmanager.NewContractLambadaCoprocessorTaskManager(taskManagerAddr, ethclient)
	if err != nil {
		return nil, fmt.Errorf("failed to create LambadaCoprocessorTaskManager binding - %s", err)
	}

	b := &AvsManagersBindings{
		ServiceManager: contractServiceManager,
		TaskManager:    contractTaskManager,
	}

	return b, nil
}
