package common

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// IssueToken issues a JWT token for use of the API
func IssueToken(ID string) (token string, err error) {
	key := []byte(JwtSecretKey)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &APIClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
		ID,
	})

	token, err = claims.SignedString(key)
	return
}

// TODO: generate this somehow or retrieve it from something
const (
	JwtSecretKey = "changeme"
)

type APIClaims struct {
	*jwt.StandardClaims
	UserID string
}
