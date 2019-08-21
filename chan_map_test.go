package toolkit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChMap(t *testing.T) {
	m := NewChMap()
	defer m.Close()
	m.Set("1", 1)
	m.Set("2", 2)
	assert.Equal(t, 2, m.Size(), "Size")

	value, ok := m.Get("2")
	assert.Equal(t, 2, value, "Get")
	assert.True(t, ok, "Get ok")

	assert.True(t, m.Has("2"), "Has true")
	assert.False(t, m.Has("no"), "Has false")

	assert.Equal(t, 2, len(m.Keys()), "Keys")

	m.Set("2", 3)
	value, ok = m.Get("2")
	assert.Equal(t, 3, value, "Set")
	assert.True(t, ok, "Set ok")

	m.Del("1")
	assert.Equal(t, 1, m.Size(), "Remove")

	m.Iterator(func(k, want interface{}) {
		got, _ := m.Get(k)
		if got != want {
			assert.Equal(t, want, got, "Iterator")
		}
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
