package core

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	tm "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
	"golang.org/x/crypto/sha3"
)

// TaskResponseSigHash computes hash, which uniquely indentifies task and its output
func TaskResponseSigHash(batchIdx TaskBatchIndex, t Task, r TaskResponse) ([32]byte, error) {
	// The order here has to match the field ordering of taskmanager.ILambadaCoprocessorTaskManagerTaskResponse
	ot, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "batchIndex",
			Type: "uint32",
		},

		{
			Name: "programId",
			Type: "bytes",
		},
		{
			Name: "inputHash",
			Type: "bytes",
		},
		{
			Name: "resultCID",
			Type: "bytes",
		},
		{
			Name: "outputHash",
			Type: "bytes32",
		},
	})
	if err != nil {
		return [32]byte{}, err
	}

	o := struct {
		BatchIndex sdktypes.TaskIndex
		ProgramId  []byte
		InputHash  []byte
		ResultCID  []byte
		OutputHash [32]byte
	}{
		BatchIndex: batchIdx,
		ProgramId:  t.ProgramID,
		InputHash:  t.InputHash,
		ResultCID:  r.ResultCID,
		OutputHash: [32]byte(r.OutputHash),
	}

	hash, err := hashObject(ot, &o)

	return hash, err
}

// TaskBatchHash returns the hash of the TaskBatch
func TaskBatchHash(b *tm.ILambadaCoprocessorTaskManagerTaskBatch) ([32]byte, error) {
	t, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "index",
			Type: "uint32",
		},
		{
			Name: "blockNumber",
			Type: "uint32",
		},
		{
			Name: "merkeRoot",
			Type: "bytes32",
		},
		{
			Name: "quorumNumbers",
			Type: "bytes",
		},
		{
			Name: "quorumThresholdPercentage",
			Type: "uint32",
		},
	})
	if err != nil {
		return [32]byte{}, err
	}

	return hashObject(t, b)
}

func hashObject(t abi.Type, value interface{}) ([32]byte, error) {
	// Pack object to bytes - abi.encode()
	arguments := abi.Arguments{
		{
			Type: t,
		},
	}
	data, err := arguments.Pack(value)
	if err != nil {
		return [32]byte{}, err
	}

	// Compute keccack256.
	var digest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	copy(digest[:], hasher.Sum(nil)[:32])

	return digest, nil
}

func Keccack256(data []byte) []byte {
	return crypto.Keccak256(data)
}
