package types

import "time"

type BaseEntity struct {
	CreatedAt  RFC3339Time `json:"created_at" datastore:"-" siren:"property"`
	ModifiedAt RFC3339Time `json:"modified_at" datastore:"-" siren:"property"`

	// appengine will not store RFC3339Time, we must store them as time.Time
	CreatedAtTime  time.Time `json:"-" datastore:"created_at"`
	ModifiedAtTime time.Time `json:"-" datastore:"modified_at"`
}

func (self *BaseEntity) SetCreatedAt(time time.Time) {
	self.CreatedAt = RFC3339Time(time)
	self.CreatedAtTime = time
}

func (self *BaseEntity) SetModifiedAt(time time.Time) {
	self.ModifiedAt = RFC3339Time(time)
	self.ModifiedAtTime = time
}

func (self *BaseEntity) SetTimes() {
	self.ModifiedAt = RFC3339Time(self.ModifiedAtTime)
	self.CreatedAt = RFC3339Time(self.CreatedAtTime)
}

type CreatedAt interface {
	SetCreatedAt(time time.Time)
}

type ModifiedAt interface {
	SetModifiedAt(time time.Time)
}

type CreatedAtModifiedAt interface {
	CreatedAt
	ModifiedAt
	SetTimes()
}
