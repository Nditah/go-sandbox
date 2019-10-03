package datastore

import (
	"go-sandbox/cache/datatype"
	"sync"
)

type DataStore struct {
	sync.RWMutex // ← this mutex protect cache below
	cache        map[string]datatype.DataType
}

func New() *DataStore {
	return &DataStore{
		cache: make(map[string]datatype.DataType),
	}
}

func (ds *DataStore) set(key string, value datatype.DataType) {
	ds.cache[key] = value
}

func (ds *DataStore) get(key string) datatype.DataType {
	return ds.cache[key]
}

func (ds *DataStore) getKeys() []string {
	keys := make([]string, 0, len(ds.cache))
	for k := range ds.cache {
		keys = append(keys, k)
	}
	return keys
}

func (ds *DataStore) delete(key string) bool {
	if ds.contains(key) {
		delete(ds.cache, key)
		return true
	}
	return false
}

func (ds *DataStore) contains(key string) bool {
	_, ok := ds.cache[key]
	return ok
}

func (ds *DataStore) count() int {
	return len(ds.cache)
}

func (ds *DataStore) Set(key string, value datatype.DataType) {
	ds.Lock()
	defer ds.Unlock()
	ds.set(key, value)
}

func (ds *DataStore) Get(key string) (*datatype.DataType, bool) {
	ds.RLock()
	defer ds.RUnlock()
	isContains := ds.contains(key)
	if isContains {
		res := ds.get(key)
		return &res, isContains
	}
	return nil, isContains
}

func (ds *DataStore) GetKeys() []string {
	ds.RLock()
	defer ds.RUnlock()
	return ds.getKeys()
}

func (ds *DataStore) Delete(key string) bool {
	ds.Lock()
	defer ds.Unlock()
	return ds.delete(key)
}

func (ds *DataStore) Contains(key string) bool {
	ds.RLock()
	defer ds.RUnlock()
	return ds.contains(key)
}

func (ds *DataStore) Count() int {
	ds.RLock()
	defer ds.RUnlock()
	return ds.count()
}
