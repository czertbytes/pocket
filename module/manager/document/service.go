package document

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

func (self *Service) Find(id t.DocumentId, user *t.User) (*t.Document, error) {
	return nil, nil
}

func (self *Service) Delete(id t.DocumentId, user *t.User) error {
	return nil
}
