package overviews

import (
	"time"

	"appengine"
	"appengine/datastore"

	"github.com/czertbytes/pocket/pkg/gae"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	kind string = "overview"
)

type Storage struct {
	AppEngineContext appengine.Context
	storage          *gae.Storage
}

func NewStorage(appEngineContext appengine.Context) *Storage {
	return &Storage{
		AppEngineContext: appEngineContext,
		storage:          NewStorage(appEngineContext, kind),
	}
}

func (self *Storage) Save(overview *t.Overview) error {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	overview.SetCreatedAt(now)
	overview.SetModifiedAt(now)

	id, err := self.storage.Save(overview)
	if err != nil {
		return err
	}

	overview.Id = t.OverviewId(id)
	overview.SetFormattedValues()

	return nil
}

func (self *Storage) FindAllByStatus(status t.OverviewStatus) (t.Overviews, error) {
	var overviews t.Overviews

	query := datastore.NewQuery(kind).
		Filter("status =", status)

	ids, err := self.storage.FindAll(query, &overviews)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Overviews, 0), nil
	}

	for i, _ := range overviews {
		overviews[i].Id = t.OverviewId(ids[i])
		overviews[i].SetFormattedValues()
	}

	return overviews, nil
}

func (self *Storage) FindAllActive() (t.Overviews, error) {
	return self.FindAllByStatus(t.OverviewStatusActive)
}

func (self *Storage) Find(id t.OverviewId) (t.Overview, error) {
	var overview t.Overview

	if _, err := self.storage.Find(int64(id), overview); err != nil {
		return t.Overview{}, err
	}

	overview.Id = id
	overview.SetFormattedValues()

	return overview, nil
}

func (self *Storage) FindMulti(ids t.OverviewIds) (t.Overviews, error) {
	var overviews t.Overviews

	if err := self.storage.FindMulti(ids.AsInt64Arr()); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return make(t.Overviews, 0), nil
	}

	for i, _ := range overviews {
		overviews[i].Id = t.OverviewId(ids[i])
		overviews[i].SetFormattedValues()
	}

	return overviews, nil
}

func (self *Storage) Update(overview t.Overview) (t.Overview, error) {
	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)
	overview.SetModifiedAt(now)

	if err := self.storage.Update(int64(overview.Id), overview); err != nil {
		return t.Overview{}, err
	}

	return self.Find(overview.Id)
}

func (self *Storage) Delete(id t.OverviewId) error {
	overview, err := self.Find(id)
	if err != nil {
		return err
	}

	return self.storage.Delete(int64(overview.Id))
}
