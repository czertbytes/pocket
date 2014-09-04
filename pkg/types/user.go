package types

type UserId int64
type UserIds []UserId
type UserStatus int8

func (self UserIds) AsInt64Arr() []int64 {
	ids := make([]int64, len(self))

	for i, id := range self {
		ids[i] = int64(id)
	}

	return ids
}

func ParseUserStatus(value int) UserStatus {
	switch value {
	case -1:
		return UserStatusDeleted
	case 1:
		return UserStatusActive
	default:
		return UserStatusUnknown
	}
}

func (self UserStatus) String() string {
	switch self {
	case UserStatusDeleted:
		return "deleted"
	case UserStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	UserStatusDeleted UserStatus = -1
	UserStatusUnknown UserStatus = 0
	UserStatusActive  UserStatus = 1
)

type User struct {
	BaseEntity
	Id              UserId     `json:"id" datastore:"-" siren:"property"`
	Status          UserStatus `json:"status" datastore:"status" siren:"property"`
	StatusFormatted string     `json:"status_formatted" datastore:"-" siren:"property"`

	// Entity fields
	FullName string `json:"full_name" datastore:"full_name" siren:"property"`
	Email    string `json:"email" datastore:"email" siren:"property"`

	// Internal fields
	GooglePlusId string `json:"-" datastore:"google_plus_id"`
	FacebookId   string `json:"-" datastore:"facebook_id"`
}

func (self *User) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
}

func (self *User) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

type Users []User
