package toolkit

import (
	"fmt"
	"testing"
)

func BenchmarkSyncMap_Put(b *testing.B) {
	m := NewSyncMap()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.Set(n, n)
	}
}

func TestSyncMap_Put(t *testing.T) {
	m := NewSyncMap()
	m.Set("key1", "value1")
	m.Set("key2", "value2")
	if m.Size() != 2 {
		t.Errorf("Size error %d.", m.Size())
	}
}

func ExampleSyncMap() {
	m := NewSyncMap()
	m.Set("key1", "value1")
	m.Set("key2", "value2")

	fmt.Println(m.Get("key1"))
	fmt.Println(m.Size())

	// Output:
	// value1 true
	// 2
}
