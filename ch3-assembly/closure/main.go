package main

import (
	"fmt"
	"unsafe"
)

type FunTwiceClosure struct {
	F uintptr
	X int
}

func ptrToFunc(p unsafe.Pointer) func() int

func asmFunTwiceClosureAddr() uintptr

func asmFunTwiceClosureBody() int

func NewTwiceFunClosure(x int) func() int {
	var p = &FunTwiceClosure{
		F: asmFunTwiceClosureAddr(),
		X: x,
	}
	return ptrToFunc(unsafe.Pointer(p))
}

func main() {
	f := NewTwiceFunClosure(2)
	fmt.Println(f())
}
