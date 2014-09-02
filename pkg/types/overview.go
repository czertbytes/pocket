package types

type OverviewId int64
type OverviewIds []OverviewId
type OverviewStatus int8

func (self OverviewIds) AsInt64Arr() []int64 {
	ids := make([]int64, len(self))

	for i, id := range self {
		ids[i] = int64(id)
	}

	return ids
}

func ParseOverviewStatus(value int) OverviewStatus {
	switch value {
	case -1:
		return OverviewStatusDeleted
	case 1:
		return OverviewStatusActive
	default:
		return OverviewStatusUnknown
	}
}

func (self OverviewStatus) String() string {
	switch self {
	case OverviewStatusDeleted:
		return "deleted"
	case OverviewStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	OverviewStatusDeleted OverviewStatus = -1
	OverviewStatusUnknown OverviewStatus = 0
	OverviewStatusActive  OverviewStatus = 1
)

type Overview struct {
	BaseEntity
	Id              OverviewId     `json:"id" datastore:"-"`
	Status          OverviewStatus `json:"status" datastore:"status"`
	StatusFormatted string         `json:"status_formatted" datastore:"-"`

	// Entity fields
	URLToken     string `json:"token" datastore:"token"`
	Name         string `json:"name" datastore:"name"`
	Description  string `json:"description" datastore:"description"`
	OwnerId      UserId `json:"-" datastore:"owner_id"`
	Owner        User   `json:"owner" datastore:"-"`
	Participants Users  `json:"participants" datastore:"-"`
	Price        `json:"total_price" datastore:"total_price"`
	Payments     Payments `json:"payments" datastore:"-"`
}

func (self *Overview) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
}

func (self *Overview) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

type Overviews []Overview
