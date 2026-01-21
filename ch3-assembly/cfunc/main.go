package main

// #include "myadd.h"
import "C"
import (
	"fmt"

	"github.com/sisyphus/go-advanced-programming/ch3-assembly/cfunc/asm"
)

func main() {
	ans := asm.AsmCallCAdd(uintptr(C.myadd), 14, 3)
	fmt.Println(ans)
}
