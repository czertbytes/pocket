package payment

import (
	"fmt"
	"mime/multipart"
	"strings"

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

func (self *Validator) Create(payment *t.Payment) error {
	return nil
}

func (self *Validator) Update(payment *t.Payment) error {
	return nil
}

func (self *Validator) Patch(payment *t.Payment, fields []string) error {
	return nil
}

func (self *Validator) CreateDocument(part *multipart.Part) error {
	fileName := part.FileName()

	hasAllowedSuffix := false
	for _, suffix := range []string{"jpg", "png"} {
		if strings.HasSuffix(fileName, suffix) {
			hasAllowedSuffix = true
		}
	}

	if !hasAllowedSuffix {
		return fmt.Errorf("File suffix is not supported!")
	}

	return nil
}

func (self *Validator) CreateComment(comment *t.Comment) error {
	return nil
}
