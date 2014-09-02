package comment

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

func (self *Notificator) Create(comment *t.Comment) error {
	return nil
}
