package main

import (
	"sync"
	"time"
)

type Value struct {
	Value      string
	Expiration time.Time
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type KVStore struct {
	data map[string]Value
	mu   sync.RWMutex
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// NewKVStore creates a new KVStore
func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]Value),
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Set sets a key-value pair in the store
func (k *KVStore) Set(key, value string) {
	k.mu.Lock()
	defer k.mu.Unlock()

	k.data[key] = Value{
		Value: value,
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Set sets a key-value pair in the store
func (k *KVStore) SetWithTTL(key, value string, ttl time.Duration) {
	k.mu.Lock()
	defer k.mu.Unlock()

	k.data[key] = Value{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Get gets a value from the store given a key
func (k *KVStore) Get(key string) (string, bool) {
	k.mu.RLock()
	defer k.mu.RUnlock()

	value, ok := k.data[key]
	if !ok {
		return "", false
	}

	if value.Expiration.IsZero() || time.Now().Before(value.Expiration) {
		return value.Value, true
	}

	delete(k.data, key)
	return "", false
}
