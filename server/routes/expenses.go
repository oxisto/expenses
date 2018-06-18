package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oxisto/track-expenses/server/common"
	"github.com/oxisto/track-expenses/server/model"
)

var expenses = []model.Expense{}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	common.JsonResponse(w, r, expenses, nil)
}

func PostExpense(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var expense model.Expense
	if err := decoder.Decode(&expense); err != nil {
		common.JsonResponse(w, r, nil, err)
	}

	expenses = append(expenses, expense)

	defer r.Body.Close()
}
