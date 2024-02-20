package core

import (
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/accounts/abi"

	taskmanager "github.com/zippiehq/cartesi-lambada-coprocessor/contracts/bindings/LambadaCoprocessorTaskManager"
)

// GetTaskResponseDigest returns the hash of the TaskResponse
func GetTaskResponseDigest(r *taskmanager.ILambadaCoprocessorTaskManagerTaskResponse) ([32]byte, error) {
	// The order here has to match the field ordering of taskmanager.ILambadaCoprocessorTaskManagerTaskResponse
	t, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "outputHash",
			Type: "bytes32",
		},
	})
	if err != nil {
		return [32]byte{}, err
	}

	return hashObject(t, r)
}

// GetTaskBatchDigest returns the hash of the TaskBatch
func GetTaskBatchDigest(b *taskmanager.ILambadaCoprocessorTaskManagerTaskBatch) ([32]byte, error) {
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
			Type: "bytes32",
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

// BINDING UTILS - conversion from contract structs to golang structs

// BN254.sol is a library, so bindings for G1 Points and G2 Points are only generated
// in every contract that imports that library. Thus the output here will need to be
// type casted if G1Point is needed to interface with another contract (eg: BLSPublicKeyCompendium.sol)
func ConvertToBN254G1Point(input *bls.G1Point) taskmanager.BN254G1Point {
	output := taskmanager.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) taskmanager.BN254G2Point {
	output := taskmanager.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
