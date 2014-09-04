package users

import (
	"fmt"
	"time"

	"appengine"
	"appengine/datastore"

	gs "github.com/czertbytes/pocket/pkg/google/storage"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "user"
)

var (
	ErrUserNotFound error = fmt.Errorf("user: Not found!")
)

type Storage struct {
	AppEngineContext appengine.Context
	storage          *gs.Datastore
}

func NewStorage(appEngineContext appengine.Context) *Storage {
	return &Storage{
		AppEngineContext: appEngineContext,
		storage:          gs.NewDatastore(appEngineContext, kind),
	}
}

func (self *Storage) Save(user *t.User) error {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	user.SetCreatedAt(now)
	user.SetModifiedAt(now)

	id, err := self.storage.Save(user)
	if err != nil {
		return err
	}

	user.Id = t.UserId(id)
	user.SetFormattedValues()

	return nil
}

func (self *Storage) FindAllByStatus(status t.UserStatus) (t.Users, error) {
	var users t.Users

	query := datastore.NewQuery(kind).
		Filter("status =", status)

	ids, err := self.storage.FindAll(query, &users)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Users, 0), nil
	}

	for i, _ := range users {
		users[i].Id = t.UserId(ids[i])
		users[i].SetFormattedValues()
	}

	return users, nil
}

func (self *Storage) FindAllActive() (t.Users, error) {
	return self.FindAllByStatus(t.UserStatusActive)
}

func (self *Storage) Find(id t.UserId) (t.User, error) {
	var user t.User

	if _, err := self.storage.Find(int64(id), user); err != nil {
		return t.User{}, err
	}

	user.Id = id
	user.SetFormattedValues()

	return user, nil
}

func (self *Storage) FindByEmail(email string) (t.User, error) {
	var users t.Users

	query := datastore.NewQuery(kind).
		Filter("email =", email).
		Filter("status =", t.UserStatusActive)

	ids, err := self.storage.FindAll(query, &users)
	if err != nil {
		return t.User{}, err
	}

	if len(ids) == 0 {
		return t.User{}, ErrUserNotFound
	}

	for i, _ := range users {
		users[i].Id = t.UserId(ids[i])
		users[i].SetFormattedValues()
	}

	return users[0], nil
}

func (self *Storage) FindMulti(ids t.UserIds) (t.Users, error) {
	var users t.Users

	if err := self.storage.FindMulti(ids.AsInt64Arr(), &users); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Users, 0), nil
	}

	for i, _ := range users {
		users[i].Id = t.UserId(ids[i])
		users[i].SetFormattedValues()
	}

	return users, nil
}

func (self *Storage) Update(user t.User) (t.User, error) {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	user.SetModifiedAt(now)

	if err := self.storage.Update(int64(user.Id), user); err != nil {
		return t.User{}, err
	}

	return self.Find(user.Id)
}

func (self *Storage) Delete(id t.UserId) error {
	user, err := self.Find(id)
	if err != nil {
		return err
	}

	return self.storage.Delete(int64(user.Id))
}
