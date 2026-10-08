package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CE map[string]cacheEntry
	MU sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.MU.Lock()
	defer c.MU.Unlock()

	c.CE[key] = cacheEntry{
		createdAt: time.Now(),
		val:       append([]byte(nil), val...),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.MU.RLock()
	defer c.MU.RUnlock()

	entry, ok := c.CE[key]
	if !ok {
		return nil, false
	}

	return append([]byte(nil), entry.val...), true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.MU.Lock()
		for k, v := range c.CE {
			if v.createdAt.Add(interval).Before(time.Now()) {
				delete(c.CE, k)
			}
		}
		c.MU.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		CE: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}
