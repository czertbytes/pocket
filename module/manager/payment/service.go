package payment

import (
	"fmt"
	"mime/multipart"

	"appengine"

	c "github.com/czertbytes/pocket/pkg/comments"
	d "github.com/czertbytes/pocket/pkg/documents"
	h "github.com/czertbytes/pocket/pkg/http"
	o "github.com/czertbytes/pocket/pkg/overviews"
	p "github.com/czertbytes/pocket/pkg/payments"
	t "github.com/czertbytes/pocket/pkg/types"
)

var (
	ErrPaymentUnauthorized error = fmt.Errorf("payment: Unauthorized access")
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
	notificator      *Notificator
	Overviews        *o.Overviews
	Payments         *p.Payments
	Documents        *d.Documents
	Comments         *c.Comments
}

func NewService(RequestContext *h.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
		Overviews:        o.NewOverviews(RequestContext.AppEngineContext),
		Payments:         p.NewPayments(RequestContext.AppEngineContext),
		Documents:        d.NewDocuments(RequestContext.AppEngineContext),
		Comments:         c.NewComments(RequestContext.AppEngineContext),
	}
}

func (self *Service) Find(id t.PaymentId, user *t.User) (t.Payment, error) {
	return self.getPayment(id, user)
}

func (self *Service) Update(payment *t.Payment, user *t.User) error {
	oldPayment, err := self.getPayment(payment.Id, user)
	if err != nil {
		return err
	}

	oldPayment.Status = payment.Status

	if err := self.Update(&oldPayment, user); err != nil {
		return err
	}

	payment = &oldPayment

	if err := self.notificator.Update(payment); err != nil {
		return err
	}

	return nil
}

func (self *Service) Patch(payment *t.Payment, fields []string, user *t.User) error {
	oldPayment, err := self.getPayment(payment.Id, user)
	if err != nil {
		return err
	}

	for _, field := range fields {
		switch field {
		case "status":
			oldPayment.Status = payment.Status
		}
	}

	if err := self.Update(&oldPayment, user); err != nil {
		return err
	}

	payment = &oldPayment

	return nil
}

func (self *Service) Delete(id t.PaymentId, user *t.User) error {
	if _, err := self.getPayment(id, user); err != nil {
		return err
	}

	if err := self.Payments.Delete(id); err != nil {
		return err
	}

	return nil
}

func (self *Service) CreateDocument(part *multipart.Part, id t.PaymentId, user *t.User) error {
	return nil
}

func (self *Service) FindAllDocuments(id t.PaymentId, user *t.User) (t.Documents, error) {
	if _, err := self.getPayment(id, user); err != nil {
		return nil, err
	}

	paymentDocuments, err := self.Documents.FindAllByPaymentId(id)
	if err != nil {
		return nil, err
	}

	return paymentDocuments, nil
}

func (self *Service) CreateComment(comment *t.Comment, id t.PaymentId, user *t.User) error {
	payment, err := self.getPayment(id, user)
	if err != nil {
		return err
	}

	comment.PaymentId = payment.Id
	if err := self.Comments.Create(comment); err != nil {
		return err
	}

	if err := self.notificator.CreateComment(comment); err != nil {
		return err
	}

	return nil
}

func (self *Service) FindAllComments(id t.PaymentId, user *t.User) (t.Comments, error) {
	if _, err := self.getPayment(id, user); err != nil {
		return nil, err
	}

	paymentComments, err := self.Comments.FindAllByPaymentId(id)
	if err != nil {
		return nil, err
	}

	return paymentComments, nil
}

func (self *Service) getPayment(id t.PaymentId, user *t.User) (t.Payment, error) {
	payment, err := self.Payments.Find(id)
	if err != nil {
		return t.Payment{}, err
	}

	overview, err := self.Overviews.Find(payment.OverviewId)
	if err != nil {
		return t.Payment{}, err
	}

	if payment.FromId == user.Id || payment.ToId == user.Id || overview.OwnerId == user.Id {
		return payment, nil
	}

	return t.Payment{}, ErrPaymentUnauthorized
}
