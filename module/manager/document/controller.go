package document

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

func (self *Controller) Get(url *url.URL) (*t.Document, error) {
	return self.service.Find(t.DocumentId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) Delete(url *url.URL) error {
	return self.service.Delete(t.DocumentId(self.RequestContext.EntityId), self.RequestContext.User)
}
