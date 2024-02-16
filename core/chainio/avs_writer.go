package chainio

import (
	"context"
	"fmt"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
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
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.EthClient
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.LambadaCoprocessorRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.EthClient, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}
func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

func (w *AvsWriter) RegisterNewTaskBatch(
	ctx context.Context,
	batchRoot [32]byte,
	quorumThresholdPercentage uint32,
	quorumNumbers []byte,
) (taskmanager.ILambadaCoprocessorTaskManagerTaskBatch, error) {
	txOpts := w.TxMgr.GetNoSendTxOpts()
	tx, err := w.AvsContractBindings.TaskManager.RegisterNewTaskBatch(
		txOpts, batchRoot, quorumThresholdPercentage, quorumNumbers,
	)
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to send registerNewTaskBatch transaction - %s", err)
	}

	receipt := w.client.WaitForTransactionReceipt(ctx, tx.Hash())
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
	tx, err := w.AvsContractBindings.TaskManager.RespondTask(
		txOpts, taskBatch, task, taskProof, taskResponse, nonSignerStakesAndSignature,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send respondTask trasaction - %s", err)
	}

	receipt := w.client.WaitForTransactionReceipt(ctx, tx.Hash())

	return receipt, nil
}
