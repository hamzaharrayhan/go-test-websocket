package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func Encrypt(data []byte, key []byte) string {
	keyHex, err := hex.DecodeString(string(key))
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher(keyHex)
	if err != nil {
		return ""
	}

	padding := aes.BlockSize - (len(data) % aes.BlockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddedData := append(data, padtext...)

	ciphertext := make([]byte, aes.BlockSize+len(paddedData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}

	base64IV := base64.StdEncoding.EncodeToString(iv)

	mode := cipher.NewCBCEncrypter(block, []byte(base64IV[:aes.BlockSize]))
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedData)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(data string, key []byte) []byte {
	keyHex, err := hex.DecodeString(string(key))
	if err != nil {
		return nil
	}

	block, err := aes.NewCipher(keyHex)
	if err != nil {
		return nil
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil
	}

	if len(ciphertext) < aes.BlockSize {
		return nil
	}

	iv := ciphertext[:aes.BlockSize]
	base64IV := base64.StdEncoding.EncodeToString(iv)

	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, []byte(base64IV[:aes.BlockSize]))

	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding
	paddingLength := int(ciphertext[len(ciphertext)-1])
	if paddingLength > aes.BlockSize || paddingLength == 0 {
		return nil
	}

	return ciphertext[:len(ciphertext)-paddingLength]
}
