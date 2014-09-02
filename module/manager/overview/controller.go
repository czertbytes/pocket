package overview

import (
	"net/url"
	"strings"

	"appengine"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Controller struct {
	AppEngineContext appengine.Context
	RequestContext   *shttp.RequestContext
	validator        *Validator
	service          *Service
}

func NewController(RequestContext *shttp.RequestContext) *Controller {
	return &Controller{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		validator:        NewValidator(RequestContext),
		service:          NewService(RequestContext),
	}
}

func (self *Controller) Post(overview *t.Overview) error {
	if err := self.validator.Create(overview); err != nil {
		return err
	}

	return self.service.Create(overview, self.RequestContext.User)
}

func (self *Controller) Get(url *url.URL) (t.Overview, error) {
	return self.service.Find(t.OverviewId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) Put(url *url.URL, overview *t.Overview) error {
	overview.Id = t.OverviewId(self.RequestContext.EntityId)

	if err := self.validator.Put(overview); err != nil {
		return err
	}

	return self.service.Update(overview, self.RequestContext.User)
}

func (self *Controller) Patch(url *url.URL, overview *t.Overview) error {
	patchFields := url.Query().Get("fields")
	fields := strings.Split(patchFields, ",")

	overview.Id = t.OverviewId(self.RequestContext.EntityId)

	if err := self.validator.Patch(overview, fields); err != nil {
		return err
	}

	return self.service.Patch(overview, fields, self.RequestContext.User)
}

func (self *Controller) Delete(url *url.URL) error {
	return self.service.Delete(t.OverviewId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetOverviewPayments(url *url.URL) (t.Payments, error) {
	return self.service.FindAllPayments(t.OverviewId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) PostOverviewParticipant(user *t.User) error {
	if err := self.validator.CreateParticipant(user); err != nil {
		return err
	}

	return self.service.CreateParticipant(user, t.OverviewId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetOverviewParticipants(url *url.URL) (t.Users, error) {
	return self.service.FindAllParticipants(t.OverviewId(self.RequestContext.EntityId), self.RequestContext.User)
}
