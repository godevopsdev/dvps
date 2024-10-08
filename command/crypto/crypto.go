package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

// generateRSAKeys generates a new RSA 2048-bit key pair
func generateRSAKeys() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// exportPrivateKeyToPEM exports the private key to PEM format
func exportPrivateKeyToPEM(privateKey *rsa.PrivateKey) string {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	return string(privateKeyPEM)
}

// base64URLEncode encodes bytes in a base64 URL format without padding
func base64URLEncode(input []byte) string {
	encoded := base64.RawURLEncoding.EncodeToString(input)
	return encoded
}

// exportPublicKeyToPEM exports the public key to standard PEM format
func exportPublicKeyToPEM(publicKey *rsa.PublicKey) string {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Error marshalling public key:", err)
		return ""
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return string(publicKeyPEM)
}

// GenerateRSA2048
func GenerateRSA2048(keyname string) {
	// Step 1: Generate RSA 2048-bit key pair
	privateKey, err := generateRSAKeys()
	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		return
	}
	// Step 2: Export private key to PEM format
	privateKeyPEM := exportPrivateKeyToPEM(privateKey)
	// Step 3: Export public key to standard PEM format
	publicKeyPEM := exportPublicKeyToPEM(&privateKey.PublicKey)
	// Create the output directory if it doesn't exist
	if err := os.MkdirAll("keys", 0755); err != nil {
		fmt.Println("Failed to create output directory:", err)
		return
	}
	// Save the keys to files

	err = os.WriteFile(fmt.Sprintf("keys/%s.pem", keyname), []byte(privateKeyPEM), 0600)
	if err != nil {
		fmt.Println("Error saving private key file:", err)
		return
	}
	err = os.WriteFile(fmt.Sprintf("keys/%s.pub", keyname), []byte(publicKeyPEM), 0600)
	if err != nil {
		fmt.Println("Error saving public pub key file:", err)
		return
	}
	fmt.Println("Keys saved successfully!")
}
