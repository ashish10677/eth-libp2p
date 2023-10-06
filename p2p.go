package main

import (
	"fmt"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

func getPeerIDFromPrivateKey(privateKeyBytes []byte) (peer.ID, error) {
	key, err := crypto.UnmarshalSecp256k1PrivateKey(privateKeyBytes)
	if err != nil {
		return "", err
	}
	return peer.IDFromPrivateKey(key)
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
	return fmt.Sprintf("0x%x", getAddressFromPublicKey(pubKey)), nil
}
