package intset

// Locate the bucket and mask for the given bit.
func locate(bit int) (bucket int, mask int) {
	bucket = bit / bitsPerInt
	mask = 1 << uint(bit % bitsPerInt)
	return
}

// Bitvector implementation for dense sets.
type Bitset struct {
	data []int
}

func (self *Bitset) Init(max int) {
	self.data = make([]int, (max / bitsPerInt) + 1)
}

func (self *Bitset) Insert(i int) {
	bucket, mask := locate(i)
	self.data[bucket] |= mask
}

func (self *Bitset) Remove(i int) {
	bucket, mask := locate(i)
	self.data[bucket] &^= mask
}

func (self *Bitset) Has(i int) bool {
	bucket, mask := locate(i)
	return (self.data[bucket] & mask) != 0
}

func (self *Bitset) iterate(c chan<- int) {
	for bucket, value := range self.data {
		t := bucket * bitsPerInt // loop invariant
		for i := 0; i < bitsPerInt && value != 0; i++ {
			if value & 1 == 1 {
				c <- t + i
			}
			value >>= 1
		}
	}
	close(c)
}

func (self *Bitset) Iter() <-chan int {
	c := make(chan int)
	go self.iterate(c)
	return c
}
