package main

/*
#include <stdio.h>
#include <stdlib.h>
int display(){
    return 3;
}
*/
import "C"

import (
    "unsafe"
    "fmt"
)

func main() {
    cstr := C.CString("Hello, world")
    C.puts(cstr)
    fmt.Println(C.display())
    C.free(unsafe.Pointer(cstr))
}
