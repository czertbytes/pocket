package types

import "sync"

type Balance struct {
	sync.Mutex
	User    User   `json:"user" datastore:"-"`
	Debits  Prices `json:"debits" datastore:"-"`
	Credits Prices `json:"credits" datastore:"-"`
	Price   `json:"balance" datastore:"-"`
}

func NewBalance(user User) *Balance {
	return &Balance{
		User:    user,
		Debits:  make(Prices, 0),
		Credits: make(Prices, 0),
	}
}

func (self *Balance) AddDebit(price Price) {
	self.Lock()
	defer self.Unlock()

	self.Debits = append(self.Debits, price)
	self.Price.Value -= price.Value
}

func (self *Balance) AddCredit(price Price) {
	self.Lock()
	defer self.Unlock()

	self.Credits = append(self.Credits, price)
	self.Price.Value += price.Value
}

type Balances []Balance
