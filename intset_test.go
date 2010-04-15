// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package intset

import (
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
	r := new(Sparse)
	r.Init(100)
	checkAll(t, r, 100)
	q := new(Hash)
	q.Init(100)
	checkAll(t, q, 100)
	m := new(Williams)
	m.Init(100)
	checkAll(t, m, 100)
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

func BenchmarkBitset(b *testing.B) {
	b.StopTimer()
	s := new(Bitset)
	b.StartTimer()
	benchIt(b, s)
}

func BenchmarkSparse(b *testing.B) {
	b.StopTimer()
	s := new(Sparse)
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

