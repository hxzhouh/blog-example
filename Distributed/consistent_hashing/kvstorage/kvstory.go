package kvstorage

import "sync"

type KVStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]string),
	}
}

func (store *KVStore) Set(key, value string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data[key] = value
}

func (store *KVStore) Get(key string) (string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	value, ok := store.data[key]
	return value, ok
}

func (store *KVStore) GetAll() map[string]string {
	return store.data
}

func (store *KVStore) Delete(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.data, key)
}
