package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator/types"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
)

type task struct {
	ProgramID string `json:"programId"`
	Input     string `json:"input"`
}

func TestIntegration(t *testing.T) {
	config, err := config.NewConfig(
		"../../config-files/aggregator.yaml",
		"../../contracts/script/output/31337/credible_squaring_avs_deployment_output.json",
		"../../contracts/script/output/31337/shared_avs_contracts_deployment_output.json",
		"0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
	)
	if err != nil {
		t.Fatalf("failed to create avs config - %s", err)
	}

	avsReader, err := chainio.NewAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("failed to create avs reader - %s", err)
	}
	avsSubscriber, err := chainio.NewAvsSubscriberFromConfig(config)
	if err != nil {
		t.Fatalf("failed to create avs subscriber - %s", err)
	}

	batches := [][]task{
		// batch 0
		{
			{
				ProgramID: "program1",
				Input:     "input1",
			},
			{
				ProgramID: "program1",
				Input:     "intput1",
			},
		},
	}

	for _, b := range batches {
		checkTaskBatch(t, b, avsReader, avsSubscriber)
	}
}

func checkTaskBatch(t *testing.T, tasks []task, avsReader *chainio.AvsReader, avsSubscriber *chainio.AvsSubscriber) {
	batchIdx, taskIdx, err := submitTasks(tasks)
	if err != nil {
		t.Fatalf("failed to submit tasks - %s", err)
	}

	// Wait for batch and comapre batch hash. Compare batch index.

	// Wait for task reponse and compare task response hash.
}

