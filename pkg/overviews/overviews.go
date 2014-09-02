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

func (self *Overviews) Create(overview *t.Overview) error {
	return self.Storage.Save(overview)
}

func (self *Overviews) FindAll(status t.OverviewStatus) (t.Overviews, error) {
	overviews, err := self.Storage.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}

	return overviews, nil
}

func (self *Overviews) FindAllActive() (t.Overviews, error) {
	overviews, err := self.Storage.FindAllActive()
	if err != nil {
		return nil, err
	}

	return overviews, nil
}

func (self *Overviews) Find(id t.OverviewId) (t.Overview, error) {
	overview, err := self.Storage.Find(id)
	if err != nil {
		return t.Overview{}, err
	}

	return overview, nil
}

func (self *Overviews) FindMulti(ids t.OverviewIds) (t.Overviews, error) {
	overviews, err := self.Storage.FindMulti(ids)
	if err != nil {
		return nil, err
	}

	return overviews, nil
}

func (self *Overviews) Update(overview t.Overview) (t.Overview, error) {
	return self.Storage.Update(overview)
}

func (self *Overviews) Delete(id t.OverviewId) error {
	return self.Storage.Delete(id)
}
