package main

import (
	"crypto/rand"
	"github.com/ethereum/go-ethereum/common"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"strings"
	"testing"
)

func TestGetPeerIDFromPrivateKey(t *testing.T) {
	// Generate a private key
	priv, _, err := crypto.GenerateSecp256k1Key(rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Marshal the private key to bytes
	privBytes, err := priv.Raw()
	if err != nil {
		t.Fatalf("Failed to marshal private key: %v", err)
	}

	// Get the peer ID from the private key
	_, err = getPeerIDFromPrivateKey(privBytes)
	if err != nil {
		t.Fatalf("Failed to get peer ID from private key: %v", err)
	}
}

func TestGetEthereumAddressFromPeerID(t *testing.T) {
	// Generate a private key
	priv, _, err := crypto.GenerateSecp256k1Key(rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Get ethereum address from this private key
	privBytes, err := priv.Raw()
	if err != nil {
		t.Fatalf("Failed to marshal private key: %v", err)
	}
	privKey, err := crypto2.ToECDSA(privBytes)
	if err != nil {
		t.Fatalf("Failed to convert private key: %v", err)
	}
	address1 := crypto2.PubkeyToAddress(privKey.PublicKey).String()

	// Get the peer ID from the private key
	id, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		t.Fatalf("Failed to get peer ID from private key: %v", err)
	}

	// Get the Ethereum address from the peer ID
	address2, err := getEthereumAddressFromPeerID(id)
	if err != nil {
		t.Fatalf("Failed to get Ethereum address from peer ID: %v", err)
	}

	checksumAddress1 := common.HexToAddress(address1).Hex()
	checksumAddress2 := common.HexToAddress(address2).Hex()

	if strings.EqualFold(address1, address2) && checksumAddress1 != checksumAddress2 {
		t.Fatalf("The addresses are not the same or not valid checksum addresses.")
	}
}
