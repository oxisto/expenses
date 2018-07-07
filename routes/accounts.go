package routes

import (
	"net/http"

	"github.com/oxisto/expenses/common"
	"github.com/oxisto/expenses/db"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(db.User)

	accounts := db.FindDelegatedAccounts(user)

	common.JsonResponse(w, r, accounts, nil)
}
