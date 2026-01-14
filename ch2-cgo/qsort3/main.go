package main

/*
#include<stdlib.h>
extern int compare(void* a,void* b);
typedef int (*qsort_cmp_func_t)(const void* a, const void* b);
*/
import "C"
import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var sort_info struct {
	base     unsafe.Pointer
	elemnum  int
	elemsize int
	less     func(a, b int) bool
	sync.Mutex
}

//export compare
func compare(a, b unsafe.Pointer) C.int {
	var (
		base     = uintptr(sort_info.base)
		elemsize = uintptr(sort_info.elemsize)
	)

	i := int((uintptr(a) - base) / elemsize)
	j := int((uintptr(b) - base) / elemsize)

	switch {
	case sort_info.less(i, j):
		return -1
	case sort_info.less(j, i):
		return 1
	default:
		return 0
	}
}

func Slice(slice any, less func(a, b int) bool) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("qsort called with non-slice value of type %T", slice))
	}
	if sv.Len() == 0 {
		return
	}

	sort_info.Lock()
	defer sort_info.Unlock()

	defer func() {
		sort_info.base = nil
		sort_info.elemnum = 0
		sort_info.elemsize = 0
		sort_info.less = nil
	}()

	sort_info.base = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	sort_info.elemnum = sv.Len()
	sort_info.elemsize = int(sv.Type().Elem().Size())
	sort_info.less = less

	C.qsort(sort_info.base, C.size_t(sort_info.elemnum), C.size_t(sort_info.elemsize), C.qsort_cmp_func_t(C.compare))
}

func main() {
	arr := []int32{23, 1, 44, 32, 3, 91}
	Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
	arr2 := []string{"sdv", "apple", "bool", "Bob"}
	Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	for _, num := range arr2 {
		fmt.Printf("%s ", num)
	}
	fmt.Println()
}
