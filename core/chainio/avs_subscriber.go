package chainio

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
)

type AvsSubscriberer interface {
	SubscribeToNewBatches(
		newBatchChan chan *tm.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered,
	) (event.Subscription, error)

	SubscribeToTaskResponses(
		taskResponseLogs chan *tm.ContractLambadaCoprocessorTaskManagerTaskResponded,
	) (event.Subscription, error)

	ParseTaskResponded(rawLog types.Log) (*tm.ContractLambadaCoprocessorTaskManagerTaskResponded, error)
}

type AvsSubscriber struct {
	Bindings *AvsManagersBindings

	log sdklogging.Logger
}

func NewAvsSubscriber(deployment AVSDeployment, ethClient eth.Client, log logging.Logger) (*AvsSubscriber, error) {
	bindings, err := NewAvsManagersBindings(deployment, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract bindings - %s", err)
	}

	s := &AvsSubscriber{
		Bindings: bindings,

		log: log,
	}

	return s, nil
}

func (s *AvsSubscriber) SubscribeToNewBatches(
	newBatchChan chan *tm.ContractLambadaCoprocessorTaskManagerTaskBatchRegistered,
) (event.Subscription, error) {
	sub, err := s.Bindings.TaskManager.WatchTaskBatchRegistered(
		&bind.WatchOpts{}, newBatchChan,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to NewTaskBatchRegistered events - %s", err)
	}

	s.log.Infof("Subscribed to new TaskManager tasks")

	return sub, nil
}

func (s *AvsSubscriber) SubscribeToTaskResponses(
	taskResponseChan chan *tm.ContractLambadaCoprocessorTaskManagerTaskResponded,
) (event.Subscription, error) {
	sub, err := s.Bindings.TaskManager.WatchTaskResponded(
		&bind.WatchOpts{}, taskResponseChan,
	)
	if err != nil {
		s.log.Error("failed to subscribe to TaskResponded events - %s", err)
	}

	s.log.Infof("Subscribed to TaskResponded events")

	return sub, nil
}

func (s *AvsSubscriber) ParseTaskResponded(
	rawLog types.Log,
) (*tm.ContractLambadaCoprocessorTaskManagerTaskResponded, error) {
	return s.Bindings.TaskManager.ContractLambadaCoprocessorTaskManagerFilterer.ParseTaskResponded(rawLog)
}
