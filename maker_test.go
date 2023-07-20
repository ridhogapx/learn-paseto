package maker

import (
	"fmt"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		t.Error(err)
	}

	token := maker.CreateToken("this is the data", time.Minute)
	fmt.Println(token)
}
