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

	// Message to encrypt
	message := []byte("Confidential data: Insurance claim #2025")

	// 2. Encrypt with PUBLIC key
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(), // hash function
		rand.Reader,  // randomness
		publicKey,    // receiver's public key
		message,      // message
		nil,          // optional label
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted (first 32 bytes): %x...\n", ciphertext[:32])

	// 3. Decrypt with PRIVATE key
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey, // owner's private key
		ciphertext,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted message:", string(plaintext))
}
