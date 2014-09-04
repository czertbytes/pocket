package types

type Balance struct {
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
	self.Debits = append(self.Debits, price)
	self.Price.Value -= price.Value
}

func (self *Balance) AddCredit(price Price) {
	self.Credits = append(self.Credits, price)
	self.Price.Value += price.Value
}

type Balances []Balance
