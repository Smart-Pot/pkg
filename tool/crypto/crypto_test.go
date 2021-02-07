package crypto

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	var text = "deneme"
	key := []byte("example_key_xxxx")
	e, err := encrypt(key, text)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	d, err := decrypt(key, e)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println("e", e, "d", d)
	if d != text {
		t.Error("Want", text, "Got", d)
		t.FailNow()
	}
}
