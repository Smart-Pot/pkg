package jwt

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


const (
	_testSecretKey = "default_token_key"
	_testJWTStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJ1c2VyIiwiZGV2aWNlSWRzIjpbImRldmljZSJdLCJhdXRob3JpemF0aW9uIjo5OCwiZXhwIjoxNjEzNjgyOTQyfQ.azdHTnN4O_QcD2zNN2s8icQAQ3QfUYi7diLTQaXRQlo"
)

var (
	_testAuthToken = &AuthToken{
		UserID: "user",
		DeviceIDs: []string{"device"},
		Authorization: 98,
	}
	_testJWT = &JWT {
		Secret: []byte(_testSecretKey),
		Timeout: 2*365*24*time.Hour,
	}
)

func TestMain(m *testing.M) {
	// Make jwt secret empty to use test secret key
	os.Setenv("JWT_SECRET","")
	// set test secret key as secret key
	_defaultJWT.Secret = []byte(_testSecretKey)
	os.Exit(m.Run())
}


func TestTokenize(t *testing.T) {
	s,err :=_testJWT.Tokenize(_testAuthToken)
	assert.Nil(t,err)
	assert.NotEqual(t,"",s)
}


func TestDecode(t *testing.T) {
	a,err := _testJWT.Verify(_testJWTStr)
	assert.Nil(t,err)
	assert.Equal(t,_testAuthToken.UserID,a.UserID)
	assert.Equal(t,_testAuthToken.DeviceIDs[0],a.DeviceIDs[0])
	assert.Equal(t,_testAuthToken.Authorization,a.Authorization)
}

