package main

import (
	"fmt"
	"reflect"
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

	/*
	* Extract encrypted data
	 */

	a := reflect.ValueOf(newJSONToken)

	for i := 0; i < a.NumField(); i++ {
		fmt.Printf("Values: %v", a.Field(i))
	}

}
