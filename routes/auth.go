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

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oxisto/expenses/common"
	"github.com/oxisto/expenses/db"

	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

// Login handles a login request
func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var (
		user    db.User
		request LoginRequest
		token   string
		err     error
	)
	if err = decoder.Decode(&request); err != nil {
		common.JsonResponse(w, r, nil, err)
	}

	// try to find the user
	err = db.Find(bson.M{"username": request.Username}, &user)
	if err == db.ErrNotFound {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		common.JsonResponse(w, r, nil, err)
		return
	}

	// authenticate
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		common.JsonResponse(w, r, nil, err)
		return
	}

	// make sure, to remove the hash before issuing the ticket
	user.PasswordHash = ""

	token, err = common.IssueToken(user)

	resp := TokenResponse{Token: token}

	// redirect to main dashboard page
	//w.Header().Add("Location", "/#?token="+token)
	w.Header().Add("Set-Cookie", "token="+token+"; Path=/")
	//w.WriteHeader(http.StatusFound)

	common.JsonResponse(w, r, resp, err)

	defer r.Body.Close()

}
