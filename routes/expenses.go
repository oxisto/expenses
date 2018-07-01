package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oxisto/expenses/common"
	"github.com/oxisto/expenses/db"
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

	user := r.Context().Value("user").(db.User)

	expense, err := db.FindExpense(user, expenseID)

	if err == db.ErrNotFound {
		common.JsonResponse(w, r, nil, nil)
	} else {
		common.JsonResponse(w, r, expense, err)
	}
}

// GetExpenses handles a GET request to the /expenses endpoint
func GetExpenses(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(db.User)

	expenses, err := db.FindExpenses(user, db.ExpensesCollectionName)

	common.JsonResponse(w, r, expenses, err)
}

// PostExpense handles a POST request to the /expenses endpoint
func PostExpense(w http.ResponseWriter, r *http.Request) {
	var (
		expense db.Expense
		err     error
	)

	user := r.Context().Value("user").(db.User)

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&expense); err != nil {
		common.JsonResponse(w, r, nil, err)
	}

	// create a new ID
	expense.ID = db.NextID()

	// TODO: support access to other accounts via delegation (https://github.com/oxisto/expenses/issues/4)
	if expense.AccountID != user.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	defer r.Body.Close()

	// save it into the database
	err = db.Insert(expense)

	common.JsonResponse(w, r, expense, err)
}

// PutExpense stores an expense at a given id
func PutExpense(w http.ResponseWriter, r *http.Request) {
	var (
		expense   db.Expense
		expenseID string
		ok        bool
		err       error
	)

	user := r.Context().Value("user").(db.User)

	decoder := json.NewDecoder(r.Body)

	if err = decoder.Decode(&expense); err != nil {
		common.JsonResponse(w, r, nil, err)
	}

	if expenseID, ok = mux.Vars(r)["expenseID"]; !ok {
		common.JsonResponse(w, r, nil, nil)
		return
	}

	// make sure, IDs match
	expense.ID = expenseID

	// TODO: support access to other accounts via delegation (https://github.com/oxisto/expenses/issues/4)
	if expense.AccountID != user.ID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	defer r.Body.Close()

	// save it into the database
	err = db.Upsert(expense)

	common.JsonResponse(w, r, expense, err)
}
