package toolkit

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChMap(t *testing.T) {
	Convey("ChMap", t, func() {
		chMap := NewChMap()
		defer chMap.Close()
		chMap.Set("1", 1)
		chMap.Set("2", 2)
		Convey("Size", func() {
			So(chMap.Size(), ShouldEqual, 2)
		})
		Convey("Get", func() {
			v, ok := chMap.Get("2")
			So(v, ShouldEqual, 2)
			So(ok, ShouldEqual, true)
		})
		Convey("Has", func() {
			So(chMap.Has("1"), ShouldEqual, true)
			So(chMap.Has("no"), ShouldEqual, false)
		})
		Convey("Keys", func() {
			keys := chMap.Keys()
			So(len(keys), ShouldEqual, 2)
		})
		Convey("Set", func() {
			chMap.Set("2", 3)
			v, ok := chMap.Get("2")
			So(v, ShouldEqual, 3)
			So(ok, ShouldEqual, true)
		})
		Convey("Remove", func() {
			chMap.Remove("1")
			So(chMap.Size(), ShouldEqual, 1)
		})
		Convey("Iterator", func() {
			chMap.Iterator(func(k, v interface{}) {
				g, _ := chMap.Get(k)
				So(g, ShouldEqual, v)
			})
		})
	})
}

func ExampleChMap() {
	chMap := NewChMap()
	defer chMap.Close()
	chMap.Set("key1", "value1")
	chMap.Set("key2", "value2")

	fmt.Println(chMap.Get("key1"))
	fmt.Println(chMap.Size())

	// Output:
	// value1 true
	// 2
}

func BenchmarkChMap_Put(b *testing.B) {
	m := NewChMap()
	defer m.Close()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.Set(n, n)
	}
}

func BenchmarkMap(b *testing.B) {
	m := map[int]int{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m[n] = n
	}
}
