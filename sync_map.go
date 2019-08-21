package toolkit

import "sync"

// SyncMap sync map
type SyncMap struct {
	lock sync.RWMutex
	data map[interface{}]interface{}
}

// Set value by key.
func (p *SyncMap) Set(key, value interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.data[key] = value
}

// Del obj by key.
func (p *SyncMap) Del(key interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.data, key)
}

// Size by map.
func (p *SyncMap) Size() int {
	return len(p.data)
}

// Get obj by key.
func (p *SyncMap) Get(key interface{}) (interface{}, bool) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	v, ok := p.data[key]
	return v, ok
}

// Has key.
func (p *SyncMap) Has(key interface{}) bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	_, ok := p.data[key]
	return ok
}

// Keys is get this map keys.
func (p *SyncMap) Keys() []interface{} {
	p.lock.RLock()
	defer p.lock.RUnlock()
	keys := make([]interface{}, len(p.data))
	i := 0
	for k := range p.data {
		keys[i] = k
		i++
	}
	return keys
}

// NewSyncMap new SyncMap.
func NewSyncMap() *SyncMap {
	syncMap := SyncMap{
		data: make(map[interface{}]interface{}),
	}
	return &syncMap
}
