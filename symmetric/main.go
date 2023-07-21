package main

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

var symmetricKey = []byte("ROSE IS RED, VIOLET IS BLUE, I'M") // Must be 32 bytes

func Encrypt() string {
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
	token, err := paseto.NewV2().Encrypt(symmetricKey, jsonToken, footer)

	if err != nil {
		panic(err)
	}

	return token
}

func main() {
	token := Encrypt()

	var j paseto.JSONToken
	var footer string

	err := paseto.NewV2().Decrypt(token, symmetricKey, &j, &footer)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(j)

	marshall, _ := j.MarshalJSON()

	fmt.Println(string(marshall))

}
