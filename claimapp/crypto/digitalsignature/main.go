package main
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 1. Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// Message to sign
	message := []byte("Doctor approved patient discharge #5678")

	// 2. Hash the message (always hash before signing)
	hashed := sha256.Sum256(message)

	// 3. Sign with private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, 0, hashed[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("🖊 Signature (first 32 bytes): %x...\n", signature[:32])

	// 4. Verify with public key
	err = rsa.VerifyPKCS1v15(publicKey, 0, hashed[:], signature)
	if err != nil {
		fmt.Println("Verification failed!")
	} else {
		fmt.Println("Signature verified successfully.")
	}
}