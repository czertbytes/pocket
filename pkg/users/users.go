package users

import (
	"appengine"

	t "github.com/czertbytes/pocket/pkg/types"
)

type Users struct {
	AppEngineContext appengine.Context
	Storage          *Storage
}

func NewUsers(appEngineContext appengine.Context) *Users {
	return &Users{
		AppEngineContext: appEngineContext,
		Storage:          NewStorage(appEngineContext),
	}
}

func (self *Users) Create(user *t.User) error {
	return self.Storage.Save(user)
}

func (self *Users) FindAll(status t.UserStatus) (t.Users, error) {
	users, err := self.Storage.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (self *Users) FindAllActive() (t.Users, error) {
	users, err := self.Storage.FindAllActive()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (self *Users) Find(id t.UserId) (t.User, error) {
	user, err := self.Storage.Find(id)
	if err != nil {
		return t.User{}, err
	}

	return user, nil
}

func (self *Users) FindByEmail(email string) (t.User, error) {
	return self.Storage.FindByEmail(email)
}

func (self *Users) FindMulti(ids t.UserIds) (t.Users, error) {
	users, err := self.Storage.FindMulti(ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (self *Users) Update(user t.User) (t.User, error) {
	return self.Storage.Update(user)
}

func (self *Users) Delete(id t.UserId) error {
	return self.Storage.Delete(id)
}
