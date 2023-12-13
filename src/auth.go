package main

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
)

func ValidateLogin(users map[string]string, tokens []string, email string, pass string) ([]string, string, error) {
	_, exists := users[email]
	if !exists {
		return tokens, "Error Returned", errors.New("invalid login")
	}

	newToken, err := GenerateToken(20, tokens)
	if err != nil {
		return tokens, "Error Returned", errors.New("error generating token: " + err.Error())
	}

	tokens = append(tokens, newToken)

	return tokens, newToken, nil
}

func GenerateToken(length int, tokens []string) (string, error) {
	// initialize new string variable for new token
	var token string

	// loops until unique token is generated
	for {
		// initialize new boolean
		unique := false

		// make a list of random bytes
		randomBytes := make([]byte, 32)

		// checks validity of bytes
		_, err := rand.Read(randomBytes)
		if err != nil {
			return "", err
		}

		// encodes bytes as a randomized string limited to a set length
		token = base32.StdEncoding.EncodeToString(randomBytes)[:length]

		// if no tokens in list, token is unique
		if len(tokens) < 1 {
			break
		}

		// if token already in list, token is not unique, try again
		for _, val := range tokens {
			if token == val {
				continue
			}
			unique = true
		}

		// break loop if unique
		if unique {
			break
		}
	}

	// return unique token and nil error
	return token, nil
}
