package overview

import (
	"appengine"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Notificator struct {
	AppEngineContext appengine.Context
	RequestContext   *shttp.RequestContext
}

func NewNotificator(RequestContext *shttp.RequestContext) *Notificator {
	return &Notificator{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
	}
}

func (self *Notificator) Create(overview *t.Overview) error {
	return nil
}

func (self *Notificator) Update(overview *t.Overview) error {
	return nil
}

func (self *Notificator) CreatePayment(payment *t.Payment) error {
	return nil
}

func (self *Notificator) CreateParticipant(participant *t.User) error {
	return nil
}
