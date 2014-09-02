package payments

import (
	"appengine"

	t "github.com/czertbytes/pocket/pkg/types"
)

type Payments struct {
	AppEngineContext appengine.Context
	Storage          *Storage
}

func NewPayments(appEngineContext appengine.Context) *Payments {
	return &Payments{
		AppEngineContext: appEngineContext,
		Storage:          NewStorage(appEngineContext),
	}
}

func (self *Payments) Create(payment *t.Payment) error {
	return self.Storage.Save(payment)
}

func (self *Payments) FindAll(status t.PaymentStatus) (t.Payments, error) {
	payments, err := self.Storage.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (self *Payments) FindAllActive() (t.Payments, error) {
	payments, err := self.Storage.FindAllActive()
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (self *Payments) FindAllByOverviewId(overviewId t.OverviewId) (t.Payments, error) {
	return self.Storage.FindAllByOverviewId(overviewId)
}

func (self *Payments) Find(id t.PaymentId) (t.Payment, error) {
	payment, err := self.Storage.Find(id)
	if err != nil {
		return t.Payment{}, err
	}

	return payment, nil
}

func (self *Payments) FindMulti(ids t.PaymentIds) (t.Payments, error) {
	payments, err := self.Storage.FindMulti(ids)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (self *Payments) Update(payment t.Payment) (t.Payment, error) {
	return self.Storage.Update(payment)
}

func (self *Payments) Delete(id t.PaymentId) error {
	return self.Storage.Delete(id)
}
