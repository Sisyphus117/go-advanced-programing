package main

import (
	"fmt"
	"sync"

	"github.com/sisyphus/go-advanced-programming/ch3-assembly/goid4/gls"
)

func main() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			defer gls.Clean()

			defer func() {
				fmt.Printf("%d number=%d\n", idx, gls.Get("number"))
			}()
			gls.Put("number", idx+100)
		}(i)
	}
	wg.Wait()
}
