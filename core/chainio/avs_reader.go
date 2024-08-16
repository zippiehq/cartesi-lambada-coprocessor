package chainio

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
)

type AvsReaderer interface {
	sdkavsregistry.AvsRegistryReader

	CheckSignatures(
		ctx context.Context,
		msgHash [32]byte,
		quorumNumbers []byte,
		referenceBlockNumber uint32,
		nonSignerStakesAndSignature tm.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (tm.IBLSSignatureCheckerQuorumStakeTotals, error)
}

type AvsReader struct {
	sdkavsregistry.AvsRegistryReader

	Bindings *AvsManagersBindings

	log logging.Logger
}

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
	referenceBlockNumber uint32, nonSignerStakesAndSignature tm.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (tm.IBLSSignatureCheckerQuorumStakeTotals, error) {
	stakeTotalsPerQuorum, _, err := r.Bindings.TaskManager.CheckSignatures(
		&bind.CallOpts{}, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature,
	)
	if err != nil {
		return tm.IBLSSignatureCheckerQuorumStakeTotals{}, err
	}
	return stakeTotalsPerQuorum, nil
}
