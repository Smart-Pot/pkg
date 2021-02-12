package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


const (
	_msg = "test_example"
	_encryptedMsg = "-L2uLGvx18N1U3wUAd22nARrOAlxspS3DXwWgw=="
)

var (
	_key = []byte("placeholder_key_")
	_chEncrypt = make(chan string,1)
)

func TestEncrypt(t *testing.T) {
	// Encrypt message
	_, err := encrypt(_key, _msg)
	assert.Nil(t,err)
}

func TestDecrypt(t *testing.T) {
	d,err := decrypt(_key,_encryptedMsg)
	assert.Nil(t,err)
	assert.Equal(t,_msg,d)
}
