package main

import (
	"fmt"
	"log"
)

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
