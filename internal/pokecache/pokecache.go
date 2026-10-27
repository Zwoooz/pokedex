// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)
	entry := make(map[string]cacheEntry)
	c := &Cache{
		entry: entry,
	}
	go c.reapLoop(ticker.C, interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entry[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(ch <-chan time.Time, interval time.Duration) {
	for range ch {
		c.mu.Lock()
		for key, entry := range c.entry {
			if entry.createdAt.Add(interval).Before(time.Now()) {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}
