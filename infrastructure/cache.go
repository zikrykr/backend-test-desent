package infrastructure

import (
	"path/filepath"
	"sync"
	"time"
)

type cacheItem struct {
	value      any
	expiration int64
}

type Cache struct {
	items map[string]cacheItem
	mu    sync.RWMutex
}

func NewCache() *Cache {
	c := &Cache{
		items: make(map[string]cacheItem),
	}
	go c.startCleanupTimer(time.Minute * 10)

	return c
}

func (c *Cache) Set(key string, value any, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var exp int64
	if duration > 0 {
		exp = time.Now().Add(duration).UnixNano()
	}

	c.items[key] = cacheItem{
		value:      value,
		expiration: exp,
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
		return nil, false
	}

	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *Cache) GetAll(pattern string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var result []any
	for key, item := range c.items {
		matched, err := filepath.Match(pattern, key)
		if err != nil || !matched {
			continue
		}

		if item.expiration > 0 && time.Now().UnixNano() > item.expiration {
			continue
		}

		result = append(result, item.value)
	}

	return result, true
}

func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]cacheItem)
}

func (c *Cache) startCleanupTimer(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now().UnixNano()
		for key, item := range c.items {
			if item.expiration > 0 && now > item.expiration {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
