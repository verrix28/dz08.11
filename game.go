package main

type Fighter interface {
	Attack() float64
	TakeDamage(dmg float64)
	IsAlive() bool
}
type Warrior struct {
	Name  string
	HP    float64
	Level int
	Armor float64
}

func (w *Warrior) Attack() float64 {
	return 20.0
}
func (w *Warrior) TakeDamage(dmg float64) {
	w.HP -= (dmg - w.Armor)
}
func (w *Warrior) IsAlive() bool {
	return w.HP > 0
}

type Mage struct {
	Name  string
	HP    float64
	Level int
	Mana  float64
}

func (m *Mage) Attack() float64 {
	if m.Mana >= 10 {
		m.Mana -= 10
		return 30.0
	}
	return 5.0
}
func (m *Mage) TakeDamage(dmg float64) {
	m.HP -= dmg
}
func (m *Mage) IsAlive() bool {
	return m.HP > 0
}

type Archer struct {
	Name   string
	HP     float64
	Level  int
	Arrows int
}

func (a *Archer) Attack() float64 {
	if a.Arrows > 0 {
		a.Arrows--
		return 25.0
	}
	return 3.0
}
func (a *Archer) TakeDamage(dmg float64) {
	a.HP -= dmg
}
func (a *Archer) IsAlive() bool {
	return a.HP > 0
}
