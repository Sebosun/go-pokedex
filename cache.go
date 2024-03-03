package main

import (
	"sync"
	"time"
)

type mapCache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// key is an url
func (c *mapCache) Add(key string, entry []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.entries[key]

	// if we already have it no need to save it
	if ok {
		return
	}

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       entry,
	}

}

func (c *mapCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return v.val, true
}

func constructCache() mapCache {
	return mapCache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
}
