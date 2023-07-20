package main

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

func main() {
	symmetricKey := []byte("ROSE IS RED, VIOLET IS BLUE, I'M BAD AT PROGRAMMING") // Must be 32 bytes
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "NPC",
		Issuer:     "Wotaku Bau Bawang",
		Jti:        "123",
		Subject:    "Apa coba",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	// Add custom claim to the token
	jsonToken.Set("data", "pesan rahasia")
	footer := "Ini footer"

	// Encrypt data
	token, err := paseto.NewV1().Encrypt(symmetricKey, jsonToken, footer)

	if err != nil {
		panic(err)
	}

	// Decrypt Token
	var newJsonToken paseto.JSONToken
	var newFooter string

	errs := paseto.NewV1().Decrypt(token, symmetricKey, &newJsonToken, &newFooter)

	if errs != nil {
		panic(errs)
	}

	fmt.Printf("Decrypt: %v", newJsonToken)
	fmt.Printf("Footer: %v", newFooter)
}
