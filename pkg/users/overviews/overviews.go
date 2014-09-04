package overviews

import (
	"appengine"

	t "github.com/czertbytes/pocket/pkg/types"
)

type Overviews struct {
	AppEngineContext appengine.Context
	Storage          *Storage
}

func NewOverviews(appEngineContext appengine.Context) *Overviews {
	return &Overviews{
		AppEngineContext: appEngineContext,
		Storage:          NewStorage(appEngineContext),
	}
}

func (self *Overviews) Create(userOverview *t.UserOverview) error {
	return self.Storage.Save(userOverview)
}

func (self *Overviews) FindAll(status t.UserOverviewStatus) (t.UserOverviews, error) {
	userOverviews, err := self.Storage.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}

	return userOverviews, nil
}

func (self *Overviews) FindAllActive() (t.UserOverviews, error) {
	userOverviews, err := self.Storage.FindAllActive()
	if err != nil {
		return nil, err
	}

	return userOverviews, nil
}

func (self *Overviews) FindAllByOverviewId(overviewId t.OverviewId) (t.UserOverviews, error) {
	return self.Storage.FindAllByOverviewId(overviewId)
}

func (self *Overviews) Find(id t.UserOverviewId) (t.UserOverview, error) {
	userOverview, err := self.Storage.Find(id)
	if err != nil {
		return t.UserOverview{}, err
	}

	return userOverview, nil
}

func (self *Overviews) FindMulti(ids t.UserOverviewIds) (t.UserOverviews, error) {
	userOverviews, err := self.Storage.FindMulti(ids)
	if err != nil {
		return nil, err
	}

	return userOverviews, nil
}

func (self *Overviews) Update(userOverview t.UserOverview) (t.UserOverview, error) {
	return self.Storage.Update(userOverview)
}

func (self *Overviews) Delete(id t.UserOverviewId) error {
	return self.Storage.Delete(id)
}
