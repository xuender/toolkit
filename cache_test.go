package toolkit

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Second * 1)
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	t.Run("Size", func(t *testing.T) {
		got := cache.Size()
		want := 2
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Get", func(t *testing.T) {
		got, ok := cache.Get("key2")
		want := "value2"
		if !ok {
			t.Errorf("got '%v' want true", ok)
		}
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("Set", func(t *testing.T) {
		cache.Set("key2", "value3")
		got, ok := cache.Get("key2")
		want := "value3"
		if !ok {
			t.Errorf("got '%v' want true", ok)
		}
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("Del", func(t *testing.T) {
		cache.Del("key1")
		got := cache.Size()
		want := 1
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Keys", func(t *testing.T) {
		got := len(cache.Keys())
		want := 1
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Keys", func(t *testing.T) {
		time.Sleep(time.Second * 2)
		got := cache.Size()
		want := 0
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
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
	cache := NewCache(time.Second)
	cache.Set("key1", "value1")
	cache.SetByDuration("key2", "value2", time.Second*3)

	fmt.Println(cache.Get("key1"))
	fmt.Println(cache.Size())
	time.Sleep(time.Second * 2)
	fmt.Println(cache.Size())

	// Output:
	// value1 true
	// 2
	// 1
}

func TestCachePut1(t *testing.T) {
	c := NewCache(time.Second * 3)
	c.Set("3", "3")
	time.Sleep(time.Second * 4)
	if c.Size() != 0 {
		t.Errorf("The expired cache still exists")
	}
}

func TestCachePut2(t *testing.T) {
	c := NewCache(time.Second * 3)
	time.Sleep(time.Second * 1)
	c.Set("3", "3")
	time.Sleep(time.Second * 4)
	if c.Size() != 0 {
		t.Errorf("The expired cache still exists")
	}
}
func TestCachePut3(t *testing.T) {
	c := NewCache(time.Second*2, true)
	c.Set(1, 1)
	time.Sleep(time.Second * 1)
	c.Get(1)
	time.Sleep(time.Second * 1)
	c.Get(1)
	time.Sleep(time.Second * 1)
	if c.Size() == 0 {
		t.Errorf("LRU error")
	}
}
