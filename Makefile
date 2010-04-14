include ../../../Make.$(GOARCH)

TARG=container/intset
GOFILES=\
	set.go\
	bitset.go\
	sparseset.go\

include ../../../Make.pkg
