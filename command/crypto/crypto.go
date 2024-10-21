package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/godevopsdev/dvps/color"
	"os"
	"path/filepath"
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

	//Validate folder and file
	// Extract the folder structure
	folders := filepath.Dir(keyname)
	if folders != "" {
		// Create the output directory if it doesn't exist
		if err := os.MkdirAll(folders, 0755); err != nil {
			fmt.Println(color.Red("Failed to create output directory:", err))
			return
		}
	}

	// Step 1: Generate RSA 2048-bit key pair
	privateKey, err := generateRSAKeys()
	if err != nil {
		fmt.Println(color.Red("Error generating RSA key:", err))
		return
	}

	// Step 2: Export private key to PEM format
	privateKeyPEM := exportPrivateKeyToPEM(privateKey)
	// Step 3: Export public key to standard PEM format
	publicKeyPEM := exportPublicKeyToPEM(&privateKey.PublicKey)

	// Save the keys to files

	err = os.WriteFile(fmt.Sprintf("%s.pem", keyname), []byte(privateKeyPEM), 0600)
	if err != nil {
		fmt.Println(color.Red("Error saving private key file:", err))
		return
	}
	err = os.WriteFile(fmt.Sprintf("%s.pub", keyname), []byte(publicKeyPEM), 0600)
	if err != nil {
		fmt.Println(color.Red("Error saving public pub key file:", err))
		return
	}
	fmt.Println(color.Green("Keys saved successfully: %s.pub", keyname))
}
