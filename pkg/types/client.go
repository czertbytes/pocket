package types

type ClientId int64
type ClientIds []ClientId
type ClientClientId string
type ClientStatus int8

func ParseClientStatus(value int) ClientStatus {
	switch value {
	case -1:
		return ClientStatusDeleted
	case 1:
		return ClientStatusActive
	default:
		return ClientStatusUnknown
	}
}

func (self ClientStatus) String() string {
	switch self {
	case ClientStatusDeleted:
		return "deleted"
	case ClientStatusActive:
		return "active"
	default:
		return "unknown"
	}
}

var (
	ClientStatusDeleted ClientStatus = -1
	ClientStatusUnknown ClientStatus = 0
	ClientStatusActive  ClientStatus = 1
)

type Client struct {
	BaseEntity
	Id              ClientId     `json:"id" datastore:"-"`
	Status          ClientStatus `json:"status" datastore:"status"`
	StatusFormatted string       `json:"status_formatted" datastore:"-"`

	// Entity fields
	ClientId    ClientClientId `json:"client_id" datastore:"client_id"`
	ClientToken ClientToken    `json:"token,omitempty" datastore:"token"`
	UserId      UserId         `json:"-" datastore:"user_id"`
	User        User           `json:"user" datastore:"-"`

	// Signup
	AuthOrigin
}

func (self *Client) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
	self.AuthOrigin.SetFormattedValues()
}

func (self *Client) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

type Clients []Client
