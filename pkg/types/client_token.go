package types

import "time"

type ClientToken struct {
	Value        string      `json:"value" datastore:"value"`
	ExpireAt     RFC3339Time `json:"expire_at" datastore:"-" siren:"property"`
	ExpireAtTime time.Time   `json:"-" datastore:"expire_at"`
}

func (self *ClientToken) SetTimes() {
	self.ExpireAt = RFC3339Time(self.ExpireAtTime)
}
