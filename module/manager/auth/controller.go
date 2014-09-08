package auth

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

func (self *Controller) Post(client *t.Client) (t.Client, error) {
	if err := self.validator.Create(client); err != nil {
		return t.Client{}, err
	}

	return self.service.Create(client, self.RequestContext.User)
}

func (self *Controller) Delete(url *url.URL) error {
	return self.service.Delete(self.RequestContext.Client.Id, self.RequestContext.User)
}
