package types

import "time"

type PaymentId int64
type PaymentIds []PaymentId
type PaymentStatus int8

func ParsePaymentStatus(value int) PaymentStatus {
	switch value {
	case -1:
		return PaymentStatusDeleted
	case 1:
		return PaymentStatusActive
	case 2:
		return PaymentStatusCancelled
	default:
		return PaymentStatusUnknown
	}
}

func (self PaymentStatus) String() string {
	switch self {
	case PaymentStatusDeleted:
		return "deleted"
	case PaymentStatusActive:
		return "active"
	case PaymentStatusCancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

var (
	PaymentStatusDeleted   PaymentStatus = -1
	PaymentStatusUnknown   PaymentStatus = 0
	PaymentStatusActive    PaymentStatus = 1
	PaymentStatusCancelled PaymentStatus = 2
)

type Payment struct {
	BaseEntity
	Id              PaymentId     `json:"id" datastore:"-"`
	Status          PaymentStatus `json:"status" datastore:"status"`
	StatusFormatted string        `json:"status_formatted" datastore:"-"`

	// Entity fields
	PaidAt        RFC3339Time `json:"paid_at" datastore:"-"`
	PaidAtTime    time.Time   `json:"-" datastore:"paid_at"`
	Location      `json:"location" datastore:"location"`
	FromId        UserId `json:"-" datastore:"from_id"`
	From          User   `json:"from" datastore:"-"`
	ToId          UserId `json:"-" datastore:"to_id"`
	To            User   `json:"to" datastore:"-"`
	PaymentMethod `json:"payment_method" datastore:"payment_method"`
	Price         `json:"price" datastore:"price"`

	// Internal fields
	OverviewId OverviewId `json:"-" datastore:"overview_id"`
}

func (self *Payment) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
}

func (self *Payment) SetTimes() {
	self.BaseEntity.SetTimes()
	self.PaidAt = RFC3339Time(self.PaidAtTime)
}

func (self *Payment) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

type Payments []Payment
