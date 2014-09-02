package payment

import (
	"net/url"

	"appengine"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Validator struct {
	AppEngineContext appengine.Context
	RequestContext   *shttp.RequestContext
}

func NewValidator(RequestContext *shttp.RequestContext) *Validator {
	return &Validator{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
	}
}

func (self *Validator) Create(payment *t.Payment) error {
	return nil
}

func (self *Validator) Update(payment *t.Payment) error {
	return nil
}

func (self *Validator) Patch(url *url.URL, payment *t.Payment) error {
	return nil
}

func (self *Validator) CreateComment(comment *t.Comment) error {
	return nil
}
