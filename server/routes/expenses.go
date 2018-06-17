package routes

import (
	"net/http"

	"github.com/oxisto/track-expenses/server/common"
	"github.com/oxisto/track-expenses/server/model"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expense := model.NewExpense()
	expense.Amount = 1.0

	expenses := []model.Expense{
		expense,
	}

	common.JsonResponse(w, r, expenses, nil)
}
