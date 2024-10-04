package eventtracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type expense struct {
	ID     int
	Name   string
	Amount int
	Date   time.Time
	Price  int
}

type List []expense

func (l *List) Add(name string, amount int, price int) {
	ls := *l
	id := len(ls) + 1
	e := expense{
		ID:     id,
		Name:   name,
		Amount: amount,
		Date:   time.Now(),
		Price:  price,
	}
	*l = append(*l, e)
}

func (l *List) UpdateAmount(id int, amount int) error {
	ls := *l
	for i := range ls {
		if ls[i].ID == id {
			ls[i].Amount = amount
			break
		}
	}
	return nil
}

func (l *List) UpdatePrice(id int, price int) error {
	ls := *l
	for i := range ls {
		if ls[i].ID == id {
			ls[i].Price = price
			break
		}
	}

	return nil
}

func (l *List) UpdateName(id int, name string) error {
	ls := *l
	for i := range ls {
		if ls[i].ID == id {
			ls[i].Name = name
			break
		}
	}

	return nil
}

func (l *List) Delete(id int) error {
	ls := *l
	if id <= 0 || id > len(ls) {
		return fmt.Errorf("expense with and id %d deosn't exist", id)
	}

	*l = append(ls[:id-1], ls[id:]...)
	return nil
}

func (l *List) View(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		log.Print("No file to display")
		return nil
	}
	return json.Unmarshal(file, l)
}

func (l *List) Summary() (int, string) {
	count := 0
	formatted := "$"
	for _, v := range *l {
		count += v.Price
	}
	return count, formatted
}

func (l *List) SummaryMonth(month int) (string, int) {
	count := 0
	formatted := "$"
	for _, v := range *l {
		if int(v.Date.Month()) == month {
			count += v.Price
		}
	}
	return formatted, count
}

func (l *List) Save(filepath string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, js, 0644)
}

func (l *List) String() string {
	formatted := ""
	for _, v := range *l {
		formatted += fmt.Sprintf("%d| %s | %s | %d | %d\n", v.ID, v.Name, v.Date.Format("2006-01-02"), v.Amount, v.Price)
	}

	return formatted
}
