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

func (self *Controller) Post(payment *t.Payment) error {
	if err := self.validator.Create(payment); err != nil {
		return err
	}

	return self.service.Create(payment, self.RequestContext.User)
}

func (self *Controller) Get(url *url.URL) (*t.Payment, error) {
	return self.service.Find(t.PaymentId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) Put(url *url.URL, payment *t.Payment) (*t.Payment, error) {
	if err := self.validator.Update(payment); err != nil {
		return nil, err
	}

	return self.service.Update(payment, self.RequestContext.User)
}

func (self *Controller) Patch(url *url.URL, payment *t.Payment) (*t.Payment, error) {
	if err := self.validator.Patch(url, payment); err != nil {
		return nil, err
	}

	return self.service.Update(payment, self.RequestContext.User)
}

func (self *Controller) Delete(url *url.URL) error {
	return self.service.Delete(t.PaymentId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) PostDocuments() error {
	return nil
}

func (self *Controller) GetDocuments(url *url.URL) (t.Documents, error) {
	return self.service.FindAllDocuments(t.PaymentId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) PostComment(comment *t.Comment) error {
	if err := self.validator.CreateComment(comment); err != nil {
		return err
	}

	return self.service.CreateComment(comment, self.RequestContext.User)
}

func (self *Controller) GetComments(url *url.URL) (t.Comments, error) {
	return self.service.FindAllComments(t.PaymentId(self.RequestContext.EntityId), self.RequestContext.User)
}
