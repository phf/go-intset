include ../../../Make.$(GOARCH)

TARG=container/intset
GOFILES=\
	set.go\
	bitset.go\

include ../../../Make.pkg
