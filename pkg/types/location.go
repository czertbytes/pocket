package types

type Location struct {
	Address    string `json:"address,omitempty" datastore:"address,noindex"`
	City       string `json:"city,omitempty" datastore:"city"`
	State      string `json:"state,omitempty" datastore:"state"`
	PostalCode string `json:"postal_code,omitempty" datastore:"postal_code"`
	Country    string `json:"country,omitempty" datastore:"country"`
	LatitudeLongitude
}

func (self *Location) SetFormattedValues() {
	self.LatitudeLongitude.SetFormattedValues()
}
