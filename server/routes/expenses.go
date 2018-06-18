package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oxisto/track-expenses/server/common"
	"github.com/oxisto/track-expenses/server/db"
	"github.com/oxisto/track-expenses/server/model"
)

// GetExpenses handles a GET request to the /expenses endpoint
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := db.FindExpenses()

	common.JsonResponse(w, r, expenses, err)
}

// PostExpense handles a POST request to the /expenses endpoint
func PostExpense(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var (
		expense model.Expense
		err     error
	)
	if err = decoder.Decode(&expense); err != nil {
		common.JsonResponse(w, r, nil, err)
	}

	// create a new ID
	expense.ID = db.NextID()

	defer r.Body.Close()

	// save it into the database
	err = db.InsertExpense(expense)

	common.JsonResponse(w, r, expense, err)
}
