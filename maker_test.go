package maker

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/o1egl/paseto"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		t.Error(err)
	}

	token := maker.CreateToken("this is the data", time.Minute)

	fmt.Println(token)
}

func TestPasetoV2(t *testing.T) {
	b, _ := hex.DecodeString("b4cbfb43df4ce210727d953e4a713307fa19bb7d9f85041438d9e11b942a37741eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	privateKey := ed25519.PrivateKey(b)

	pub, _ := hex.DecodeString("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2")
	publicKey := ed25519.PublicKey(pub)

	jsonToken := paseto.JSONToken{
		Expiration: time.Now().Add(25 * time.Hour),
	}

	jsonToken.Set("data", "there's kids in my basement")
	footer := "Schizo"

	token, _ := paseto.NewV2().Sign(privateKey, jsonToken, footer)

	fmt.Println("Token:", token)

	// Verifying data
	var newJsonToken paseto.JSONToken
	var newFooter string

	err := paseto.NewV2().Verify(token, publicKey, &newJsonToken, &newFooter)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data:", newJsonToken)
	fmt.Println("Footer:", newFooter)

}
