package payment

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

func (self *Controller) Post(overview *t.Overview) error {
	if err := self.validator.Create(overview); err != nil {
		return err
	}

	return self.service.Create(overview, self.RequestContext.User)
}

func (self *Controller) Get(url *url.URL) (*t.Overview, error) {
	return self.service.Find(t.OverviewId(self.RequestContext.EntityId), self.RequestContext.User)
}
