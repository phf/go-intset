package intset

type Hash struct {
	data map[int]bool
}

func (self *Hash) Init(max int) {
	self.data = make(map[int]bool)
}

func (self *Hash) Insert(i int) {
	self.data[i] = true
}

func (self *Hash) Remove(i int) {
	self.data[i] = false
}

func (self *Hash) Has(i int) (b bool) {
	return self.data[i]
}
