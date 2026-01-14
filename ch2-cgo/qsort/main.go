package main

//extern int compare(void* a,void* b);
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/sisyphus/go-advanced-programming/ch2-cgo/qsort/qsort"
)

//export compare
func compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func main() {
	arr := []int32{23, 1, 44, 32, 3, 91}
	qsort.Sort(unsafe.Pointer(&arr[0]), len(arr), int(unsafe.Sizeof(arr[0])), qsort.SortFunc(C.compare))
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
