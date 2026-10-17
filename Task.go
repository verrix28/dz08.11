package main

import (
	"fmt"
	"sort"
)

type Tovary struct {
	ID       int
	Name     string
	Price    int
	Quantity int
}

func main() {
	tovarys = []Tovary{
		{1, "bread", 59, 198},
		{2, "egg", 159, 18},
		{3, "coffee", 1159, 734},
		{4, "milk", 89, 189},
	}
	PrintAll()
	var Name string
	var Price int
	var Quantity int
	fmt.Scan(&Name, &Price, &Quantity)

	newId := 1
	for _, a := range tovarys {
		if a.ID >= newId {
			newId = a.ID + 1
		}
	}

	tovarys = append(tovarys, Tovary{newId, Name, Price, Quantity})
	fmt.Println("Новый товар добавлен:", Name)
	PrintAll()

	if a := poiskId(2); a != nil {
		fmt.Println("Товар найдет:", a.Name, a.Price)
	}
	change(4, 100)
	fmt.Println("Отсортированный список")
	sortedAll()
	PrintAll()
	maxTovar()
	AllPrice()
	delete(1)
	PrintAll()
}

var tovarys []Tovary

func poiskId(id int) *Tovary {
	for i := range tovarys {
		if tovarys[i].ID == id {
			return &tovarys[i]
		}
	}
	return nil
}
func delete(id int) {
	for i, a := range tovarys {
		if a.ID == id {
			tovarys = append(tovarys[:i], tovarys[i+1:]...)
			fmt.Println("Товар с ID удален: ", id)
			return
		}
	}
}
func change(id int, NewPrice int) {
	a := poiskId(id)
	if a != nil {
		a.Price = NewPrice
		fmt.Println("Цена товара изменена:", a.Name, NewPrice)
	}
}
func PrintAll() {
	for _, a := range tovarys {
		fmt.Println(a.ID, a.Name, a.Price, a.Quantity)
	}
}
func sortedAll() {
	sort.Slice(tovarys, func(i, j int) bool {
		return tovarys[i].Price < tovarys[j].Price
	})
}
func maxTovar() {
	if len(tovarys) == 0 {
		return
	}
	max := tovarys[0]
	for _, a := range tovarys {
		if a.Price > max.Price {
			max = a
		}
	}
	fmt.Println("Самый дорогой товар: ", max.Name, max.Price)
}
func AllPrice() {
	summa := 0
	for _, a := range tovarys {
		summa += a.Price * a.Quantity
	}
	fmt.Println("Общая сумма", summa)

}
