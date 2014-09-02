package clients

import (
	"fmt"
	"time"

	"appengine"
	"appengine/datastore"

	"github.com/czertbytes/pocket/pkg/gae"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "client"
)

var (
	ErrClientNotFound error = fmt.Errorf("client: Client not found!")
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

func (self *Storage) FindAllByStatus(status t.ClientStatus) (t.Clients, error) {
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
	return self.FindAllByStatus(t.ClientStatusActive)
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

func (self *Storage) FindByClientIdAndToken(clientId t.ClientClientId, token string) (t.Client, error) {
	var clients t.Clients

	query := datastore.NewQuery(kind).
		Filter("client_id =", clientId).
		Filter("status =", t.ClientStatusActive)

	ids, err := self.storage.FindAll(query, &clients)
	if err != nil {
		return t.Client{}, err
	}

	if len(ids) == 0 {
		return t.Client{}, ErrClientNotFound
	}

	for i, client := range clients {
		if client.Token == token {
			clients[i].Id = t.ClientId(ids[i])
			clients[i].SetFormattedValues()

			return clients[i], nil
		}
	}

	return t.Client{}, ErrClientNotFound
}

func (self *Storage) FindMulti(ids t.ClientIds) (t.Clients, error) {
	var clients t.Clients

	if err := self.storage.FindMulti(ids.AsInt64Arr(), &clients); err != nil {
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
