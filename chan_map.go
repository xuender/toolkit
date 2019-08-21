package toolkit

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

// ChMap is channel map.
type ChMap struct {
	data       map[interface{}]interface{}
	chCallBack chan callBack
}

// Set value by key.
func (p ChMap) Set(key, value interface{}) {
	ch := make(chan callBack, 1)
	defer close(ch)
	p.chCallBack <- callBack{
		Key:    key,
		Value:  value,
		Route:  _Put,
		ChBack: ch,
	}
	<-ch
}

// Del obj by key.
func (p ChMap) Del(key interface{}) {
	ch := make(chan callBack, 1)
	defer close(ch)
	p.chCallBack <- callBack{
		Key:    key,
		Route:  _Remove,
		ChBack: ch,
	}
	<-ch
}

// Close this ChMap.
func (p ChMap) Close() {
	p.chCallBack <- callBack{
		Route: _Close,
	}
}

// Size ChMap.
func (p ChMap) Size() int {
	return len(p.data)
}

// Get obj by key.
func (p ChMap) Get(key interface{}) (interface{}, bool) {
	v, ok := p.data[key]
	return v, ok
}

// Has key.
func (p ChMap) Has(key interface{}) bool {
	_, ok := p.data[key]
	return ok
}

// Keys is get this map keys.
func (p ChMap) Keys() []interface{} {
	keys := make([]interface{}, len(p.data))
	i := 0
	for k := range p.data {
		keys[i] = k
		i++
	}
	return keys
}

// Iterator map.
func (p ChMap) Iterator(callBack func(k, v interface{})) {
	for k, v := range p.data {
		if p.Has(k) {
			callBack(k, v)
		}
	}
}

func (p ChMap) run() {
	for {
		cb := <-p.chCallBack
		switch cb.Route {
		case _Put:
			p.data[cb.Key] = cb.Value
		case _Remove:
			if _, ok := p.data[cb.Key]; ok {
				delete(p.data, cb.Key)
			}
		case _Close:
			close(p.chCallBack)
			return
		}
		cb.ChBack <- cb
	}
}

// NewChMap new ChMap.
func NewChMap() *ChMap {
	chMap := ChMap{
		data:       make(map[interface{}]interface{}),
		chCallBack: make(chan callBack, 3),
	}
	go chMap.run()
	return &chMap
}
