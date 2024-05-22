package chainio

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	erc20mock "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/ERC20Mock"
	taskmanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
)

type AvsReaderer interface {
	sdkavsregistry.AvsRegistryReader

	CheckSignatures(
		ctx context.Context,
		msgHash [32]byte,
		quorumNumbers []byte,
		referenceBlockNumber uint32,
		nonSignerStakesAndSignature taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (taskmanager.IBLSSignatureCheckerQuorumStakeTotals, error)

	GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error)
}

type AvsReader struct {
	sdkavsregistry.AvsRegistryReader

	AvsManagersBindings *AvsManagersBindings
	log                 logging.Logger
}

func BuildAvsReaderFromConfig(c *config.AggregatorConfig) (*AvsReader, error) {
	return BuildAvsReader(c.LambadaCoprocessorRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsReader(registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsReader, error) {
	avsManagersBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}
	avsRegistryReader, err := sdkavsregistry.BuildAvsRegistryChainReader(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}

	reader := AvsReader{
		AvsRegistryReader:   avsRegistryReader,
		AvsManagersBindings: avsManagersBindings,
		log:                 logger,
	}

	return &reader, nil
}

func (r *AvsReader) CheckSignatures(
	ctx context.Context, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (taskmanager.IBLSSignatureCheckerQuorumStakeTotals, error) {
	stakeTotalsPerQuorum, _, err := r.AvsManagersBindings.TaskManager.CheckSignatures(
		&bind.CallOpts{}, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature,
	)
	if err != nil {
		return taskmanager.IBLSSignatureCheckerQuorumStakeTotals{}, err
	}
	return stakeTotalsPerQuorum, nil
}

func (r *AvsReader) GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error) {
	erc20Mock, err := r.AvsManagersBindings.GetErc20Mock(tokenAddr)
	if err != nil {
		r.log.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return erc20Mock, nil
}
