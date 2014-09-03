package user

import (
	"net/url"

	"appengine"

	h "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Controller struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
	validator        *Validator
	service          *Service
}

func NewController(RequestContext *h.RequestContext) *Controller {
	return &Controller{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		validator:        NewValidator(RequestContext),
		service:          NewService(RequestContext),
	}
}

func (self *Controller) Get(url *url.URL) (t.User, error) {
	return self.service.Find(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) Delete(url *url.URL) error {
	return self.service.Delete(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetOverviews(url *url.URL) (t.Overviews, error) {
	return self.service.FindAllOverviews(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetPayments(url *url.URL) (t.Payments, error) {
	return self.service.FindAllPayments(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetComments(url *url.URL) (t.Comments, error) {
	return self.service.FindAllComments(t.UserId(self.RequestContext.EntityId), self.RequestContext.User)
}
