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
	URLToken     string       `json:"token" datastore:"token"`
	Name         string       `json:"name" datastore:"name"`
	Description  string       `json:"description" datastore:"description"`
	BaseCurrency string       `json:"base_currency" datastore:"base_currency"`
	OwnerId      UserId       `json:"-" datastore:"owner_id"`
	Owner        User         `json:"owner" datastore:"-"`
	Participants Users        `json:"participants" datastore:"-"`
	Payments     Payments     `json:"payments" datastore:"-"`
	UserAccounts UserAccounts `json:"user_accounts" datastore:"-"`
}

func (self *Overview) SetFormattedValues() {
	self.SetTimes()
	self.SetStatusFormatted()
}

func (self *Overview) SetStatusFormatted() {
	self.StatusFormatted = self.Status.String()
}

func (self *Overview) ComputeUserAccounts() {
	userAccounts := make(map[UserId]*UserAccount)
	for _, participant := range self.Participants {
		userAccounts[participant.Id] = NewUserAccount(participant)
	}
	userAccounts[self.OwnerId] = NewUserAccount(self.Owner)

	for _, payment := range self.Payments {
		from := payment.From
		to := payment.To

		fromUserAccount, found := userAccounts[from.Id]
		if !found {
			continue
		}
		toUserAccount, found := userAccounts[to.Id]
		if !found {
			continue
		}

		fromToBalance, found := fromUserAccount.Balances[to.Id]
		if !found {
			fromToBalance = NewBalance(to)
			fromUserAccount.Balances[to.Id] = fromToBalance
		}

		toFromBalance, found := toUserAccount.Balances[from.Id]
		if !found {
			toFromBalance = NewBalance(from)
			toUserAccount.Balances[from.Id] = toFromBalance
		}

		fromToBalance.AddCredit(payment.Price)
		toFromBalance.AddDebit(payment.Price)
	}

	self.UserAccounts = make(UserAccounts, len(self.Participants)+1)
	i := 0
	for _, userAccount := range userAccounts {
		self.UserAccounts[i] = *userAccount
		i++
	}
}

type Overviews []Overview
