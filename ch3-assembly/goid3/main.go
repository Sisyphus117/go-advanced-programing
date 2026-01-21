package main

import (
	"fmt"
	"reflect"
)

func getg() any

func GetGoid() int64 {
	g := getg()
	goid := reflect.ValueOf(g).FieldByName("goid").Int()
	return goid
}

func main() {
	goid := GetGoid()
	fmt.Println(goid)
}
