package lib

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
)

func readPublicKeyFile(file string) (*rsa.PublicKey, error) {
	publicF, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(publicF *os.File) {
		_ = publicF.Close()
	}(publicF)
	publicKeyData, err := io.ReadAll(publicF)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(publicKeyData)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key.(*rsa.PublicKey), nil
}

func readPrivateKeyFile(file string) (*rsa.PrivateKey, error) {
	privateF, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(privateF *os.File) {
		_ = privateF.Close()
	}(privateF)
	privateKeyData, err := io.ReadAll(privateF)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKeyData)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
