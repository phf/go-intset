package intset

import (
	"unsafe"
)

// Kludge to figure out how many bits an int has. If Go ever
// runs on machines where a byte is *not* 8 bits... Trouble!

var bits_per_int int = unsafe.Sizeof(bits_per_int)*8
