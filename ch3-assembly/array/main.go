package main

import (
	"fmt"

	"github.com/sisyphus/go-advanced-programming/ch3-assembly/array/pkg"
)

func main() {
	for _, num := range pkg.Num {
		fmt.Println(num)
	}
}
