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
	checkRemove(t, s, max)
}

func TestSet(t *testing.T) {
	s := new(Bitset)
	s.Init(100)
	checkAll(t, s, 100)
}
