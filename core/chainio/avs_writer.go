package chainio

import (
	"context"
	"fmt"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	taskmanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	RegisterNewTaskBatch(
		ctx context.Context,
		batchRoot [32]byte,
		quorumThresholdPercentage uint32,
		quorumNumbers []byte,
	) (taskmanager.ILambadaCoprocessorTaskManagerTaskBatch, error)

	RespondTask(
		ctx context.Context,
		taskBatch taskmanager.ILambadaCoprocessorTaskManagerTaskBatch,
		task taskmanager.ILambadaCoprocessorTaskManagerTask,
		taskProof [][32]byte,
		taskResponse taskmanager.ILambadaCoprocessorTaskManagerTaskResponse,
		nonSignerStakesAndSignature taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter

	AvsContractBindings *AvsManagersBindings
	log                 logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.EthClient
}

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.LambadaCoprocessorRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(
	txMgr txmgr.TxManager,
	registryCoordinatorAddr,
	operatorStateRetrieverAddr gethcommon.Address,
	ethHttpClient eth.EthClient,
	log logging.Logger,
) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, log)
	if err != nil {
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, log, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}

	writer := AvsWriter{
		AvsRegistryWriter: avsRegistryWriter,

		AvsContractBindings: avsServiceBindings,
		log:                 log,
		TxMgr:               txMgr,
		client:              ethHttpClient,
	}

	return &writer, nil
}

func (w *AvsWriter) RegisterNewTaskBatch(
	ctx context.Context,
	batchRoot [32]byte,
	quorumThresholdPercentage uint32,
	quorumNumbers []byte,
) (taskmanager.ILambadaCoprocessorTaskManagerTaskBatch, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to create opts for RegisterNewTaskBatch tx - %s", err)
	}

	tx, err := w.AvsContractBindings.TaskManager.RegisterNewTaskBatch(
		txOpts, batchRoot, quorumThresholdPercentage, quorumNumbers,
	)
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to create RegisterNewTaskBatch tx - %s", err)
	}

	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to send RegisterNewTaskBatch tx - %s", err)
	}

	event, err := w.AvsContractBindings.TaskManager.ContractLambadaCoprocessorTaskManagerFilterer.
		ParseTaskBatchRegistered(*receipt.Logs[0])
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to parse TaskBatchRegistered event - %S", err)
	}

	return event.Batch, nil
}

func (w *AvsWriter) RespondTask(
	ctx context.Context,
	taskBatch taskmanager.ILambadaCoprocessorTaskManagerTaskBatch,
	task taskmanager.ILambadaCoprocessorTaskManagerTask,
	taskProof [][32]byte,
	taskResponse taskmanager.ILambadaCoprocessorTaskManagerTaskResponse,
	nonSignerStakesAndSignature taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, fmt.Errorf("failed create opts for RespondTask tx - %s", err)
	}

	tx, err := w.AvsContractBindings.TaskManager.RespondTask(
		txOpts, taskBatch, task, taskProof, taskResponse, nonSignerStakesAndSignature,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RespondTask tx - %s", err)
	}

	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to send CreateNewTask tx - %s", err)
	}

	return receipt, nil
}
