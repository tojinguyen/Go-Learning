package main

import "sync"

type SafeCache[V any] struct {
	data map[string]V
	mu   sync.RWMutex
}

func (c *SafeCache[V]) Get(key string) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *SafeCache[V]) Set(key string, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *SafeCache[V]) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
