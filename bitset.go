package intset

import (
	"unsafe"
)

var bits_per_int int = unsafe.Sizeof(bits_per_int)*8

func locate(i int) (bucket int, mask int) {
	bucket = i/bits_per_int
	mask = 1 << uint(i%bits_per_int)
	return
}

// Bitvector implementation for dense sets.
type Bitset struct {
	data []int
}

func (self *Bitset) Init(max int) {
	self.data = make([]int, (max/bits_per_int)+1)
}

func (self *Bitset) Insert(i int) {
	bucket, mask := locate(i)
	self.data[bucket] |= mask
}

func (self *Bitset) Remove(i int) {
	bucket, mask := locate(i)
	self.data[bucket] &^= mask
}

func (self *Bitset) Has(i int) (b bool) {
	bucket, mask := locate(i)
	return (self.data[bucket] & mask) != 0
}
