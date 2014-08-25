package types

import "appengine"

type DocumentId int64
type DocumentIds []DocumentId
type DocumentStatus int8

func ParseDocumentStatus(value int) DocumentStatus {
	switch value {
	case -1:
		return DocumentStatusDeleted
	case 1:
		return DocumentStatusActive
	default:
		return DocumentStatusUnknown
	}
}

func (self DocumentStatus) String() string {
	switch self {
	case DocumentStatusDeleted:
		return "deleted"
	case DocumentStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	DocumentStatusDeleted DocumentStatus = -1
	DocumentStatusUnknown DocumentStatus = 0
	DocumentStatusActive  DocumentStatus = 1
)

type Document struct {
	BaseEntity
	Id              DocumentId     `json:"id" datastore:"-"`
	Status          DocumentStatus `json:"status" datastore:"status"`
	StatusFormatted string         `json:"status_formatted" datastore:"-"`

	// Entity fields
	URL     string `json:"url" datastore:"url"`
	OwnerId UserId `json:"-" datastore:"owner_id"`
	Owner   User   `json:"owner" datastore:"-"`

	// Internal fields
	PaymentId PaymentId         `json:"-" datastore:"payment_id"`
	BlobKey   appengine.BlobKey `json:"-" datastore:"blob_key"`
}

func (self *Document) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
}

func (self *Document) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

type Documents []Document
