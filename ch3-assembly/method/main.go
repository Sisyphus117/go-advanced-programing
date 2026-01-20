package main

import "fmt"

type myInt int

func (v myInt) twice() int

func main() {
	a := myInt(3)
	a.twice()
	fmt.Println(a)
}
