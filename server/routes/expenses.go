package routes

import (
	"net/http"

	"github.com/oxisto/track-expenses/server/common"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses := []int32{1, 2, 4}

	common.JsonResponse(w, r, expenses, nil)
}
