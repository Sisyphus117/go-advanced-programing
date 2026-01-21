package main

import (
	"fmt"

	"github.com/klauspost/cpuid/v2"
)

func CopySlice_AVX2(dst, src []byte, len int)

func main() {
	if cpuid.CPU.Has(cpuid.AVX2) {
		dst := make([]byte, 5)
		src := []byte("abcde")
		CopySlice_AVX2(dst, src, len(src))

		fmt.Println(string(dst))
	}
}
