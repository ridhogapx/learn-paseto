package main

import (
	"testing"

	"github.com/o1egl/paseto"
)

func TestDecrypt(t *testing.T) {
	var newJSONToken paseto.JSONToken
	var newFooter string

	token := Encrypt()

	err := paseto.NewV1().Decrypt(token, symmetricKey, &newJSONToken, &newFooter)

	/*
	* Case: Error should be nill
	* Or there should no error
	 */
	if err != nil {
		t.Error("There is error", err)
	}

}
