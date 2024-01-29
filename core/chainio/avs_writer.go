package chainio

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signer"

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
	AvsContractBindings *AvsServiceBindings
	logger              logging.Logger
	Signer              signer.Signer
	client              eth.EthClient
}

var _ AvsWriterer = (*AvsWriter)(nil)

func NewAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return NewAvsWriter(c.Signer, c.LambadaCoprocessorServiceManagerAddr, c.BlsOperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func NewAvsWriter(signer signer.Signer, serviceManagerAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.EthClient, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsServiceBindings(serviceManagerAddr, blsOperatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	blsRegistryCoordinatorAddr, err := avsServiceBindings.ServiceManager.RegistryCoordinator(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := avsServiceBindings.ServiceManager.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	blsPubkeyRegistryAddr, err := avsServiceBindings.ServiceManager.BlsPubkeyRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	avsRegistryContractClient, err := sdkclients.NewAvsRegistryContractsChainClient(
		blsRegistryCoordinatorAddr, blsOperatorStateRetrieverAddr, stakeRegistryAddr, blsPubkeyRegistryAddr, ethHttpClient, logger,
	)
	if err != nil {
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.NewAvsRegistryWriter(avsRegistryContractClient, logger, signer, ethHttpClient)
	if err != nil {
		return nil, err
	}

	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		Signer:              signer,
		client:              ethHttpClient,
	}, nil
}

func (w *AvsWriter) RegisterNewTaskBatch(
	ctx context.Context,
	batchRoot [32]byte,
	quorumThresholdPercentage uint32,
	quorumNumbers []byte,
) (taskmanager.ILambadaCoprocessorTaskManagerTaskBatch, error) {
	txOpts := w.Signer.GetTxOpts()
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
	txOpts := w.Signer.GetTxOpts()
	tx, err := w.AvsContractBindings.TaskManager.RespondTask(
		txOpts, taskBatch, task, taskProof, taskResponse, nonSignerStakesAndSignature,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send respondTask trasaction - %s", err)
	}

	receipt := w.client.WaitForTransactionReceipt(ctx, tx.Hash())

	return receipt, nil
}
