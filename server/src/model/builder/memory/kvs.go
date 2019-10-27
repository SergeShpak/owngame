package memory

import (
	"sync"
)

type keyValStore struct {
	mux  sync.RWMutex
	vals map[string]interface{}
}

func newKeyValStore() *keyValStore {
	kvs := &keyValStore{
		mux:  sync.RWMutex{},
		vals: make(map[string]interface{}),
	}
	return kvs
}

func (kv *keyValStore) Set(key string, val interface{}) {
	kv.mux.Lock()
	kv.vals[key] = val
	kv.mux.Unlock()
}

func (kv *keyValStore) Get(key string) (interface{}, bool) {
	kv.mux.RLock()
	val, ok := kv.vals[key]
	kv.mux.RUnlock()
	return val, ok
}

func (kv *keyValStore) Put(key string, val interface{}) bool {
	kv.mux.Lock()
	if _, ok := kv.vals[key]; ok {
		kv.mux.Unlock()
		return false
	}
	kv.vals[key] = val
	kv.mux.Unlock()
	return true
}

func (kv *keyValStore) Delete(key string) bool {
	kv.mux.Lock()
	if _, ok := kv.vals[key]; !ok {
		kv.mux.Unlock()
		return false
	}
	delete(kv.vals, key)
	kv.mux.Unlock()
	return true
}

func (kv *keyValStore) Alter(key string, alter func(interface{}, bool) (interface{}, bool)) bool {
	kv.mux.Lock()
	val, ok := kv.vals[key]
	newVal, shouldAlter := alter(val, ok)
	if !shouldAlter {
		kv.mux.Unlock()
		return false
	}
	kv.vals[key] = newVal
	kv.mux.Unlock()
	return true
}
