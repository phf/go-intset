package intset

// Sparse implementation following Peter Williams' code.
// We were trying to see what the overhead for maps in
// his code would roughly be, that's why I added this.
type Williams struct {
	data map[int]int
}

func (self *Williams) Init(max int) {
	self.data = make(map[int]int)
}

func (self *Williams) Insert(i int) {
	bucket, mask := locate(i)
	self.data[bucket] |= mask
}

func (self *Williams) Remove(i int) {
	bucket, mask := locate(i)
	self.data[bucket] &^= mask
}

func (self *Williams) Has(i int) (b bool) {
	bucket, mask := locate(i)
	return (self.data[bucket] & mask) != 0
}

func (self *Williams) iterate(c chan<- int) {
	for bucket, value := range self.data {
		for i := 0; i < bits_per_int; i++ {
			if value & 1 == 1 {
				c <- bucket*bits_per_int + i
			}
			value >>= 1
		}
	}
	close(c)
}

func (self *Williams) Iter() <-chan int {
	c := make(chan int)
	go self.iterate(c)
	return c
}

