package model

import (
	"time"
)

const CurrencyEuro = "EUR"
const DefaultCurrency = CurrencyEuro

// Category is a struct to categorize your expenses
type Category struct {
	Name     string
	Icon     string
	Keywords string
}

// Expense holds all the information about an expense you made
type Expense struct {
	// The amount you spent
	Amount float64
	// The currency
	Currency string
	// Optinal, a comment
	Comment string
	// The time of the expense
	Timestamp time.Time
	// The categories
	Categories []Category
	// Optional, if you track expenses of more than one person in one account
	Person string
}

// NewExpense creates a new expense with default values
func NewExpense() Expense {
	return Expense{
		Timestamp: time.Now(),
		Currency:  CurrencyEuro,
	}
}
