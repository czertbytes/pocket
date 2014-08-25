package comment

import (
	"appengine"

	shttp "github.com/czertbytes/pocket/pkg/http"
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
