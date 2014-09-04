package types

type UserAccount struct {
	User     User                `json:"user" datastore:"-"`
	Balances map[UserId]*Balance `json:"balances" datastore:"-"`
}

func NewUserAccount(user User) *UserAccount {
	return &UserAccount{
		User:     user,
		Balances: make(map[UserId]*Balance),
	}
}

type UserAccounts []UserAccount
