package routes

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {

	/*middleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(model.JwtSecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})*/

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/expenses", GetExpenses)

	return router
}
