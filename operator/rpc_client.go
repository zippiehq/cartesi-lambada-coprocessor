package operator

import (
	"fmt"
	"net/rpc"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core"
	"github.com/zippiehq/cartesi-lambada-coprocessor/metrics"
)

type AggregatorRpcClienter interface {
	GetBatchTasks(batchIdx core.TaskBatchIndex) ([]core.Task, error)
	SendSignedTaskResponseToAggregator(signedTaskResponse *aggregator.SignedTaskResponse)
}
type AggregatorRpcClient struct {
	rpcClient            *rpc.Client
	metrics              metrics.Metrics
	log                  logging.Logger
	aggregatorIpPortAddr string
}

func NewAggregatorRpcClient(aggregatorIpPortAddr string, logger logging.Logger, metrics metrics.Metrics) (*AggregatorRpcClient, error) {
	return &AggregatorRpcClient{
		// set to nil so that we can create an rpc client even if the aggregator is not running
		rpcClient:            nil,
		metrics:              metrics,
		log:                  logger,
		aggregatorIpPortAddr: aggregatorIpPortAddr,
	}, nil
}

func (c *AggregatorRpcClient) dialAggregatorRpcClient() error {
	client, err := rpc.DialHTTP("tcp", c.aggregatorIpPortAddr)
	if err != nil {
		return err
	}
	c.rpcClient = client
	return nil
}

func (c *AggregatorRpcClient) GetBatchTasks(batchIdx core.TaskBatchIndex) ([]core.Task, error) {
	if c.rpcClient == nil {
		c.log.Info("dialing aggregator rpc client")
		if err := c.dialAggregatorRpcClient(); err != nil {
			return nil, fmt.Errorf("failed to dial aggregator rpc client -%s", err)
		}
	}

	tasks := aggregator.BatchTasks{}
	for i := 0; i < 5; i++ {
		if err := c.rpcClient.Call("Aggregator.GetBatchTasks", batchIdx, &tasks); err != nil {
			c.log.Errorf("failed to fetch task batch %d from aggregator - %s, retrying in 2 seconds", batchIdx, err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	return tasks.Tasks, nil
}

// SendSignedTaskResponseToAggregator sends a signed task response to the aggregator.
// it is meant to be ran inside a go thread, so doesn't return anything.
// this is because sending the signed task response to the aggregator is time sensitive,
// so there is no point in retrying if it fails for a few times.
// Currently hardcoded to retry sending the signed task response 5 times, waiting 2 seconds in between each attempt.
func (c *AggregatorRpcClient) SendSignedTaskResponseToAggregator(signedTaskResponse *aggregator.SignedTaskResponse) {
	if c.rpcClient == nil {
		c.log.Info("rpc client is nil. Dialing aggregator rpc client")
		err := c.dialAggregatorRpcClient()
		if err != nil {
			c.log.Error("Could not dial aggregator rpc client. Not sending signed task response header to aggregator. Is aggregator running?", "err", err)
			return
		}
	}
	// we don't check this bool. It's just needed because rpc.Call requires rpc methods to have a return value
	var reply bool
	// We try to send the response 5 times to the aggregator, waiting 2 times in between each attempt.
	// This is mostly only necessary for local testing, since the aggregator sometimes is not ready to process task responses
	// before the operator gets the new task created log from anvil (because blocks are mined instantly)
	// the aggregator needs to read some onchain data related to quorums before it can accept operator signed task responses.
	c.log.Info("Sending signed task response header to aggregator", "signedTaskResponse", fmt.Sprintf("%#v", signedTaskResponse))
	for i := 0; i < 5; i++ {
		err := c.rpcClient.Call("Aggregator.ProcessSignedTaskResponse", signedTaskResponse, &reply)
		if err != nil {
			c.log.Info("Received error from aggregator", "err", err)
		} else {
			c.log.Info("Signed task response header accepted by aggregator.", "reply", reply)
			c.metrics.IncNumTasksAcceptedByAggregator()
			return
		}
		c.log.Infof("Retrying in 2 seconds")
		time.Sleep(2 * time.Second)
	}
	c.log.Errorf("Could not send signed task response to aggregator. Tried 5 times.")
}
