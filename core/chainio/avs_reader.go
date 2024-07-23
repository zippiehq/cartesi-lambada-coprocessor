package chainio

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	erc20mock "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/ERC20Mock"
	taskmanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
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

	Bindings *AvsManagersBindings

	log logging.Logger
}

//!!!
/*
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
*/

func NewAvsReader(deployment AVSDeployment, ethClient eth.Client, log logging.Logger) (*AvsReader, error) {
	bindings, err := NewAvsManagersBindings(deployment, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract bindings - %s", err)
	}
	reader, err := sdkavsregistry.BuildAvsRegistryChainReader(
		common.HexToAddress(deployment.Addresses.RegistryCoordinator),
		common.HexToAddress(deployment.Addresses.OperatorStateRetriever),
		ethClient,
		log,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create AVS registy chain reader - %s", err)
	}

	r := &AvsReader{
		AvsRegistryReader: reader,

		Bindings: bindings,

		log: log,
	}

	return r, nil
}

func (r *AvsReader) CheckSignatures(ctx context.Context, msgHash [32]byte, quorumNumbers []byte,
	referenceBlockNumber uint32, nonSignerStakesAndSignature taskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (taskmanager.IBLSSignatureCheckerQuorumStakeTotals, error) {
	stakeTotalsPerQuorum, _, err := r.Bindings.TaskManager.CheckSignatures(
		&bind.CallOpts{}, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature,
	)
	if err != nil {
		return taskmanager.IBLSSignatureCheckerQuorumStakeTotals{}, err
	}
	return stakeTotalsPerQuorum, nil
}
