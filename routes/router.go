package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	/*middleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(model.JwtSecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})*/

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/api/expenses").HandlerFunc(GetExpenses)
	router.Methods("GET").Path("/api/expenses/{expenseID}").HandlerFunc(GetExpense)
	router.Methods("POST").Path("/api/expenses").HandlerFunc(PostExpense)
	router.Methods("PUT").Path("/api/expenses/{expenseID}").HandlerFunc(PutExpense)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist/frontend")))

	return router
}
