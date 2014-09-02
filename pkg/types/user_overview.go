package types

type UserOverviewId int64
type UserOverviewIds []UserOverviewId
type UserOverviewStatus int8

func (self UserOverviewIds) AsInt64Arr() []int64 {
	ids := make([]int64, len(self))

	for i, id := range self {
		ids[i] = int64(id)
	}

	return ids
}

func ParseUserOverviewStatus(value int) UserOverviewStatus {
	switch value {
	case -1:
		return UserOverviewStatusDeleted
	case 1:
		return UserOverviewStatusActive
	default:
		return UserOverviewStatusUnknown
	}
}

func (self UserOverviewStatus) String() string {
	switch self {
	case UserOverviewStatusDeleted:
		return "deleted"
	case UserOverviewStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	UserOverviewStatusDeleted UserOverviewStatus = -1
	UserOverviewStatusUnknown UserOverviewStatus = 0
	UserOverviewStatusActive  UserOverviewStatus = 1
)

type UserOverview struct {
	BaseEntity
	Id              UserOverviewId     `json:"id" datastore:"-" siren:"property"`
	Status          UserOverviewStatus `json:"status" datastore:"status" siren:"property"`
	StatusFormatted string             `json:"status_formatted" datastore:"-" siren:"property"`

	// Internal fields
	OverviewId OverviewId `json:"overview_id" datastore:"overview_id"`
	UserId     UserId     `json:"user_id" datastore:"user_id"`
}

func (self *UserOverview) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
}

func (self *UserOverview) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

type UserOverviews []UserOverview
