package maker

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {

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
