package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	crypto2 "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"golang.org/x/crypto/sha3"
)

const (
	defaultPassword = "your-strong-password"
	keystoreDir     = "./keystore"
)

func generatePrivateKey() ([]byte, *ecdsa.PrivateKey, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	privateKeyBytes := crypto.FromECDSA(privKey)
	fmt.Printf("Private Key: 0x%x\n", privateKeyBytes)
	return privateKeyBytes, privKey, nil
}

func createKeystore(privateKey *ecdsa.PrivateKey, password string) error {
	if password == "" {
		password = defaultPassword
	}

	err := os.MkdirAll(keystoreDir, 0700)
	if err != nil {
		return err
	}

	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		return err
	}
	fmt.Printf("Keystore file created for address: %s\n", account.Address.Hex())
	return nil
}

func getPeerIDFromPrivateKey(privateKeyBytes []byte) (peer.ID, error) {
	key, err := crypto2.UnmarshalSecp256k1PrivateKey(privateKeyBytes)
	if err != nil {
		return "", err
	}
	return peer.IDFromPrivateKey(key)
}

func getPublicKeyFromCompressedKey(compressedKey []byte) (*btcec.PublicKey, error) {
	return btcec.ParsePubKey(compressedKey)
}

func getAddressFromPublicKey(pubKey *btcec.PublicKey) []byte {
	uncompressedKey := pubKey.SerializeUncompressed()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(uncompressedKey[1:])
	return hash.Sum(nil)[12:]
}

func getEthereumAddressFromPeerID(id peer.ID) (string, error) {
	key, err := id.ExtractPublicKey()
	if err != nil {
		return "", err
	}
	rawPublicKey, err := key.Raw()
	if err != nil {
		return "", err
	}
	fmt.Printf("Length of public key: %d\n", len(rawPublicKey))
	fmt.Printf("Compressed Public Key: 0x%x\n", rawPublicKey)
	pubKey, err := getPublicKeyFromCompressedKey(rawPublicKey)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(getAddressFromPublicKey(pubKey)), nil
}

func main() {
	privateKeyBytes, privateKey, err := generatePrivateKey()
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}
	err = createKeystore(privateKey, "")
	if err != nil {
		log.Fatalf("Error generating keystore: %v", err)
	}

	id, err := getPeerIDFromPrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error generating peer ID: %v", err)
	}
	fmt.Printf("Peer ID: %s\n", id)
	address, err := getEthereumAddressFromPeerID(id)
	if err != nil {
		log.Fatalf("Error generating address: %v", err)
	}
	fmt.Printf("Address from ID: 0x%s\n", address)
}
