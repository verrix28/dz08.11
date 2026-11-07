package main

import "fmt"

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
