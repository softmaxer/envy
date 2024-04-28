package vault

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func encrypt(key, fileContentsReader io.Reader) ([]byte, error) {
	keyText, err := io.ReadAll(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyText)
	if err != nil {
		return nil, err
	}
	fileContents, err := io.ReadAll(fileContentsReader)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(fileContents)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, fileContentsReader io.Reader) ([]byte, error) {
	keyText, err := io.ReadAll(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyText)
	if err != nil {
		return nil, err
	}
	fileContents, err := io.ReadAll(fileContentsReader)
	if err != nil {
		return nil, err
	}
	if len(fileContents) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := fileContents[:aes.BlockSize]
	fileContents = fileContents[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(fileContents, fileContents)
	data, err := base64.StdEncoding.DecodeString(string(fileContents))
	if err != nil {
		return nil, err
	}
	return data, nil
}
