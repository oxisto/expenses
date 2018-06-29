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
	Username string
	Password string
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

	token, err = common.IssueToken(user.ID)

	// redirect to main dashboard page
	w.Header().Add("Location", "/#?token="+token)
	w.Header().Add("Set-Cookie", "token="+token+"; Path=/")
	w.WriteHeader(http.StatusFound)

	defer r.Body.Close()

}
