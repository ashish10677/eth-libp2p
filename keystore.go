package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"os"
)

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
