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
	"context"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/oxisto/expenses/common"
	"github.com/oxisto/expenses/db"
	"github.com/urfave/negroni"
)

const (
	UserContext = "user"
)

func NewRouter() *mux.Router {
	middleware := common.CreateMiddleware()

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/api/expenses").Handler(WithMiddleware(middleware, GetExpenses))
	router.Methods("GET").Path("/api/expenses/{expenseID}").Handler(WithMiddleware(middleware, GetExpense))
	router.Methods("POST").Path("/api/expenses").Handler(WithMiddleware(middleware, PostExpense))
	router.Methods("DELETE").Path("/api/expenses/{expenseID}").Handler(WithMiddleware(middleware, DeleteExpense))
	router.Methods("PUT").Path("/api/expenses/{expenseID}").Handler(WithMiddleware(middleware, PutExpense))
	router.Methods("GET").Path("/api/accounts").Handler(WithMiddleware(middleware, GetAccounts))
	router.Methods("POST").Path("/auth/login").HandlerFunc(Login)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	return router
}

func WithMiddleware(middleware *jwtmiddleware.JWTMiddleware, handlerFunc http.HandlerFunc) *negroni.Negroni {
	return negroni.New(
		negroni.HandlerFunc(middleware.HandlerWithNext),
		negroni.HandlerFunc(HandleFetchUserWithNext),
		negroni.Wrap(handlerFunc),
	)
}

func HandleFetchUserWithNext(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, ok := r.Context().Value("auth").(*jwt.Token)
	if !ok {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	m, ok := claims["user"]
	if !ok {
		return
	}

	var user db.User
	mapstructure.Decode(m, &user)

	request := r.WithContext(context.WithValue(r.Context(), UserContext, user))

	*r = *request
	next(w, r)
}
