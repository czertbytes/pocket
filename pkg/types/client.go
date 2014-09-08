package types

import "time"

type ClientId int64
type ClientIds []ClientId
type ClientClientId string
type ClientStatus int8

func (self ClientIds) AsInt64Arr() []int64 {
	ids := make([]int64, len(self))

	for i, id := range self {
		ids[i] = int64(id)
	}

	return ids
}

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

func (self *Client) SetTimes() {
	self.BaseEntity.SetTimes()
	self.ClientToken.SetTimes()
}

func (self *Client) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

func (self *Client) GenerateClientId() {
	clientId := Hash(
		"53088f0c8f7e74488e07807dfc022542",
		self.Id,
	)

	self.ClientId = ClientClientId(clientId)
}

var (
	ClientTokenExpirationTime time.Duration = 8 * time.Hour
)

func (self *Client) RegenerateToken() {
	tokenValue := Hash(
		"eb3b8da16fda97e84f6a13b2f94cda93",
		self.ClientId,
		self.Id,
	)

	location, _ := time.LoadLocation(DefaultLocation)
	now := time.Now().In(location)

	self.ClientToken = ClientToken{
		Value:        tokenValue,
		ExpireAtTime: now.Add(ClientTokenExpirationTime),
	}
	self.ClientToken.SetTimes()
}

type Clients []Client
