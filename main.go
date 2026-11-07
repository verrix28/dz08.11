package main

import (
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Stock    int
}
type User struct {
	ID      int
	Name    string
	Balance float64
}
type Payment interface {
	Pay(amount float64)
}
type Card struct {
}

func (c *Card) Pay(amount float64) {
	fmt.Println("Оплата картой:", amount)
}

type Cash struct {
}

func (c *Cash) Pay(amount float64) {
	fmt.Println("Оплата наличными:", amount)
}

type Crypto struct {
}

func (c *Crypto) Pay(amount float64) {
	fmt.Println("Оплата криптой:", amount)
}
func main() {
	catalog := map[int]Product{
		1: {1, "Ноутбук", "Электроника", 80000, 5},
		2: {2, "Мышь", "Электроника", 1500, 20},
		3: {3, "Книга", "Книги", 700, 10},
	}
	catalog[4] = Product{4, "Клавиатура", "Электроника", 3000, 15}
	p2 := catalog[2]
	p2.Price = 1200
	catalog[2] = p2
	fmt.Println("Товар по ID 1:", catalog[1].Name)
	for _, p := range catalog {
		if p.Name == "Книга" {
			fmt.Println("Найдено:", p.Name)
		}
	}
	delete(catalog, 4)

	totalStock := 0.0
	for _, p := range catalog {
		totalStock += p.Price * float64(p.Stock)
	}
	fmt.Println("Стоимость склада:", totalStock)

	user := User{1, "Иван", 100000}
	cart := make(map[int]int)
	cart[1] = 1
	cart[2] = 3
	cart[2] = 2
	delete(cart, 1)

	var orderTotal float64
	canBuy := true
	for id, qty := range cart {
		if catalog[id].Stock < qty {
			canBuy = false
		}
		orderTotal += catalog[id].Price * float64(qty)
	}
	if user.Balance < orderTotal {
		canBuy = false
	}

	if canBuy {
		pay := new(Card)
		pay.Pay(orderTotal)
		for id, qty := range cart {
			p := catalog[id]
			p.Stock -= qty
			catalog[id] = p
		}
		user.Balance -= orderTotal
		fmt.Println("Остаток:", user.Balance)
	}
}
