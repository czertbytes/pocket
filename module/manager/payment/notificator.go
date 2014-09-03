package payment

import (
	"appengine"

	h "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Notificator struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
}

func NewNotificator(RequestContext *h.RequestContext) *Notificator {
	return &Notificator{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
	}
}

func (self *Notificator) Create(payment *t.Payment) error {
	return nil
}

func (self *Notificator) Update(payment *t.Payment) error {
	return nil
}

func (self *Notificator) CreateComment(comment *t.Comment) error {
	return nil
}
