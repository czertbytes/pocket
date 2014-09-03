package storage

import (
	"appengine"
	"appengine/datastore"
)

type Datastore struct {
	AppEngineContext appengine.Context
	kind             string
}

func NewDatastore(appEngineContext appengine.Context, kind string) *Datastore {
	return &Datastore{
		AppEngineContext: appEngineContext,
		kind:             kind,
	}
}

func (self *Datastore) Save(entity interface{}) (int64, error) {
	incKey := datastore.NewIncompleteKey(self.AppEngineContext, self.kind, nil)
	key, err := datastore.Put(self.AppEngineContext, incKey, entity)
	if err != nil {
		return 0, err
	}

	return key.IntID(), nil
}

func (self *Datastore) Find(id int64, entity interface{}) (int64, error) {
	key := datastore.NewKey(self.AppEngineContext, self.kind, "", id, nil)
	if err := datastore.Get(self.AppEngineContext, key, entity); err != nil {
		return 0, err
	}

	return key.IntID(), nil
}

func (self *Datastore) FindAll(query *datastore.Query, entities interface{}) ([]int64, error) {
	keys, err := query.GetAll(self.AppEngineContext, entities)
	if err != nil {
		return nil, err
	}

	ids := make([]int64, len(keys))
	for i, key := range keys {
		ids[i] = key.IntID()
	}

	return ids, nil
}

func (self *Datastore) FindMulti(ids []int64, entities interface{}) error {
	keys := make([]*datastore.Key, len(ids))
	for i, id := range ids {
		keys[i] = datastore.NewKey(self.AppEngineContext, self.kind, "", id, nil)
	}

	return datastore.GetMulti(self.AppEngineContext, keys, entities)
}

func (self *Datastore) Update(id int64, entity interface{}) error {
	key := datastore.NewKey(self.AppEngineContext, self.kind, "", id, nil)
	if _, err := datastore.Put(self.AppEngineContext, key, entity); err != nil {
		return err
	}

	return nil
}

func (self *Datastore) Delete(id int64) error {
	key := datastore.NewKey(self.AppEngineContext, self.kind, "", id, nil)
	return datastore.Delete(self.AppEngineContext, key)
}
