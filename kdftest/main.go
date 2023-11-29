package main

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
	"crypto/md5"
	"encoding/hex"
	//"io"
)

func calculateMD5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func deriveKey(password []byte, salt []byte, N int, r int, p int, keyLength int) ([]byte, error) {
	derivedKey, err := scrypt.Key(password, salt, N, r, p, keyLength)
	if err != nil {
		return nil, err
	}
	return derivedKey, nil
}

func main() {
	md5 := calculateMD5("123456")

	fmt.Printf("md5: %x\n",md5)

	password := []byte(md5)
	salt := make([]byte, 524) // Generate a random salt for each user and store it securely.

	// Parameters for scrypt
	N := 16384   // CPU/memory cost factor (higher values make the function slower)
	r := 8       // block size
	p := 1       // parallelization factor
	keyLength := 524 // length of the derived key

	derivedKey, err := deriveKey(password, salt, N, r, p, keyLength)
	if err != nil {
		fmt.Println("Key derivation error:", err)
		return
	}

	fmt.Printf("Derived Key: %x\n", derivedKey)
}
