package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nikolalohinski/gonja"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
)

func GenerateOperator(ctx *cli.Context) error {
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
	blsKey, err := GenerateBLSKey(operatorDir)
	if err != nil {
		return fmt.Errorf("failed to generate BLS key - %s", err)
	}
	ecdsaKey, err := GenerateECDSAKey(operatorDir)
	if err != nil {
		return fmt.Errorf("failed to generate ECDSDA key - %s", err)
	}

	// Generate configuration file
	configTmpl, err := gonja.FromFile("./tests/jinja/operator-config.j2")
	if err != nil {
		return err
	}
	config, err := configTmpl.Execute(gonja.Context{
		"avs_deployment_output_path":        "/path/to/avs/deployment/output",
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

func GenerateBLSKey(path string) (OperatorBLSKey, error) {
	key, err := bls.GenRandomBlsKeys()
	if err != nil {
		return OperatorBLSKey{}, err
	}

	password := generateRandomPassword()

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

func GenerateECDSAKey(path string) (OperatorECDSAKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return OperatorECDSAKey{}, err
	}
	privateKeyHex := "0x" + hex.EncodeToString(key.D.Bytes())

	password := generateRandomPassword()

	keyPath := filepath.Join(path, "ecdsa.key.json")
	if err = ecdsa.WriteKey(keyPath, key, password); err != nil {
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

func generateRandomPassword() string {
	// Seed the random number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define character sets for the password
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	//specialSymbols := "!@#$%^&*()-_+=[]{}|;:,.<>?/\\"

	// Combine character sets into one
	//allCharacters := uppercaseLetters + lowercaseLetters + digits + specialSymbols
	allCharacters := uppercaseLetters + lowercaseLetters + digits

	// Length of the password you want
	passwordLength := 20

	// Generate the password
	password := make([]byte, passwordLength)
	for i := range password {
		password[i] = allCharacters[random.Intn(len(allCharacters))]
	}
	return string(password)
}
