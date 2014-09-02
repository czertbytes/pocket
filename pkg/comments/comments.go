package comments

import (
	"appengine"

	t "github.com/czertbytes/pocket/pkg/types"
)

type Comments struct {
	AppEngineContext appengine.Context
	Storage          *Storage
}

func NewComments(appEngineContext appengine.Context) *Comments {
	return &Comments{
		AppEngineContext: appEngineContext,
		Storage:          NewStorage(appEngineContext),
	}
}

func (self *Comments) Create(comment *t.Comment) error {
	return self.Storage.Save(comment)
}

func (self *Comments) FindAll(status t.CommentStatus) (t.Comments, error) {
	comments, err := self.Storage.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}

	// set users

	return comments, nil
}

func (self *Comments) FindAllActive() (t.Comments, error) {
	comments, err := self.Storage.FindAllActive()
	if err != nil {
		return nil, err
	}

	// set users

	return comments, nil
}

func (self *Comments) FindAllByPaymentId(paymentId t.PaymentId) (t.Comments, error) {
	return self.Storage.FindAllByPaymentId(paymentId)
}

func (self *Comments) Find(id t.CommentId) (t.Comment, error) {
	comment, err := self.Storage.Find(id)
	if err != nil {
		return t.Comment{}, err
	}

	// set user

	return comment, nil
}

func (self *Comments) FindMulti(ids t.CommentIds) (t.Comments, error) {
	comments, err := self.Storage.FindMulti(ids)
	if err != nil {
		return nil, err
	}

	// set users

	return comments, nil
}

func (self *Comments) Update(comment t.Comment) (t.Comment, error) {
	return self.Storage.Update(comment)
}

func (self *Comments) Delete(id t.CommentId) error {
	return self.Storage.Delete(id)
}
