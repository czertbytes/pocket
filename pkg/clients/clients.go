package clients

import (
	"appengine"

	t "github.com/czertbytes/pocket/pkg/types"
)

type Clients struct {
	AppEngineContext appengine.Context
	Storage          *Storage
}

func NewClients(appEngineContext appengine.Context) *Clients {
	return &Clients{
		AppEngineContext: appEngineContext,
		Storage:          NewStorage(appEngineContext),
	}
}

func (self *Clients) Create(client *t.Client) error {
	return self.Storage.Save(client)
}

func (self *Clients) FindAll(status t.ClientStatus) (t.Clients, error) {
	clients, err := self.Storage.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}

	// set users

	return clients, nil
}

func (self *Clients) FindAllActive() (t.Clients, error) {
	clients, err := self.Storage.FindAllActive()
	if err != nil {
		return nil, err
	}

	// set users

	return clients, nil
}

func (self *Clients) Find(id t.ClientId) (t.Client, error) {
	client, err := self.Storage.Find(id)
	if err != nil {
		return t.Client{}, err
	}

	// set user

	return client, nil
}

func (self *Clients) FindByClientIdAndToken(clientId t.ClientClientId, token string) (t.Client, error) {
	return self.Storage.FindByClientIdAndToken(clientId, token)
}

func (self *Clients) FindMulti(ids t.ClientIds) (t.Clients, error) {
	clients, err := self.Storage.FindMulti(ids)
	if err != nil {
		return nil, err
	}

	// set users

	return clients, nil
}

func (self *Clients) Update(client t.Client) (t.Client, error) {
	return self.Storage.Update(client)
}

func (self *Clients) Delete(id t.ClientId) error {
	return self.Storage.Delete(id)
}
