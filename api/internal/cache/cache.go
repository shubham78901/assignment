// internal/cache/cache.go
package cache

import (
	"assignment/api/internal/model"
	"context"
	"sync"
)

type CacheInterface interface {
	Get(ctx context.Context, key string) (model.Country, bool)
	Set(ctx context.Context, key string, value model.Country)
}

type Cache struct {
	data map[string]model.Country
	mu   sync.RWMutex
}

func NewCache() CacheInterface {
	return &Cache{data: make(map[string]model.Country)}
}

func (c *Cache) Get(ctx context.Context, key string) (model.Country, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, exists := c.data[key]
	return val, exists
}

func (c *Cache) Set(ctx context.Context, key string, value model.Country) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
