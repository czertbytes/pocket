package auth

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

func (self *Notificator) Create(client *t.Client) error {
	return nil
}
