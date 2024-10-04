package main

import (
	"flag"
	"fmt"
	"github.com/Mensurui/expenseTracker"
	"os"
)

var expenseFile = ".expense.json"

func main() {
	if os.Getenv("EXPENSE_FILENAME") != "" {
		expenseFile = os.Getenv("EXPENSE_FILENAME")
	}

	add := flag.String("add", "", "Use to add an expense item.")
	amount := flag.Int("amount", 0, "Amount of the item.")
	price := flag.Int("price", 0, "Price of the item.")
	view := flag.Bool("view", false, "Use to view expense items.")
	del := flag.Bool("del", false, "Use to remove an expense item from your list.")
	id := flag.Int("id", 0, "Combine with '-del' to delete an item from the list.")
	summary := flag.Bool("summary", false, "Use this flag to get your total price.")
	month := flag.Int("month", 0, "Use this to further filter the summary result.")

	flag.Parse()

	l := &eventtracker.List{}
	if err := l.View(expenseFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		if *amount <= 0 || *price <= 0 {
			fmt.Fprintln(os.Stderr, "Amount and price must be greater than zero")
			os.Exit(1)
		}
		l.Add(*add, *amount, *price)
		if err := l.Save(expenseFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *view:
		fmt.Print(l)

	case *del:
		l.Delete(*id)
		if err := l.Save(expenseFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *summary:
		if *month == 0 {
			fmt.Println(l.Summary())
		} else {
			fmt.Println(l.SummaryMonth(*month))
		}

	default:
		fmt.Fprintln(os.Stderr, "No valid operation provided. Use -add or -view.")
	}
}
