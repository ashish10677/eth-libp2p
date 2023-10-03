package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"log"
	"os"
)

func generateEthCompatiblePrivateKey() ([]byte, error) {
	privKey, _, err := crypto.GenerateSecp256k1Key(rand.Reader)
	if err != nil {
		return nil, err
	}
	privateKeyBytes, err := privKey.Raw()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Private Key: 0x%x\n", privateKeyBytes)
	return privateKeyBytes, nil
}

func generateKeystore(privateKey *ecdsa.PrivateKey, password string) error {
	dir := "./keystore"
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return err
	}

	// Set a passphrase for the keystore. This is what you'll use to unlock it later.
	if password == "" {
		password = "your-strong-password"
	}

	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return err
	}
	fmt.Printf("Keystore file created for address: %s\n", account.Address.Hex())
	return nil
}

func getIdFromPrivateKey(privateKeyBytes []byte) (peer.ID, error) {
	key, err := crypto.UnmarshalSecp256k1PrivateKey(privateKeyBytes)
	if err != nil {
		return "", err
	}
	id, err := peer.IDFromPrivateKey(key)
	if err != nil {
		return "", err
	}
	return id, nil
}

func main() {
	privateKeyBytes, err := generateEthCompatiblePrivateKey()
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}
	privateKey, err := crypto2.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error converting private key to ECDSA: %v", err)
	}
	err = generateKeystore(privateKey, "")
	if err != nil {
		log.Fatalf("Error generating keystore: %v", err)
	}
	id, err := getIdFromPrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error generating peer ID: %v", err)
	}
	fmt.Printf("Peer ID: %s\n", id)
}
