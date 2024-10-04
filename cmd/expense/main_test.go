package main

import (
	"github.com/Mensurui/expenseTracker"
	"testing"
)

func TestView(t *testing.T) {
	l := &eventtracker.List{}
	l.View(".expense.json")
}
