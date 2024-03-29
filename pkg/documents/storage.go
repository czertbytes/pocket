package documents

import (
	"mime/multipart"
	"time"

	"appengine"
	"appengine/datastore"

	gs "github.com/czertbytes/pocket/pkg/google/storage"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "document"
)

type Storage struct {
	AppEngineContext appengine.Context
	storage          *gs.Datastore
	cloud            *gs.Cloud
}

func NewStorage(appEngineContext appengine.Context) *Storage {
	return &Storage{
		AppEngineContext: appEngineContext,
		storage:          gs.NewDatastore(appEngineContext, kind),
		cloud:            gs.NewCloud(appEngineContext),
	}
}

func (self *Storage) Save(document *t.Document) error {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	document.SetCreatedAt(now)
	document.SetModifiedAt(now)

	id, err := self.storage.Save(document)
	if err != nil {
		return err
	}

	document.Id = t.DocumentId(id)
	document.SetFormattedValues()

	return nil
}

func (self *Storage) SaveFile(part *multipart.Part) (string, error) {
	name := t.Hash(
		"a3ebf9b990440119c8dbc27cf1ba81be",
		part.FileName(),
	)

	cloudObject := gs.NewCloudObject(part).
		WithBucket("document").
		WithObjectName(name).
		WithContentType(part.Header["Content-Type"][0])

	if err := self.cloud.Save(cloudObject); err != nil {
		return "", err
	}

	return cloudObject.URLPath(), nil
}

func (self *Storage) FindAllByStatus(status t.DocumentStatus) (t.Documents, error) {
	var documents t.Documents

	query := datastore.NewQuery(kind).
		Filter("status =", status)

	ids, err := self.storage.FindAll(query, &documents)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Documents, 0), nil
	}

	for i, _ := range documents {
		documents[i].Id = t.DocumentId(ids[i])
		documents[i].SetFormattedValues()
	}

	return documents, nil
}

func (self *Storage) FindAllActive() (t.Documents, error) {
	return self.FindAllByStatus(t.DocumentStatusActive)
}

func (self *Storage) FindAllByPaymentId(paymentId t.PaymentId) (t.Documents, error) {
	var documents t.Documents

	query := datastore.NewQuery(kind).
		Filter("payment_id=", paymentId).
		Filter("status =", t.DocumentStatusActive)

	ids, err := self.storage.FindAll(query, &documents)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Documents, 0), nil
	}

	for i, _ := range documents {
		documents[i].Id = t.DocumentId(ids[i])
		documents[i].SetFormattedValues()
	}

	return documents, nil
}

func (self *Storage) Find(id t.DocumentId) (t.Document, error) {
	var document t.Document

	if _, err := self.storage.Find(int64(id), document); err != nil {
		return t.Document{}, err
	}

	document.Id = id
	document.SetFormattedValues()

	return document, nil
}

func (self *Storage) FindMulti(ids t.DocumentIds) (t.Documents, error) {
	var documents t.Documents

	if err := self.storage.FindMulti(ids.AsInt64Arr(), documents); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Documents, 0), nil
	}

	for i, _ := range documents {
		documents[i].Id = t.DocumentId(ids[i])
		documents[i].SetFormattedValues()
	}

	return documents, nil
}

func (self *Storage) Update(document t.Document) (t.Document, error) {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	document.SetModifiedAt(now)

	if err := self.storage.Update(int64(document.Id), document); err != nil {
		return t.Document{}, err
	}

	return self.Find(document.Id)
}

func (self *Storage) Delete(id t.DocumentId) error {
	document, err := self.Find(id)
	if err != nil {
		return err
	}

	return self.storage.Delete(int64(document.Id))
}
