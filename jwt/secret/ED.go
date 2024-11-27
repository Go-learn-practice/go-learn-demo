package secret

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

type EdGenerator struct {
}

func (ed *EdGenerator) Generate() (*OutSecret, error) {
	out := &OutSecret{}
	var err error
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	x509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	privateBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509PrivateKey,
	}
	privateFilePath := KEYPATH + "/ed/private.pem"
	file, err := os.OpenFile(privateFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	pem.Encode(file, privateBlock)

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}
	publicFilePath := KEYPATH + "/ed/public.pem"
	file, err = os.OpenFile(publicFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	pem.Encode(file, publicBlock)
	out.PrivateKeyFile = privateFilePath
	out.PublicKeyFile = publicFilePath

	return out, nil
}
