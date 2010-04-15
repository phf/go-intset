include ../../../Make.$(GOARCH)

TARG=container/intset
GOFILES=\
	low.go\
	set.go\
	bitset.go\
	briggsset.go\
	hashset.go\
	williamsset.go\
	simpleset.go\

include ../../../Make.pkg
