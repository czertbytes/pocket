package auth

import (
	"errors"

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

var (
	ErrAuthOriginNameRequired error = errors.New("client: Field 'name' is required!")
	ErrAuthOriginNameNotValid error = errors.New("client: Field 'name' is not valid!")
)

func (self *Validator) Create(client *t.Client) error {
	if err := self.createAuthOriginName(client.AuthOrigin.Name); err != nil {
		return err
	}

	return nil
}

func (self *Validator) createAuthOriginName(name string) error {
	if len(Name) < 1 {
		return ErrAuthOriginNameRequired
	}

	switch name {
	case "googleplus":
		return nil
	case "facebook":
		return nil
	default:
		return ErrAuthOriginNameNotValid
	}

	return ErrAuthOriginNameNotValid
}
