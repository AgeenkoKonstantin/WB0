package cache

import (
	"errors"
	"sync"
)

var (
	ErrNotFoundInCache = errors.New("order not found in cache")
)

type Cache struct {
	mutex sync.RWMutex
	data  map[string]string
}

func (cache *Cache) Get(uid string) (string, error) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	if order, ok := cache.data[uid]; ok {
		return order, nil
	}
	return "", ErrNotFoundInCache
}

func (cache *Cache) Add(uid string, order string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.data[uid] = order
}
