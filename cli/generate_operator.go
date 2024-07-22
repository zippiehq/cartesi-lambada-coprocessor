package main

import (
	"crypto/ecdsa"
	cryptorand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/nikolalohinski/gonja"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigensdk_ecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	eigensdk_utils "github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/zippiehq/cartesi-lambada-coprocessor/core/config"
)

func GenerateOperator(ctx *cli.Context) error {
	deploymentOutputPath := ctx.String(deploymentOutputFlag.Name)
	var deploymentOutput config.AVSDeployment
	if err := eigensdk_utils.ReadJsonConfig(deploymentOutputPath, &deploymentOutput); err != nil {
		return fmt.Errorf("failed to read deployment output - %s", err)
	}

	operatorDir := ctx.String("operator-dir")

	// Clear operator directory
	if _, err := os.Stat(operatorDir); err != nil {
		return err
	} else {
		if err = os.RemoveAll(operatorDir); err != nil {
			return err
		}
	}

	// Generate BLS and ECDSA keys
	blsKey, err := GenerateBLSKey(operatorDir, cryptorand.Reader)
	if err != nil {
		return fmt.Errorf("failed to generate BLS key - %s", err)
	}
	ecdsaKey, err := GenerateECDSAKey(operatorDir, cryptorand.Reader)
	if err != nil {
		return fmt.Errorf("failed to generate ECDSDA key - %s", err)
	}

	// Generate configuration file
	configTmpl, err := gonja.FromFile("./tests/jinja/operator-docker-compose.j2")
	if err != nil {
		return err
	}
	config, err := configTmpl.Execute(gonja.Context{
		"address":                           ecdsaKey.Address,
		"registry_coordinator_address":      deploymentOutput.Addresses.RegistryCoordinator,
		"operator_state_retriever_address":  deploymentOutput.Addresses.OperatorStateRetriever,
		"ecdsa_private_key_store_path":      ecdsaKey.FilePath,
		"bls_private_key_store_path":        blsKey.FilePath,
		"eth_rpc_url":                       "http://127.0.0.1:8545",
		"eth_ws_url":                        "ws://127.0.0.1:8545",
		"aggregator_server_ip_port_address": "127.0.0.1:8090",
		"ipfs_ip_port_address":              "127.0.0.1:5001",
		"lambada_ip_port_address":           "127.0.0.1:3033",
		"eigen_metrics_ip_port_address":     "127.0.0.1:9090",
		"node_api_ip_port_address":          "127.0.0.1:9010",
	})
	if err != nil {
		return err
	}
	configPath := filepath.Join(operatorDir, "config.yaml")
	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()
	if _, err = configFile.WriteString(config); err != nil {
		return err
	}

	fmt.Printf("Generated operator configuration template in %s\n", operatorDir)

	// Print password and private key
	fmt.Printf("Address: %s\n", ecdsaKey.Address)
	fmt.Printf("Private key: %s\n", ecdsaKey.PrivateKey)
	fmt.Printf("BLS key password: %s\n", blsKey.Password)
	fmt.Printf("ECDSA key password: %s\n", ecdsaKey.Password)

	return nil
}

type OperatorBLSKey struct {
	FilePath   string
	Password   string
	PrivateKey string
}

func GenerateBLSKey(path string, rand io.Reader) (OperatorBLSKey, error) {
	// Max random value is order of the curve
	max := new(big.Int)
	max.SetString(fr.Modulus().String(), 10)
	// Generate cryptographically strong pseudo-random between 0 - max
	n, err := cryptorand.Int(rand, max)
	if err != nil {
		return OperatorBLSKey{}, err
	}
	sk := new(fr.Element).SetBigInt(n)
	key := bls.NewKeyPair(sk)

	password, err := generateRandomPassword(rand)
	if err != nil {
		return OperatorBLSKey{}, err
	}

	keyPath := filepath.Join(path, "bls.key.json")
	if err = key.SaveToFile(keyPath, password); err != nil {
		return OperatorBLSKey{}, err
	}

	operatorKey := OperatorBLSKey{
		FilePath:   keyPath,
		Password:   password,
		PrivateKey: key.PrivKey.String(),
	}

	return operatorKey, nil
}

type OperatorECDSAKey struct {
	FilePath   string
	Password   string
	PrivateKey string
	Address    string
}

func GenerateECDSAKey(path string, rand io.Reader) (OperatorECDSAKey, error) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand)
	if err != nil {
		return OperatorECDSAKey{}, err
	}
	privateKeyHex := "0x" + hex.EncodeToString(key.D.Bytes())

	password, err := generateRandomPassword(rand)
	if err != nil {
		return OperatorECDSAKey{}, err
	}

	keyPath := filepath.Join(path, "ecdsa.key.json")
	if err = eigensdk_ecdsa.WriteKey(keyPath, key, password); err != nil {
		return OperatorECDSAKey{}, err
	}

	operatorKey := OperatorECDSAKey{
		FilePath:   keyPath,
		Password:   password,
		PrivateKey: privateKeyHex,
		Address:    crypto.PubkeyToAddress(key.PublicKey).Hex(),
	}

	return operatorKey, nil
}

func generateRandomPassword(rand io.Reader) (string, error) {
	// Define character sets for the password
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	allCharacters := uppercaseLetters + lowercaseLetters + digits

	passwordLength := 20

	// Generate the password
	password := make([]byte, passwordLength)
	for i := range password {
		idx := big.NewInt(int64(len(allCharacters)))
		charIdx, err := cryptorand.Int(rand, idx)
		if err != nil {
			return "", err
		}
		password[i] = allCharacters[charIdx.Int64()]
	}

	//!!!
	fmt.Println(string(password))

	return string(password), nil
}
