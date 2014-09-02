package payments

import (
	"time"

	"appengine"
	"appengine/datastore"

	"github.com/czertbytes/pocket/pkg/gae"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "payment"
)

type Storage struct {
	AppEngineContext appengine.Context
	storage          *gae.Storage
}

func NewStorage(appEngineContext appengine.Context) *Storage {
	return &Storage{
		AppEngineContext: appEngineContext,
		storage:          gae.NewStorage(appEngineContext, kind),
	}
}

func (self *Storage) Save(payment *t.Payment) error {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	payment.SetCreatedAt(now)
	payment.SetModifiedAt(now)

	id, err := self.storage.Save(payment)
	if err != nil {
		return err
	}

	payment.Id = t.PaymentId(id)
	payment.SetFormattedValues()

	return nil
}

func (self *Storage) FindAllByStatus(status t.PaymentStatus) (t.Payments, error) {
	var payments t.Payments

	query := datastore.NewQuery(kind).
		Filter("status =", status)

	ids, err := self.storage.FindAll(query, &payments)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Payments, 0), nil
	}

	for i, _ := range payments {
		payments[i].Id = t.PaymentId(ids[i])
		payments[i].SetFormattedValues()
	}

	return payments, nil
}

func (self *Storage) FindAllActive() (t.Payments, error) {
	return self.FindAllByStatus(t.PaymentStatusActive)
}

func (self *Storage) FindAllByOverviewId(overviewId t.OverviewId) (t.Payments, error) {
	var payments t.Payments

	query := datastore.NewQuery(kind).
		Filter("overview_id =", overviewId).
		Filter("status =", t.PaymentStatusActive)

	ids, err := self.storage.FindAll(query, &payments)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Payments, 0), nil
	}

	for i, _ := range payments {
		payments[i].Id = t.PaymentId(ids[i])
		payments[i].SetFormattedValues()
	}

	return payments, nil
}

func (self *Storage) FindAllByUserId(userId t.UserId) (t.Payments, error) {
	var payments t.Payments

	query := datastore.NewQuery(kind).
		Filter("user_id=", userId).
		Filter("status =", t.PaymentStatusActive)

	ids, err := self.storage.FindAll(query, &payments)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Payments, 0), nil
	}

	for i, _ := range payments {
		payments[i].Id = t.PaymentId(ids[i])
		payments[i].SetFormattedValues()
	}

	return payments, nil
}

func (self *Storage) Find(id t.PaymentId) (t.Payment, error) {
	var payment t.Payment

	if _, err := self.storage.Find(int64(id), payment); err != nil {
		return t.Payment{}, err
	}

	payment.Id = id
	payment.SetFormattedValues()

	return payment, nil
}

func (self *Storage) FindMulti(ids t.PaymentIds) (t.Payments, error) {
	var payments t.Payments

	if err := self.storage.FindMulti(ids.AsInt64Arr(), &payments); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Payments, 0), nil
	}

	for i, _ := range payments {
		payments[i].Id = t.PaymentId(ids[i])
		payments[i].SetFormattedValues()
	}

	return payments, nil
}

func (self *Storage) Update(payment t.Payment) (t.Payment, error) {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	payment.SetModifiedAt(now)

	if err := self.storage.Update(int64(payment.Id), payment); err != nil {
		return t.Payment{}, err
	}

	return self.Find(payment.Id)
}

func (self *Storage) Delete(id t.PaymentId) error {
	payment, err := self.Find(id)
	if err != nil {
		return err
	}

	return self.storage.Delete(int64(payment.Id))
}
