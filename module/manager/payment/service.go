package payment

import (
	"appengine"

	c "github.com/czertbytes/pocket/pkg/comments"
	d "github.com/czertbytes/pocket/pkg/documents"
	shttp "github.com/czertbytes/pocket/pkg/http"
	p "github.com/czertbytes/pocket/pkg/payments"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *shttp.RequestContext
	notificator      *Notificator
	Payments         *p.Payments
	Documents        *d.Documents
	Comments         *c.Comments
}

func NewService(RequestContext *shttp.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
		Payments:         p.NewPayments(RequestContext.AppEngineContext),
		Documents:        d.NewDocuments(RequestContext.AppEngineContext),
		Comments:         c.NewComments(RequestContext.AppEngineContext),
	}
}

func (self *Service) Create(payment *t.Payment, user *t.User) error {

	if err := self.notificator.Create(payment); err != nil {
		return err
	}

	return nil
}

func (self *Service) Find(id t.PaymentId, user *t.User) (*t.Payment, error) {
	return nil, nil
}

func (self *Service) Update(payment *t.Payment, user *t.User) (*t.Payment, error) {
	return nil, nil
}

func (self *Service) Delete(id t.PaymentId, user *t.User) error {
	return nil
}

func (self *Service) CreateDocument(user *t.User) error {
	return nil
}

func (self *Service) FindAllDocuments(id t.PaymentId, user *t.User) (t.Documents, error) {
	return nil, nil
}

func (self *Service) CreateComment(comment *t.Comment, user *t.User) error {

	if err := self.notificator.CreateComment(comment); err != nil {
		return err
	}

	return nil
}

func (self *Service) FindAllComments(id t.PaymentId, user *t.User) (t.Comments, error) {
	return nil, nil
}
