package toolkit

import (
	"sync"
	"time"
)

// Cache support LRU (Least Recently Used).
type Cache struct {
	lock   sync.RWMutex
	data   map[interface{}]interface{}
	access map[interface{}]time.Time

	Expire   time.Duration
	LRU      bool
	Callback func(key, value interface{})
}

// Set value by key.
func (c *Cache) Set(key, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = value
	c.access[key] = time.Now().Add(c.Expire)
}

// SetByTime value by key.
func (c *Cache) SetByTime(key, value interface{}, expire time.Time) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = value
	c.access[key] = expire
}

// SetByDuration value by key.
func (c *Cache) SetByDuration(key, value interface{}, expire time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = value
	c.access[key] = time.Now().Add(expire)
}

// Get value by key.
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.lock.RLock()
	value, ok := c.data[key]
	c.lock.RUnlock()
	if c.LRU {
		c.Reset(key)
	}
	return value, ok
}

// GetString by key.
func (c *Cache) GetString(key interface{}) (string, bool) {
	if value, ok := c.Get(key); ok {
		return value.(string), ok
	}
	return "", false
}

// Reset expire by key.
func (c *Cache) Reset(key interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if a, ok := c.access[key]; ok {
		e := time.Now().Add(c.Expire)
		if e.After(a) {
			c.access[key] = e
		}
	}
}

// Keys by map.
func (c *Cache) Keys() []interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	keys := make([]interface{}, len(c.data))
	i := 0
	for k := range c.data {
		keys[i] = k
		i++
	}
	return keys

}

// Del key.
func (c *Cache) Del(key interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if value, ok := c.data[key]; ok {
		c.Callback(key, value)
	}
	delete(c.data, key)
	delete(c.access, key)
}

// Clean overdue.
func (c *Cache) Clean(overdue time.Time) {
	for _, key := range c.Overdue(overdue) {
		c.Del(key)
	}
}

// Overdue keys.
func (c *Cache) Overdue(overdue time.Time) []interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	keys := []interface{}{}
	for key, v := range c.access {
		if overdue.After(v) {
			keys = append(keys, key)
		}
	}
	return keys
}

// Size by Cache.
func (c *Cache) Size() int {
	return len(c.data)
}

var caches = []*Cache{}

// NewCache new cache.
func NewCache(expire time.Duration, LRU ...bool) *Cache {
	cache := &Cache{
		data:     make(map[interface{}]interface{}),
		access:   make(map[interface{}]time.Time),
		Expire:   expire,
		Callback: func(key, value interface{}) {},
	}
	if len(LRU) > 0 {
		cache.LRU = LRU[0]
	}
	if len(caches) == 0 {
		go func() {
			ticker := time.NewTicker(1 * time.Second)
			for now := range ticker.C {
				for _, c := range caches {
					c.Clean(now)
				}
			}
		}()
	}
	caches = append(caches, cache)
	return cache
}

// NewCallbackCache new have del callback cache.
func NewCallbackCache(expire time.Duration, callback func(key, value interface{}), LRU ...bool) *Cache {
	cache := NewCache(expire, LRU...)
	cache.Callback = callback
	return cache
}
