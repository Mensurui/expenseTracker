package eventtracker_test

import (
	"github.com/Mensurui/expenseTracker"
	"testing"
)

func TestAdd(t *testing.T) {
	l := eventtracker.List{}
	expenseName := "Expense One"
	amount := 2
	price := 200
	l.Add(expenseName, amount, price)
	if l[0].Name != expenseName || l[0].Amount != amount || l[0].Price != price {
		t.Errorf("Expeceted: %s | %d  | %d \n Got: %s | %d | %d", expenseName, amount, price, l[0].Name, l[0].Amount, l[0].Price)
	}
}

func TestView(t *testing.T) {
	l := eventtracker.List{}
	expenseName := "Expense One"
	amount := 2
	price := 200
	l.Add(expenseName, amount, price)
	l.View()
}
