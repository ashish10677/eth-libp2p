package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"testing"
)

const (
	testKeystoreDir     = "./test_keystore"
	testDefaultPassword = "test_password"
)

func TestCreateKeystore(t *testing.T) {
	// Generate a random private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Test with a password
	err = createKeystore(privateKey, testDefaultPassword)
	if err != nil {
		t.Fatalf("Failed to create keystore with password: %v", err)
	}

	privateKey2, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	// Test with an empty password
	err = createKeystore(privateKey2, "")
	if err != nil {
		t.Fatalf("Failed to create keystore with empty password: %v", err)
	}

	// Cleanup
	err = os.RemoveAll(testKeystoreDir)
	if err != nil {
		t.Fatalf("Failed to remove test keystore directory: %v", err)
	}
}
