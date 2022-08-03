package cryptox

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

// RSAGenerateKey generates a new RSA key pair.
func RSAGenerateKey(bits int, out io.Writer) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: X509PrivateKey,
	}

	return pem.Encode(out, privateBlock)
}

// RSAGeneratePublicKey generates a public key from a private key
func RSAGeneratePublicKey(priKey []byte, out io.Writer) error {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return errors.New("private key error")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	publicKey := &privateKey.PublicKey

	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	return pem.Encode(out, publicBlock)
}

// RSAEncrypt encrypts data.tmpl with rsa public key
func RSAEncrypt(src, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("the  key is not a rsa public key")
	}

	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
}

// RSADecrypt decrypts data.tmpl using RSA private key.
func RSADecrypt(src, priKey []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
}

// RSASign signs the given data.tmpl using the given private key. The hash defaults to crypto.SHA256
func RSASign(src []byte, priKey []byte, hash ...crypto.Hash) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hashType := crypto.SHA256
	if len(hash) > 0 {
		hashType = hash[0]
	}

	return rsa.SignPKCS1v15(rand.Reader, privateKey, hashType, src)
}

// RSAVerify verifies the given data.tmpl using the given public key. The hash defaults to crypto.SHA256
func RSAVerify(src, sign, pubKey []byte, hash ...crypto.Hash) error {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return errors.New("public key error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return errors.New("the  key is not a rsa public key")
	}

	hashType := crypto.SHA256
	if len(hash) > 0 {
		hashType = hash[0]
	}

	return rsa.VerifyPKCS1v15(publicKey, hashType, src, sign)
}
