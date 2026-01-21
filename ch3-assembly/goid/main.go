package main

import "runtime"

func main() {
	//panic("goid")
	var buf = make([]byte, 64)
	var stk = buf[:runtime.Stack(buf, false)]
	print(string(stk))
}
