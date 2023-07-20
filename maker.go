package learn

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key size")
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(data string, duration time.Duration) string {
	token, err := maker.paseto.Encrypt(maker.symmetricKey, paseto.JSONToken{
		Audience: "Users",
		Issuer:   "RageNeko26",
		Jti:      "QWERTY",
		Subject:  "This is subject",
	}, "this is footer")

	if err != nil {
		fmt.Println(err)
	}

	return token
}
