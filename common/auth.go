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
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"time"

	"github.com/oxisto/expenses/db"
	"github.com/sirupsen/logrus"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var (
	log *logrus.Entry
	key *ecdsa.PrivateKey
)

func init() {
	var err error

	log = logrus.WithField("component", "common")

	// check, if a private.pem is the current path
	if key, err = loadKeyFromFile("private.pem"); err != nil {
		log.Warnf("Could not read private key from file: %v", err)

		log.Info("Generating new random private key...")
		key, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	} else {
		log.Info("Loaded private key from file..")
	}

}

func loadKeyFromFile(filename string) (key *ecdsa.PrivateKey, err error) {
	var (
		b     []byte
		block *pem.Block
	)
	if b, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}

	if block, _ = pem.Decode(b); block == nil {
		return nil, errors.New("File contains invalid PEM structure")
	}

	key, err = x509.ParseECPrivateKey(block.Bytes)

	return
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
