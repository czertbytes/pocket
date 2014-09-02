package documents

import (
	"appengine"

	t "github.com/czertbytes/pocket/pkg/types"
)

type Documents struct {
	AppEngineContext appengine.Context
	Storage          *Storage
}

func NewDocuments(appEngineContext appengine.Context) *Documents {
	return &Documents{
		AppEngineContext: appEngineContext,
		Storage:          NewStorage(appEngineContext),
	}
}

func (self *Documents) Create(document *t.Document) error {
	return self.Storage.Save(document)
}

func (self *Documents) FindAll(status t.DocumentStatus) (t.Documents, error) {
	documents, err := self.Storage.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (self *Documents) FindAllActive() (t.Documents, error) {
	documents, err := self.Storage.FindAllActive()
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (self *Documents) FindAllByPaymentId(paymentId t.PaymentId) (t.Documents, error) {
	return self.Storage.FindAllByPaymentId(paymentId)
}

func (self *Documents) Find(id t.DocumentId) (t.Document, error) {
	document, err := self.Storage.Find(id)
	if err != nil {
		return t.Document{}, err
	}

	return document, nil
}

func (self *Documents) FindMulti(ids t.DocumentIds) (t.Documents, error) {
	documents, err := self.Storage.FindMulti(ids)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (self *Documents) Update(document t.Document) (t.Document, error) {
	return self.Storage.Update(document)
}

func (self *Documents) Delete(id t.DocumentId) error {
	return self.Storage.Delete(id)
}
