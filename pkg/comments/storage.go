package comments

import (
	"time"

	"appengine"
	"appengine/datastore"

	"github.com/czertbytes/pocket/pkg/gae"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "comment"
)

type Storage struct {
	AppEngineContext appengine.Context
	storage          *gae.Storage
}

func NewStorage(appEngineContext appengine.Context) *Storage {
	return &Storage{
		AppEngineContext: appEngineContext,
		storage:          gae.NewStorage(appEngineContext, kind),
	}
}

func (self *Storage) Save(comment *t.Comment) error {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	comment.SetCreatedAt(now)
	comment.SetModifiedAt(now)

	id, err := self.storage.Save(comment)
	if err != nil {
		return err
	}

	comment.Id = t.CommentId(id)
	comment.SetFormattedValues()

	return nil
}

func (self *Storage) FindAllByStatus(status t.CommentStatus) (t.Comments, error) {
	var comments t.Comments

	query := datastore.NewQuery(kind).
		Filter("status =", status)

	ids, err := self.storage.FindAll(query, &comments)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Comments, 0), nil
	}

	for i, _ := range comments {
		comments[i].Id = t.CommentId(ids[i])
		comments[i].SetFormattedValues()
	}

	return comments, nil
}

func (self *Storage) FindAllActive() (t.Comments, error) {
	return self.FindAllByStatus(t.CommentStatusActive)
}

func (self *Storage) Find(id t.CommentId) (t.Comment, error) {
	var comment t.Comment

	if _, err := self.storage.Find(int64(id), comment); err != nil {
		return t.Comment{}, err
	}

	comment.Id = id
	comment.SetFormattedValues()

	return comment, nil
}

func (self *Storage) FindAllByPaymentId(paymentId t.PaymentId) (t.Comments, error) {
	var comments t.Comments

	query := datastore.NewQuery(kind).
		Filter("payment_id=", paymentId).
		Filter("status =", t.CommentStatusActive)

	ids, err := self.storage.FindAll(query, &comments)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Comments, 0), nil
	}

	for i, _ := range comments {
		comments[i].Id = t.CommentId(ids[i])
		comments[i].SetFormattedValues()
	}

	return comments, nil
}

func (self *Storage) FindMulti(ids t.CommentIds) (t.Comments, error) {
	var comments t.Comments

	if err := self.storage.FindMulti(ids.AsInt64Arr(), comments); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Comments, 0), nil
	}

	for i, _ := range comments {
		comments[i].Id = t.CommentId(ids[i])
		comments[i].SetFormattedValues()
	}

	return comments, nil
}

func (self *Storage) Update(comment t.Comment) (t.Comment, error) {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	comment.SetModifiedAt(now)

	if err := self.storage.Update(int64(comment.Id), comment); err != nil {
		return t.Comment{}, err
	}

	return self.Find(comment.Id)
}

func (self *Storage) Delete(id t.CommentId) error {
	comment, err := self.Find(id)
	if err != nil {
		return err
	}

	return self.storage.Delete(int64(comment.Id))
}
