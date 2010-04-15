package intset

// Set of integers between 0 and max.
type Set interface {
	Init(max int)
	Insert(i int)
	Remove(i int)
	Has(i int) bool
	Iter() <-chan int
//	Union(other Set) (result Set)
}