func submitTasks(tasks []task) (types.TaskBatchIndex, sdktypes.TaskIndex, error) {
	body, err := json.Marshal(tasks)
	if err != nil {
		return 0, 0, err
	}

	resp, err := http.Post("http://localhost:8080/createTask", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	respJSON := struct {
		BatchIndex types.TaskBatchIndex `json:"batchIndex"`
		TaskIndex  sdktypes.TaskIndex   `json:"taskIndex"`
	}{}
	if err = json.Unmarshal(respData, &respJSON); err != nil {
		return 0, 0, err
	}

	return respJSON.BatchIndex, respJSON.TaskIndex, nil
}

// TODO: update tests
/*
import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	clientconstructor "github.com/Layr-Labs/eigensdk-go/chainio/constructor"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signer"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/zippiehq/cartesi-lambada-coprocessor/aggregator"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/chainio"
	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
	"github.com/zippiehq/cartesi-lambada-coprocessor/operator"
	"github.com/zippiehq/cartesi-lambada-coprocessor/types"
)

type IntegrationClients struct {
	Sdkclients clientconstructor.Clients
}

func TestIntegration(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

	// Start the anvil chain
	anvilC := startAnvilTestContainer()
	// Not sure why but deferring anvilC.Terminate() causes a panic when the test finishes...
	// so letting it terminate silently for now
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	// Prepare the config file for aggregator
	var aggConfigRaw config.ConfigRaw
	aggConfigFilePath := "../../config-files/aggregator.yaml"
	sdkutils.ReadYamlConfig(aggConfigFilePath, &aggConfigRaw)
	aggConfigRaw.EthRpcUrl = "http://" + anvilEndpoint
	aggConfigRaw.EthWsUrl = "ws://" + anvilEndpoint

	var credibleSquaringDeploymentRaw config.CredibleSquaringDeploymentRaw
	credibleSquaringDeploymentFilePath := "../../contracts/script/output/31337/credible_squaring_avs_deployment_output.json"
	sdkutils.ReadJsonConfig(credibleSquaringDeploymentFilePath, &credibleSquaringDeploymentRaw)

	var sharedAvsContractsDeploymentRaw config.SharedAvsContractsRaw
	sharedAvsContractsDeploymentFilePath := "../../contracts/script/output/31337/shared_avs_contracts_deployment_output.json"
	sdkutils.ReadJsonConfig(sharedAvsContractsDeploymentFilePath, &sharedAvsContractsDeploymentRaw)

	logger, err := sdklogging.NewZapLogger(aggConfigRaw.Environment)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}
	ethRpcClient, err := eth.NewClient(aggConfigRaw.EthRpcUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}
	ethWsClient, err := eth.NewClient(aggConfigRaw.EthWsUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}

	ecdsaPrivateKeyString := "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	if ecdsaPrivateKeyString[:2] == "0x" {
		ecdsaPrivateKeyString = ecdsaPrivateKeyString[2:]
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivateKeyString)
	if err != nil {
		t.Fatalf("Cannot parse ecdsa private key: %s", err.Error())
	}

	operatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(ecdsaPrivateKey)
	if err != nil {
		t.Fatalf("Cannot get operator address: %s", err.Error())
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		t.Fatalf("Cannot get chainId: %s", err.Error())
	}

	privateKeySigner, err := signer.NewPrivateKeySigner(ecdsaPrivateKey, chainId)
	if err != nil {
		t.Fatalf("Cannot create signer: %s", err.Error())
	}

	config := &config.Config{
		EcdsaPrivateKey:                      ecdsaPrivateKey,
		Logger:                               logger,
		EthRpcUrl:                            aggConfigRaw.EthRpcUrl,
		EthHttpClient:                        ethRpcClient,
		EthWsClient:                          ethWsClient,
		BlsOperatorStateRetrieverAddr:        common.HexToAddress(sharedAvsContractsDeploymentRaw.BlsOperatorStateRetrieverAddr),
		LambadaCoprocessorServiceManagerAddr: common.HexToAddress(credibleSquaringDeploymentRaw.Addresses.LambadaCoprocessorServiceManagerAddr),
		SlasherAddr:                          common.HexToAddress(""),
		AggregatorServerIpPortAddr:           aggConfigRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:            aggConfigRaw.RegisterOperatorOnStartup,
		Signer:                               privateKeySigner,
		OperatorAddress:                      operatorAddr,
		BlsPublicKeyCompendiumAddress:        common.HexToAddress(aggConfigRaw.BLSPubkeyCompendiumAddr),
	}

	// Prepare the config file for operator
	nodeConfig := types.NodeConfig{}
	nodeConfigFilePath := "../../config-files/operator.anvil.yaml"
	err = sdkutils.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}
	// Register operator
	// log.Println("registering operator for integration tests")
	// we need to do this dynamically and can't just hardcode a registered operator into the anvil
	// state because the anvil state dump doesn't also dump the receipts tree so we lose events,
	// and the aggregator thus can't get the operator's pubkey
	// operatorRegistrationCmd := exec.Command("bash", "./operator-registration.sh")
	// err = operatorRegistrationCmd.Run()
	// if err != nil {
	// 	t.Fatalf("Failed to register operator: %s", err.Error())
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
	defer cancel()
	// start operator
	// the passwords are set to empty strings
	log.Println("starting operator for integration tests")
	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")
	nodeConfig.BlsPrivateKeyStorePath = "../keys/test.bls.key.json"
	nodeConfig.EcdsaPrivateKeyStorePath = "../keys/test.ecdsa.key.json"
	nodeConfig.RegisterOperatorOnStartup = true
	nodeConfig.EthRpcUrl = "http://" + anvilEndpoint
	nodeConfig.EthWsUrl = "ws://" + anvilEndpoint
	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		t.Fatalf("Failed to create operator: %s", err.Error())
	}
	go operator.Start(ctx)
	log.Println("Started operator. Sleeping 15 seconds to give it time to register...")
	time.Sleep(15 * time.Second)

	// start aggregator
	log.Println("starting aggregator for integration tests")
	agg, err := aggregator.NewAggregator(config)
	if err != nil {
		t.Fatalf("Failed to create aggregator: %s", err.Error())
	}
	go agg.Start(ctx)
	log.Println("Started aggregator. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	// get avsRegistry client to interact with the chain
	avsReader, err := chainio.NewAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	// check if the task is recorded in the contract for task index 1
	taskHash, err := avsReader.AvsServiceBindings.TaskManager.AllTaskHashes(&bind.CallOpts{}, 1)
	if err != nil {
		t.Fatalf("Cannot get task hash: %s", err.Error())
	}
	if taskHash == [32]byte{} {
		t.Fatalf("Task hash is empty")
	}

	// check if the task response is recorded in the contract for task index 1
	taskResponseHash, err := avsReader.AvsServiceBindings.TaskManager.AllTaskResponses(&bind.CallOpts{}, 1)
	log.Printf("taskResponseHash: %v", taskResponseHash)
	if err != nil {
		t.Fatalf("Cannot get task response hash: %s", err.Error())
	}
	if taskResponseHash == [32]byte{} {
		t.Fatalf("Task response hash is empty")
	}

}

func startAnvilTestContainer() testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "ghcr.io/foundry-rs/foundry:latest",
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: filepath.Join(integrationDir, "avs-and-eigenlayer-deployed-anvil-state.json"),
				},
				Target: "/root/.anvil/state.json",
			},
		},
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--load-state", "/root/.anvil/state.json"},
		ExposedPorts: []string{"8545/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
	}
	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	return anvilC
}
*/
