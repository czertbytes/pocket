package types

type CommentId int64
type CommentIds []CommentId
type CommentStatus int8

func (self CommentIds) AsInt64Arr() []int64 {
	ids := make([]int64, len(self))

	for i, id := range self {
		ids[i] = int64(id)
	}

	return ids
}

func ParseCommentStatus(value int) CommentStatus {
	switch value {
	case -1:
		return CommentStatusDeleted
	case 1:
		return CommentStatusActive
	default:
		return CommentStatusUnknown
	}
}

func (self CommentStatus) String() string {
	switch self {
	case CommentStatusDeleted:
		return "deleted"
	case CommentStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	CommentStatusDeleted CommentStatus = -1
	CommentStatusUnknown CommentStatus = 0
	CommentStatusActive  CommentStatus = 1
)

type Comment struct {
	BaseEntity
	Id              CommentId     `json:"id" datastore:"-" siren:"property"`
	Status          CommentStatus `json:"status" datastore:"status" siren:"property"`
	StatusFormatted string        `json:"status_formatted" datastore:"-" siren:"property"`

	// Entity fields
	UserId UserId `json:"-" datastore:"user_id"`
	User   User   `json:"user" datastore:"-"`
	Value  string `json:"value" datastore:"value" siren:"property"`

	// Internal fields
	PaymentId PaymentId `json:"-" datastore:"payment_id"`
}

func (self *Comment) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
}

func (self *Comment) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

type Comments []Comment
