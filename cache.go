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

func (c *mapCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *mapCache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}

func constructCache(interval time.Duration) mapCache {
	cache := mapCache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
