package main

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

func main() {
	symmetricKey := []byte("YELLOW SUBMARINE, BLACK WIZARDRY") // Must be 32 bytes
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "test",
		Issuer:     "test_service",
		Jti:        "123",
		Subject:    "test_subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	// Add custom claim    to the token
	jsonToken.Set("data", "this is a signed message")
	footer := "some footer"

	// Encrypt data
	token, err := paseto.NewV1().Encrypt(symmetricKey, jsonToken, footer)

	if err != nil {
		panic(err)
	}

	fmt.Println(token)
}
