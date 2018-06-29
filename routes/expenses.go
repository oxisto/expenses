package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oxisto/expenses/common"
	"github.com/oxisto/expenses/db"
	"github.com/oxisto/expenses/model"
)

// GetExpense retrieves a single expense
func GetExpense(w http.ResponseWriter, r *http.Request) {
	var (
		expenseID string
		ok        bool
	)

	if expenseID, ok = mux.Vars(r)["expenseID"]; !ok {
		common.JsonResponse(w, r, nil, nil)
		return
	}

	var expense model.Expense
	err := db.Find(expenseID, &expense)

	if err == db.ErrNotFound {
		common.JsonResponse(w, r, nil, nil)
	} else {
		common.JsonResponse(w, r, expense, err)
	}
}

// GetExpenses handles a GET request to the /expenses endpoint
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := db.FindExpenses(model.ExpensesCollectionName)

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
	err = db.Insert(expense)

	common.JsonResponse(w, r, expense, err)
}

// PutExpense stores an expense at a given id
func PutExpense(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var (
		expense   model.Expense
		expenseID string
		ok        bool
		err       error
	)

	if err = decoder.Decode(&expense); err != nil {
		common.JsonResponse(w, r, nil, err)
	}

	if expenseID, ok = mux.Vars(r)["expenseID"]; !ok {
		common.JsonResponse(w, r, nil, nil)
		return
	}

	// make sure, IDs match
	expense.ID = expenseID

	defer r.Body.Close()

	// save it into the database
	err = db.Upsert(expense)

	common.JsonResponse(w, r, expense, err)
}
