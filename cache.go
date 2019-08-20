package toolkit

import "time"

const (
	_Put = iota
	_Get
	_Remove
	_Close
	_Error
)

type callBack struct {
	Key    interface{}
	Value  interface{}
	ChBack chan callBack
	Route  int
}

// Closer have Close function.
type Closer interface {
	Close()
}

// Cache is LRU.
type Cache struct {
	data       map[interface{}]interface{}
	access     map[interface{}]time.Time
	chCallBack chan callBack
}

// Put vlaue by key.
func (c *Cache) Put(key, value interface{}) {
	ch := make(chan callBack, 1)
	defer close(ch)
	c.chCallBack <- callBack{
		Key:    key,
		Value:  value,
		Route:  _Put,
		ChBack: ch,
	}
	<-ch
}

// Get value by key.
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	ch := make(chan callBack, 1)
	defer close(ch)
	c.chCallBack <- callBack{
		Key:    key,
		Route:  _Get,
		ChBack: ch,
	}
	ret := <-ch
	if ret.Route == _Error {
		return nil, false
	}
	return ret.Value, true
}

// Remove key.
func (c *Cache) Remove(key interface{}) {
	ch := make(chan callBack, 1)
	defer close(ch)
	c.chCallBack <- callBack{
		Key:    key,
		Route:  _Remove,
		ChBack: ch,
	}
	<-ch
}

// Count Cache.
func (c *Cache) Count() int {
	return len(c.data)
}

// Close cache.
func (c *Cache) Close() {
	c.chCallBack <- callBack{
		Route: _Close,
	}
}

func (c *Cache) run() {
	for {
		cb := <-c.chCallBack
		switch cb.Route {
		case _Put:
			c.data[cb.Key] = cb.Value
			c.access[cb.Key] = time.Now()
		case _Get:
			if value, ok := c.data[cb.Key]; ok {
				cb.Value = value
				c.access[cb.Key] = time.Now()
			} else {
				cb.Route = _Error
			}
		case _Remove:
			if value, ok := c.data[cb.Key]; ok {
				if v, o := value.(Closer); o {
					v.Close()
				}
				delete(c.data, cb.Key)
				delete(c.access, cb.Key)
			}
		case _Close:
			for key, value := range c.data {
				if v, ok := value.(Closer); ok {
					v.Close()
				}
				delete(c.data, key)
				delete(c.access, key)
			}
			close(c.chCallBack)
			return
		}
		cb.ChBack <- cb
	}
}

// NewCache new cache.
func NewCache(expire time.Duration) *Cache {
	cache := &Cache{
		data:       make(map[interface{}]interface{}),
		access:     make(map[interface{}]time.Time),
		chCallBack: make(chan callBack, 3),
	}
	go cache.run()
	t := time.Minute
	if expire < t {
		t = expire
	}
	ticker := time.NewTicker(t)
	go func() {
		for range ticker.C {
			now := time.Now()
			for key, v := range cache.access {
				if now.Sub(v) > expire {
					cache.Remove(key)
				}
			}
		}
	}()
	return cache
}
