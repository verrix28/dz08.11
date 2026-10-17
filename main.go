package main

import "fmt"

type Tovar struct {
	name  string
	price int
}

func main() {
	tovars := []Tovar{
		{"egg", 109},
		{"coffee", 1120},
		{"meat", 1455},
		{"bread", 59},
	}
	var name string
	var price int
	fmt.Scan(&name, &price)

	tovars = append(tovars, Tovar{name, price})

	for _, tovary := range tovars {
		fmt.Printf("%s: %d\n", tovary.name, tovary.price)
	}

	fmt.Println("Товары больше 1000: ")
	for _, tovary := range tovars {
		if tovary.price > 1000 {
			fmt.Printf("%s: %d\n", tovary.name, tovary.price)
		}
	}
}
