package payment

import (
	"net/url"
	"strings"

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

func (self *Controller) Get(url *url.URL) (t.Payment, error) {
	return self.service.Find(t.PaymentId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) Put(url *url.URL, payment *t.Payment) error {
	payment.Id = t.PaymentId(self.RequestContext.EntityId)

	if err := self.validator.Update(payment); err != nil {
		return err
	}

	return self.service.Update(payment, self.RequestContext.User)
}

func (self *Controller) Patch(url *url.URL, payment *t.Payment) error {
	patchFields := url.Query().Get("fields")
	fields := strings.Split(patchFields, ",")

	payment.Id = t.PaymentId(self.RequestContext.EntityId)

	if err := self.validator.Patch(payment, fields); err != nil {
		return err
	}

	return self.service.Patch(payment, fields, self.RequestContext.User)
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

	return self.service.CreateComment(comment, t.PaymentId(self.RequestContext.EntityId), self.RequestContext.User)
}

func (self *Controller) GetComments(url *url.URL) (t.Comments, error) {
	return self.service.FindAllComments(t.PaymentId(self.RequestContext.EntityId), self.RequestContext.User)
}
