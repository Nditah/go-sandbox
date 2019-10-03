package datastore

import (
	"github.com/stretchr/testify/assert"
	"go-sandbox/cache/datatype"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	dataStore := New()
	assert.Equal(t, len(dataStore.cache), 0, "map should be empty")
}

func TestDataStore_set(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	expectedMap := make(map[string]datatype.DataType)
	expectedMap[key] = value

	dataStore.set(key, value)
	if !reflect.DeepEqual(dataStore.cache, expectedMap) {
		t.Errorf("Inner map should contain item")
	}
}

func TestDataStore_get(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	dataStore.cache[key] = value
	assert.Equal(t, dataStore.get(key), value)
}

func TestDataStore_getKeys(t *testing.T) {
	dataStore := New()
	dataStore.cache["key 1"] = datatype.New("value 1", time.Minute)
	dataStore.cache["key 2"] = datatype.New("value 2", time.Minute)

	keys := dataStore.getKeys()
	assert.Equal(t, len(keys), 2)
	assert.Contains(t, keys, "key 1")
	assert.Contains(t, keys, "key 2")
}

func TestDataStore_delete(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	dataStore.cache[key] = value

	assert.Equal(t, dataStore.delete(key), true, "existing key")
	assert.Equal(t, len(dataStore.cache), 0, "map should be empty")
	assert.Equal(t, dataStore.delete(key), false, "absent key")
	assert.Equal(t, dataStore.delete("another key"), false, "absent key 2")
}

func TestDataStore_contains(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	dataStore.cache[key] = value

	assert.Equal(t, dataStore.contains(key), true, "existing key")
	assert.Equal(t, dataStore.contains("another key"), false, "absent key")
}

func TestDataStore_count(t *testing.T) {
	dataStore := New()
	dataStore.cache["key 1"] = datatype.New("value 1", time.Minute)
	dataStore.cache["key 2"] = datatype.New("value 2", time.Minute)

	assert.Equal(t, dataStore.count(), 2)
}

func TestDataStore_Set(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	expectedMap := make(map[string]datatype.DataType)
	expectedMap[key] = value

	dataStore.Set(key, value)
	if !reflect.DeepEqual(dataStore.cache, expectedMap) {
		t.Errorf("Inner map should contain item")
	}
}

func TestDataStore_Get(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	dataStore.cache[key] = value

	res, ok := dataStore.Get(key)
	assert.Equal(t, value, *res)
	assert.Equal(t, true, ok)

	res, ok = dataStore.Get("another key")
	if res != nil {
		t.Errorf("nil should be returned")
	}
	assert.Equal(t, false, ok)
}

func TestDataStore_GetKeys(t *testing.T) {
	dataStore := New()
	dataStore.cache["key 1"] = datatype.New("value 1", time.Minute)
	dataStore.cache["key 2"] = datatype.New("value 2", time.Minute)

	keys := dataStore.GetKeys()
	assert.Equal(t, len(keys), 2)
	assert.Contains(t, keys, "key 1")
	assert.Contains(t, keys, "key 2")
}

func TestDataStore_Delete(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	dataStore.cache[key] = value

	assert.Equal(t, dataStore.Delete(key), true, "existing key")
	assert.Equal(t, len(dataStore.cache), 0, "map should be empty")
	assert.Equal(t, dataStore.Delete(key), false, "absent key")
	assert.Equal(t, dataStore.Delete("another key"), false, "absent key 2")
}

func TestDataStore_Contains(t *testing.T) {
	dataStore := New()
	key := "Some key"
	value := datatype.New("Some value", time.Minute)
	dataStore.cache[key] = value

	assert.Equal(t, dataStore.Contains(key), true, "existing key")
	assert.Equal(t, dataStore.Contains("another key"), false, "absent key")
}

func TestDataStore_Count(t *testing.T) {
	dataStore := New()
	dataStore.cache["key 1"] = datatype.New("value 1", time.Minute)
	dataStore.cache["key 2"] = datatype.New("value 2", time.Minute)

	assert.Equal(t, dataStore.Count(), 2)
}
