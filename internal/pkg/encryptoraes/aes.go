package encryptoraes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type KeySize int

const (
	Size16 KeySize = 16
	Size24 KeySize = 24
	Size32 KeySize = 32
)

func GenerateKey(s KeySize) ([]byte, error) {
	key := make([]byte, s)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}

	return key, nil
}

type Encryptor struct {
	Key      []byte
	MaxPlain int
}

func (e *Encryptor) SetMaxPlain(max int) {
	e.MaxPlain = max
}

func New(key []byte) (*Encryptor, error) {
	if len(key) == 0 {
		return nil, errors.New("key must not be empty")
	}

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("key length must be 16, 24, or 32")
	}

	result := &Encryptor{
		Key: key,
	}

	if result.MaxPlain == 0 {
		result.SetMaxPlain(10 * 1024 * 1024)
	}

	return result, nil
}

func (e *Encryptor) Encrypt(input string) (string, error) {
	if input == "" {
		return "", errors.New("input must not be empty")
	}
	plaintext := []byte(input)
	if len(plaintext) > e.MaxPlain {
		return "", errors.New("plaintext too long")
	}

	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e *Encryptor) Decrypt(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
