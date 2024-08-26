package chainio

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"

	sm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/CoprocessorServiceManager"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/CoprocessorTaskManager"
)

type AvsManagersBindings struct {
	TaskManager    *tm.ContractCoprocessorTaskManager
	ServiceManager *sm.ContractCoprocessorServiceManager
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
	contractServiceManager, err := sm.NewContractCoprocessorServiceManager(serviceManagerAddr, ethclient)
	if err != nil {
		return nil, fmt.Errorf("failed to create CoprocessorServiceManager binding - %s", err)
	}

	taskManagerAddr, err := contractServiceManager.CoprocessorTaskManager(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get LambadaCorpocessorTaskManager address - %s", err)
	}
	contractTaskManager, err := tm.NewContractCoprocessorTaskManager(taskManagerAddr, ethclient)
	if err != nil {
		return nil, fmt.Errorf("failed to create CoprocessorTaskManager binding - %s", err)
	}

	b := &AvsManagersBindings{
		ServiceManager: contractServiceManager,
		TaskManager:    contractTaskManager,
	}

	return b, nil
}
