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

const defaultKey = "placeholder_key_"

var (
	// ForgotPwdCip :
	ForgotPwdCip Cipher
	// VerifyMailCip :
	VerifyMailCip Cipher
)


func init(){
	forgotCipSecret := os.Getenv("FOROT_PWD_SECRET")
	if forgotCipSecret == "" {
		forgotCipSecret = defaultKey
	}
	
	verifyMailSecret := os.Getenv("VERIFY_MAIL_SECRET")
	if verifyMailSecret  == "" {
		verifyMailSecret = defaultKey
	}	
	
	ForgotPwdCip = &_cipher{
		key: []byte(forgotCipSecret),
	}
	VerifyMailCip = &_cipher {
		key:[]byte(verifyMailSecret),
	}
}



// Cipher :
type Cipher interface {
	Encrypt(text string) (string , error)
	Decrypt(hash string) (string,error)
}

type _cipher  struct {
	key []byte
}

// NewCipher :
func NewCipher(key string) Cipher {
	return &_cipher{
		key : []byte(key),
	}
}

// Encrypt text with private key using aes algorithm
func (c *_cipher) Encrypt(text string) (string, error) {
	return encrypt(c.key, text)
}
// Decrypt text with private key using aes algorithm
func (c *_cipher) Decrypt(hash string) (string, error) {
	return decrypt(c.key, hash)
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


