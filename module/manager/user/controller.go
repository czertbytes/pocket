package user

import (
	"net/url"

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

func (self *Controller) Get(url *url.URL) (*t.User, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) Delete(url *url.URL) error {
	return self.service.Delete(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetOverviews(url *url.URL) (t.Overviews, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetPayments(url *url.URL) (t.Payments, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetComments(url *url.URL) (t.Comments, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetMeOverviews(url *url.URL) (t.Overviews, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetMePayments(url *url.URL) (t.Payments, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetMeComments(url *url.URL) (t.Comments, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}
