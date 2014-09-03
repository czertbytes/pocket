package auth

import (
	"errors"

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

var (
	ErrAuthOriginServiceRequired error = errors.New("client: Field 'auth_origin_service' is required!")
	ErrAuthOriginServiceNotValid error = errors.New("client: Field 'auth_origin_service' is not valid!")
)

func (self *Validator) Create(client *t.Client) error {
	if err := self.createAuthOriginService(client.AuthOrigin.Service); err != nil {
		return err
	}

	return nil
}

func (self *Validator) createAuthOriginService(service t.AuthOriginService) error {
	if service == t.AuthOriginServiceUnknown {
		return ErrAuthOriginServiceNotValid
	}

	return nil
}
