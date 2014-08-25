package clients

import (
	"time"

	"appengine"
	"appengine/datastore"

	"github.com/czertbytes/pocket/pkg/gae"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "client"
)

type Storage struct {
	AppEngineContext appengine.Context
	storage          *gae.Storage
}

func NewStorage(appEngineContext appengine.Context) *Storage {
	return &Storage{
		AppEngineContext: appEngineContext,
		storage:          NewStorage(appEngineContext, kind),
	}
}

func (self *Storage) Save(client *t.Client) error {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	client.SetCreatedAt(now)
	client.SetModifiedAt(now)

	id, err := self.storage.Save(client)
	if err != nil {
		return err
	}

	client.Id = t.ClientId(id)
	client.SetFormattedValues()

	return nil
}

func (self *Storage) FindAllByStatus(status t.clientstatus) (t.Clients, error) {
	var clients t.Clients

	query := datastore.NewQuery(kind).
		Filter("status =", status)

	ids, err := self.storage.FindAll(query, &clients)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Clients, 0), nil
	}

	for i, _ := range clients {
		clients[i].Id = t.ClientId(ids[i])
		clients[i].SetFormattedValues()
	}

	return clients, nil
}

func (self *Storage) FindAllActive() (t.Clients, error) {
	return self.FindAllByStatus(t.clientstatusActive)
}

func (self *Storage) Find(id t.ClientId) (t.Client, error) {
	var client t.Client

	if _, err := self.storage.Find(int64(id), client); err != nil {
		return t.Client{}, err
	}

	client.Id = id
	client.SetFormattedValues()

	return client, nil
}

func (self *Storage) FindMulti(ids t.ClientIds) (t.Clients, error) {
	var clients t.Clients

	if err := self.storage.FindMulti([]int64(ids)); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Clients, 0), nil
	}

	for i, _ := range clients {
		clients[i].Id = t.ClientId(ids[i])
		clients[i].SetFormattedValues()
	}

	return clients, nil
}

func (self *Storage) Update(client t.Client) (t.Client, error) {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	client.SetModifiedAt(now)

	if err := self.storage.Update(int64(client.Id), client); err != nil {
		return t.Client{}, err
	}

	return self.Find(client.Id)
}

func (self *Storage) Delete(id t.ClientId) error {
	client, err := self.Find(id)
	if err != nil {
		return err
	}

	return self.storage.Delete(int64(client.Id))
}
