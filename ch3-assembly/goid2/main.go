package main

import (
	"fmt"
	"unsafe"
)

const g_goid_offset = 152

func getg() unsafe.Pointer

func GetGoroutineId() int64 {
	g := getg()
	p := (*int64)(unsafe.Add(g, g_goid_offset))
	return *p
}

func main() {
	fmt.Println(GetGoroutineId())
}
