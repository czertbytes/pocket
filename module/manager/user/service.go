package user

import (
	"appengine"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *shttp.RequestContext
	notificator      *Notificator
}

func NewService(RequestContext *shttp.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
	}
}

func (self *Service) Find(id t.UserId, user *t.User) (*t.User, error) {
	return nil, nil
}

func (self *Service) Delete(id t.UserId, user *t.User) error {
	return nil
}

func (self *Service) FindAllOverviews(id t.UserId, user *t.User) (t.Overviews, error) {
	return nil, nil
}

func (self *Service) FindAllPayments(id t.UserId, user *t.User) (t.Payments, error) {
	return nil, nil
}

func (self *Service) FindAllComments(id t.UserId, user *t.User) (t.Comments, error) {
	return nil, nil
}
