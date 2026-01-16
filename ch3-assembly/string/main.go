package main

import "github.com/sisyphus/go-advanced-programming/ch3-assembly/string/pkg"

func main() {
	println(pkg.Name)

	pkg.NameData[0] = '?'
	println(pkg.Name)
}
