package main
// TODO Update package name here

import (
	"fmt"
	"runtime"
)

/*
#include <stdlib.h>
// Passing "-lwaku" to the Go linker through "-extldflags" is not enough. We need it in here, for some reason.
#cgo LDFLAGS: -Wl,-rpath,'$ORIGIN' -L${SRCDIR}/../build -lwaku

#include "libwaku.h"
*/
import "C"

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

// Wrapper that uses C library instead

// Just call info here?
func Start() {
	C.NimMain()
	var str = C.info("hello there")
	fmt.Println("String", str)
}

func main() {
	fmt.Println("Hi main")
	Start()
}
