package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


const (
	_keyStr = "placeholder_key_"
	_msg = "test_example"
	_encryptedMsg = "-L2uLGvx18N1U3wUAd22nARrOAlxspS3DXwWgw=="
)

var (
	_cip = NewCipher(_keyStr)
	_key = []byte(_keyStr)
	_chEncrypt = make(chan string,1)
)

func TestEncrypt(t *testing.T) {
	// Encrypt message
	_, err := encrypt(_key, _msg)
	assert.Nil(t,err)

	cips := []Cipher {
		_cip,
		ForgotPwdCip,
		VerifyMailCip,
	}
	
	for _, c := range cips {
		_ ,err := c.Encrypt(_msg)
		assert.Nil(t,err)
	}

}

func TestDecrypt(t *testing.T) {
	
	d,err := decrypt(_key,_encryptedMsg)
	assert.Nil(t,err)
	assert.Equal(t,_msg,d)

	d, err = _cip.Decrypt(_encryptedMsg)
	assert.Nil(t,err)
	assert.Equal(t,_msg,d)

}
