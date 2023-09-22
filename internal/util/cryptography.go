package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

const (
	iterationCount int = 1000
)

type Cryptography struct {
	encryption_key string
	salt           []byte
}

func NewCryptography() Cryptography {
	return Cryptography{
		encryption_key: "A60934D8C1A2AC3A69642A3902198",
		salt:           []byte{99, 52, 2, 24, 51, 67, 22, 88},
	}
}

func (c Cryptography) EncryptUrl(originalUrl string, maxLength int) string {
	plainText := c.padInput(originalUrl)

	secretKey := pbkdf2.Key([]byte(c.encryption_key), c.salt, iterationCount, 32, sha1.New)
	IV := pbkdf2.Key([]byte(c.encryption_key), c.salt, iterationCount, 48, sha1.New)[32:]

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(err)
	}

	plainText, _ = c.pkcs7Pad(plainText, block.BlockSize())

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	bm := cipher.NewCBCEncrypter(block, IV)
	bm.CryptBlocks(ciphertext[aes.BlockSize:], plainText)

	cipherStr := base64.StdEncoding.EncodeToString(ciphertext[aes.BlockSize:])

	if len(cipherStr) > maxLength {
		return cipherStr[:maxLength]
	}

	return cipherStr
}

func (c Cryptography) padInput(input string) []byte {
	b := []byte(input)
	pb := make([]byte, len(b)*2)
	for i, value := range b {
		j := i * 2
		pb[j] = value
		pb[j+1] = 0
	}
	return pb
}

func (c Cryptography) pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, errors.New("invalid blocksize")
	}
	if len(b) == 0 {
		return nil, errors.New("invalid PKCS7 data (empty or not padded)")
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}
