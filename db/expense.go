package db

import (
	"time"
)

const CurrencyEuro = "EUR"
const DefaultCurrency = CurrencyEuro

// Category is a struct to categorize your expenses
type Category struct {
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Keywords string `json:"keywords"`
}

const ExpensesCollectionName = "expenses"

// Expense holds all the information about an expense you made
type Expense struct {
	// Expense id
	ID string `json:"id" bson:"_id"`
	// Account id
	AccountID string `json:"accountID" bson:"accountID"`
	// The amount you spent
	Amount float64 `json:"amount"`
	// The currency
	Currency string `json:"currency"`
	// Optional, a comment
	Comment string `json:"comment"`
	// The time of the expense
	Timestamp time.Time `json:"timestamp"`
	// The categories
	Categories []Category `json:"categories"`
}

func (e Expense) Collection() string {
	return ExpensesCollectionName
}

func (e Expense) Identifier() string {
	return e.ID
}

// DBObject represents objects that can be stored in a database
type DBObject interface {
	// The name of the db collection
	Collection() string

	// The id of the object
	Identifier() string
}

// NewExpense creates a new expense with default values
func NewExpense() Expense {
	return Expense{
		ID:        NextID(),
		Timestamp: time.Now(),
		Currency:  CurrencyEuro,
	}
}
