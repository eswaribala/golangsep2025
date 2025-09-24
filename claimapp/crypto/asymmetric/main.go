package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 1. Generate RSA key pair (2048 bits is safe; 3072/4096 for higher security)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	message := []byte("Top secret: Insurance claim #2025")

	// 2. Encrypt with PUBLIC key (sender does this)
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(), // hash function
		rand.Reader,  // randomness
		publicKey,    // recipient's public key
		message,      // message to encrypt
		nil,          // optional label (can be nil)
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("🔒 Ciphertext (first 32 bytes): %x...\n", ciphertext[:32])

	// 3. Decrypt with PRIVATE key (recipient does this)
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey, // recipient’s private key
		ciphertext,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted message:", string(plaintext))
}
