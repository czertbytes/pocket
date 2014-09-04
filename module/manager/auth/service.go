package auth

import (
	"appengine"

	h "github.com/czertbytes/pocket/pkg/http"
	p "github.com/czertbytes/pocket/pkg/proof"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
	notificator      *Notificator
	Proof            *p.Proof
}

func NewService(RequestContext *h.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
		Proof:            p.NewProof(RequestContext.AppEngineContext),
	}
}

func (self *Service) Create(client *t.Client, user *t.User) error {
	if err := self.Proof.Login(client); err != nil {
		return err
	}

	if err := self.notificator.Create(client); err != nil {
		return err
	}

	return nil
}

func (self *Service) Delete(id t.ClientId, user *t.User) error {
	return self.Proof.Logout(self.RequestContext.Client)
}
