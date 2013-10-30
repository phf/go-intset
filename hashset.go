package intset

type Hash struct {
	data map[int]bool
}

func (self *Hash) Init(max int) {
	self.data = make(map[int]bool, max+1)
}

func (self *Hash) Insert(i int) {
	self.data[i] = true
}

func (self *Hash) Remove(i int) {
	delete(self.data, i)
}

func (self *Hash) Has(i int) bool {
	return self.data[i]
}

func (self *Hash) iterate(c chan<- int) {
	for k, _ := range self.data {
		c <- k
	}
	close(c)
}

func (self *Hash) Iter() <-chan int {
	c := make(chan int)
	go self.iterate(c)
	return c
}
