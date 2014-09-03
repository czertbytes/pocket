package document

import (
	"appengine"

	h "github.com/czertbytes/pocket/pkg/http"
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
