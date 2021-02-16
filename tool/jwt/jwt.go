package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)



var (
	// ErrTokenExpired :
	ErrTokenExpired = errors.New("token expired")
	// ErrParseClaims :
	ErrParseClaims  = errors.New("Could not parse claims")
	// ErrSignToken :
	ErrSignToken  = errors.New("Failed to sign token")

)


var (
	_defaultJWT *JWT = getDefaultJWT()
)

func getSecretKey() []byte {
	r := os.Getenv("JWT_SECRET_KEY")
	if len(r) == 0 {
		fmt.Println("2")
		return []byte("default_token_key") // Default JWT Secret 
	}
	return []byte(r)
}

func getDefaultJWT() *JWT {
	return &JWT{
		Secret: getSecretKey(),
		Timeout: 48*time.Hour,
		Issuer: "Auth",
	}
}


// AuthToken :
type AuthToken struct {
	UserID string
	DeviceIDs []string
	Authorization int
}

type authTokenClaims struct {
	UserID string `json:"userId"`
	DeviceIDs []string `json:"deviceIds"`
	Authorization int `json:"authorization"`
	*jwt.StandardClaims
}

// Tokenize :
func Tokenize(t *AuthToken) (string,error) {
	return _defaultJWT.Tokenize(t)
} 

// Verify :
func Verify(token string) (*AuthToken,error) {
	return _defaultJWT.Verify(token)
}

// JWT :
type JWT struct {
	Timeout time.Duration
	Secret []byte
	Issuer string 
}

// Tokenize :
func (j *JWT) Tokenize(t *AuthToken) (string,error) {
	claims := authTokenClaims{
		UserID: t.UserID,
		DeviceIDs: t.DeviceIDs,
		Authorization: t.Authorization,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(j.Timeout).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	st,err := token.SignedString(j.Secret)
	if err != nil {
		return "", ErrSignToken
	}
	return st,nil
}

// Verify :
func (j *JWT) Verify(token string) (*AuthToken,error) {
	c,err := jwt.ParseWithClaims(
				token,
				&authTokenClaims{},
				func(token *jwt.Token) (interface{}, error) {
					return j.Secret, nil
				},
			)
	if err != nil {
		return nil,err
	}

	ac , ok := c.Claims.(*authTokenClaims)

	if !ok {
		return nil, ErrParseClaims
	}
	fmt.Println("ac ",ac.ExpiresAt)
	if ac.ExpiresAt < time.Now().Local().Unix() {
		return nil, ErrTokenExpired
	}

	return &AuthToken{
		UserID: ac.UserID,
		DeviceIDs: ac.DeviceIDs,
		Authorization: ac.Authorization,
	},nil
}