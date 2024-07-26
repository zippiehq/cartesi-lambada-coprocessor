package aggregator

import (
	"context"
	"errors"
	"net/http"
	"net/rpc"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator/types"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer(ctx context.Context) error {

	err := rpc.Register(agg)
	if err != nil {
		agg.log.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.apiServerAddr, nil)
	if err != nil {
		agg.log.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type BatchTasks struct {
	Tasks []types.Task
}

func (agg *Aggregator) GetBatchTasks(batchIdx types.TaskBatchIndex, tasks *BatchTasks) error {
	batchTasks, err := agg.getBatchTasks(batchIdx)
	if err != nil {
		return err
	}

	*tasks = BatchTasks{
		Tasks: batchTasks,
	}

	return nil
}

type SignedTaskResponse struct {
	tm.ILambadaCoprocessorTaskManagerTaskResponse
	TaskIndex    sdktypes.TaskIndex
	BlsSignature bls.Signature
	OperatorId   sdktypes.OperatorId
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedTaskResponse(resp *SignedTaskResponse, reply *bool) error {
	agg.log.Infof("Received signed task response: %#v", resp)
	return agg.processTaskResponse(
		resp.TaskIndex, resp.OperatorId, resp.ILambadaCoprocessorTaskManagerTaskResponse, resp.BlsSignature,
	)
}
