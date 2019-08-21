package toolkit

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Second * 1)
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	assert.Equal(t, 2, cache.Size(), "Size")

	value, ok := cache.Get("key2")
	assert.Equal(t, "value2", value, "Get")
	assert.True(t, ok, "Get ok")

	cache.Set("key2", "value3")
	value, ok = cache.Get("key2")
	assert.Equal(t, "value3", value, "Get")
	assert.True(t, ok, "Get ok")

	cache.Del("key1")
	assert.Equal(t, 1, cache.Size(), "Size")

	assert.Equal(t, 1, len(cache.Keys()), "Keys")

	time.Sleep(time.Second * 2)
	assert.Equal(t, 0, cache.Size(), "time")
}

func BenchmarkCacheCount(b *testing.B) {
	cache := NewCache(time.Second * 20)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Size()
	}
}

func BenchmarkCachePut(b *testing.B) {
	cache := NewCache(time.Second * 20)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Set("key1", "value1")
	}
}

func BenchmarkMapPut(b *testing.B) {
	cache := map[string]string{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache["key1"] = "value1"
	}
}

func ExampleCache() {
	cache := NewCache(time.Second*3, true)
	cache.Set("key1", "value1")
	cache.SetByDuration("key2", "value2", time.Second)
	cache.Set("key3", "value3")

	fmt.Println(cache.Get("key1"))
	fmt.Println("init size:", cache.Size())
	time.Sleep(time.Second * 2)
	cache.Get("key3") // reset expire time.
	fmt.Println("2 second:", cache.Size())
	time.Sleep(time.Second * 2)
	fmt.Println("4 second:", cache.Size())

	// Output:
	// value1 true
	// init size: 3
	// 2 second: 2
	// 4 second: 1
}

func TestCachePut1(t *testing.T) {
	c := NewCache(time.Second * 3)
	c.Set("3", "3")
	time.Sleep(time.Second * 4)
	assert.Equal(t, 0, c.Size(), "exists")
}

func TestCachePut2(t *testing.T) {
	c := NewCache(time.Second * 3)
	time.Sleep(time.Second * 1)
	c.Set("3", "3")
	time.Sleep(time.Second * 4)
	assert.Equal(t, 0, c.Size(), "exists")
}
func TestCachePut3(t *testing.T) {
	c := NewCache(time.Second*2, true)
	c.Set(1, 1)
	time.Sleep(time.Second * 1)
	c.Get(1)
	time.Sleep(time.Second * 1)
	c.Get(1)
	time.Sleep(time.Second * 1)
	assert.Equal(t, 1, c.Size(), "LRU")
}
