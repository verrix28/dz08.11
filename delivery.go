package main

import "fmt"

type Delivery interface {
	Cost() float64
	Info() string
}
type Courier struct {
	km float64
}

func (c *Courier) Cost() float64 {
	return c.km * 50
}
func (c *Courier) Info() string {
	return "Курьер"
}

type Pickup struct {
	price float64
}

func (p *Pickup) Cost() float64 {
	return p.price
}
func (p *Pickup) Info() string {
	return "Пункт выдачи"
}

type Post struct {
	kg float64
}

func (p *Post) Cost() float64 {
	return p.kg * 300
}
func (p *Post) Info() string {
	return "Почта"
}

type Order struct {
	Number       int
	Address      string
	Weight       float64
	ProductsCost float64
	Delivery     Delivery
}

func PrintOrderInfo(o Order) {
	total := o.ProductsCost + o.Delivery.Cost()
	fmt.Println(o.Delivery.Info(), "Стоимость доставки:", o.Delivery.Cost(), "Итого:", total)
}
