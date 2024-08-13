package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache           map[string]cacheEntry
	mu              *sync.Mutex
	cleanupInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache:           make(map[string]cacheEntry),
		mu:              &sync.Mutex{},
		cleanupInterval: interval,
	}
	go c.readLoop() // Start the readLoop in a separate goroutine
	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) readLoop() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.cleanupInterval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
