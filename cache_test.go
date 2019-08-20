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
		defer cache.Close()
		cache.Put(1, 1)
		cache.Put(2, 2)
		Convey("Count", func() {
			So(cache.Count(), ShouldEqual, 2)
		})
		Convey("Get", func() {
			v, ok := cache.Get(2)
			So(v, ShouldEqual, 2)
			So(ok, ShouldEqual, true)
		})
		/*
			Convey("Keys", func() {
				keys := chMap.Keys()
				So(len(keys), ShouldEqual, 2)
			})
		*/
		Convey("Put", func() {
			cache.Put(2, 3)
			v, ok := cache.Get(2)
			So(v, ShouldEqual, 3)
			So(ok, ShouldEqual, true)
		})
		Convey("Remove", func() {
			cache.Remove(1)
			So(cache.Count(), ShouldEqual, 1)
		})
		Convey("Time", func() {
			time.Sleep(time.Second * 2)
			So(cache.Count(), ShouldEqual, 0)
		})

	})
}
func BenchmarkCacheCount(b *testing.B) {
	cache := NewCache(time.Second * 20)
	defer cache.Close()
	cache.Put(1, 1)
	cache.Put(2, 2)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Count()
	}
}

func BenchmarkCachePut(b *testing.B) {
	cache := NewCache(time.Second * 20)
	defer cache.Close()
	cache.Put(1, 1)
	cache.Put(2, 2)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache.Put(2, 2)
	}
}

func BenchmarkMapPut(b *testing.B) {
	cache := map[int]int{}
	cache[1] = 1
	cache[2] = 2
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cache[2] = 2
	}
}

func ExampleCache() {
	cache := NewCache(1 * time.Second)
	defer cache.Close()
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	fmt.Println(cache.Get("key1"))
	fmt.Println(cache.Count())
	time.Sleep(time.Second * 2)
	fmt.Println(cache.Count())

	// Output:
	// value1 true
	// 2
	// 0
}
