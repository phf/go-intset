package intset

// Briggs and Torzcon's sparse sets.
type Sparse struct {
	dense []int
	sparse []int
	next int
}

func (self *Sparse) Init(max int) {
	self.dense = make([]int, max+1)
	self.sparse = make([]int, max+1)
}

func (self *Sparse) Insert(i int) {
	if !self.Has(i) {
		self.sparse[i] = self.next;
		self.dense[self.next] = i;
		self.next++
	}
}

func (self *Sparse) Remove(i int) {
	if self.Has(i) {
		to := self.sparse[i]
		self.next--;
		element := self.dense[self.next]
		self.dense[to] = element
		self.sparse[element] = to
		self.dense[self.next] = -1
	}
}

func (self *Sparse) Has(i int) (b bool) {
	return self.next > 0 && self.dense[self.sparse[i]] == i
}

func (self *Sparse) iterate(c chan<- int) {
	for i := 0; i < self.next; i++ {
		c <- self.dense[i]
	}
}

func (self *Sparse) Iter() <-chan int {
	c := make(chan int)
	go self.iterate(c)
	return c
}

