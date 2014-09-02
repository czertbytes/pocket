package proof

import (
	"appengine"

	c "github.com/czertbytes/pocket/pkg/clients"
	t "github.com/czertbytes/pocket/pkg/types"
	u "github.com/czertbytes/pocket/pkg/user"
)

type Proof struct {
	AppEngineContext appengine.Context
	Clients          *c.Clients
	User             *u.User
}

func NewProof(appEngineContext appengine.Context) *Proof {
	return &Proof{
		AppEngineContext: appEngineContext,
		Clients:          c.NewClients(appEngineContext),
		Users:            u.NewUsers(appEngineContext),
	}
}

func (self *Proof) Logout(client *t.Client) error {
	newClient, err := self.Clients.FindByClientIdAndToken(client.ClientId, client.Token.Value)
	if err != nil {
		return err
	}

	newClient.RegenerateToken()
	return self.Clients.Update(newClient)
}
