package storage

import (
	"fmt"
	"io"
)

const (
	createURL  string = "https://www.googleapis.com/upload/storage/v1/b/%s/o?uploadType=multipart&name=%s&predefinedAcl=publicRead"
	contentURL string = "https://%s.storage.googleapis.com/%s"
)

type CloudObject struct {
	Bucket      string
	ObjectName  string
	ContentType string
	Content     io.Reader
}

func NewCloudObject(content io.Reader) *CloudObject {
	return &CloudObject{
		Content: content,
	}
}

func (self *CloudObject) clone() *CloudObject {
	x := *self

	return &x
}

func (self *CloudObject) WithBucket(bucket string) *CloudObject {
	self = self.clone()
	self.Bucket = bucket

	return self
}

func (self *CloudObject) WithObjectName(objectName string) *CloudObject {
	self = self.clone()
	self.ObjectName = objectName

	return self
}

func (self *CloudObject) WithContentType(contentType string) *CloudObject {
	self = self.clone()
	self.ContentType = contentType

	return self
}

func (self *CloudObject) CreateURLPath() string {
	return fmt.Sprintf(createURL, self.Bucket, self.ObjectName)
}

func (self *CloudObject) URLPath() string {
	return fmt.Sprintf(contentURL, self.Bucket, self.ObjectName)
}
