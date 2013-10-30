package intset

import (
	"unsafe"
)

// Kludge to figure out how many bits an int has. If Go ever
// ran on machines where a byte is *not* 8 bits wide... We'd
// be screwed... :-/

var bitsPerInt int = int(unsafe.Sizeof(bitsPerInt) * 8)
