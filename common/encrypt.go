package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

func Encrypt(secret string, target string) (string, error) {
	secretBytes := []byte(secret)
	targetBytes := []byte(target)

	c, err := aes.NewCipher(secretBytes)
	if err != nil {
		return "", errors.New("Cipher Failed")
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return "", errors.New("GCM Failed")
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.New("Error creating nonce")
	}
	encrypted := string(gcm.Seal(nonce, nonce, targetBytes, nil))
	return encrypted, nil
}

func Decrypt() {
}

func CreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyPassword(hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
