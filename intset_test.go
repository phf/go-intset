// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package intset

import (
	"rand"
	"testing"
)

func checkEmpty(t *testing.T, s Set, max int) {
	for i := 0; i < max; i++ {
		if s.Has(i) {
			t.Errorf("empty set contains %v", i)
		}
	}
}

func checkInsert(t *testing.T, s Set, max int) {
	for i := 0; i < max; i++ {
		s.Insert(i)
		for j := 0; j <= i; j++ {
			if !s.Has(j) {
				t.Errorf("insert of %v didn't work", i)
			}
		}
		for j := i+1; j < max; j++ {
			if s.Has(j) {
				t.Errorf("insert of %v caused %v to be set", i, j)
			}
		}
	}
}

func checkIter(t *testing.T, s Set, max int) {
	var id map[int]bool = make(map[int]bool)
	for i := 0; i < max; i++ {
		id[i] = true
	}
	for x := range s.Iter() {
		id[x] = false, false
	}
	if len(id) > 0 {
		t.Errorf("not all values in set returned by Iter()")
	}
}

func checkRemove(t *testing.T, s Set, max int) {
	for i := 0; i < max; i++ {
		s.Remove(i)
		for j := 0; j <= i; j++ {
			if s.Has(j) {
				t.Errorf("remove of %v didn't work", i)
			}
		}
		for j := i+1; j < max; j++ {
			if !s.Has(j) {
				t.Errorf("remove of %v caused %v to be cleared", i, j)
			}
		}
	}
}

func checkAll(t *testing.T, s Set, max int) {
	checkEmpty(t, s, max)
	checkInsert(t, s, max)
	checkIter(t, s, max)
	checkRemove(t, s, max)
}

func TestSet(t *testing.T) {
	s := new(Bitset)
	s.Init(100)
	checkAll(t, s, 100)
	r := new(Briggs)
	r.Init(100)
	checkAll(t, r, 100)
	q := new(Hash)
	q.Init(100)
	checkAll(t, q, 100)
	m := new(Williams)
	m.Init(100)
	checkAll(t, m, 100)
	e := new(Simple)
	e.Init(100)
	checkAll(t, e, 100)
}

func benchIt(b *testing.B, s Set) {
	s.Init(b.N)
	for i := 0; i < b.N; i++ {
		s.Insert(i)
	}
	for i := 0; i < b.N; i++ {
		s.Has(i)
	}
	for i := 0; i < b.N; i++ {
		s.Remove(i)
	}
}

func benchRandom(b *testing.B, s Set, max int) {
	s.Init(max)
	for i := 0; i < b.N; i++ {
		s.Insert(rand.Int() % (max+1))
		s.Remove(rand.Int() % (max+1))
	}
}

func BenchmarkBitset(b *testing.B) {
	b.StopTimer()
	s := new(Bitset)
	b.StartTimer()
	benchIt(b, s)
}

func BenchmarkBriggs(b *testing.B) {
	b.StopTimer()
	s := new(Briggs)
	b.StartTimer()
	benchIt(b, s)
}

func BenchmarkHash(b *testing.B) {
	b.StopTimer()
	s := new(Hash)
	b.StartTimer()
	benchIt(b, s)
}

func BenchmarkWilliams(b *testing.B) {
	b.StopTimer()
	s := new(Williams)
	b.StartTimer()
	benchIt(b, s)
}

func BenchmarkSimple(b *testing.B) {
	b.StopTimer()
	s := new(Simple)
	b.StartTimer()
	benchIt(b, s)
}

const (
	SMALL = 1000
	LARGE = 100000000
)

func BenchmarkBitsetRandomDense(b *testing.B) {
	b.StopTimer()
	s := new(Bitset)
	b.StartTimer()
	benchRandom(b, s, SMALL)
}

func BenchmarkBitsetRandomSparse(b *testing.B) {
	b.StopTimer()
	s := new(Bitset)
	b.StartTimer()
	benchRandom(b, s, LARGE)
}

func BenchmarkWilliamsRandomDense(b *testing.B) {
	b.StopTimer()
	s := new(Williams)
	b.StartTimer()
	benchRandom(b, s, SMALL)
}

func BenchmarkWilliamsRandomSparse(b *testing.B) {
	b.StopTimer()
	s := new(Williams)
	b.StartTimer()
	benchRandom(b, s, LARGE)
}

func BenchmarkBriggsRandomDense(b *testing.B) {
	b.StopTimer()
	s := new(Briggs)
	b.StartTimer()
	benchRandom(b, s, SMALL)
}

func BenchmarkBriggsRandomSparse(b *testing.B) {
	b.StopTimer()
	s := new(Briggs)
	b.StartTimer()
	benchRandom(b, s, LARGE)
}
