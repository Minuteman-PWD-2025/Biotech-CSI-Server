package main

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
)

// Add password encryption
func ValidateLogin(users map[string]string, tokens []string, email string, pass string) ([]string, error) {
	_, exists := users[email]
	if !exists {
		return tokens, errors.New("invalid login")
	}

	newToken, err := GenerateToken(10)
	if err != nil {
		return tokens, errors.New("error generating token: " + err.Error())
	}

	tokens = append(tokens, newToken)

	return tokens, nil
}

func GenerateToken(length int) (string, error) {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.EncodeToString(randomBytes)[:length], nil
}
