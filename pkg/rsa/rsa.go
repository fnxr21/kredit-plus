package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateRsaPem() {
	// Generate RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Save private key to file
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	privatePathKey := "pkg/rsa/" + "gate-sap-private.pem"
	err = os.WriteFile(privatePathKey, privateKeyPEM, 0600)
	if err != nil {
		fmt.Println("Error saving private key to file:", err)
		return
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Error decoding private key")
		// return
	}
	privateKeyFromPEM, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		// return
	}

	// Generate RSA public key
	publicKey := &privateKeyFromPEM.PublicKey

	pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	publicPathKey := "pkg/rsa/" + "gate-sap-public.pem"

	err = os.WriteFile(publicPathKey, publicKeyPEM, 0600)
	if err != nil {
		fmt.Println("Error saving private key to file:", err)
		return
	}

	fmt.Println("Private key saved to private.pem")

}

// input from file private
func GeneratePublicByFilePrv() error {

	// Load private key from file
	privateKeyFile, err := os.ReadFile("pkg/rsa/gate-sap-private.pem")

	// pemData, err := os.ReadFile("/app/your_pem_file.pem")
	if err != nil {
		fmt.Println("Error reading private key from file:", err)

		return err
	}

	// fmt.Println(privateKeyFile)
	block, _ := pem.Decode(privateKeyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Error decoding private key")
		return err
	}
	privateKeyFromPEM, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return err
	}

	// Generate RSA public key
	publicKey := &privateKeyFromPEM.PublicKey

	pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	err = os.WriteFile("gate-sap-public.pem", publicKeyPEM, 0600)
	if err != nil {
		fmt.Println("Error saving private key to file:", err)
		return err
	}

	fmt.Println("Private key saved to private.pem")

	return nil
}

func RSAEncryptMesagge(messageSTR string) ([]byte, *rsa.PrivateKey) {

	// Load private key from file
	privateKeyFile, err := os.ReadFile("pkg/rsa/gate-sap-private.pem")

	// pemData, err := os.ReadFile("/app/your_pem_file.pem")
	if err != nil {
		fmt.Println("Error reading private key from file:", err)

		privateKeyFile, err = os.ReadFile("app/gate-sap-private.pem")
		if err != nil {
			fmt.Println("pem not found check again")
		}
		// return
	}

	// fmt.Println(privateKeyFile)
	block, _ := pem.Decode(privateKeyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Error decoding private key")
		// return
	}
	privateKeyFromPEM, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		// return
	}

	// Generate RSA public key
	publicKey := &privateKeyFromPEM.PublicKey

	pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	// publicKeyPEM := pem.EncodeToMemory(&pem.Block{
	// 	Type:  "RSA PUBLIC KEY",
	// 	Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	// })

	// fmt.Println("Public key:")
	// fmt.Println(string(publicKeyPEM))

	// Encrypt a message with the public key
	message := messageSTR
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	if err != nil {
		fmt.Println("Error encrypting message:", err)
		// return
	}

	return ciphertext, privateKeyFromPEM
}

func RSADecryptMessage(ciphertext []byte) string {
	privateKeyFile, _ := os.ReadFile("pkg/rsa/gate-sap-private.pem")
	// fmt.Println(privateKeyFile)
	block, _ := pem.Decode(privateKeyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Error decoding private key")
		// return
	}
	privateKeyFromPEM, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		// return
	}

	// Decrypt the message with the private key
	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKeyFromPEM, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting message:", err)
		return ""
	}

	fmt.Println("Original message:", string(decryptedMessage))
	return string(decryptedMessage)
}
