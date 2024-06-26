package actions

import (
	"encoding/hex"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
)

func GenerateOperator(ctx *cli.Context) error {
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
