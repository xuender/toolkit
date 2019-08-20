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
		chMap.Put("1", 1)
		chMap.Put("2", 2)
		Convey("Count", func() {
			So(chMap.Count(), ShouldEqual, 2)
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
		Convey("Put", func() {
			chMap.Put("2", 3)
			v, ok := chMap.Get("2")
			So(v, ShouldEqual, 3)
			So(ok, ShouldEqual, true)
		})
		Convey("Remove", func() {
			chMap.Remove("1")
			So(chMap.Count(), ShouldEqual, 1)
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
	chMap.Put("key1", "value1")
	chMap.Put("key2", "value2")

	fmt.Println(chMap.Get("key1"))
	fmt.Println(chMap.Count())

	// Output:
	// value1 true
	// 2
}
