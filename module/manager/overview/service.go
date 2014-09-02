package overview

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

func (self *Service) Create(overview *t.Overview, user *t.User) error {

	if err := self.notificator.Create(overview); err != nil {
		return err
	}

	return nil
}

func (self *Service) Find(id t.OverviewId, user *t.User) (*t.Overview, error) {
	return nil, nil
}

func (self *Service) Update(overview *t.Overview, user *t.User) error {

	if err := self.notificator.Update(overview); err != nil {
		return err
	}

	return nil
}

func (self *Service) Patch(overview *t.Overview, fields []string, user *t.User) error {
	return nil
}

func (self *Service) Delete(id t.OverviewId, user *t.User) error {
	return nil
}

func (self *Service) FindAllPayments(id t.OverviewId, user *t.User) (t.Payments, error) {
	return nil, nil
}

func (self *Service) CreateParticipant(participant, user *t.User) error {

	if err := self.notificator.CreateParticipant(participant); err != nil {
		return err
	}

	return nil
}

func (self *Service) FindAllParticipants(id t.OverviewId, user *t.User) (t.Users, error) {
	return nil, nil
}
