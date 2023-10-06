package main

import (
	"bytes"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKeyBytes, privKey, err := generatePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	if privateKeyBytes == nil {
		t.Fatal("Failed to generate private key bytes, got nil")
	}
	if privKey == nil {
		t.Fatal("Failed to generate private key, got nil")
	}
}

func TestGetPublicKeyFromCompressedKey(t *testing.T) {
	// Generate a private key
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Get the public key and compress it
	pubKey := privKey.PubKey()
	compressedKey := pubKey.SerializeCompressed()

	// Use the function to get the public key from the compressed key
	resultPubKey, err := getPublicKeyFromCompressedKey(compressedKey)
	if err != nil {
		t.Fatalf("Failed to get public key from compressed key: %v", err)
	}

	// Compare the original public key with the one obtained from the function
	if !bytes.Equal(pubKey.SerializeUncompressed(), resultPubKey.SerializeUncompressed()) {
		t.Fatal("Original public key and public key obtained from function do not match")
	}
}

func TestGetAddressFromPublicKey(t *testing.T) {
	// Generate a private key
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Get the public key
	pubKey := privKey.PubKey()

	// Get the address from the public key
	address := getAddressFromPublicKey(pubKey)

	// Check if the address is not nil
	if address == nil {
		t.Fatalf("Failed to get address from public key")
	}
}
