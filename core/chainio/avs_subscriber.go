package chainio

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	taskmanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
)

type AvsSubscriberer interface {
	SubscribeToNewBatches(
		newBatchChan chan *taskmanager.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered,
	) (event.Subscription, error)

	SubscribeToTaskResponses(
		taskResponseLogs chan *taskmanager.ContractLambadaCoprocessorTaskManagerTaskResponded,
	) (event.Subscription, error)

	ParseTaskResponded(rawLog types.Log) (*taskmanager.ContractLambadaCoprocessorTaskManagerTaskResponded, error)
}

type AvsSubscriber struct {
	AvsContractBindings *AvsManagersBindings
	logger              sdklogging.Logger
}

func BuildAvsSubscriberFromConfig(config *config.Config) (*AvsSubscriber, error) {
	return BuildAvsSubscriber(
		config.LambadaCoprocessorRegistryCoordinatorAddr,
		config.OperatorStateRetrieverAddr,
		config.EthWsClient,
		config.Logger,
	)
}

func BuildAvsSubscriber(registryCoordinatorAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethclient eth.EthClient, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, blsOperatorStateRetrieverAddr, ethclient, logger)
	if err != nil {
		logger.Errorf("Failed to create contract bindings", "err", err)
		return nil, err
	}
	return NewAvsSubscriber(avsContractBindings, logger), nil
}

func NewAvsSubscriber(avsContractBindings *AvsManagersBindings, logger sdklogging.Logger) *AvsSubscriber {
	return &AvsSubscriber{
		AvsContractBindings: avsContractBindings,
		logger:              logger,
	}
}

func (s *AvsSubscriber) SubscribeToNewBatches(
	newBatchChan chan *taskmanager.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered,
) (event.Subscription, error) {
	sub, err := s.AvsContractBindings.TaskManager.WatchTaskBatchRegistered(
		&bind.WatchOpts{}, newBatchChan,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to NewTaskBatchRegistered events - %s", err)
	}

	s.logger.Infof("Subscribed to new TaskManager tasks")

	return sub, nil
}

func (s *AvsSubscriber) SubscribeToTaskResponses(
	taskResponseChan chan *taskmanager.ContractLambadaCoprocessorTaskManagerTaskResponded,
) (event.Subscription, error) {
	sub, err := s.AvsContractBindings.TaskManager.WatchTaskResponded(
		&bind.WatchOpts{}, taskResponseChan,
	)
	if err != nil {
		s.logger.Error("failed to subscribe to TaskResponded events - %s", err)
	}

	s.logger.Infof("Subscribed to TaskResponded events")

	return sub, nil
}

func (s *AvsSubscriber) ParseTaskResponded(
	rawLog types.Log,
) (*taskmanager.ContractLambadaCoprocessorTaskManagerTaskResponded, error) {
	return s.AvsContractBindings.TaskManager.ContractLambadaCoprocessorTaskManagerFilterer.ParseTaskResponded(rawLog)
}
