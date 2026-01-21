package main

import "fmt"

func main() {
	nums := make([]int, 5)
	for i := range nums {
		nums[i] = i * i
	}
	fmt.Println(nums)
}
