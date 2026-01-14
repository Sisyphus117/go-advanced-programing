package main

/*
#include<stdlib.h>
extern int compare(void* a,void* b);
typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
*/
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

var sort_info struct {
	fn func(a, b unsafe.Pointer) int
	sync.Mutex
}

func Sort(base unsafe.Pointer, num, size int, cmp func(a, b unsafe.Pointer) int) {
	sort_info.Lock()
	defer sort_info.Unlock()

	sort_info.fn = cmp
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(C.compare))
}

//export compare
func compare(a, b unsafe.Pointer) C.int {
	return C.int(sort_info.fn(a, b))
}

func main() {
	arr := []int32{23, 1, 44, 32, 3, 91}
	Sort(unsafe.Pointer(&arr[0]), len(arr), int(unsafe.Sizeof(arr[0])), func(a, b unsafe.Pointer) int {
		pa, pb := (*int32)(a), (*int32)(b)
		return int(*pa) - int(*pb)
	})
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
