package crypto

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	var text = "682efc3f-6dac-46eb-804b-86970550bf31"
	key := privateKey
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

func TestDecrypt(t *testing.T) {
	m := "QxKUJqrQZX1mQIEZ"
	res, err := Decrypt(m)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println("Res", res)
}
