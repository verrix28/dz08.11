package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Магазин:")
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

	fmt.Println("Доставка:")
	courier := new(Courier)
	courier.km = 15.0
	pickup := new(Pickup)
	pickup.price = 200.0
	post := new(Post)
	post.kg = 2.5
	o1 := Order{101, "ул. Ленина 1", 5.0, 1500, courier}
	o2 := Order{102, "ул. Пушкина 2", 2.5, 800, pickup}
	o3 := Order{103, "ул. Гоголя 3", 2.5, 1200, post}
	PrintOrderInfo(o1)
	PrintOrderInfo(o2)
	PrintOrderInfo(o3)

	fmt.Println("Банк:")
	acc1 := &Account{Number: "111", Balance: 1000, Status: "active"}
	acc2 := &Account{Number: "222", Balance: 500, Status: "active"}
	client := Client{ID: 1, Name: "Иван", Accounts: []*Account{acc1, acc2}}
	fmt.Println(client.Name)
	dep := &Deposit{Acc: acc1, Sum: 500}
	dep.Execute()
	wit := &Withdraw{Acc: acc1, Sum: 2000}
	wit.Execute()
	wit2 := &Withdraw{Acc: acc1, Sum: 200}
	wit2.Execute()
	trans := &Transfer{From: acc1, To: acc2, Sum: 300}
	trans.Execute()
	acc1.Status = "blocked"
	wit3 := &Withdraw{Acc: acc1, Sum: 100}
	wit3.Execute()
	fmt.Println("Баланс 111:", acc1.Balance, "Баланс 222:", acc2.Balance)

	fmt.Println("Игра:")
	w := &Warrior{Name: "Воин", HP: 100, Level: 1, Armor: 5}
	m := &Mage{Name: "Маг", HP: 80, Level: 1, Mana: 20}
	a := &Archer{Name: "Лучник", HP: 90, Level: 1, Arrows: 5}
	for w.IsAlive() && m.IsAlive() {
		m.TakeDamage(w.Attack())
		if !m.IsAlive() {
			break
		}
		w.TakeDamage(m.Attack())
	}
	if w.IsAlive() {
		fmt.Println("Победил:", w.Name)
	} else {
		fmt.Println("Победил:", m.Name)
	}
	w.HP = 100
	fmt.Println("Воин вылечен, HP:", w.HP)
	w.Level++
	m.Level++

	chars := []Fighter{w, m, a}
	sort.Slice(chars, func(i, j int) bool {
		var hpI, hpJ float64
		var lvlI, lvlJ int
		if c, ok := chars[i].(*Warrior); ok {
			hpI, lvlI = c.HP, c.Level
		}
		if c, ok := chars[i].(*Mage); ok {
			hpI, lvlI = c.HP, c.Level
		}
		if c, ok := chars[i].(*Archer); ok {
			hpI, lvlI = c.HP, c.Level
		}

		if c, ok := chars[j].(*Warrior); ok {
			hpJ, lvlJ = c.HP, c.Level
		}
		if c, ok := chars[j].(*Mage); ok {
			hpJ, lvlJ = c.HP, c.Level
		}
		if c, ok := chars[j].(*Archer); ok {
			hpJ, lvlJ = c.HP, c.Level
		}
		if lvlI != lvlJ {
			return lvlI > lvlJ
		}
		return hpI > hpJ
	})
}
