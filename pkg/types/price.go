package types

import "fmt"

type Price struct {
	Value          int64  `json:"value" datastore:"value"`
	ValueFormatted string `json:"value_formatted" datastore:"-"`
	Currency       string `json:"currency" datastore:"currency"`
}

func (self *Price) SetValueFormatted() {
	self.ValueFormatted = fmt.Sprintf("%d %s", self.Value, self.Currency)
}

type Prices []Price
