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
