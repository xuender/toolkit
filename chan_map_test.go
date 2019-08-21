package toolkit

import (
	"fmt"
	"testing"
)

func TestChMap(t *testing.T) {
	m := NewChMap()
	defer m.Close()
	m.Set("1", 1)
	m.Set("2", 2)
	t.Run("Size", func(t *testing.T) {
		got := m.Size()
		want := 2
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Get", func(t *testing.T) {
		got, ok := m.Get("2")
		want := 2
		if !ok {
			t.Errorf("got '%v' want true", ok)
		}
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Has", func(t *testing.T) {
		got := m.Has("2")
		if !got {
			t.Errorf("got '%v' want true", got)
		}
		got = m.Has("no")
		if got {
			t.Errorf("got '%v' want false", got)
		}
	})
	t.Run("Keys", func(t *testing.T) {
		got := len(m.Keys())
		want := 2
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Set", func(t *testing.T) {
		m.Set("2", 3)
		got, ok := m.Get("2")
		want := 3
		if !ok {
			t.Errorf("got '%v' want true", ok)
		}
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Del", func(t *testing.T) {
		m.Remove("1")
		got := m.Size()
		want := 1
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
	t.Run("Iterator", func(t *testing.T) {
		m.Iterator(func(k, want interface{}) {
			got, _ := m.Get(k)
			if got != want {
				t.Errorf("got '%d' want '%d'", got, want)
			}
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
