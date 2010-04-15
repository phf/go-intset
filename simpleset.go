package intset

// The simplest (and fastest) way of doing sets,
// at the price of memory *and* Iter() time.
type Simple struct {
	data []bool
}

func (self *Simple) Init(max int) {
	self.data = make([]bool, max+1)
}

func (self *Simple) Insert(i int) {
	self.data[i] = true
}

func (self *Simple) Remove(i int) {
	self.data[i] = false
}

func (self *Simple) Has(i int) (b bool) {
	return self.data[i]
}

func (self *Simple) iterate(c chan<- int) {
	for i, x := range self.data {
		if x {
			c <- i
		}
	}
	close(c)
}

func (self *Simple) Iter() <-chan int {
	c := make(chan int)
	go self.iterate(c)
	return c
}

