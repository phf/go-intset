include ../../../Make.$(GOARCH)

TARG=container/intset
GOFILES=\
	low.go\
	set.go\
	bitset.go\
	sparseset.go\
	hashset.go\
	willset.go\
	simpleset.go\

include ../../../Make.pkg
