package payment

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

func (self *Service) Create(payment *t.Payment, user *t.User) error {

	if err := self.notificator.Create(payment); err != nil {
		return err
	}

	return nil
}

func (self *Service) CreateDocument(user *t.User) error {
	return nil
}

func (self *Service) FindDocument(id t.DocumentId, user *t.User) (*t.Payment, error) {
	return nil, nil
}

func (self *Service) CreateComment(comment *t.Comment, user *t.User) error {

	if err := self.notificator.Create(comment); err != nil {
		return err
	}

	return nil
}

func (self *Service) FindComment(id t.CommentId, user *t.User) (*t.Comment, error) {
	return nil, nil
}
