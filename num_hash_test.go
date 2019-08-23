package toolkit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	h := NewNumHash(1000)
	h.Write([]byte("test"))
	assert.Equal(t, 571, h.Sum1(), "Hash")
}

func ExampleNewNumHash() {
	h := NewNumHash(10)
	for i := 0; i < 10; i++ {
		h.Reset()
		h.Write([]byte(fmt.Sprintf("value:%d", i)))
		fmt.Println(h.Sum1(), h.Sum2(), h.Sum3(), h.Sum4())
	}

	// Output:
	// 3 2 0 3
	// 4 5 4 7
	// 5 2 7 7
	// 6 5 1 1
	// 7 2 4 7
	// 8 5 8 1
	// 3 8 6 9
	// 4 5 9 5
	// 5 8 3 9
	// 6 5 6 9
}
