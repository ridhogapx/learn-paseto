package main

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

func main() {
	symmetric := []byte("ROSE IS RED, VIOLET IS BLUE")
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "NPC",
		Issuer:     "Developer",
		Jti:        "123",
		Subject:    "Testing",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}

	jsonToken.Set("data", "example of signed message")
	footer := "some footer"

	token, err := paseto.NewV2().Encrypt(symmetric, jsonToken, footer)

	if err != nil {
		panic(err)
	}

	fmt.Println(token)
}
