package main

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
)

// ValidateLogin validates the login credentials of a user.
// It takes a map of users, a slice of tokens, an email, and a password as input.
// It returns the updated slice of tokens and an error if the login is invalid.
func ValidateLogin(users map[string]string, tokens []string, email string, pass string) ([]string, error) {
	_, exists := users[email]
	if !exists {
		return tokens, errors.New("invalid login")
	}

	newToken, err := GenerateToken(20, tokens)
	if err != nil {
		return tokens, errors.New("error generating token: " + err.Error())
	}

	tokens = append(tokens, newToken)

	return tokens, nil
}

// GenerateToken generates a unique token of the specified length.
// It takes the length of the token and a slice of existing tokens as input.
// It returns the generated token and an error if the token generation fails.
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
