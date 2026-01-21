package main

import (
	"fmt"
	"runtime"
)

func SyscallWrite(fd int, msg string) int

func main() {
	fmt.Printf("Current OS: %s\n", runtime.GOOS)

	SyscallWrite(1, "hello syscall from "+runtime.GOOS+"!\n")
}
