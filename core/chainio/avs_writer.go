package chainio

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"

	taskmanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	RegisterNewTaskBatch(
		ctx context.Context,
		batchRoot [32]byte,
		quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
		quorumNumbers sdktypes.QuorumNums,
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

	Bindings *AvsManagersBindings

	log logging.Logger

	txmgr txmgr.TxManager
}

func NewAvsWriter(privKey *ecdsa.PrivateKey, deployment AVSDeployment, ethClient eth.Client, log logging.Logger) (*AvsWriter, error) {
	addr, err := sdkutils.EcdsaPrivateKeyToAddress(privKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get address - %s", err)
	}
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID - %s", err)
	}

	signer, _, err := signerv2.SignerFromConfig(
		signerv2.Config{PrivateKey: privKey},
		chainID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer - %s", err)
	}
	wallet, err := wallet.NewPrivateKeyWallet(ethClient, signer, addr, log)
	if err != nil {
		return nil, fmt.Errorf("failed to create signer wallet - %s", err)
	}
	txMgr := txmgr.NewSimpleTxManager(wallet, ethClient, log, addr)

	bindings, err := NewAvsManagersBindings(deployment, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract bindings - %s", err)
	}
	writer, err := avsregistry.BuildAvsRegistryChainWriter(
		common.HexToAddress(deployment.Addresses.RegistryCoordinator),
		common.HexToAddress(deployment.Addresses.OperatorStateRetriever),
		log,
		ethClient,
		txMgr,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create AVS registry chain writer -%s", err)
	}

	w := &AvsWriter{
		AvsRegistryWriter: writer,

		Bindings: bindings,

		log: log,

		txmgr: txMgr,
	}

	return w, nil
}

func (w *AvsWriter) RegisterNewTaskBatch(
	ctx context.Context,
	batchRoot [32]byte,
	quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
	quorumNumbers sdktypes.QuorumNums,
) (taskmanager.ILambadaCoprocessorTaskManagerTaskBatch, error) {
	txOpts, err := w.txmgr.GetNoSendTxOpts()
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to create opts for RegisterNewTaskBatch tx - %s", err)
	}

	tx, err := w.Bindings.TaskManager.RegisterNewTaskBatch(
		txOpts, batchRoot, uint32(quorumThresholdPercentage), quorumNumbers.UnderlyingType(),
	)
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to create RegisterNewTaskBatch tx - %s", err)
	}

	receipt, err := w.txmgr.Send(ctx, tx)
	if err != nil {
		return taskmanager.ILambadaCoprocessorTaskManagerTaskBatch{},
			fmt.Errorf("failed to send RegisterNewTaskBatch tx - %s", err)
	}

	event, err := w.Bindings.TaskManager.ContractLambadaCoprocessorTaskManagerFilterer.
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
	txOpts, err := w.txmgr.GetNoSendTxOpts()
	if err != nil {
		return nil, fmt.Errorf("failed create opts for RespondTask tx - %s", err)
	}

	tx, err := w.Bindings.TaskManager.RespondTask(
		txOpts, taskBatch, task, taskProof, taskResponse, nonSignerStakesAndSignature,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create RespondTask tx - %s", err)
	}

	receipt, err := w.txmgr.Send(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to send CreateNewTask tx - %s", err)
	}

	return receipt, nil
}
