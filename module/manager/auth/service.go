package auth

import (
	"appengine"

	h "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
	notificator      *Notificator
}

func NewService(RequestContext *h.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
	}
}

func (self *Service) Create(client *t.Client, user *t.User) error {

	if err := self.notificator.Create(client); err != nil {
		return err
	}

	return nil
}

func (self *Service) Delete(id t.ClientId, user *t.User) error {
	return nil
}
