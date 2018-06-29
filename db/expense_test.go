package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewExpense(t *testing.T) {
	expense := NewExpense()
	expense.Amount = 2.0
	expense.Person = "Christian"

	// assert default currency
	assert.Equal(t, expense.Currency, DefaultCurrency, "A new expense should have the default currency")

	fmt.Printf("%s spent %f %s", expense.Person, expense.Amount, expense.Currency)
}
