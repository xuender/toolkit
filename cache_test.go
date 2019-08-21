package toolkit

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCache(t *testing.T) {
	Convey("Cache", t, func() {
		cache := NewCache(time.Second * 1)
		cache.Set("key1", "value1")
		cache.Set("key2", "value2")
		Convey("Size", func() {
			So(cache.Size(), ShouldEqual, 2)
		})
		Convey("Get", func() {
			v, ok := cache.Get("key2")
			So(v, ShouldEqual, "value2")
			So(ok, ShouldEqual, true)
		})
		Convey("Set", func() {
			cache.Set("key2", "value3")
			v, ok := cache.Get("key2")
			So(v, ShouldEqual, "value3")
			So(ok, ShouldEqual, true)
		})
		Convey("Del", func() {
			cache.Del("key1")
			So(cache.Size(), ShouldEqual, 1)
		})
		Convey("Keys", func() {
			So(len(cache.Keys()), ShouldEqual, 2)
		})
		Convey("Time", func() {
			time.Sleep(time.Second * 2)
			So(cache.Size(), ShouldEqual, 0)
		})

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
