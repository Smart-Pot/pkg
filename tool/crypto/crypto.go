/*
Package crypto implements functions for simple encrypt and decrypt process.
In encryption and decryption process, package uses aes algorithm */
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	// ErrShortCipher codes returned by failures when key is shorter than expected value.
	ErrShortCipher = errors.New("cipher text is too short")
)

var privateKey []byte

// PrivateKey returns private key that using in encrpytion and decryption process
// And makes private key access to read-only.
func PrivateKey() string {
	return string(privateKey)
}

const defaultKey = "placeholder_key_"

func init() {

	k := os.Getenv("VERIF_PRIVATE_KEY")
	if k == "" {
		k = defaultKey
	}
	if len(k) != 16 {
		panic("Crypto key is too short:")
	}
	privateKey = []byte(k)
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) (string, error) {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		//panic("ciphertext too short")
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}

// Encrypt text with private key using aes algorithm
func Encrypt(text string) (string, error) {
	return encrypt(privateKey, text)
}
// Decrypt text with private key using aes algorithm
func Decrypt(hash string) (string, error) {
	return decrypt(privateKey, hash)
}
