package main

type Account struct {
	Number  string
	Balance float64
	Status  string
	History []string
}
type Client struct {
	ID       int
	Name     string
	Accounts []*Account
}
type Transaction interface {
	Execute()
}
type Deposit struct {
	Acc *Account
	Sum float64
}

func (d *Deposit) Execute() {
	if d.Acc.Status == "active" {
		d.Acc.Balance += d.Sum
		d.Acc.History = append(d.Acc.History, "Пополнение")
	}
}

type Withdraw struct {
	Acc *Account
	Sum float64
}

func (w *Withdraw) Execute() {
	if w.Acc.Status == "active" && w.Acc.Balance >= w.Sum {
		w.Acc.Balance -= w.Sum
		w.Acc.History = append(w.Acc.History, "Снятие")
	}
}

type Transfer struct {
	From *Account
	To   *Account
	Sum  float64
}

func (t *Transfer) Execute() {
	if t.From.Status == "active" && t.To.Status == "active" && t.From.Balance >= t.Sum {
		t.From.Balance -= t.Sum
		t.To.Balance += t.Sum
		t.From.History = append(t.From.History, "Перевод отправлен")
		t.To.History = append(t.To.History, "Перевод получен")
	}
}
