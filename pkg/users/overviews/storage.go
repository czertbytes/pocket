package overviews

import (
	"time"

	"appengine"
	"appengine/datastore"

	gs "github.com/czertbytes/pocket/pkg/google/storage"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "user_overview"
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

func (self *Storage) Save(userOverview *t.UserOverview) error {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	userOverview.SetCreatedAt(now)
	userOverview.SetModifiedAt(now)

	id, err := self.storage.Save(userOverview)
	if err != nil {
		return err
	}

	userOverview.Id = t.UserOverviewId(id)
	userOverview.SetFormattedValues()

	return nil
}

func (self *Storage) FindAllByStatus(status t.UserOverviewStatus) (t.UserOverviews, error) {
	var userOverviews t.UserOverviews

	query := datastore.NewQuery(kind).
		Filter("status =", status)

	ids, err := self.storage.FindAll(query, &userOverviews)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.UserOverviews, 0), nil
	}

	for i, _ := range userOverviews {
		userOverviews[i].Id = t.UserOverviewId(ids[i])
		userOverviews[i].SetFormattedValues()
	}

	return userOverviews, nil
}

func (self *Storage) FindAllActive() (t.UserOverviews, error) {
	return self.FindAllByStatus(t.UserOverviewStatusActive)
}

func (self *Storage) FindAllByOverviewId(overviewId t.OverviewId) (t.UserOverviews, error) {
	var userOverviews t.UserOverviews

	query := datastore.NewQuery(kind).
		Filter("overview_id =", overviewId).
		Filter("status =", t.UserOverviewStatusActive)

	ids, err := self.storage.FindAll(query, &userOverviews)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.UserOverviews, 0), nil
	}

	for i, _ := range userOverviews {
		userOverviews[i].Id = t.UserOverviewId(ids[i])
		userOverviews[i].SetFormattedValues()
	}

	return userOverviews, nil
}

func (self *Storage) Find(id t.UserOverviewId) (t.UserOverview, error) {
	var userOverview t.UserOverview

	if _, err := self.storage.Find(int64(id), userOverview); err != nil {
		return t.UserOverview{}, err
	}

	userOverview.Id = id
	userOverview.SetFormattedValues()

	return userOverview, nil
}

func (self *Storage) FindMulti(ids t.UserOverviewIds) (t.UserOverviews, error) {
	var userOverviews t.UserOverviews

	if err := self.storage.FindMulti(ids.AsInt64Arr(), &userOverviews); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.UserOverviews, 0), nil
	}

	for i, _ := range userOverviews {
		userOverviews[i].Id = t.UserOverviewId(ids[i])
		userOverviews[i].SetFormattedValues()
	}

	return userOverviews, nil
}

func (self *Storage) Update(userOverview t.UserOverview) (t.UserOverview, error) {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	userOverview.SetModifiedAt(now)

	if err := self.storage.Update(int64(userOverview.Id), userOverview); err != nil {
		return t.UserOverview{}, err
	}

	return self.Find(userOverview.Id)
}

func (self *Storage) Delete(id t.UserOverviewId) error {
	userOverview, err := self.Find(id)
	if err != nil {
		return err
	}

	return self.storage.Delete(int64(userOverview.Id))
}
