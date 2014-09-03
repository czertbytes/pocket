package overview

import (
	"appengine"

	h "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Validator struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
}

func NewValidator(RequestContext *h.RequestContext) *Validator {
	return &Validator{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
	}
}

func (self *Validator) Create(overview *t.Overview) error {
	return nil
}

func (self *Validator) Put(overview *t.Overview) error {
	return nil
}

func (self *Validator) Patch(overview *t.Overview, fields []string) error {
	return nil
}

func (self *Validator) CreatePayment(payment *t.Payment) error {
	return nil
}

func (self *Validator) CreateParticipant(user *t.User) error {
	return nil
}
