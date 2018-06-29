package common

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var key *rsa.PrivateKey

func init() {
	// TODO: support loading the key from a file or Kubernetes secret
	key, _ = rsa.GenerateKey(rand.Reader, 4048)
}

// IssueToken issues a JWT token for use of the API
func IssueToken(ID string) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodRS512, &APIClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
		ID,
	})

	token, err = claims.SignedString(key)
	return
}

type APIClaims struct {
	*jwt.StandardClaims
	UserID string
}
