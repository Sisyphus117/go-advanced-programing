package main

import "fmt"

func If(ok bool, a, b int) int

func main() {
	x := If(false, 1, 5)
	fmt.Println(x)
	y := If(true, 1, 5)
	fmt.Println(y)
}
