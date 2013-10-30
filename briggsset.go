package intset

// Briggs and Torzcon's sparse sets.
type Briggs struct {
	dense []int
	sparse []int
	next int
}

func (self *Briggs) Init(max int) {
	self.dense = make([]int, max+1)
	self.sparse = make([]int, max+1)
}

func (self *Briggs) Insert(i int) {
	if !self.Has(i) {
		self.sparse[i] = self.next;
		self.dense[self.next] = i;
		self.next++
	}
}

func (self *Briggs) Remove(i int) {
	if self.Has(i) {
		to := self.sparse[i]
		self.next--;
		element := self.dense[self.next]
		self.dense[to] = element
		self.sparse[element] = to
		self.dense[self.next] = -1
	}
}

func (self *Briggs) Has(i int) bool {
	return self.next > 0 && self.dense[self.sparse[i]] == i
}

func (self *Briggs) iterate(c chan<- int) {
	for i := 0; i < self.next; i++ {
		c <- self.dense[i]
	}
	close(c)
}

func (self *Briggs) Iter() <-chan int {
	c := make(chan int)
	go self.iterate(c)
	return c
}
