package toolkit

import (
	"hash"
	"hash/fnv"
)

// NumHash is hash to num.
type NumHash struct {
	Max  int
	Hash hash.Hash32
}

// NewNumHash by max num.
func NewNumHash(max int) *NumHash {
	return &NumHash{Max: max, Hash: fnv.New32a()}
}

// Writer is the interface that wraps the basic Write method.
func (h *NumHash) Write(data []byte) (int, error) {
	return h.Hash.Write(data)
}

// Reset resets the Hash to its initial state.
func (h *NumHash) Reset() {
	h.Hash.Reset()
}

// Sum1 bytes hash to int by max.
func (h *NumHash) Sum1() int {
	return h.sum([]int{0, 1, 2, 3})
}

// Sum2 bytes hash to int by max.
func (h *NumHash) Sum2() int {
	return h.sum([]int{3, 2, 1, 0})
}

// Sum3 bytes hash to int by max.
func (h *NumHash) Sum3() int {
	return h.sum([]int{2, 3, 0, 1})
}

// Sum4 bytes hash to int by max.
func (h *NumHash) Sum4() int {
	return h.sum([]int{1, 0, 3, 2})
}

func (h *NumHash) sum(nums []int) int {
	bs := h.Hash.Sum(nil)
	ret := 0
	for i, num := range nums {
		ret += int(bs[num]<<uint(i)) % h.Max
	}
	return ret % h.Max
}
