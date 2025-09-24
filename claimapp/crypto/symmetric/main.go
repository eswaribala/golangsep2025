package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	// Exactly 32 bytes (AES-256). Hyphen removed.
	key := []byte("thisIs32ByteSecretKeyForAESGCM!!")
	fmt.Println("key length:", len(key)) // should print 32

	plaintext := []byte("Patient report: Blood pressure normal")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic("decryption failed (tampered?)")
	}
	fmt.Println("Decrypted:", string(decrypted))
}
