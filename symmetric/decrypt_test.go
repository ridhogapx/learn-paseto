package main

import (
	"fmt"
	"testing"

	"github.com/o1egl/paseto"
)

func TestDecrypt(t *testing.T) {
	var newJSONToken paseto.JSONToken
	var newFooter string

	token := Encrypt()
	decode := newJSONToken.Get(token)

	err := paseto.NewV1().Decrypt(token, symmetricKey, &newJSONToken, &newFooter)

	/*
	* Case: Error should be nill
	* Or there should no error
	 */
	if err != nil {
		t.Error("There is error", err)
	}

	/*
	* Extract encrypted data
	 */
	fmt.Println("Audience:", newJSONToken.Audience)
	fmt.Println("Expiration:", newJSONToken.Expiration)
	fmt.Println("Issued:", newJSONToken.IssuedAt)
	fmt.Println("Issuer:", newJSONToken.Issuer)
	fmt.Println("JTI:", newJSONToken.Jti)
	fmt.Println("Not Before:", newJSONToken.NotBefore)
	fmt.Println("Subject:", newJSONToken.Subject)
	fmt.Println("Claims:", decode)

}
