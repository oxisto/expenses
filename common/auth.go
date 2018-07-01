/*
Copyright 2018 Christian Banse

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"

	"github.com/oxisto/expenses/db"
)

var key *ecdsa.PrivateKey

func init() {
	// TODO: support loading the key from a file or Kubernetes secret
	key, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func CreateMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return &key.PublicKey, nil
		},
		UserProperty:  "auth",
		SigningMethod: jwt.SigningMethodES256,
	})
}

// IssueToken issues a JWT token for use of the API
func IssueToken(user db.User) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, &APIClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
		user,
	})

	token, err = claims.SignedString(key)
	return
}

type APIClaims struct {
	*jwt.StandardClaims
	User db.User `json:"user"`
}
