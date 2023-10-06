package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
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

func getPublicKeyFromCompressedKey(compressedKey []byte) (*btcec.PublicKey, error) {
	return btcec.ParsePubKey(compressedKey)
}

func getAddressFromPublicKey(pubKey *btcec.PublicKey) []byte {
	uncompressedKey := pubKey.SerializeUncompressed()
	hash := sha3.NewLegacyKeccak256()
	hash.Write(uncompressedKey[1:])
	return hash.Sum(nil)[12:]
}
