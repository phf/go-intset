include ../../../Make.$(GOARCH)

TARG=container/intset
GOFILES=\
	set.go\
	bitset.go\
	sparseset.go\
	hashset.go\

include ../../../Make.pkg
