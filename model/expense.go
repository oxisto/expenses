package model

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

// Expense holds all the information about an expense you made
type Expense struct {
	// Expense id
	ID string `json:"id" bson:"_id"`
	// Account id
	AccountID int32 `json:"accountID" bson:"accountID"`
	// The amount you spent
	Amount float64 `json:"amount"`
	// The currency
	Currency string `json:"currency"`
	// Optinal, a comment
	Comment string `json:"comment"`
	// The time of the expense
	Timestamp time.Time `json:"timestamp"`
	// The categories
	Categories []Category `json:"categories"`
}

// NewExpense creates a new expense with default values
func NewExpense() Expense {
	return Expense{
		Timestamp: time.Now(),
		Currency:  CurrencyEuro,
	}
}
